package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Initializing system...")
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		log.Fatalln("listen pot has not been set!")
	}
	http.HandleFunc("/", postHandler)
	log.Println("Server initialized. Listening at :" + port)
	http.ListenAndServe(":"+port, nil)
}
