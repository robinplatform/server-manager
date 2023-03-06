package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/getHelloWorld", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"type":"success","result":{"message": "Hello World"}}`))
	})
	router.HandleFunc("/api/RunAppMethod", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"type":"success","result":{"message": "Hello World"}}`))
	})
	router.HandleFunc("/api/health", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"status": "ok"}`))
	})

	server := http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			fmt.Printf("%s %s\n", req.Method, req.URL.Path)

			res.Header().Set("Content-Type", "application/json")
			router.ServeHTTP(res, req)
		}),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
