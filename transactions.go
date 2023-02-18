package upgopher

const transBase = baseURL + "/transactions"

type TransactionList struct {
	Data []struct {
		Type string
		ID   string

		Attributes struct {
			Status          string
			RawText         string
			Description     string
			Message         string
			IsCategorizable bool
			HoldInfo        string
			RoundUp         string
			Cashback        string

			Amount struct {
				CurrencyCode     string
				Value            string
				ValueInBaseUnits int
			}

			ForeignAmount      string
			CardPurchaseMethod string
			SettledAt          string
			CreatedAt          string
		}

		Relationships struct {
			Account struct {
				Data struct {
					Type string
					ID   string
				}
				Links struct {
					Related string
				}

				Category struct {
					Data  string
					Links struct {
						Self string
					}

					ParentCategory struct {
						Data string
					}

					Tags struct {
						Data []struct {
							Type string
							ID   string
						}

						Links struct {
							Self string
						}
					}
				}
			}
		}

		Links struct {
			Self string
		}
	}

	Links struct {
		Prev string
		Next string
	}
}
