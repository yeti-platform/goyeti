package goyeti

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yeti-platform/goyeti/models"
)

func TestNewYetiClient(t *testing.T) {
	type args struct {
		host   string
		port   int
		apikey string
	}
	tests := []struct {
		name string
		args args
		want *YetiClient
	}{
		{"TestNewYetiClientSuccess", args{"127.0.0.1", 1234, "apikey"}, &YetiClient{ClientConfig: models.ClientConfig{Host: "127.0.0.1", Port: 1234, Timeout: time.Second * 5, Proto: "https"}, apikey: "apikey"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewYetiClient(tt.args.host, tt.args.port, tt.args.apikey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewYetiClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseResponse(t *testing.T) {
	type args struct {
		res      *http.Response
		receiver map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestParseResponseSuccess",
			args{
				&http.Response{
					Body: io.NopCloser(bytes.NewBufferString("{\"access_token\":\"token\",\"token_type\":\"type\"}")),
				},
				map[string]interface{}{
					"access_token": "token",
					"token_type":   "type",
				},
			},
			false,
		},
		{
			"TestParseResponseError",
			args{
				&http.Response{
					Body: io.NopCloser(bytes.NewBufferString(":\"type\"}")),
				}, nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseResponse(tt.args.res, &tt.args.receiver); (err != nil) != tt.wantErr {
				t.Errorf("parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestYetiClient_doRequest(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	type args struct {
		method   string
		endpoint string
		header   http.Header
		data     []byte
	}
	tests := []struct {
		name     string
		args     args
		scenario scenario
		want     *http.Response
		wantErr  bool
	}{
		{
			"TestDoRequestSuccess",
			args{
				http.MethodGet,
				"api/v2/auth/me",
				http.Header{},
				nil,
			},
			scenario{
				http.StatusOK,
				"body",
			},
			&http.Response{
				Body:   io.NopCloser(bytes.NewBufferString("body")),
				Status: http.StatusText(http.StatusOK),
			},
			false,
		},
		{
			"TestDoRequestError",
			args{
				http.MethodGet,
				"api/v2/auth/me",
				http.Header{},
				nil,
			},
			scenario{
				http.StatusInternalServerError,
				"body",
			},
			&http.Response{
				Body:   io.NopCloser(bytes.NewBufferString("body")),
				Status: http.StatusText(http.StatusInternalServerError),
			},
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(tt.scenario.expectedRespStatus)
			res.Write([]byte(tt.scenario.expectedRespBody))
		}))
		defer func() { testServer.Close() }()

		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			got, err := c.doRequest(tt.args.method, tt.args.endpoint, tt.args.header, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.scenario.expectedRespStatus, got.StatusCode, "status code should match the expected response")
		})
	}
}

func TestYetiClient_getToken(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	tests := []struct {
		name     string
		scenario scenario
		want     string
		wantErr  bool
	}{
		{
			"TestGetTokenSuccess",
			scenario{
				http.StatusOK,
				"{\"access_token\":\"token\",\"token_type\":\"type\"}",
			},
			"token",
			false,
		},
		{
			"TestGetTokenError500",
			scenario{
				http.StatusInternalServerError,
				"error",
			},
			"",
			true,
		},
		{
			"TestGetTokenErrorMalformedResponse",
			scenario{
				http.StatusOK,
				"{",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(tt.scenario.expectedRespStatus)
			res.Write([]byte(tt.scenario.expectedRespBody))
		}))
		defer func() { testServer.Close() }()

		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			got, err := c.getToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.getToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YetiClient.getToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYetiClient_checkAuth(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name     string
		scenario scenario
		args     args
		want     bool
		wantErr  bool
	}{
		{
			"TestCheckAuthSuccess",
			scenario{
				http.StatusOK,
				"{\"username\":\"test\"}",
			},
			args{
				"token",
			},
			true,
			false,
		},
		{
			"TestCheckAuthError",
			scenario{
				http.StatusUnauthorized,
				"error",
			},
			args{
				"token",
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(tt.scenario.expectedRespStatus)
			res.Write([]byte(tt.scenario.expectedRespBody))
		}))
		defer func() { testServer.Close() }()
		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			got, err := c.checkAuth(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.checkAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YetiClient.checkAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYetiClient_init(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	tests := []struct {
		name     string
		scenario []scenario
		wantErr  bool
	}{
		{
			"TestInitSuccess",
			[]scenario{
				{
					http.StatusOK,
					"{\"access_token\":\"token\",\"token_type\":\"type\"}",
				},
				{
					http.StatusOK,
					"{\"username\":\"test\"}",
				},
			},
			false,
		},
		{
			"TestInitErrorFetchingToken",
			[]scenario{
				{
					http.StatusInternalServerError,
					"error",
				},
				{
					http.StatusOK,
					"",
				},
			},
			true,
		},
		{
			"TestInitErrorCheckingAuth",
			[]scenario{
				{
					http.StatusOK,
					"{\"access_token\":\"token\",\"token_type\":\"type\"}",
				},
				{
					http.StatusUnauthorized,
					"error",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/api/v2/auth/api-token" {
				res.WriteHeader(tt.scenario[0].expectedRespStatus)
				res.Write([]byte(tt.scenario[0].expectedRespBody))
			} else {
				res.WriteHeader(tt.scenario[1].expectedRespStatus)
				res.Write([]byte(tt.scenario[1].expectedRespBody))
			}
		}))
		defer func() { testServer.Close() }()
		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			if err := c.init(); (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestYetiClient_ObservablesSearch(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	type args struct {
		search models.QuerySearch
	}
	tests := []struct {
		name     string
		scenario []scenario
		args     args
		want     models.ServerResponseObservables
		wantErr  bool
	}{
		{
			"TestObservablesSearchSuccess",
			[]scenario{
				{
					http.StatusOK,
					"{\"observables\":[{\"id\":\"id\",\"type\":\"hostname\",\"value\":\"fake\"}],\"total\":1}",
				},
			},
			args{
				models.QuerySearch{
					Query: map[string]interface{}{
						"value": "fake",
					},
					Type:  "hostname",
					Count: 50,
					Page:  1,
				},
			},
			models.ServerResponseObservables{
				Observables: []models.Observable{
					{
						ID:    "id",
						Type:  "hostname",
						Value: "fake",
					},
				},
				Total: 1,
			},
			false,
		},
		{
			"TestObservablesSearchSuccessMultiplePages",
			[]scenario{
				{
					http.StatusOK,
					"{\"observables\":[{\"id\":\"id1\",\"type\":\"hostname\",\"value\":\"fake1\"}],\"total\":2}",
				},
				{
					http.StatusOK,
					"{\"observables\":[{\"id\":\"id2\",\"type\":\"hostname\",\"value\":\"fake2\"}],\"total\":2}",
				},
			},
			args{
				models.QuerySearch{
					Query: map[string]interface{}{
						"value": "fake",
					},
					Type:    "hostname",
					Sorting: []struct{}{},
					Count:   1,
					Page:    1,
				},
			},
			models.ServerResponseObservables{
				Observables: []models.Observable{
					{
						ID:    "id1",
						Type:  "hostname",
						Value: "fake1",
					},
					{
						ID:    "id2",
						Type:  "hostname",
						Value: "fake2",
					},
				},
				Total: 2,
			},
			false,
		},
		{
			"TestObservablesSearchError",
			[]scenario{
				{
					http.StatusInternalServerError,
					"error",
				},
			},
			args{
				models.QuerySearch{
					Query: map[string]interface{}{
						"value": "fake",
					},
					Type:    "hostname",
					Sorting: []struct{}{},
					Count:   1,
					Page:    1,
				},
			},
			models.ServerResponseObservables{},
			true,
		},
		{
			"TestObservablesSearchErrorUnmarshal",
			[]scenario{
				{
					http.StatusOK,
					"{error",
				},
			},
			args{
				models.QuerySearch{
					Query: map[string]interface{}{
						"value": "fake",
					},
					Type:    "hostname",
					Sorting: []struct{}{},
					Count:   1,
					Page:    1,
				},
			},
			models.ServerResponseObservables{},
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				t.Errorf("error reading request body: %v", err)
			}
			parsedBody := models.QuerySearch{}
			err = json.Unmarshal(body, &parsedBody)
			if err != nil {
				t.Errorf("error unmarshalling request body: %v", err)
			}
			page := parsedBody.Page
			res.WriteHeader(tt.scenario[page-1].expectedRespStatus)
			res.Write([]byte(tt.scenario[page-1].expectedRespBody))
		}))
		defer func() { testServer.Close() }()
		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			got, err := c.ObservablesSearch(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.ObservablesSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YetiClient.ObservablesSearch() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestYetiClient_Query(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	type args struct {
		endpoint string
		method   string
		data     []byte
	}
	tests := []struct {
		name     string
		args     args
		scenario []scenario
		want     map[string]interface{}
		wantErr  bool
	}{
		{
			"TestQuerySuccess",
			args{
				"api/v2/auth/me",
				http.MethodGet,
				nil,
			},
			[]scenario{
				{
					http.StatusOK,
					"{\"username\":\"test\"}",
				},
			},
			map[string]interface{}{
				"username": "test",
			},
			false,
		},
		{
			"TestQueryError",
			args{
				"api/v2/auth/me",
				http.MethodGet,
				nil,
			},
			[]scenario{
				{
					http.StatusInternalServerError,
					"error",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(tt.scenario[0].expectedRespStatus)
			res.Write([]byte(tt.scenario[0].expectedRespBody))
		}))
		defer func() { testServer.Close() }()
		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			got, err := c.Query(tt.args.endpoint, tt.args.method, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YetiClient.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYetiClient_Close(t *testing.T) {
	type scenario struct {
		expectedRespStatus int
		expectedRespBody   string
	}
	tests := []struct {
		name     string
		scenario []scenario
		wantErr  bool
	}{
		{
			"TestCloseSuccess",
			[]scenario{
				{
					http.StatusOK,
					"body",
				},
			},
			false,
		},
		{
			"TestCloseError",
			[]scenario{
				{
					http.StatusInternalServerError,
					"error",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(tt.scenario[0].expectedRespStatus)
			res.Write([]byte(tt.scenario[0].expectedRespBody))
		}))
		defer func() { testServer.Close() }()
		t.Run(tt.name, func(t *testing.T) {
			c := &YetiClient{
				url:    testServer.URL,
				client: testServer.Client(),
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("YetiClient.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
