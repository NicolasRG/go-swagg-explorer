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

	schemas := generateMarshalRefs(config.refs)
	openApiConfig.Components.Schemas = schemas

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
		//FIXME: remove the * of pointers as this is out of spec
		actualType := reflect.TypeOf(responseBody).String()
		//check that the type key is not in the map
		_, ok := config.refs[actualType]

		if ok {
			continue
		} else {
			if config.refs == nil {
				config.refs = make(map[string]interface{})
			}
			fmt.Printf("adding type : %s to refs map \n", actualType)
			config.refs[actualType] = responseBody
		}
	}
}

// TODO: return a object that can be attached to the config, so does the it need this arguement????
// returns a map of parameters based on types given
func generateMarshalRefs(refs map[string]interface{}) map[string]*Schema {
	schemaMap := make(map[string]*Schema)
	//loop through all the refs and create schemas out of all them
	for typeName, typeRef := range refs {
		fmt.Printf("generating type docs: %s to refs map \n", typeName)
		//tarverse through each field and generate the docs for each type
		if isBaseType(reflect.TypeOf(typeRef)) { //FIXME: this will not work for ptr types FYI
			schema := &Schema{}
			// deserialize as normal

			//determine the type
			schema.Type = mapTypeToJsonType(reflect.TypeOf(typeRef))

			//TODO: Idea use tags to add descriptions or things of nature

			//attacht to refs at the end
			schemaMap[typeName] = schema

		} else {
			schema := &Schema{}
			schema.Type = "object"

			interfaceMap := structToMap(typeRef)

			// this definitly breaks
			schema.Properties = generateMarshalRefs(interfaceMap)

			schemaMap[typeName] = schema
		}
	}
	return schemaMap
}

func structToMap(obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not a struct")
	}

	result := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Handle nested structs recursively
		if value.Kind() == reflect.Struct && field.Anonymous == false {
			result[field.Name] = structToMap(value.Interface())
		} else {
			result[field.Name] = value.Interface()
		}
	}

	return result
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
