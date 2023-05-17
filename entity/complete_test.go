package entity_test

import (
	"testing"

	"github.com/p1ck0/multiset/entity"
	"github.com/stretchr/testify/assert"
)

func TestComplete(t *testing.T) {
	tests := []struct {
		name   string
		config entity.Config
		want   entity.Config
	}{
		{
			name:   "empty",
			config: entity.Config{},
			want:   entity.Config{},
		},
		{
			name: "simple",
			config: entity.Config{
				AutoComplete: entity.AutoComplete{
					Method: "GET",
					Headers: map[string]string{
						"Accept":          "text/html,application/xhtml+xml",
						"Accept-Language": "en-US,en;q=0.9",
						"Cache-Control":   "max-age=0",
						"Connection":      "keep-alive",
						"User-Agent":      "Mozilla/5.0",
					},
					Body: map[string]any{
						"key": "value",
					},
				},
				RequestsDesc: []entity.RequestDesc{
					{
						Method: "POST",
						URL:    "https://google.com",
						Headers: map[string]string{
							"Application": "application/json",
						},
						Body: map[string]any{
							"key1": "value1",
						},
					},
				},
			},
			want: entity.Config{
				AutoComplete: entity.AutoComplete{
					Method: "GET",
					Headers: map[string]string{
						"Accept":          "text/html,application/xhtml+xml",
						"Accept-Language": "en-US,en;q=0.9",
						"Cache-Control":   "max-age=0",
						"Connection":      "keep-alive",
						"User-Agent":      "Mozilla/5.0",
					},
					Body: map[string]any{
						"key": "value",
					},
				},
				RequestsDesc: []entity.RequestDesc{
					{
						Method: "POST",
						URL:    "https://google.com",
						Headers: map[string]string{
							"Application":     "application/json",
							"Accept":          "text/html,application/xhtml+xml",
							"Accept-Language": "en-US,en;q=0.9",
							"Cache-Control":   "max-age=0",
							"Connection":      "keep-alive",
							"User-Agent":      "Mozilla/5.0",
						},
						Body: map[string]any{
							"key1": "value1",
							"key":  "value",
						},
					},
				},
			},
		},
		{
			name: "full complete",
			config: entity.Config{
				AutoComplete: entity.AutoComplete{
					Method: "GET",
					Headers: map[string]string{
						"Accept":          "text/html,application/xhtml+xml",
						"Accept-Language": "en-US,en;q=0.9",
						"Cache-Control":   "max-age=0",
						"Connection":      "keep-alive",
						"User-Agent":      "Mozilla/5.0",
					},
					Body: map[string]any{
						"key": "value",
					},
				},
				RequestsDesc: []entity.RequestDesc{
					{
						URL: "https://google.com",
					},
					{
						URL: "https://bing.com",
					},
				},
			},
			want: entity.Config{
				AutoComplete: entity.AutoComplete{
					Method: "GET",
					Headers: map[string]string{
						"Accept":          "text/html,application/xhtml+xml",
						"Accept-Language": "en-US,en;q=0.9",
						"Cache-Control":   "max-age=0",
						"Connection":      "keep-alive",
						"User-Agent":      "Mozilla/5.0",
					},
					Body: map[string]any{
						"key": "value",
					},
				},
				RequestsDesc: []entity.RequestDesc{
					{
						Method: "GET",
						Headers: map[string]string{
							"Accept":          "text/html,application/xhtml+xml",
							"Accept-Language": "en-US,en;q=0.9",
							"Cache-Control":   "max-age=0",
							"Connection":      "keep-alive",
							"User-Agent":      "Mozilla/5.0",
						},
						Body: map[string]any{
							"key": "value",
						},
						URL: "https://google.com",
					},
					{
						Method: "GET",
						Headers: map[string]string{
							"Accept":          "text/html,application/xhtml+xml",
							"Accept-Language": "en-US,en;q=0.9",
							"Cache-Control":   "max-age=0",
							"Connection":      "keep-alive",
							"User-Agent":      "Mozilla/5.0",
						},
						Body: map[string]any{
							"key": "value",
						},
						URL: "https://bing.com",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, r := range tt.config.RequestsDesc {
				req := entity.RequestDesc{
					Method:  tt.config.AutoComplete.Method,
					Headers: tt.config.AutoComplete.Headers,
					Body:    tt.config.AutoComplete.Body,
				}

				req.Complete(r.Method, r.URL, r.Headers, r.Body)
				assert.Equal(t, tt.want.RequestsDesc[i], req)
			}
		})
	}
}
