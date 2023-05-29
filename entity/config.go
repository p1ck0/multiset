package entity

type Config struct {
	AutoComplete AutoComplete  `json:"auto_complete"`
	RequestsDesc []RequestDesc `json:"requests"`
}

type AutoComplete struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    map[string]any    `json:"body"`
}

type RequestDesc struct {
	Method string `json:"method"`
	URL    string `json:"url"`

	Headers map[string]string `json:"headers"`
	Body    map[string]any    `json:"body"`
}

func (r *RequestDesc) Complete(method string, url string, headers map[string]string, body map[string]any) {
	if r.Headers == nil {
		r.Headers = make(map[string]string, len(headers))
	}
	for k, v := range headers {
		r.Headers[k] = v
	}
	if r.Body == nil {
		r.Body = make(map[string]any, len(body))
	}
	for k, v := range body {
		r.Body[k] = v
	}
	if len(method) > 0 {
		r.Method = method
	}
	r.URL = url
}
