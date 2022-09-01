package main

import (
	"context"
	"fmt"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

const min = 0

const max = 100

var ctx = context.TODO()

func main() {
	somethingWithMongoDB()
	somethingWithES()
	somethingWithOwnMetrics()
}

func somethingWithMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = client.Database("test").Collection("test")
}

func somethingWithES() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
}

func somethingWithOwnMetrics() {
	config := &statsd.ClientConfig{
		Address: "telegraf:8125",
		Prefix:  "projector",
	}
	client, err := statsd.NewClientWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(3 * time.Second)
		client.Inc("users", rand.Int63n(max-min)+min, 1.0)
		fmt.Println("Metric has been written")
	}
}
