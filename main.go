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

	respRef := &example.Response{Message: "Successfully api'd"}
	errResp := &example.ErrorResponse{Message: "The big failure 😡", Code: 400}
	var mapish = make(map[string]interface{})

	mapish["200"] = respRef
	mapish["400"] = errResp
	libs.AddToSwaggerAndRegister(
		libs.PathItem{
			Summary: "I am the home endpoint",
			Get: &libs.Operation{
				Summary:     "The Home Endpoint",
				Description: "I AM DESCRIPTION",
				Responses:   mapish,
			},
		}, config, "/", homeHandler)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	libs.GenerateDocs(config)

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
