package configs

import (
	"api/pkg/file"
	"bytes"
	_ "embed"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var config = new(Config)

type Config struct {
	Redis struct {
		Addr         string `toml:"addr"`
		DB           int    `toml:"db"`
		Pass         string `toml:"pass"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"cache"`

	Server struct {
		Env                    string `toml:"env"`
		Addr                   string `toml:"addr"`
		Port                   int    `toml:"port"`
		GracefulShutdownPeriod int    `toml:"gracefulShutdownPeriod"`
	} `toml:"server"`

	Auth struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"auth"`
}

//go:embed configs.toml
var configs []byte

func init() {
	var r = bytes.NewReader(configs)

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	// 测试时，在当前目录写入./config/config.toml
	viper.SetConfigName("configs")
	viper.AddConfigPath("./configs")
	configFile := "./configs/configs.toml"
	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}
		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

// Get 获取配置信息
func Get() Config {
	return *config
}
