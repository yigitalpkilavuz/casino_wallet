package conf

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var vp *viper.Viper

func GetConfig() (Config, error) {
	vp = viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./conf")

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
