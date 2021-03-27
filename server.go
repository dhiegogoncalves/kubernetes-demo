package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Healthz(response http.ResponseWriter, request *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		response.WriteHeader(500)
		response.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		response.WriteHeader(200)
		response.Write([]byte("ok"))
	}
}

func Secret(response http.ResponseWriter, request *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(response, "User: %s, Password: %s.", user, password)
}

func ConfigMap(response http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(response, "My Family: %s.", string(data))
}

func Hello(response http.ResponseWriter, request *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(response, "Hello, I'm %s. I'm %s years old.", name, age)
}
