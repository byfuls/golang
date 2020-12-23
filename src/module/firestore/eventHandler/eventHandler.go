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

type snapshotsDocs struct {
	docs  []interface{}
	count int
}

type EventHandler struct {
	ctx  context.Context
	ctm  context.Context
	sa   option.ClientOption
	app  *firebase.App
	conn *firestore.Client

	SnapShotDocs chan []EventA
}

func (e *EventHandler) Snapshots(coll string) {
	for {
		snapIter := e.conn.Collection(coll).Snapshots(e.ctx)
		var snapCount int
		var docs []EventA
		defer snapIter.Stop()
		for {
			snap, err := snapIter.Next()
			if err != nil {
				if err == iterator.Done {
					fmt.Println("get snapshot done: ", err)
				}
				fmt.Println("snapshot iter err: ", err)
			}
			snapCount = len(snap.Changes)
			fmt.Printf("change size: %d\n", snapCount)

			for _, diff := range snap.Changes {
				//fmt.Printf("diff: %+v\n", diff)
				//fmt.Printf("doc: %+v\n", diff.Doc)
				//fmt.Printf("doc: %+v\n", diff.Doc.Ref)

				fmt.Println("1111")

				/* get document */
				var eventAData EventA
				if err := diff.Doc.DataTo(&eventAData); err != nil {
					if err == iterator.Done {
						fmt.Println("get document done: ", err)
						break
					}
					fmt.Println("get document err: ", err)
				} else {
					eventAData.Id = diff.Doc.Ref.ID
					fmt.Println(eventAData)

					docs = append(docs, eventAData)
					fmt.Println("put=>")
					fmt.Println(docs)
				}
			}
			e.SnapShotDocs <- docs
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
