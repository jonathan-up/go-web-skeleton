package response

type JsonResponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func New(message string, result interface{}) JsonResponse {
	return JsonResponse{
		Message: message,
		Result:  result,
	}
}
