package database

import (
	"context"
	"log"
	"time"

	"github.com/Damilola99-web/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://localhost:27017/"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{client: client}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database("job").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var JobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&JobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &JobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var JobListings []*model.JobListing
	return JobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	var JobListing model.JobListing
	return &JobListing
}

func (db *DB) UpdateJobListing(id string, jobInfo model.UpdateJobListing) *model.JobListing {
	var JobListing model.JobListing
	return &JobListing
}

func (db *DB) DeleteJobListing(id string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{JobID: id}
}
