package config

import "github.com/spf13/viper"

const (
	Env			= "ENV"
	Port		= "PORT"
	GRPCPort	= "GRPC_PORT"
)

func init()  {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetDefault(Env, "dev")
	viper.SetDefault(Port, 8080)
	viper.SetDefault(GRPCPort, 8090)
}

func ReadConfig(env string) error  {
	viper.SetConfigFile("app-" + env + ".yaml")
	return viper.ReadInConfig();
}