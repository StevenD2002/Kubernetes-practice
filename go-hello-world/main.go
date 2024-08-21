// this is going to write a simple message to a redis server
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {
	port := ":8080"
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()

		red := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // No password for local development
			DB:       0,
		})

		switch r.Method {
		case "GET":
			// Connect to Redis

			// Example: Get data from Redis
			val, err := red.Get(ctx, "mykey").Result()
			if err != nil {
				log.Println("Error fetching data from Redis:", err)
			}
			w.Write([]byte("Data from Redis: " + val))
		case "POST":
			// Example: Store data in Redis
			err := red.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
			if err != nil {
				log.Println("Error storing data in Redis:", err)
			}
			w.Write([]byte("Data stored in Redis"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
