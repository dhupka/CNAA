package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = "mongodb://172.17.0.2:27017" // Find this from the Mongo container
)


type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Body      string             `bson:"body"`
	Tags      []string           `bson:"tags"`
	Comments  uint64             `bson:"comments"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongodbEndpoint)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var userDB = db().Database("blog").Collection("posts")

func main() {
	http.HandleFunc("/list", list)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


//curl "http://localhost:8000/create?title=recipes&body=recipeSteps"
func create(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	body := req.URL.Query().Get("body")
	fmt.Println(title)

	res, _ := userDB.InsertOne(context.TODO(), &Post{
		ID:        primitive.NewObjectID(),
		Title:     title,
		Tags:      []string{"foods"},
		Body:      body,
		CreatedAt: time.Now(),
	})
	fmt.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())

}

//curl "http://localhost:8000/list"
func list(w http.ResponseWriter, req *http.Request) {
	cursor, err := userDB.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var results bson.M
		if err = cursor.Decode(&results); err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
	}
	
}

//curl "http://localhost:8000/update?title=recipe&body=soupRecipes"
func update(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	body := req.URL.Query().Get("body")
	fmt.Println(title)
	
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	filter := bson.M{"title": title}
	update := bson.M{ "$set":bson.M{"body": body}}
	res := userDB.FindOneAndUpdate(context.TODO(), filter,update, &opt)
	_ = res
}

//curl "http://localhost:8000/delete?title=socks"
func delete(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	res, _ := userDB.DeleteOne(context.TODO(), bson.D{{"title", title}})
	fmt.Printf("deleted docs: %s\n", res.DeletedCount)
}
// func checkError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
