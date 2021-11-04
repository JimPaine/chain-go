package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/go", handle)

	log.Println("** Service Started on Port 80 **")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	responseMessage := os.Getenv("RESPONSE_MESSAGE")
	if len(responseMessage) == 0 {
		responseMessage = "Hello from Go link"
	}

	nextHop := os.Getenv("NEXT_HOP")
	nestedMessage := ""

	if len(nextHop) > 0 {
		log.Println("nextHop value:" + nextHop)
		resp, err := http.Get(nextHop)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		nestedMessage = string(body[:])
	}

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, "{\"response\":\""+responseMessage+"\"")
	if len(nestedMessage) > 0 {
		io.WriteString(w, ", \"nestedResponse\":\"{\"response\":\""+nestedMessage+"\"}}")
	} else {
		io.WriteString(w, "}")
	}
}
