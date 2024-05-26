package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
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

func init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New(jwa.HS256.String(), []byte(cfg.JWTSecret), nil)

}

func LoadConfig() *conf {
	return cfg
}
