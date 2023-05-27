package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Database struct {
		StorageType      string `mapstructure:"storageType"`
		ConnectionString string `mapstructure:"connectionString"`
	} `mapstructure:"database"`
	Authorization struct {
		JwtKey string `mapstructure:"jwtKey"`
	} `mapstructure:"authorization"`
}

var vp *viper.Viper

func InitConfig() (Config, error) {
	vp = viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")

	err := vp.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetConfig(key string) string {
	return vp.GetString(key)
}
