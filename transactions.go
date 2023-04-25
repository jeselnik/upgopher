package upgopher

import "fmt"

const transBase = baseURL + "/transactions"

type TransactionList struct {
	Transaction []Transaction `json:"data"`
	Links       NavLinks      `json:"links"`
}

type TransactionRes struct {
	Data Transaction `json:"data"`
}

type Transaction struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status          string `json:"status"`
		RawText         string `json:"rawText"`
		Description     string `json:"description"`
		Message         string `json:"message"`
		IsCategorizable bool   `json:"isCategorizable"`

		HoldInfo struct {
			Amount        MoneyObject `json:"amount"`
			ForeignAmount MoneyObject `json:"foreignAmount"`
		} `json:"holdInfo"`

		RoundUp struct {
			Amount       MoneyObject `json:"amount"`
			BoostPortion MoneyObject `json:"boostPortion"`
		} `json:"roundUp"`

		CashBack struct {
			Description string      `json:"description"`
			Amount      MoneyObject `json:"amount"`
		} `json:"cashback"`

		Amount        MoneyObject `json:"amount"`
		ForeignAmount MoneyObject `json:"foreignAmount"`

		CardPurchaseMethod struct {
			Method           string `json:"method"`
			CardNumberSuffix string `json:"cardNumberSuffix"`
		} `json:"cardPurchaseMethod"`

		SettledAt string `json:"settledAt"`
		CreatedAt string `json:"createdAt"`
	} `json:"attributes"`

	Relationships struct {
		Account struct {
			Data  DataObject    `json:"data"`
			Links LinkToRelated `json:"related"`
		} `json:"account"`

		TransferAccount struct {
			Data  DataObject    `json:"data"`
			Links LinkToRelated `json:"related"`
		} `json:"transferAccount"`

		Category struct {
			Data DataObject `json:"data"`

			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"category"`

		ParentCategory struct {
			Data  DataObject    `json:"data"`
			Links LinkToRelated `json:"links"`
		} `json:"parentCategory"`

		Tags struct {
			Data  []DataObject `json:"data"`
			Links LinkToSelf   `json:"links"`
		} `json:"tags"`
	} `json:"relationships"`

	Links LinkToSelf `json:"links"`
}

func GetTransactions(b Bearer) (TransactionList, error) {
	list := TransactionList{}
	err := get(&list, transBase, b)

	return list, err
}

func GetTransactionsByAccount(id string, b Bearer) (TransactionList, error) {
	list := TransactionList{}
	url := fmt.Sprintf("%s/%s/transactions", accountsBase, id)
	err := get(&list, url, b)
	return list, err
}

/* Working */
func GetTransactionById(id string, b Bearer) (TransactionRes, error) {
	var transaction TransactionRes
	url := fmt.Sprintf("%s/%s", transBase, id)
	err := get(&transaction, url, b)
	return transaction, err
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
