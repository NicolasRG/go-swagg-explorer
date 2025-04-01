package libs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//funciton to go through the registered http endpoints and generate swagger docs

type SwaggerServerConfig struct {
	Server    *http.ServeMux
	Info      OpenAPIInfo
	routeDocs map[string]PathItem
}

//shelve this idea for now
// type RouteConfig struct {
// 	Summary     string
// 	Description string
// 	Servers     []OpenAPIServer
// 	Parameters  []Parameter
// 	Methods     string
// }

func AddToSwaggerAndRegister(pathItem PathItem, config *SwaggerServerConfig, path string, passthrough http.HandlerFunc) {
	if config.routeDocs == nil {
		config.routeDocs = map[string]PathItem{}
	}
	//convert this route config to path item
	//add
	config.routeDocs[path] = pathItem

	//jury is still out on this
	config.Server.Handle(path, passthrough)
}

func GenerateDocs(config *SwaggerServerConfig) {
	fmt.Println(config.routeDocs)
	// attach over all swagger details

	openApiConfig := OpenAPIConfig{}

	//hard coded for now
	openApiConfig.OpenAPI = "3.0.0"
	openApiConfig.Info = config.Info
	openApiConfig.Paths = config.routeDocs

	jsonData, _ := json.MarshalIndent(openApiConfig, "", "")

	err := os.WriteFile("docs.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//feed in data to swagger endpoint
}

// TODO:  use this to parse path for related params
// func parsePath(path string, config *SwaggerServerConfig) {
// 	config.routeDocs
// }
