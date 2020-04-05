package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Go"))
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
