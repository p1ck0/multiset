package parser

import (
	"encoding/json"
	"net/http"

	"github.com/p1ck0/multiset/entity"
)

func Parse(r []byte) ([]entity.Request, error) {
	config := entity.Config{}
	if err := json.Unmarshal(r, &config); err != nil {
		return nil, err
	}

	result := make([]entity.Request, 0, len(config.RequestsDesc))

	for _, r := range config.RequestsDesc {
		reqD := entity.RequestDesc{
			Method:  config.AutoComplete.Method,
			Headers: config.AutoComplete.Headers,
			Body:    config.AutoComplete.Body,
		}
		reqD.Complete(r.Method, r.URL, r.Headers, r.Body)
		body, err := json.Marshal(reqD.Body)
		if err != nil {
			return nil, err
		}

		req := entity.Request{
			Body:    body,
			URL:     reqD.URL,
			Method:  reqD.Method,
			Headers: make(http.Header),
		}

		for k, v := range reqD.Headers {
			req.Headers.Set(k, v)
		}

		result = append(result, req)
	}
	return result, nil
}
