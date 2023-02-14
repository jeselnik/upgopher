package upgopher

const baseURL = "https://api.up.com.au/api/v1"

type Bearer struct {
	val string
}

func Up(token string) Bearer {
	return Bearer{"Bearer:" + "token"}
}
