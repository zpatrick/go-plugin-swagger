package swagger

type Spec struct {
	SwaggerVersion string          `json:"swagger"`
	Host           string          `json:"host,omitempty"`
	BasePath       string          `json:"basePath,omitempty"`
	Info           Info            `json:"info"`
	Schemes        []string        `json:"schemes"`
	ExternalDocs   ExternalDocs    `json:"externalDocs"`
	Tags           []Tag           `json:"tags"`
	Paths          map[string]Path `json:"paths"`
}
