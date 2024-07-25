package main 

import(
	"fmt"
	"context"
	firebase"firebase.google.com/go"
	firestore"cloud.google.com/go/firestore"
	storage"firebase.google.com/go/storage"
	bucket"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func initializeApp() (*firebase.App, context.Context){
	opt := option.WithCredentialsFile("./sf6r-db-key.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil{
		fmt.Printf("error initializing app: %v", err)
	}
	return app, ctx
}

func initializeDBClient(app *firebase.App, ctx context.Context) *firestore.Client{
	db, err := app.Firestore(ctx)
	if err != nil{
		fmt.Printf("error initializing database: %v", err)
	}

	return db
}

func getCollection(collectionName string, dbClient *firestore.Client) *firestore.CollectionRef{
	coll := dbClient.Collection(collectionName)

	return coll
}

func readDocument(coll *firestore.CollectionRef, docName string) *firestore.DocumentRef{
	doc := coll.Doc(docName)
	return doc

	// getting data from doc
	// just do doc.Get(context)
	//
}

func getDocInfo(ctx context.Context, doc *firestore.DocumentRef) map[string]interface{}{
	
	snap, err := doc.Get(ctx)
	
	if err != nil{
		fmt.Printf("Couldnt read document info:", err)
	}

	return snap.Data()
}

func initializeStorage(app *firebase.App, ctx context.Context) *storage.Client{
	storage, err := app.Storage(ctx)
	if err != nil{
		fmt.Printf("Couldnt initialize storage:", err)		
	}

	return storage
}

func getBucket(client *storage.Client) *bucket.BucketHandle{
	bucket, _ := client.DefaultBucket()

	return bucket
}

func main(){
	app, ctx := initializeApp()
	dbClient := initializeDBClient(app, ctx)

	users := getCollection("users", dbClient) 
	doc := readDocument(users, "S3ZCvhf7dSPLSJdkc9p3")
	dataMap := getDocInfo(ctx, doc)

	fmt.Println(dataMap)

}
