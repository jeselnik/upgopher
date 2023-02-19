package upgopher

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://api.up.com.au/api/v1"
const headerKey = "Authorization"
const pingURL = baseURL + "/util/ping"

type Bearer struct {
	Val string
}

func Up(token string) Bearer {
	return Bearer{"Bearer " + token}
}

func newRequest(url string, b Bearer) ([]byte, error) {
	client := &http.Client{}
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

type PingRes struct {
	Meta struct {
		ID          string `json:"id"`
		StatusEmoji string `json:"statusEmoji"`
	} `json:"meta"`
}

type response interface {
	PingRes | AccountsList | Account | TransactionList
}

func get[T response](list *T, url string, b Bearer) error {
	res, err := newRequest(url, b)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(res, list)
	return jsonErr
}

func Ping(b Bearer) (PingRes, error) {
	ping := PingRes{}
	err := get(&ping, pingURL, b)

	return ping, err
}
