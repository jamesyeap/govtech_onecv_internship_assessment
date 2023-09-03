package handlers

type ParamError struct {
	Issues map[string]interface{} `json:"Invalid params received"`
}