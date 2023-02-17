package config

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"os"
)

var Path = "config.toml"

type Config struct {
	Ip       string `toml:"ip"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func (c *Config) Init() {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		c.Ip = "192.168.1.1"
		c.Username = "admin"
		c.Password = ""

		c.WriteConfigFile(Path)
	} else {
		config, err := GetConfig()
		if err != nil {
			return
		}

		c.Ip = config.Ip
		c.Username = config.Username
		c.Password = config.Password
	}
}

func (c *Config) WriteConfigFile(s string) error {
	f, err := os.Create(s)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := toml.NewEncoder(f).Encode(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) ToJson() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GetConfig() (Config, error) {
	var c Config
	if _, err := toml.DecodeFile(Path, &c); err != nil {
		return c, err
	}
	return c, nil
}

func (c *Config) UpdateConfigFile(path string, ip string, username string, password string) error {
	if ip != "" && ip != c.Ip {
		c.Ip = ip
	}
	if username != "" {
		c.Username = username
	}
	if password != "" {
		c.Password = password
	}

	return c.WriteConfigFile(path)
}
