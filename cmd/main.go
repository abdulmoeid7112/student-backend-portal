package main

import (
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	"github.com/abdulmoeid7112/student-backend-portal/config"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi"
	"github.com/abdulmoeid7112/student-backend-portal/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}

