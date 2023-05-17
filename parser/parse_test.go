package parser_test

import (
	"net/http"
	"testing"

	"github.com/p1ck0/multiset/entity"
	"github.com/p1ck0/multiset/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		want    []entity.Request
	}{
		{
			name:    "empty",
			content: []byte(``),
			want:    nil,
		},
		{
			name: "simple",
			content: []byte(`
				{
					"auto_complete": {
						"method": "GET",
						"headers": {
							"Accept": "text/html,application/xhtml+xml",
							"Accept-Language": "en-US,en;q=0.9",
							"Cache-Control": "max-age=0",
							"Connection": "keep-alive",
							"User-Agent": "Mozilla/5.0"
						},
						"body": {
							"key": "value"
						}
					},
					"requests": [
						{
							"url": "https://google.com"
						},
						{
							"method": "POST",
							"url": "https://bing.com",
							"headers": {
								"Application": "application/json"
							},
							"body": {
								"key1": "value1"
							}
						}
					]
				}
			`),
			want: []entity.Request{
				{
					Method: "GET",
					URL:    "https://google.com",
					Headers: http.Header{
						"Accept":          []string{"text/html,application/xhtml+xml"},
						"Accept-Language": []string{"en-US,en;q=0.9"},
						"Cache-Control":   []string{"max-age=0"},
						"Connection":      []string{"keep-alive"},
						"User-Agent":      []string{"Mozilla/5.0"},
					},
					Body: []byte(`{"key":"value"}`),
				},
				{
					Method: "POST",
					URL:    "https://bing.com",
					Headers: http.Header{
						"Application":     []string{"application/json"},
						"Accept":          []string{"text/html,application/xhtml+xml"},
						"Accept-Language": []string{"en-US,en;q=0.9"},
						"Cache-Control":   []string{"max-age=0"},
						"Connection":      []string{"keep-alive"},
						"User-Agent":      []string{"Mozilla/5.0"},
					},
					Body: []byte(`{"key":"value","key1":"value1"}`),
				},
			},
		},
		{
			name: "full auto complete",
			content: []byte(`
					{
						"auto_complete": {
							"body": {
								"enable": true,
								"type": "auto"
							},
							"headers": {
								"Content-Type": "application/json"
							},
							"method": "POST"
						},
						"requests": [
							{
								"url": "http://localhost:8080/api/test"
							},
							{
								"url": "http://localhost:8081/api/test"
							}
						]
					}`),
			want: []entity.Request{
				{
					Method: "POST",
					URL:    "http://localhost:8080/api/test",
					Headers: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: []byte(`{"enable":true,"type":"auto"}`),
				},
				{
					Method: "POST",
					URL:    "http://localhost:8081/api/test",
					Headers: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: []byte(`{"enable":true,"type":"auto"}`),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse(tt.content)
			if err != nil && len(tt.content) != 0 {
				t.Errorf("Parse() error = %v", err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
