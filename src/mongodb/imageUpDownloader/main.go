package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	mongoIP   = "localhost"
	mongoPort = "27017"
)

func initMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client

	uri := "mongodb://" + mongoIP + ":" + mongoPort
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(10)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Fatalln(err)
		return nil
	}
	return client
}

func uploadFile(databaseName, file, filename string) bool {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return false
	}
	conn := initMongoClient()
	if conn == nil {
		log.Fatalln("connect mongodb error")
		return false
	}
	defer conn.Disconnect(context.Background())
	bucket, err := gridfs.NewBucket(
		conn.Database(databaseName),
	)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	opts := options.GridFSUpload()
	opts.SetMetadata(bsonx.Doc{{Key: "key test", Value: bsonx.String("value test")}})

	uploadStream, err := bucket.OpenUploadStream(filename, opts)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	log.Printf("write file to DB was successful. file size: %d bytes\n", fileSize)
	return true
}

func downloadFile(database, fileName string) bool {
	conn := initMongoClient()

	db := conn.Database(database)
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	log.Println(results)

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	log.Printf("file size to download: %v\n", dStream)
	ioutil.WriteFile(fileName, buf.Bytes(), 0600)
	return true
}

func main() {
	databaseName := os.Args[1]
	file := os.Args[2]
	fileName := path.Base(file)

	//if !uploadFile(databaseName, file, fileName) {
	//	log.Fatalln("upload fail")
	//}

	if !downloadFile(databaseName, fileName) {
		log.Fatalln("download fail")
	}
}
