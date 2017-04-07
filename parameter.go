package swagger

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	In          string `json:"in"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Format      string `json:"format,omitempty"`
}

func NewIntPathParam(name, description string, required bool) Parameter {
	return Parameter{
		Name:        name,
		Description: description,
		Required:    required,
		In:          "path",
		Type:        "integer",
		Format:      "int64",
	}
}

func NewStringPathParam(name, description string, required bool) Parameter {
	return Parameter{
		Name:        name,
		Description: description,
		Required:    required,
		In:          "path",
		Type:        "string",
	}
}
