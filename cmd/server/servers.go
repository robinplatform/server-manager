package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/GetServers", GetServers)
}

func GetServers(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("servers: %#v\n", serverManager.Servers)
	sendJson(res, serverManager.Servers)
}
