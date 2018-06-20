// Package storage loads the configuration file with only storage information.
package storage

import (
	"encoding/json"
	"os"
	"regexp"

	"github.com/blue-jay-fork/core/jsonconfig"
	"github.com/blue-jay-fork/core/storage/driver/mysql"
	"github.com/blue-jay-fork/core/storage/driver/postgresql"
)

// Info contains the database connection information for the different storage.
type Info struct {
	MySQL      mysql.Info      `json:"MySQL"`
	PostgreSQL postgresql.Info `json:"PostgreSQL"`
}

// ParseJSON unmarshals bytes to structs.
func (c *Info) ParseJSON(b []byte) error {
	// w is always "env:*" here, hence [4:]
	Getenv := func(w string) string {
		return os.Getenv(w[4:])
	}

	// Looking for all the "env:*" strings to replace with ENV vars
	r := regexp.MustCompile(`env:[A-Z_]+`)
	newB := r.ReplaceAllStringFunc(string(b), Getenv)

	return json.Unmarshal([]byte(newB), &c)
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Configuration
	config := &Info{}

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}
