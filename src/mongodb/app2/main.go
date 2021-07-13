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
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type id_t struct {
	Id string `bson:"id"`
}

type UserInfo struct {
	Id           id_t      `bson:"_id"`
	Uuid         string    `bson:"uuid"`
	Tid          string    `bson:"tid"`
	Imsi         string    `bson:"imsi"`
	Reg_date     time.Time `bson:"reg_date"`
	Last_in_time time.Time `bson:"last_in_time"`
	Mobile_os    string    `bson:"mobile_os"`
	//Apns_key     string    `bson:"apns_key"` /* will be removed */
	Voip_key     string `bson:"voip_key"`
	Fcm_key      string `bson:"fcm_key"`
	Mobile_no    string `bson:"mobile_no"`
	Firm_version string `bson:"firm_version"`
	Code         int    `bson:"code"`
	Call         int    `bson:"call"`
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
	collection := conn.Database("ttgo").Collection("user_info")
	fmt.Println(collection)

	// [FIND] one
	filter := bson.D{{"_id.id", "0038"}}
	var result UserInfo
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		fmt.Printf("err: %s\n", err)
		log.Fatal(err)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		fmt.Println("________________________")
		fmt.Printf("id           : (%s:%03d)[%s]\n", reflect.TypeOf(result.Id.Id), len(result.Id.Id), result.Id.Id)
		fmt.Printf("uuid         : (%s:%03d)[%s]\n", reflect.TypeOf(result.Uuid), len(result.Uuid), result.Uuid)
		fmt.Printf("tid          : (%s:%03d)[%s]\n", reflect.TypeOf(result.Tid), len(result.Tid), result.Tid)
		fmt.Printf("imsi         : (%s:%03d)[%s]\n", reflect.TypeOf(result.Imsi), len(result.Imsi), result.Imsi)
		fmt.Printf("reg_date     : (%s) [%s]\n", reflect.TypeOf(result.Reg_date), result.Reg_date)
		fmt.Printf("last_in_time : (%s) [%s]\n", reflect.TypeOf(result.Last_in_time), result.Last_in_time)
		fmt.Printf("mobile_os    : (%s:%03d)[%s]\n",
			reflect.TypeOf(result.Mobile_os), len(result.Mobile_os), result.Mobile_os)
		fmt.Printf("voip_key     : (%s:%03d)[%s]\n",
			reflect.TypeOf(result.Voip_key), len(result.Voip_key), result.Voip_key)
		fmt.Printf("fcm_key      : (%s:%03d)[%s]\n",
			reflect.TypeOf(result.Fcm_key), len(result.Fcm_key), result.Fcm_key)
		fmt.Printf("mobile_no    : (%s:%03d)[%s]\n",
			reflect.TypeOf(result.Mobile_no), len(result.Mobile_no), result.Mobile_no)
		fmt.Printf("firm_version : (%s:%03d)[%s]\n",
			reflect.TypeOf(result.Firm_version), len(result.Firm_version), result.Firm_version)
		fmt.Printf("code         : (%s)       (%d)\n", reflect.TypeOf(result.Code), result.Code)
		fmt.Printf("call         : (%s)       (%d)\n", reflect.TypeOf(result.Call), result.Call)
	}
}
