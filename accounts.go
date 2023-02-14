package upgopher

const accountsBase = baseURL + "/accounts"

type AccountsList struct {
	Data []struct {
		Type string
		ID   string

		Attributes struct {
			DisplayName   string
			AccountType   string
			OwnershipType string

			Balance struct {
				CurrencyCode     string
				Value            string
				ValueInBaseUnits string
			}

			CreatedAt string
		}

		Relationships struct {
			Transactions struct {
				Links struct {
					Related string
				}
			}

			Links struct {
				Self string
			}
		}
	}

	Links struct {
		Prev string
		Next string
	}
}

func GetAccounts(b Bearer) AccountsList {
	return *new(AccountsList)
}

func GetAccount(b Bearer, id string) AccountsList {
	return *new(AccountsList)
}
