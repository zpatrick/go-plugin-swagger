package swagger

type Schema struct {
	Type  string            `json:"type"`
	Items map[string]string `json:"items"`
}
