package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Instance represents a single RDS information from configuration file.
type Instance struct {
	Region       string `yaml:"region"`
	Instance     string `yaml:"instance"`
	AWSAccessKey string `yaml:"aws_access_key"` // may be empty
	AWSSecretKey string `yaml:"aws_secret_key"` // may be empty
}

// Config contains configuration file information.
type Config struct {
	Instances          []Instance `yaml:"instances"`
	CredentialsProcess string     `yaml:"credentials-process"`
}

// Load loads configuration from file.
func Load(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
