package upgopher

import "fmt"

const accountsBase = baseURL + "/accounts"

type AccountsList struct {
	Data  []Account `json:"data"`
	Links NavLinks  `json:"links"`
}

type AccountRes struct {
	Data Account `json:"data"`
}

type Account struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DisplayName   string      `json:"displayName"`
		AccountType   string      `json:"accountType"`
		OwnershipType string      `json:"ownershipType"`
		Balance       MoneyObject `json:"balance"`
		CreatedAt     string
	} `json:"attributes"`

	Relationships struct {
		Transactions struct {
			Links LinkToRelated `json:"related"`
		} `json:"transactions"`
	} `json:"relationships"`

	Links LinkToSelf `json:"links"`
}

func GetAccounts(b Bearer) (AccountsList, error) {
	list := AccountsList{}
	err := get(&list, accountsBase, b)

	return list, err
}

func (l AccountsList) FollowPrev(b Bearer) error {
	temp := AccountsList{}
	err := get(&temp, l.Links.Prev, b)
	if err != nil {
		return err
	}

	listAddr := &l
	*listAddr = temp

	return nil
}

func (l AccountsList) FollowNext(b Bearer) error {
	temp := AccountsList{}
	err := get(&temp, l.Links.Next, b)
	if err != nil {
		return err
	}

	listAddr := &l
	*listAddr = temp

	return nil
}

func GetAccount(id string, b Bearer) (AccountRes, error) {
	acc := AccountRes{}
	accUrl := fmt.Sprintf("%s/%s", accountsBase, id)
	err := get(&acc, accUrl, b)
	return acc, err
}
