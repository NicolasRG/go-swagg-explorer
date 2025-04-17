package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nicolasrg/go-swagg-explorer/example"
	"github.com/nicolasrg/go-swagg-explorer/libs"
)

func main() {
	// Define routes
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/health", healthHandler)

	config := &libs.SwaggerServerConfig{
		Server: http.DefaultServeMux,
		Info: libs.OpenAPIInfo{
			Title:          "Demo App",
			Description:    "This a demo app",
			Version:        "-0.0.1",
			TermsOfService: "https://swagger.io/terms/",
			Contact: &libs.OpenAPIContact{
				Name:  "John S Tea",
				URL:   "TEST.com",
				Email: "john@TEST.com",
			},
		},
	}

	// GET ENDPOINT EXAMPLE
	respRef := &example.Response{Message: "Successfully api'd"}
	errResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var getResponses = make(map[string]interface{})

	getResponses["200"] = respRef
	getResponses["400"] = errResp
	// POST ENDPOINT EXAMPLE
	postRespRef := &example.Response{Message: "Successfully api'd"}
	postErrResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var postResponses = make(map[string]interface{})

	postResponses["200"] = postRespRef
	postResponses["400"] = postErrResp

	// PUT ENDPOINT EXAMPLE
	putRespRef := &example.Response{Message: "Successfully api'd"}
	putErrResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var putResponses = make(map[string]interface{})

	putResponses["200"] = putRespRef
	putResponses["400"] = putErrResp

	// PATCH ENDPOINT EXAMPLE
	patchRespRef := &example.Response{Message: "Successfully api'd"}
	patchErrResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var patchResponses = make(map[string]interface{})

	patchResponses["200"] = patchRespRef
	patchResponses["400"] = patchErrResp

	// DELETE ENDPOINT EXAMPLE
	deleteRespRef := &example.Response{Message: "Successfully api'd"}
	deleteErrResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var deleteResponses = make(map[string]interface{})

	deleteResponses["200"] = deleteRespRef
	deleteResponses["400"] = deleteErrResp

	libs.AddToSwaggerAndRegister(
		libs.PathItem{
			Summary: "I am the home endpoint",
			Get: &libs.Operation{
				Summary:     "The Home Endpoint",
				Description: "I AM DESCRIPTION",
				Responses:   getResponses,
			},
			Post: &libs.Operation{
				Summary:     "The POST Endpoint",
				Description: "POST DESCRIPTION",
				Responses:   postResponses,
			},
			Put: &libs.Operation{
				Summary:     "The PUT Endpoint",
				Description: "PUT DESCRIPTION",
				Responses:   putResponses,
			},
			Patch: &libs.Operation{
				Summary:     "The PATCH Endpoint",
				Description: "PATCH DESCRIPTION",
				Responses:   patchResponses,
			},
			Delete: &libs.Operation{
				Summary:     "The DELETE Endpoint",
				Description: "DELETE DESCRIPTION",
				Responses:   deleteResponses,
			},
		}, config, "/", homeHandler)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	libs.GenerateDocs(config)

	//TODO: add swagger endpoint here

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler for the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the simple Go HTTP service!")
}

// Handler for /hello path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// Handler for /health path
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy"}`)
}
