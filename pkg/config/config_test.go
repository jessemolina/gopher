package config

import (
	"fmt"
	"testing"
)

// Test function for makeInfo.
func TestMakeInfo(t *testing.T) {
	cfg := struct {
		APIPort   string `config:"default:3000"`
		DebugPort string `config:"default:4000"`
		Database struct {
			Host string `config:"default:localhost"`
			Port string `config:"default:5432"`
		}
	}{}

	fmt.Println(cfg)

	results := makeInfo(&cfg, "Gopher")
	for _, r := range results {
		fmt.Println(r.name, r.OSEnv())
	}

}
