package manager

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type ServerManagerConfig struct {
	FilePath string `json:"-"`

	Name string `json:"name"`
	DevServers map[string]string `json:"devServers"`
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

	if err := json.Unmarshal(buf, config); err != nil {
		return fmt.Errorf("failed to unmarshal server config from %s: %w", config.FilePath, err)
	}

	if config.Name == "" {
		config.Name = filepath.Base(filepath.Dir(config.FilePath))
	}

	return nil
}
 
type ServerManager struct {
	Servers []ServerManagerConfig
}

func (manager *ServerManager) DiscoverServers(projectPath string) error {
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
