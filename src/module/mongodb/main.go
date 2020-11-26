package main

// REF
// : go get go.mongodb.org/mongo-driver

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func mongoConnect() (client *mongo.Client) {
	// set client options
	uri := "mongodb+srv://ttgo-200318-c1.jchgw.mongodb.net/ttgo?authSource=%24external&tlsCertificateKeyFile=./X509-cert-734204199706946844.pem"
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
	collection := conn.Database("ttgo").Collection("training")
	fmt.Println(collection)

	// [INSERT] one
	ash := Trainer{"Ash", 10, "Pallet Town"}
	if insertResult, err := collection.InsertOne(context.TODO(), ash); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	// [INSERT} many]
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	trainers := []interface{}{misty, brock}
	if insertManyResult, err := collection.InsertMany(context.TODO(), trainers); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	}

	// [UPDATE] one
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	if updateResult, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Matched %v documents and updated %v documents. \n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}

	// [FIND] one
	var result Trainer
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		fmt.Println(result.Name)
		fmt.Println(result.Age)
		fmt.Println(result.City)
	}

	// [FIND]
	// pass these options to the find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// here's an array in which you can store the decoded documents
	var results []*Trainer

	// passing bson.D{{}} as the filter matches all documents in the collection
	if cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions); err != nil {
		log.Fatal(err)
	} else {
		// finding multiple documents returns a cursor
		// iterating through the cursor allows us to decode documents one at a time
		for cur.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			var elem Trainer
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// close the cursor once finished
		cur.Close(context.TODO())
		fmt.Printf("found multiple documents (array of pointers): %+v (count: %d)\n", results, len(results))
		for i := 0; i < len(results); i++ {
			fmt.Println("found index: ", i, " ", results[i].Name)
			fmt.Println("found index: ", i, " ", results[i].Age)
			fmt.Println("found index: ", i, " ", results[i].City)
		}
	}

	// [DELETE] one
	filter := bson.D{{"name", "Ash"}}
	if deleteResult, err := collection.DeleteOne(context.TODO(), filter); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Delete %v document in the trainers collection\n", deleteResult.DeletedCount)
	}

	// [DELETE] many
	filter := bson.D{{}}
	if deleteResult, err := collection.DeleteMany(context.TODO(), filter); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Delete %v documents in the trainers collection\n", deleteResult.DeletedCount)
	}
}
