package upgopher

const transBase = baseURL + "/transactions"

type TransactionList struct {
	Data []struct {
		Type string `json:"type"`
		ID   string `json:"id"`

		Attributes struct {
			Status          string `json:"status"`
			RawText         string `json:"rawText"`
			Description     string `json:"description"`
			Message         string `json:"message"`
			IsCategorizable bool   `json:"isCategorizable"`
			HoldInfo        string `json:"holdInfo"`
			RoundUp         string `json:"roundUp"`
			Cashback        string `json:"cashback"`

			Amount struct {
				CurrencyCode     string `json:"currencyCode"`
				Value            string `json:"value"`
				ValueInBaseUnits int    `json:"valueInBaseUnits"`
			} `json:"amount"`

			ForeignAmount      string `json:"foreignAmount"`
			CardPurchaseMethod string `json:"cardPurchaseMethod"`
			SettledAt          string `json:"settledAt"`
			CreatedAt          string `json:"createdAt"`
		} `json:"attributes"`

		Relationships struct {
			Account struct {
				Data struct {
					Type string
					ID   string
				} `json:"data"`
				Links struct {
					Related string
				} `json:"links"`

				Category struct {
					Data  string `json:"data"`
					Links struct {
						Self string `json:"self"`
					} `json:"links"`

					ParentCategory struct {
						Data string `json:"data"`
					} `json:"parentCategory"`

					Tags struct {
						Data []struct {
							Type string `json:"type"`
							ID   string `json:"id"`
						} `json:"data"`

						Links struct {
							Self string `json:"self"`
						} `json:"links"`
					} `json:"tags"`
				} `json:"category"`
			} `json:"account"`
		} `json:"relationships"`

		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`

	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

func GetTransactions(b Bearer) (TransactionList, error) {
	list := TransactionList{}
	err := get(&list, transBase, b)

	return list, err
}

func (l TransactionList) FollowNext(b Bearer) error {
	temp := TransactionList{}
	err := get(&temp, l.Links.Next, b)
	if err != nil {
		return err
	}

	listAddr := &l
	*listAddr = temp
	return nil
}

func (l TransactionList) FollowPrev(b Bearer) error {
	temp := TransactionList{}
	err := get(&temp, l.Links.Prev, b)
	if err != nil {
		return err
	}

	listAddr := &l
	*listAddr = temp
	return nil
}
