package helpers

// ResponseSchema generic struct response
type ResponseSchema struct {
	StatusCode int16       `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

// Response generic struct response
type Response interface {
	Header() (interface{}, error)
	Body() (interface{}, error)
}

// Header returns headers which need to be set
func (r ResponseSchema) Header() (interface{}, error) {
	return nil, nil
}

// Body returns response body
func (r ResponseSchema) Body() (interface{}, error) {
	return nil, nil
}
