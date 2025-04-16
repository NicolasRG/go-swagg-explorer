package libs

type SchemaType string

const (
	STRING_TYPE  SchemaType = "string"
	NUMBER_TYPE  SchemaType = "number"
	INTEGER_TYPE SchemaType = "integer"
	BOOLEAN_TYPE SchemaType = "boolean"
	ARRAY        SchemaType = "array"
	OBJECT       SchemaType = "object"
)

// Schema represents the OpenAPI 3.0 Schema Object.
type Schema struct {
	Type                 SchemaType         `json:"type,omitempty"`                 // "string", "number", "integer", "boolean", "array", "object"
	Format               string             `json:"format,omitempty"`               // e.g., "date-time", "email", "int32" TODO: should be an enum
	Description          string             `json:"description,omitempty"`          // Description of the schema
	Default              interface{}        `json:"default,omitempty"`              // Default value
	Nullable             bool               `json:"nullable,omitempty"`             // If true, allows null values
	Deprecated           bool               `json:"deprecated,omitempty"`           // If true, marks as deprecated
	Required             []string           `json:"required,omitempty"`             // Required properties (for objects)
	Properties           map[string]*Schema `json:"properties,omitempty"`           // Nested properties (for objects)
	AdditionalProperties *Schema            `json:"additionalProperties,omitempty"` // Schema for additional properties (if object allows extra fields)
	Items                *Schema            `json:"items,omitempty"`                // Schema for array items
	Enum                 []interface{}      `json:"enum,omitempty"`                 // Allowed values
	Minimum              float64            `json:"minimum,omitempty"`              // Min value (for numbers)
	Maximum              float64            `json:"maximum,omitempty"`              // Max value (for numbers)
	ExclusiveMinimum     bool               `json:"exclusiveMinimum,omitempty"`     // If true, value > Minimum
	ExclusiveMaximum     bool               `json:"exclusiveMaximum,omitempty"`     // If true, value < Maximum
	MultipleOf           float64            `json:"multipleOf,omitempty"`           // Value must be a multiple of this
	MinLength            int                `json:"minLength,omitempty"`            // Min string length
	MaxLength            int                `json:"maxLength,omitempty"`            // Max string length
	Pattern              string             `json:"pattern,omitempty"`              // Regex pattern for strings
	MinItems             int                `json:"minItems,omitempty"`             // Min array length
	MaxItems             int                `json:"maxItems,omitempty"`             // Max array length
	UniqueItems          bool               `json:"uniqueItems,omitempty"`          // If true, array must have unique items
	AllOf                []*Schema          `json:"allOf,omitempty"`                // Combines multiple schemas (logical AND)
	AnyOf                []*Schema          `json:"anyOf,omitempty"`                // Combines multiple schemas (logical OR)
	OneOf                []*Schema          `json:"oneOf,omitempty"`                // Must match exactly one schema
	Not                  *Schema            `json:"not,omitempty"`                  // Must NOT match this schema
	Discriminator        *Discriminator     `json:"discriminator,omitempty"`        // Used for polymorphism
	ReadOnly             bool               `json:"readOnly,omitempty"`             // If true, property is read-only
	WriteOnly            bool               `json:"writeOnly,omitempty"`            // If true, property is write-only
	Example              interface{}        `json:"example,omitempty"`              // Example value (deprecated in favor of `Examples`)
	Examples             []interface{}      `json:"examples,omitempty"`             // Multiple examples
	ExternalDocs         *ExternalDocs      `json:"externalDocs,omitempty"`         // Link to external docs
	XML                  *XML               `json:"xml,omitempty"`                  // XML-specific metadata
	Ref                  string             `json:"$ref,omitempty"`
}

// Discriminator is used for polymorphism in OpenAPI 3.0.
type Discriminator struct {
	PropertyName string            `json:"propertyName"`      // Name of the discriminator property
	Mapping      map[string]string `json:"mapping,omitempty"` // Mapping of values to schemas
}

// ExternalDocs represents external documentation.
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// XML defines XML-specific metadata.
type XML struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}
