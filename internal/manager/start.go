package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func StartServer(config DevServerConfig) error {
	type Input struct {
		AppId      string   `json:"appId"`
		ProcessKey string   `json:"processKey"`
		Command    string   `json:"command"`
		Args       []string `json:"args"`
	}

	input := Input{
		AppId:      "server-manager",
		ProcessKey: config.Name,
		Command:    "/bin/bash",
		Args:       []string{"-c", config.Command},
	}

	jsonValue, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/StartProcess", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("StartServer resp body: %s\n", string(body))

	return nil
}

func StopServer(config DevServerConfig) error {
	type Input struct {
		AppId      string `json:"appId"`
		ProcessKey string `json:"processKey"`
	}

	input := Input{
		AppId:      "server-manager",
		ProcessKey: config.Name,
	}

	jsonValue, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/StopProcess", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("resp body: %s\n", string(body))

	return nil
}

func CheckServerHealth(config DevServerConfig) error {
	type Input struct {
		AppId      string `json:"appId"`
		ProcessKey string `json:"processKey"`
	}

	input := Input{
		AppId:      "server-manager",
		ProcessKey: config.Name,
	}

	jsonValue, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/CheckProcessHealth", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("CheckServer resp body: %s\n", string(body))

	return nil
}
