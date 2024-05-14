package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config", ConfigMap)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	envVarEx1 := os.Getenv("ENV1")
	envVarEx2 := os.Getenv("ENV2")
	w.Write([]byte("<h1>Hello People!!!</h1>"))

	w.Write([]byte("<h1>This is an example of using environment variables</h1>"))

	fmt.Fprintf(w, "Env1 %s and Env2 %s!!!", envVarEx1, envVarEx2)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/config/config.txt")

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Fprintf(w, "Configs: %s ", string(data))
}
