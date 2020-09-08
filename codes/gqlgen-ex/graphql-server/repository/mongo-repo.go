package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redstone-wb/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "graphql"
	COLLECTION = "videos"
)

// monbodb connect 목적
func New() VideoRepository {

	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(MONGODB) // URL

	clientOptions = clientOptions.SetMaxPoolSize(50) // Connection Pool

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second) // Connection timeout

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(("Connected to MongoDB!"))
	return &database{
		client: dbClient,
	}
}

// 저장하는 함수,, 중요!
func (db *database) Save(video *model.Video) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video) // video : document로 저장될 내용, 일종의 filter
	if err != nil {
		log.Fatal(err)
	}

}

func (db *database) FindAll() []*model.Video {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{}) // bson.D{}은 filter to get all the existing videos
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}
