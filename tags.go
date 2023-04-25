package upgopher

const tagBase = baseURL + "/tags"

type TagResource struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Relationships struct {
		Transactions struct {
			LinkToRelated `json:"links"`
		} `json:"transactions"`
	} `json:"relationships"`
}

type TagList struct {
	Tag   []TagResource `json:"data"`
	Links NavLinks      `json:"links"`
}

func GetTags(b Bearer) (TagList, error) {
	list := new(TagList)
	err := get(list, tagBase, b)
	return *list, err
}
