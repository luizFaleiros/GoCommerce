package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbName        string `mapstructure:"DB_NAME"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbPAssword    string `mapstructure:"DB_PASSWORD"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExperiTime int64  `mapstructure:"EXPIRE_TIME"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&cfg)
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
