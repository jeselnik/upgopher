package upgopher

import "encoding/json"

const accountsBase = baseURL + "/accounts"

type AccountsList struct {
	Data []struct {
		Type string `json:"type"`
		ID   string `json:"id"`

		Attributes struct {
			DisplayName   string `json:"displayName"`
			AccountType   string `json:"accountType"`
			OwnershipType string `json:"ownershipType"`

			Balance struct {
				CurrencyCode     string `json:"currencyCode"`
				Value            string `json:"value"`
				ValueInBaseUnits int    `json:"valueInBaseUnits"`
			} `json:"balance"`

			CreatedAt string `json:"createdAt"`
		} `json:"attributes"`

		Relationships struct {
			Transactions struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"transactions"`

			Links struct {
				Self string
			}
		} `json:"relationships"`
	} `json:"data"`

	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next`
	} `json:"links"`
}

func GetAccounts(b Bearer) (AccountsList, error) {
	list := AccountsList{}
	res, err := newRequest(accountsBase, b)
	if err != nil {
		return list, err
	}

	jsonErr := json.Unmarshal(res, &list)

	return list, jsonErr
}

func GetAccount(b Bearer, id string) AccountsList {
	return *new(AccountsList)
}
