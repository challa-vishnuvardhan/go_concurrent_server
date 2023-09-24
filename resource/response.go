package resource

type Request struct {
	Message string `json:"message"`
}

// Response represents the JSON response structure sent to clients.
type Response struct {
	Message string `json:"message"`
}
