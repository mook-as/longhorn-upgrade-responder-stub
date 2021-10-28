package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening on :8000...")
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			body = []byte(fmt.Sprintf("Error: %s", err))
		}
		fmt.Printf("%s %s\n", req.URL.Path, string(body))

		input, err := os.Open("response.json")
		if err != nil {
			fmt.Printf("Error reading response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(w, input)
		if err != nil {
			fmt.Printf("Error copying response: %s", err)
		}
	}))
}
