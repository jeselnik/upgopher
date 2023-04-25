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

type tagInputResource struct {
	tag string `json:"type"`
	id  string `json:"id"`
}

func newTagInputRes(id string) tagInputResource {
	return tagInputResource{"tags", id}
}

func GetTags(b Bearer) (TagList, error) {
	list := new(TagList)
	err := get(list, tagBase, b)
	return *list, err
}

func AddTagsToTransaction(transactionID string, tagID []string,
	b Bearer) error {

	tagInput := make([]tagInputResource, len(tagID))
	for i, t := range tagID {
		tagInput[i] = newTagInputRes(t)
	}

	return nil
}
