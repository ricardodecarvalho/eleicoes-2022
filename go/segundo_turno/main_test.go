package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockHTTPClient func(r *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockHTTPClient
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestApuracao(t *testing.T) {

	jsonResponse := `{
		"dg": "19/10/2022",
		"hg": "20:31:31",
		"pst": "0,00",
		"cand": [{
				"seq": "1",
				"sqcand": "123456789",
				"n": "99",
				"nm": "Schmidt",
				"e": "n",
				"vap": "0",
				"pvap": "00,00"
			},
			{
				"seq": "2",
				"sqcand": "987654321",
				"n": "98",
				"nm": "Müller",
				"e": "n",
				"vap": "0",
				"pvap": "00,00"
			}
		]
	}`

	body := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       body,
			}, nil
		},
	}

	err := Apuracao()
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}
