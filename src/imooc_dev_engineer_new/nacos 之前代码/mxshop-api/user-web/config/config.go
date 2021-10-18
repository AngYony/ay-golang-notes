package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key"`
	ApiSecrect string `mapstructure:"secrect"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Expire int    `mapstructure:"expire"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	Port        int           `mapstructure:"port"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
	AliSmsInfo  AliSmsConfig  `mapstructure:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul"`
}
