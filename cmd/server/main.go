package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"robinplatform.dev/smgr/internal/manager"
)

var (
	port = os.Getenv("PORT")
	projectPath = os.Getenv("ROBIN_PROJECT_PATH")

	serverManager = manager.ServerManager{}
)

func writeRes(res http.ResponseWriter, statusCode int, body any) {
	res.WriteHeader(statusCode)
	if buf, err := json.Marshal(body); err != nil {
		res.Write([]byte(fmt.Sprintf(`{"type":"error","result":{"message":%q}}`, err.Error())))
	} else {
		res.Write(buf)
	}
}

func sendError(res http.ResponseWriter, statusCode int, err error) {
	writeRes(res, statusCode, map[string]string{
		"error": err.Error(),
	})
}

func sendJson(res http.ResponseWriter, body any) {
	writeRes(res, http.StatusOK, body)
}

func main() {
	if err := serverManager.DiscoverServers(projectPath); err != nil {
		panic(err)
	}

	lastRequestTime := new(int64)

	http.HandleFunc("/api/health", func(res http.ResponseWriter, req *http.Request) {
		sendJson(res, map[string]string{
			"status": "ok",
		})
	})
	server := http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			atomic.StoreInt64(lastRequestTime, time.Now().Unix())
			fmt.Printf("%s %s\n", req.Method, req.URL.Path)

			res.Header().Set("Content-Type", "application/json")
			http.DefaultServeMux.ServeHTTP(res, req)
		}),
	}

	go func() {
		for {
			<-time.After(5 * time.Second)

			if time.Now().Unix()-atomic.LoadInt64(lastRequestTime) > 60 {
				fmt.Printf("No requests for 1 min, shutting down\n")
				server.Shutdown(context.Background())
			}
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
	fmt.Printf("Server has stopped, shutting down\n")
}
