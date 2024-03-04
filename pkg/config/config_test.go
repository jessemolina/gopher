package config

import (
	"testing"
)

// Test function for makeInfo.
func TestMakeInfo(t *testing.T) {
	cfg := struct {
		APIPort   string `config:"default:3000"`
		DebugPort string `config:"default:4000"`
		Database  struct {
			Host string `config:"default:localhost"`
			Port string `config:"default:5432"`
		}
	}{}

	results := makeInfo(&cfg, "Gopher")
	for _, r := range results {
		if r.pointer == nil {
			t.Fatalf("\nfieldInfo pointer value is nil")
		}
	}

}
