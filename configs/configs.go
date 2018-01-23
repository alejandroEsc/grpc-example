package configs

import(
    "github.com/spf13/viper"
)


var (
    envPrefix                     = "grpc_hello"
    envServiceAddress             = "service_address"
    envServicePort                = "service_port"
    envServiceKnockFailure        = "service_knock_failure"
    envGateWayServiceAddress      = "gw_service_address"
    envGateWayPort                = "gw_port"
    envGateWaySwaggerDir          = "gw_swagger_dir"
    addressDefault                = "localhost"
    portDefault                   = 8501
    gwAddressDefault              = "localhost"
    gwPortDefault                 = 8502
    knockFailureDefault           = "You should try and knock"
    gwSwaggerDirDefault           = "swagger"
)

func InitEnvVars() {
    viper.SetEnvPrefix(envPrefix)
    viper.BindEnv(envServicePort)
    viper.BindEnv(envServiceKnockFailure)
    viper.BindEnv(envServiceAddress)
    viper.BindEnv(envGateWayServiceAddress)
    viper.BindEnv(envGateWayPort)
    viper.BindEnv(envGateWaySwaggerDir)
}

func ParseEnvVars() (int, string, string) {
    port := viper.GetInt(envServicePort);
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

func ParseGWSwaggerEnvVars() string {
    gwSwaggerDir := viper.GetString(envGateWaySwaggerDir)
    if gwSwaggerDirDefault == "" {
        gwSwaggerDirDefault = gwSwaggerDirDefault
    }
    return gwSwaggerDir
}

func ParseGateWayEnvVars() (int, int, string, string) {
    gwPort := viper.GetInt(envGateWayPort);
    if gwPort == 0 {
        gwPort = gwPortDefault
    }

    port := viper.GetInt(envServicePort);
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
