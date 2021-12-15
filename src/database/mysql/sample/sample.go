package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	id       string
	pw       string
	database string
	addr     string
	db       *sql.DB

	openConns int
	idleConns int
}

func initMySQL(id, pw, database, addr string, openConns, idleConns int) (*MySQL, error) {
	if 0 >= len(id) || 0 >= len(pw) || 0 >= len(database) || 0 >= len(addr) ||
		0 >= openConns || 0 >= idleConns {
		return nil, errors.New("check parameter again")
	}

	return &MySQL{
		id:        id,
		pw:        pw,
		database:  database,
		addr:      addr,
		openConns: openConns,
		idleConns: idleConns,
	}, nil
}

func (db *MySQL) conn() error {
	t_conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.id, db.pw, db.addr, db.database)
	t_db, err := sql.Open("mysql", t_conn)
	if err != nil || t_db.Ping() != nil {
		return err
	}
	db.db = t_db
	db.db.SetConnMaxLifetime(time.Minute * 3)
	db.db.SetMaxOpenConns(10)
	db.db.SetMaxIdleConns(10)
	if _, err := db.db.Exec("set time_zone='Asia/Seoul'"); err != nil {
		return err
	}
	if err := db.prepare(); err != nil {
		return err
	}
	return nil
}

func (db *MySQL) prepare() error {
	conn := db.db

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	query := `create table if not exists test(
				idx int NOT NULL AUTO_INCREMENT PRIMARY KEY,
				date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				min int NOT NULL,
				max int NOT NULL,
				avg float NOT NULL,
				pid varchar(12) NOT NULL
			  )`
	_, err := conn.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// Test Table , Select All
type Test struct {
	Idx  int
	Date string
	Min  int
	Max  int
	Avg  float32
	Pid  string
}

func (db *MySQL) Test_selectAll() ([]Test, error) {
	conn := db.db
	// defer conn.Close()

	rows, err := conn.Query("select idx, date, min, max, avg, pid from test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tests := []Test{}
	test := Test{}
	for rows.Next() {
		err := rows.Scan(&test.Idx, &test.Date, &test.Min, &test.Max, &test.Avg, &test.Pid)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tests, nil
}

func (db *MySQL) Test_selectOne(idx int) (*Test, error) {
	conn := db.db

	test := &Test{}
	err := conn.QueryRow("select idx, date, min, max, avg, pid from test where idx = ?", idx).
		Scan(&test.Idx, &test.Date, &test.Min, &test.Max, &test.Avg, &test.Pid)
	if err != nil {
		return nil, err
	}
	return test, nil
}

func (db *MySQL) Test_insertOne() (int64, error) {
	conn := db.db
	// defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	res, err := tx.Exec("insert into test (min, max, avg, pid) values (?, ?, ?, ?)", 2, 3, 4.0, "pid")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return lastId, nil
}

func (db *MySQL) Test_update(idx int, min int) error {
	conn := db.db

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("update test set min = ? where idx = ?", min, idx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (db *MySQL) Test_delete(idx int) error {
	conn := db.db

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("delete from test where idx = ?", idx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func main() {
	db, err := initMySQL("watchers", "watchers", "watchers", "127.0.0.1:3306", 10, 10)
	if err != nil {
		panic(err)
	}
	if err := db.conn(); err != nil {
		panic(err)
	}

	// INSERT
	if lastId, err := db.Test_insertOne(); err != nil {
		panic(err)
	} else {
		fmt.Printf("last id: %v\n", lastId)
	}

	// DELETE
	// if err := db.Test_delete(1); err != nil {
	// 	panic(err)
	// }

	// UPDATE
	// if err := db.Test_update(2, 7); err != nil {
	// 	panic(err)
	// }

	// SELECT - Single
	// selected, err := db.Test_selectOne(3)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("[selectedOne] Idx  : %d\n", selected.Idx)
	// 	fmt.Printf("[selectedOne] Date : %s\n", selected.Date)
	// 	fmt.Printf("[selectedOne] Min  : %d\n", selected.Min)
	// 	fmt.Printf("[selectedOne] Max  : %d\n", selected.Max)
	// 	fmt.Printf("[selectedOne] Avg  : %f\n", selected.Avg)
	// 	fmt.Printf("[selectedOne] Pid  : %s\n", selected.Pid)
	// }

	// SELECT - Multi
	selected, err := db.Test_selectAll()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(selected); i++ {
		fmt.Printf("[selected:%d] Idx  : %d\n", i, selected[i].Idx)
		fmt.Printf("[selected:%d] Date : %s\n", i, selected[i].Date)
		fmt.Printf("[selected:%d] Min  : %d\n", i, selected[i].Min)
		fmt.Printf("[selected:%d] Max  : %d\n", i, selected[i].Max)
		fmt.Printf("[selected:%d] Avg  : %f\n", i, selected[i].Avg)
		fmt.Printf("[selected:%d] Pid  : %s\n", i, selected[i].Pid)
	}
}
