package quickemailverification

import "encoding/json"

// BuildResponseModel struct creates the actual Response struct (the response from QuickEmailVerification)
type BuildResponseModel struct{}

// Defines interface for creating Response objects
type ResponseModel interface {
        NewResponseModel(response []byte) (*Response, error)
}

// NewResponseModel creates a new Response object from an JSON API response
func (b BuildResponseModel) NewResponseModel(response []byte) (*Response, error) {
	result := &Response{}
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}

	return result, nil
}
