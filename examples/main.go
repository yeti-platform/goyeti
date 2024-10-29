package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yeti-platform/goyeti"
	"golang.org/x/oauth2"
)

func main() {
	domain := "email2.elibria.io"
	auth := context.WithValue(
		context.Background(),
		goyeti.ContextAPIKeys,
		map[string]goyeti.APIKey{
			"APIKeyHeader": {Key: os.Getenv("YETI_API_KEY")},
		},
	)
	cfg := goyeti.NewConfiguration()
	cfg.Host = os.Getenv("YETI_URL")
	cfg.Scheme = "https"
	client := goyeti.NewAPIClient(cfg)
	body, resp, err := client.AuthAPI.LoginApiApiV2AuthApiTokenPost(auth).Execute()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("Failed to authenticate with status code %d", resp.StatusCode))
	}

	access_token := body.(map[string]interface{})["access_token"].(string)
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: access_token,
		TokenType:   "Bearer",
	})
	queryType := goyeti.HOSTNAME
	auth = context.WithValue(oauth2.NoContext, goyeti.ContextOAuth2, tokenSource)
	observableSearchRequest := goyeti.NewObservableSearchRequest()
	observableSearchRequest.SetType(
		goyeti.ObservableSearchRequestType{
			ObservableTypeInput: &queryType,
		},
	)
	query := map[string]goyeti.QueryValue{
		"value": goyeti.QueryValue{
			String: &domain,
		},
	}
	observableSearchRequest.SetQuery(query)

	foo, response, err := client.ObservablesAPI.SearchApiV2ObservablesSearchPost(auth).ObservableSearchRequest(*observableSearchRequest).Execute()
	fmt.Println(foo)
	fmt.Println(response)
}
