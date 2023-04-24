package upgopher

const (
	baseURL   = "https://api.up.com.au/api/v1"
	headerKey = "Authorization"
	pingURL   = baseURL + "/util/ping"
)

type Bearer struct {
	Val string
}

func Up(token string) Bearer {
	return Bearer{"Bearer " + token}
}

type PingRes struct {
	Meta struct {
		ID          string `json:"id"`
		StatusEmoji string `json:"statusEmoji"`
	} `json:"meta"`
}

func Ping(b Bearer) (PingRes, error) {
	ping := PingRes{}
	err := get(&ping, pingURL, b)

	return ping, err
}
