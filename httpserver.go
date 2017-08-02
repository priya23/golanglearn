package main

import (
	"io"
	"net/http"
	"encoding/json"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
func getJsonResponse()([]byte, error) {
    fruits := make(map[string]int)
    fruits["Apples"] = 25
    fruits["Oranges"] = 10

    vegetables := make(map[string]int)
    vegetables["Carrats"] = 10
    vegetables["Beets"] = 0

    d := Data{fruits, vegetables}
    p := Payload{d}

    return json.MarshalIndent(p, "", "  ")
}