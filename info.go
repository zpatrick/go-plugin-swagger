package swagger

type Info struct {
	Title          string  `json:"title"`
	Version        string  `json:"version"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
}
