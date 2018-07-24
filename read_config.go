package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JuggleConfig is the struct that will hold juggle configuration data.
type JuggleConfig struct {
	Protocol  string   `json:"protocol"`
	Hosts     []string `json:"hosts"`
	Algorithm string   `json:"algorithm"`
	Port      string   `json:"port"`
}

// ReadJuggleConfig will read in the juggle config from disk and return
// a new config object with all host information that will be used by the load-balancer.
func ReadJuggleConfig(path string) (*JuggleConfig, error) {
	// Open the juggle configuration file for reading.
	jsonConfig, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Defer the closing of the juggle configuration file.
	defer jsonConfig.Close()
	// Create a new Config object to be used when unmarshalling the json config data.
	config := &JuggleConfig{}
	// Read all of the JSON data from the config file into memory.
	jsonData, err := ioutil.ReadAll(jsonConfig)
	fmt.Println(string(jsonData))
	if err != nil {
		return nil, err
	}
	// Unmarshall the JSON configuration data and insert into the config object.
	json.Unmarshal(jsonData, config)
	// Return the configuration object to be used by juggle.
	return config, nil
}
