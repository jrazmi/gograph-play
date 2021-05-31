package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/jrazmi/gograph-play/internal/schema"
)

func init() {
	u1 := schema.User{ID: "foo", FirstName: "Josh", Admin: true}
	u2 := schema.User{ID: "bar", FirstName: "Baz", Admin: false}
	schema.UserList = append(schema.UserList, u1, u2)

}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	h := handler.New(&handler.Config{
		Schema:   &schema.UserSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	port := os.Getenv("PORT")
	if "" == port {
		port = "8080"
	}
	http.Handle("/graphql", h)
	fmt.Printf("Listening on localhost:%v", port)
	http.ListenAndServe(":"+port, nil)

}
