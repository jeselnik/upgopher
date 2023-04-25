package upgopher

import "fmt"

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
	Category CategoryResource `json:"data"`
}

type CategoryList struct {
	Category []CategoryResource `json:"data"`
}

func GetCategories(b Bearer) (CategoryList, error) {
	list := new(CategoryList)
	err := get(list, catBase, b)
	return *list, err
}

func GetCategoryById(b Bearer, id string) (Category, error) {
	category := new(Category)
	url := fmt.Sprintf(catBase+"/%s", id)
	err := get(category, url, b)
	return *category, err
}
