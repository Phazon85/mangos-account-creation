package pg

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	psqlConnectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

//SQLConnectionInfo holds connection info for SQL implementations
type SQLConnectionInfo struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	DBName   string `yaml:"dbname"`
}

func newConfig(file string) string {
	cfg := &SQLConnectionInfo{}
	if err := load(cfg, file); err != nil {
		return ""
	}

	psql := fmt.Sprintf(psqlConnectionString, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	return psql
}

func load(config interface{}, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
