package manager

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func StartServer(config DevServerConfig) error {
	jsonValue := []byte(`{
		"appId":"server-manager",
		"processKey":"blarg",
		"command":"/bin/ls"
	}`)

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/StartProcess", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("resp body: %s\n", string(body))

	return nil
}

func StopServer(config DevServerConfig) error {
	jsonValue := []byte(`{
		"appId":"server-manager",
		"processKey":"blarg"
	}`)

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/StopProcess", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("resp body: %s\n", string(body))

	return nil
}
