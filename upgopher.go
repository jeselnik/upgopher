package upgopher

const baseURL = "https://api.up.com.au/api/v1"

type upAPI struct {
	headerVal string
}

func NewUp(token string) upAPI {
	return upAPI{"Bearer:" + "token"}
}
