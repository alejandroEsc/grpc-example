package configs

import (
	"github.com/spf13/viper"
)

var (
	envPrefix                = "grpc_hello"
	envServiceAddress        = "service_address"
	envServicePort           = "service_port"
	envServiceKnockFailure   = "service_knock_failure"
	envGateWayServiceAddress = "gw_service_address"
	envGateWayPort           = "gw_port"
	envGateWaySwaggerDir     = "gw_swagger_dir"
	addressDefault           = "localhost"
	portDefault              = 8501
	gwAddressDefault         = "localhost"
	gwPortDefault            = 8502
	knockFailureDefault      = "You should try and knock"
	gwSwaggerDirDefault      = "swagger"
)

// InitEnvVars allows you to initiate gathering environment variables
func InitEnvVars() error {
	var err error
	viper.SetEnvPrefix(envPrefix)

	err = viper.BindEnv(envServicePort)
	if err != nil {
		return err
	}

	err = viper.BindEnv(envServiceKnockFailure)
	if err != nil {
		return err
	}

	err = viper.BindEnv(envServiceAddress)
	if err != nil {
		return err
	}

	err = viper.BindEnv(envGateWayServiceAddress)
	if err != nil {
		return err
	}

	err = viper.BindEnv(envGateWayPort)
	if err != nil {
		return err
	}

	err = viper.BindEnv(envGateWaySwaggerDir)
	if err != nil {
		return err
	}

	return err
}

// ParseEnvVars allows you to parse variables consumed by the grpc service
func ParseEnvVars() (int, string, string) {
	port := viper.GetInt(envServicePort)
	if port == 0 {
		port = portDefault
	}

	knockFailure := viper.GetString(envServiceKnockFailure)
	if knockFailure == "" {
		knockFailure = knockFailureDefault
	}

	serviceAddress := viper.GetString(envServiceAddress)
	if serviceAddress == "" {
		serviceAddress = addressDefault
	}

	return port, knockFailure, serviceAddress
}

// ParseGWSwaggerEnvVars parses environment variables consumed by swagger server
func ParseGWSwaggerEnvVars() string {
	gwSwaggerDir := viper.GetString(envGateWaySwaggerDir)
	if gwSwaggerDir == "" {
		gwSwaggerDir = gwSwaggerDirDefault
	}
	return gwSwaggerDir
}

// ParseGateWayEnvVars parses environment variables consumed by the gateway service
func ParseGateWayEnvVars() (int, int, string, string) {
	gwPort := viper.GetInt(envGateWayPort)
	if gwPort == 0 {
		gwPort = gwPortDefault
	}

	port := viper.GetInt(envServicePort)
	if port == 0 {
		port = portDefault
	}

	gwServiceAddress := viper.GetString(envGateWayServiceAddress)
	if gwServiceAddress == "" {
		gwServiceAddress = gwAddressDefault
	}

	serviceAddress := viper.GetString(envServiceAddress)
	if serviceAddress == "" {
		serviceAddress = addressDefault
	}

	return gwPort, port, gwServiceAddress, serviceAddress
}
