package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	EmailSuperUsuario string `mapstructure:"EMAIL_SUPER_USUARIO"`
	SenhaSuperUsuario string `mapstructure:"SENHA_SUPER_USUARIO"`
	JWTSecret         string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn      int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth         *jwtauth.JWTAuth
}

func LoadConfigs(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
