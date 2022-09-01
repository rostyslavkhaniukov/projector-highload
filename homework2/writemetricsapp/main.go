package main

import (
	"fmt"
	"log"
)
import "github.com/cactus/go-statsd-client/v5/statsd"

func main() {
	fmt.Println("Hello")

	config := &statsd.ClientConfig{
		Address: "127.0.0.1:8125",
		Prefix:  "test-client",
	}
	client, err := statsd.NewClientWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	client.Inc("stat1", 42, 1.0)
}
