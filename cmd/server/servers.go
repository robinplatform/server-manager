package main

import (
	"fmt"
	"net/http"

	"robinplatform.dev/smgr/internal/manager"
)

func init() {
	http.HandleFunc("/api/GetServers", GetServers)
	http.HandleFunc("/api/StartServer", StartServer)
	http.HandleFunc("/api/CheckServerHealth", CheckServerHealth)
}

func GetServers(res http.ResponseWriter, req *http.Request) {
	if err := serverManager.DiscoverServers(projectPath); err != nil {
		sendError(res, 500, fmt.Errorf("failed to discover servers: %w", err))
	} else {
		sendJson(res, serverManager.Servers)
	}
}

func StartServer(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("StartServer Running\n")
	config := manager.DevServerConfig{
		Name:    "hello",
		Command: "ls",
	}

	if err := manager.StartServer(config); err != nil {
		sendError(res, 500, fmt.Errorf("failed to discover servers: %w", err))
	} else {
		sendJson(res, map[string]any{
			"success": true,
		})
	}
}

func CheckServerHealth(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("CheckHealth Running\n")
	config := manager.DevServerConfig{
		Name:    "hello",
		Command: "/bin/ls",
	}

	if err := manager.CheckServerHealth(config); err != nil {
		sendError(res, 500, fmt.Errorf("failed to discover servers: %w", err))
	} else {
		sendJson(res, map[string]any{
			"success": true,
		})
	}
}
