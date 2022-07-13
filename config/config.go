package config

import (
	"context"
	"log"

	"github.com/BurntSushi/toml"
)

// Config 全局配置类
type Config struct {
	GlobalCof GLobalConfig `toml:"global"`
	RedisConf RedisConfig  `toml:"redis"`
	MySQLConf MySQLConfig  `toml:"mysql"`
}

// GLobalConfig 全局配置类
type GLobalConfig struct {
	Port int `toml:"Port"`
}

// MySQLConfig mysql配置类
type MySQLConfig struct {
	URL string `toml:"URL"`
}

// RedisConfig Redis配置类
type RedisConfig struct {
	Addr   string `toml:"Addr"`
	PassWd string `toml:"Passwd"`
	DB     int    `toml:"DB"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	config := &Config{}
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return config, nil
}
