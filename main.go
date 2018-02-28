package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	ipListen := getEnv("IP_LISTEN", "")
	port := getEnv("PORT", "8080")

	address := fmt.Sprintf("%s:%s", ipListen, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	log.Println("Starting the server on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
	fmt.Fprintln(w, "Updated:", "Tue Feb 27 18:11:19 PST 2018")
}

func getEnv(envVar, defaultValue string) string {
	value, ok := os.LookupEnv(envVar)
	if ok {
		return value
	}
	return defaultValue
}
