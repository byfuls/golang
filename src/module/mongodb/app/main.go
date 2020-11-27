package main

// DOWNLOAD
// : go get go.mongodb.org/mongo-driver
// REF
// : https://github.com/tfogo/mongodb-go-tutorial
// : https://github.com/mongodb/mongo-go-driver/blob/master/mongo/client.go#L98
// : https://godoc.org/go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"log"
	//"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

type Test struct {
	Name string
}

type SimInfo struct {
	//_id.imsi			string
	//sim_pid			number
	//user_pid			number
	//user_id			string
	//user_tid			string
	//lur_date			number
	//mcc				string
	//mnc				string
	//lac				string
	//cell_id			string
	//bsic				string
	//arfcn				number
	//tmsi				string
	//kc				string
	//cksn				number
	//msisdn			number
	//sim_id			string
	//sim_serial_no		number
	//simbank_name		string
	//connector			string
	//imei				string
	//mobile_type		number
	//doing_promo		number
	//sim_balance		number
	//out_call_time		number
	//expire_match_date	number
	//lur_fail_cnt		number
	//send_sms_cnt		number
	//sim_state			number
	//sim_expire_date	number
	//charging			number
	//sim_type			number
	//lur_check			number
	//etc_balance_flag	number
	//etc_msisdn_flag	number
	//etc_charge_flag	number

	Id                string
	Imsi              string
	Sim_pid           int
	User_pid          int
	User_id           string
	User_tid          string
	Lur_date          int
	Mcc               string
	Mnc               string
	Lac               string
	Cell_id           string
	Bsic              string
	Arfcn             int
	Tmsi              string
	Kc                string
	Cksn              int
	Msisdn            int
	Sim_id            string
	Sim_serial_no     int
	Simbank_name      string
	Connector         string
	Imei              string
	Mobile_type       int
	Doing_promo       int
	Sim_balance       int
	Out_call_time     int
	Expire_match_date int
	Lur_fail_cnt      int
	Send_sms_cnt      int
	Sim_state         int
	Sim_expire_date   int
	Charging          int
	Sim_type          int
	Lur_check         int
	Etc_balance_flag  int
	Etc_msisdn_flag   int
	Etc_charge_flag   int
}

func mongoConnect() (client *mongo.Client) {
	// set client options
	uri := "mongodb+srv://ttgo-200318-c1.jchgw.mongodb.net/ttgo?authSource=%24external&tlsCertificateKeyFile=../X509-cert-734204199706946844.pem"
	credential := options.Credential{
		Username:      "ttgo",
		Password:      "ttgo",
		AuthMechanism: "MONGODB-X509",
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")

	return client
}

func mongoDisconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func main() {
	fmt.Println("start")

	conn := mongoConnect()
	collection := conn.Database("ttgo").Collection("sim_info")
	fmt.Println(collection)

	// [FIND] one
	filter := bson.D{{"_id.imsi", "5101025824020177"}}
	//filter := bson.D{{"name", "Ash"}}
	var result SimInfo
	//var result Trainer
	//var result interface{}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		//fmt.Println("===", result)
		//fmt.Println("===!!!===", reflect.TypeOf(result.Sim_pid), reflect.ValueOf(result.Sim_pid))
		//fmt.Println("===!!!===", reflect.TypeOf(result.Id), reflect.ValueOf(result.Id))
		//fmt.Println("===!!!===", reflect.TypeOf(result.Imsi), reflect.ValueOf(result.Imsi))
		//fmt.Println("===!!!===", reflect.TypeOf(result.User_tid), reflect.ValueOf(result.User_tid))
	}
}
