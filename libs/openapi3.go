package libs

// SwaggerParameter defines a parameter for the endpoint
// OpenAPIConfig represents the root-level OpenAPI/Swagger documentation
type OpenAPIConfig struct {
	// Required fields
	OpenAPI string              `json:"openapi" yaml:"openapi"` // e.g., "3.0.0"
	Info    OpenAPIInfo         `json:"info" yaml:"info"`
	Paths   map[string]PathItem `json:"paths" yaml:"paths"`

	// Optional fields
	Servers      []OpenAPIServer       `json:"servers,omitempty" yaml:"servers,omitempty"`
	Components   OpenAPIComponents     `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []OpenAPITag          `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *OpenAPIExternalDocs  `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// OpenAPIInfo contains metadata about the API
type OpenAPIInfo struct {
	Title          string          `json:"title" yaml:"title"`
	Description    string          `json:"description,omitempty" yaml:"description,omitempty"`
	Version        string          `json:"version" yaml:"version"`
	TermsOfService string          `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *OpenAPIContact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *OpenAPILicense `json:"license,omitempty" yaml:"license,omitempty"`
}

// PathItem represents operations available on a single path
type PathItem struct {
	Ref         string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string          `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation      `json:"get,omitempty" yaml:"get,omitempty"`
	Post        *Operation      `json:"post,omitempty" yaml:"post,omitempty"`
	Put         *Operation      `json:"put,omitempty" yaml:"put,omitempty"`
	Delete      *Operation      `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch       *Operation      `json:"patch,omitempty" yaml:"patch,omitempty"`
	Head        *Operation      `json:"head,omitempty" yaml:"head,omitempty"`
	Options     *Operation      `json:"options,omitempty" yaml:"options,omitempty"`
	Trace       *Operation      `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []OpenAPIServer `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []Parameter     `json:"parameters,omitempty" yaml:"parameters,omitempty"` //
}

// Operation represents a single API operation
type Operation struct {
	Tags        []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary     string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters  []Parameter            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"` //TODO: Figure out how to handle refs/schemas
	Responses   map[string]interface{} `json:"responses" yaml:"responses"`                         //TODO: Figoure out how to handle ref/schemas
	Deprecated  bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security    []map[string][]string  `json:"security,omitempty" yaml:"security,omitempty"`
	Servers     []OpenAPIServer        `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// Supporting types
type OpenAPIServer struct {
	URL         string `json:"url" yaml:"url"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// TODO: define how schemas should work
type OpenAPIComponents struct {
	Schemas         map[string]interface{} `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]interface{} `json:"responses,omitempty" yaml:"responses,omitempty"` //TOD
	Parameters      map[string]Parameter   `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBodies   map[string]interface{} `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	SecuritySchemes map[string]interface{} `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"` //TODO: secutiry schmea should have a def
}

type OpenAPITag struct {
	Name         string               `json:"name" yaml:"name"`
	Description  string               `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *OpenAPIExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type OpenAPIExternalDocs struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}

type OpenAPIContact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type OpenAPILicense struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Parameter struct {
	In          string      `json:"in" yaml:"in"`
	Name        string      `json:"name" yaml:"name"`
	Schema      interface{} `json:"schema" yaml:"schema"` //TODO: still need to decide schema
	Type        string      `json:"type" yaml:"type"`
	Required    bool        `json:"required" yaml:"required"`
	Description string      `json:"description" yaml:"description"`
}
