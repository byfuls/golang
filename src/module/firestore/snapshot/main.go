package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("../ttgo-b29bf-firebase-adminsdk-qg1jz-73e3e1bf64.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	/* snapshot */
	snapIter := client.Collection("test").Snapshots(ctx)
	defer snapIter.Stop()
	for {
		snap, err := snapIter.Next()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("change size: %d\n", len(snap.Changes))
		for _, diff := range snap.Changes {
			fmt.Printf("diff: %+v\n", diff)
			fmt.Printf("doc: %+v\n", diff.Doc)
			fmt.Printf("doc: %+v\n", diff.Doc.Ref)

			/* get document */
			type test struct {
				Id      string
				Capital string `firestore:"Capital"`
			}
			var s test
			if err := diff.Doc.DataTo(&s); err != nil {
				fmt.Println("conv error")
			} else {
				s.Id = diff.Doc.Ref.ID
				fmt.Println(s)
			}

			/* update new document */
			wr, err := client.Collection("test1").Doc("log").Collection("tt1").Doc("history").Update(
				ctx, []firestore.Update{
					{Path: "flag", Value: "go"},
				})
			if err != nil {
				fmt.Println("update err: ", err)
			} else {
				fmt.Println("update ok: ", wr)
			}

			/* delete before document */
			dl, err := client.Collection("test").Doc(s.Id).Delete(ctx)
			if err != nil {
				fmt.Println("delete err: ", err)
			} else {
				fmt.Println("delete ok: ", dl)
			}
		}
		fmt.Println("")
	}

	/* read data */
	//iter := client.Collection("test").Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}
}
