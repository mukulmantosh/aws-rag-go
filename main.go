package main

import (
	"log"
	"net/http"
)

func main() {
	bedrockAgent := NewBedrock()
	http.HandleFunc("/send-message", ProcessLLMModel(bedrockAgent))
	log.Println("Server started, listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
