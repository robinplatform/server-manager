package manager

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"robinplatform.dev/smgr/internal/manager/health"
)

type DevServerConfig struct {
	Name         string               `json:"-"`
	HealthChecks []health.HealthCheck `json:"healthChecks,omitempty"`
	Command      string               `json:"command"`
}

func (config *DevServerConfig) UnmarshalJSON(data []byte) error {
	type tmpConfigType DevServerConfig
	var tmpConfig tmpConfigType

	if err := json.Unmarshal(data, &tmpConfig); err == nil {
		config.HealthChecks = tmpConfig.HealthChecks
		config.Command = tmpConfig.Command

		return nil
	}

	// first try to unmarshal as a string
	if err := json.Unmarshal(data, &config.Command); err == nil {
		return nil
	}

	return fmt.Errorf("failed to unmarshal dev server config: %s", string(data))
}

type ServerManagerConfig struct {
	FilePath string `json:"-"`

	Name       string            `json:"name"`
	DevServers []DevServerConfig `json:"devServers"`
}

func (config *ServerManagerConfig) load() error {
	fd, err := os.Open(config.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open server config from %s: %w", config.FilePath, err)
	}

	buf, err := io.ReadAll(fd)
	if err != nil {
		return fmt.Errorf("failed to read server config from %s: %w", config.FilePath, err)
	}

	type ServerConfigJSON struct {
		HealthChecks []health.HealthCheck `json:"healthChecks,omitempty"`
		Command      string               `json:"command"`
	}

	type ConfigJSON struct {
		Name       string                      `json:"name"`
		DevServers map[string]ServerConfigJSON `json:"devServers"`
	}

	var configTmp ConfigJSON
	if err := json.Unmarshal(buf, &configTmp); err != nil {
		return fmt.Errorf("failed to unmarshal server config from %s: %w", config.FilePath, err)
	}

	config.Name = configTmp.Name
	config.DevServers = make([]DevServerConfig, 0, len(configTmp.DevServers))

	for name, value := range configTmp.DevServers {
		server := DevServerConfig{
			Name:         name,
			HealthChecks: value.HealthChecks,
			Command:      value.Command,
		}

		config.DevServers = append(config.DevServers, server)
	}

	if config.Name == "" {
		config.Name = filepath.Base(filepath.Dir(config.FilePath))
	}

	return nil
}

type ServerManager struct {
	mux *sync.Mutex

	Servers []ServerManagerConfig
}

func New() ServerManager {
	return ServerManager{
		mux: &sync.Mutex{},
	}
}

func (manager *ServerManager) DiscoverServers(projectPath string) error {
	defer manager.mux.Unlock()
	manager.mux.Lock()

	servers := make([]ServerManagerConfig, 0, len(manager.Servers))

	err := fs.WalkDir(os.DirFS(projectPath), ".", func(filename string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(filename) == "node_modules" {
			return filepath.SkipDir
		}
		if dirEntry.IsDir() {
			return nil
		}

		if filepath.Base(filename) == "robin.servers.json" {
			fmt.Printf("Discovered server config: %s\n", filename)

			serverManagerConfig := ServerManagerConfig{
				FilePath: filepath.Join(projectPath, filename),
			}
			if err := serverManagerConfig.load(); err != nil {
				return err
			}
			servers = append(servers, serverManagerConfig)
		}

		return nil
	})
	if err != nil {
		return err
	}

	manager.Servers = servers
	return nil
}
