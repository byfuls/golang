package eventHandler

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type EventHandler struct {
	ctx  context.Context
	ctm  context.Context
	sa   option.ClientOption
	app  *firebase.App
	conn *firestore.Client

	SnapShotDocs chan []EventA
}

func (e *EventHandler) AddDoc(collFullPath string, doc EventA) {
	docRef, writeResult, error := e.conn.Collection(collFullPath).Add(e.ctx, doc)
	if error != nil {
		fmt.Println("add doc error: ", error)
		fmt.Println("add doc error doc ref: ", docRef)
		fmt.Println("add doc error write result: ", writeResult)
	} else {
		fmt.Println("add doc ok doc ref: ", docRef)
		fmt.Println("add doc ok write result: ", writeResult)
	}
}

func (e *EventHandler) UpdateDoc(docFullPath string, result string) {
	ul, err := e.conn.Doc(docFullPath).Update(e.ctx, []firestore.Update{
		{Path: "status", Value: result},
	})
	if err != nil {
		fmt.Println("update error: ", err)
	} else {
		fmt.Println("update ok: ", ul)
	}
}

func (e *EventHandler) Snapshots(coll string) {
	for {
		fmt.Println("start...")
		snapIter := e.conn.Collection(coll).Snapshots(e.ctx)
		var snapCount int
		var docs []EventA
		defer snapIter.Stop()
		for {
			fmt.Println("go...")
			snap, err := snapIter.Next()
			if err != nil {
				if err == iterator.Done {
					fmt.Println("get snapshot done: ", err)
				}
				fmt.Println("snapshot iter err: ", err)
			}
			snapCount = len(snap.Changes)
			fmt.Printf("change size: %d\n", snapCount)
			if 0 >= snapCount {
				fmt.Println("again...")
				continue
			}

			for _, diff := range snap.Changes {
				//switch diff.Kind {
				//case firestore.DocumentAdded:
				//	fmt.Println("DOCUMENT ADDED")
				//case firestore.DocumentRemoved:
				//	fmt.Println("DOCUMENT REMOVED")
				//case firestore.DocumentModified:
				//	fmt.Println("DOCUMENT MODIFIED")
				//default:
				//	fmt.Println("UNKNOWN")
				//}
				//fmt.Printf("diff: %+v\n", diff)
				//fmt.Printf("doc: %+v\n", diff.Doc)
				//fmt.Printf("doc: %+v\n", diff.Doc.Ref)

				if diff.Kind != firestore.DocumentAdded {
					fmt.Println("only added document will be proceed")
					continue
				}

				/* get document */
				var doc EventA
				if err := diff.Doc.DataTo(&doc); err != nil {
					if err == iterator.Done {
						fmt.Println("get document done: ", err)
						break
					}
					fmt.Println("get document err: ", err)
				} else {
					doc.Id = diff.Doc.Ref.ID
					fmt.Println(doc)

					/* stack the got documents */
					docs = append(docs, doc)
					fmt.Println("put=>")
					fmt.Println(docs)

					/* delete after stacking documents */
					dl, err := e.conn.Collection(coll).Doc(doc.Id).Delete(e.ctx)
					if err != nil {
						fmt.Printf("delete error[%s]: %s", doc.Id, err)
					} else {
						fmt.Printf("delete ok[%s]: %+v", doc.Id, dl)
					}
				}
			}
			e.SnapShotDocs <- docs
			docs = []EventA{}
		}
	}
}

func (e *EventHandler) Term() {
	defer e.conn.Close()
}

func (e *EventHandler) Init(credentialsFile string) error {
	ctx := context.Background()
	e.ctx = ctx
	ctm, _ := context.WithTimeout(context.Background(), 3*time.Second)
	e.ctm = ctm
	e.sa = option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(e.ctx, nil, e.sa)
	if err != nil {
		return err
	} else {
		e.app = app
	}
	conn, err := e.app.Firestore(e.ctx)
	if err != nil {
		return err
	} else {
		e.conn = conn
	}

	e.SnapShotDocs = make(chan []EventA)
	return nil
}
