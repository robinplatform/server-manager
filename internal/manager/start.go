package manager

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartServer(config DevServerConfig) error {
	jsonValue := []byte(`{
		"appId":"server-manager",
		"processKey":"blarg",
		"command":"/bin/ls"
	}`)

	resp, err := http.Post("http://localhost:9010/api/apps/rpc/StartProcessForApp", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("resp body: %s\n", string(body))

	return nil
}
