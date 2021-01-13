package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	go log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir("/Users/ap7u5/IdeaProjects/vm-docker-comparison/startTimeSimulator/data"))))
	mux := http.NewServeMux()
	mux.HandleFunc("/docker", getHandlerFunc(1072))
	mux.HandleFunc("/vmware", getHandlerFunc(9648))
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

type handlerFunc = func(w http.ResponseWriter, r *http.Request)

func getHandlerFunc(delayInMs time.Duration) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayInMs * time.Millisecond)
		w.Header().Add("Location", "http://localhost:8081"+r.URL.Path+".png")
		w.WriteHeader(302)
	}
}
