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
	port        = os.Getenv("PORT")
	projectPath = os.Getenv("ROBIN_PROJECT_PATH")

	serverManager = manager.New()
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

			// TODO: This should eventually always print, but everything is a little
			// borked right now, and since this gets hit every quarter-second right now,
			// this would be a pain in the ass scroll through in the logs.
			if req.URL.Path != "/api/health" {
				fmt.Printf("%s %s\n", req.Method, req.URL.Path)
			}

			res.Header().Set("Content-Type", "application/json")
			http.DefaultServeMux.ServeHTTP(res, req)
		}),
	}

	go func() {
		// TODO: This shouldn't be necessary, but we don't have auto-kill written yet.
		for {
			<-time.After(1 * time.Second)

			if time.Now().Unix()-atomic.LoadInt64(lastRequestTime) > 1 {
				fmt.Printf("No requests for a second, shutting down\n")
				server.Shutdown(context.Background())
			}
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
	fmt.Printf("Server has stopped, shutting down\n")
}
