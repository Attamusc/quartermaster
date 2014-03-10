package config

import (
    "encoding/json"
    "errors"
    . "github.com/attamusc/quartermaster/package"
    "io/ioutil"
    "os"
    "path/filepath"
)

const configName string = "quartermaster.json"

type Config struct {
    Name         string    `json:"name"`
    Location     string    `json:"location"`
    Dependencies []Package `json:"dependencies"`
}

func findConfigFile() (string, error) {
    workingDirectory, err := os.Getwd()

    if err != nil {
        return "", errors.New("[Error] Can't find current directory")
    }

    configLocation := filepath.Join(workingDirectory, configName)
    return configLocation, nil
}

// read the config file and unmarshal it into a Config object
func readConfig(configLocation string) (*Config, error) {
    config := &Config{}
    configData, err := ioutil.ReadFile(configLocation)

    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(configData, &config); err != nil {
        return nil, err
    }

    return config, nil
}

func Read() (*Config, error) {
    configLocation, err := findConfigFile()

    if err != nil {
        return nil, err
    }

    return readConfig(configLocation)
}

func (c *Config) InstallDependencies() {
    // Install each of the dependencies mentioned in the config file
    for _, p := range c.Dependencies {
        p.Install()
    }
}
