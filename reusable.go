package upgopher

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

/* Reusable Objects */
type MoneyObject struct {
	CurrencyCode     string `json:"currencyCode"`
	Value            string `json:"value"`
	ValueInBaseUnits int    `json:"valueInBaseUnits"`
}

/* Functions */
func newRequest(url string, b Bearer) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerKey, b.Val)
	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

/* Generics */
type response interface {
	PingRes | AccountsList | Account | TransactionList | Transaction
}

func get[T response](list *T, url string, b Bearer) error {
	res, err := newRequest(url, b)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(res, list)
	return jsonErr
}
