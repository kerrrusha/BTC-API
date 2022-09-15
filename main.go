package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kerrrusha/btc-api/api/rest"
)

func getPort() string {
	PORT, presented := os.LookupEnv("PORT")
	if !presented {
		PORT = "8000"
	}
	return PORT
}
func handleRequests() {
	PORT := getPort()
	log.Println("Started server at " + PORT + " port")

	http.HandleFunc("/rate/", rest.Rate)
	http.HandleFunc("/subscribe/", rest.Subscribe)
	http.HandleFunc("/sendEmails/", rest.SendRateEmails)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

func main() {
	handleRequests()
}
