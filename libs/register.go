package libs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

//funciton to go through the registered http endpoints and generate swagger docs

type SwaggerServerConfig struct {
	Server    *http.ServeMux
	Info      OpenAPIInfo
	routeDocs map[string]PathItem
	refs      map[string]interface{}
}

func AddToSwaggerAndRegister(pathItem PathItem, config *SwaggerServerConfig, path string, passthrough http.HandlerFunc) {
	if config.routeDocs == nil {
		config.routeDocs = map[string]PathItem{}
	}

	//TODO: traverse parameters and desearilze? or make them do it proberlyy????
	if pathItem.Get != nil {
		parseResponse("GET", pathItem.Get, config)
	}
	//check for all request types in responses and requestBodies and regs

	//convert this route config to path item
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

	generateMarshalRefs(config)

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

func parseResponse(method string, op *Operation, config *SwaggerServerConfig) {
	//get the responsebodies out of map
	res := op.Responses

	for statusCode, responseBody := range res {
		fmt.Print(statusCode, responseBody)
		//get the type of the responseBody
		typeOf := reflect.TypeOf(responseBody).Name()
		//check that the type key is not in the map
		_, ok := config.refs[typeOf]
		if ok {
			continue
		} else {
			fmt.Printf("adding type : %s to refs map \n", typeOf)
			config.refs[typeOf] = responseBody
		}
	}
}

// todo return a object that can be attached to the config, so does the it need this arguement????
func generateMarshalRefs(config *SwaggerServerConfig) {
	//loop through all the refs and create schemas out of all them
	for typeName, typeRef := range config.refs {
		fmt.Printf("generating type docs: %s to refs map \n", typeName)
		//tarverse through each field and generate the docs for each type
		if isBaseType(reflect.TypeOf(typeRef)) { //FIXME: this will not work for ptr types FYI
			schema := Schema{}
			// deserialize as normal

			//determine the type
			schema.Type = mapTypeToJsonType(reflect.TypeOf(typeRef))

			//TODO: Idea use tags to add descriptions or things of nature

			//attacht to refs at the end

		} else {
			// go through each child elem and parse out
			keys := reflect.VisibleFields(reflect.TypeOf(typeRef))

			for field := range keys {
				//fill out the normal fields or let the down stream call handle this?? i think let the down stream call handle this

				//todo: split this out into a recursive process

			}
		}
	}
}

func isBaseType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		//reflect.Ptr:
		return true
	default:
		return false
	}
}

func mapTypeToJsonType(dataType reflect.Type) SchemaType {
	switch dataType.Kind() {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Struct:
		//FIXME: Handle time correctly
		// Handle time.Time as a special case (OpenAPI format: date-time)
		if dataType == reflect.TypeOf(time.Time{}) {
			return "string"
		}
		return "object"
	default:
		return "string" // Default fallback
	}
}

func getIntFormat() {

}
