package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	envVarEx1 := os.Getenv("ENV1")
	envVarEx2 := os.Getenv("ENV2")
	w.Write([]byte("<h1>Hello People!!!</h1>"))

	w.Write([]byte("<h1>This is an example of using environment variables</h1>"))

	fmt.Fprintf(w, "Env1 %s and Env2 %s!!!", envVarEx1, envVarEx2)
}
