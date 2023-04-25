package upgopher

const catBase = baseURL + "/categories"

type CategoryResource struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Relationships struct {
		Parent struct {
			Data  DataObject    `json:"data"`
			Links LinkToRelated `json:"related"`
		} `json:"parent"`
		Children struct {
			Data  []DataObject  `json:"data"`
			Links LinkToRelated `json:"related"`
		} `json:"children"`
	} `json:"relationships"`

	Links LinkToSelf `json:"links"`
}

type Category struct {
	Data CategoryResource `json:"data"`
}

type CategoryList struct {
	Data []CategoryResource `json:"data"`
}
