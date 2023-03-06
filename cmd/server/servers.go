package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/api/GetServers", GetServers)
}

func GetServers(res http.ResponseWriter, req *http.Request) {
	sendJson(res, serverManager.Servers)
}
