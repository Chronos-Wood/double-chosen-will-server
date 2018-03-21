package status

type StatusCode int

const (
	SUCCESS        StatusCode = iota
	UNAUTHORIZED
	INTERNAL_ERROR
)

type ResponseMessage struct {
	Status  StatusCode `json:"status"`
	Message string     `json:"message"`
	Data    interface{}
}
