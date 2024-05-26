package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var statedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config", ConfigMap)
	http.HandleFunc("/secrets", Secrets)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	envVarEx1 := os.Getenv("ENV1")
	envVarEx2 := os.Getenv("ENV2")
	w.Write([]byte("<h1>Hello People!!!</h1>"))

	fmt.Fprintf(w, "This is an example of using environment variables")

	fmt.Fprintf(w, "Env1 %s and Env2 %s!!!", envVarEx1, envVarEx2)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/config/config.txt")

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Fprintf(w, "Configs: %s ", string(data))
}

func Secrets(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "This is an example of using secrets on Kubernetes")
	fmt.Fprintf(w, "User %s and Password %s!!!", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(statedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
