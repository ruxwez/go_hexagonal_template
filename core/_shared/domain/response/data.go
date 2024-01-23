package response

type Data struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
