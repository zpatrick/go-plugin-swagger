package swagger

type Tag struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
}
