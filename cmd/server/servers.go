package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/GetServers", GetServers)
}

func GetServers(res http.ResponseWriter, req *http.Request) {
	if err := serverManager.DiscoverServers(projectPath); err != nil {
		sendError(res, 500, fmt.Errorf("failed to discover servers: %w", err))
	} else {
		sendJson(res, serverManager.Servers)
	}
}
