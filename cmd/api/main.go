package main

import (
	"broker/cmd/api/infrastruture"
	"broker/cmd/api/infrastruture/cassandra"
	"broker/cmd/api/routes"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {

	maybeSession := cassandra.ConnectDB("127.0.0.1", "my_keyspace")
	if maybeSession.IsError() {
		log.Panic(maybeSession.Error())
		return
	}
	userRepository := infrastruture.CreateRepository(maybeSession.MustGet())
	app := routes.Config{UserRepository: userRepository}

	log.Printf("starting broker on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
