package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yeti-platform/goyeti"
	"golang.org/x/oauth2"
)

func ExampleObservable() {
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
	auth = context.WithValue(context.Background(), goyeti.ContextOAuth2, tokenSource)

	whoami, response, err := client.AuthAPI.MeApiV2AuthMeGet(auth).Execute()
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(fmt.Errorf("Failed to get user details with status code %d", response.StatusCode))
	}
	_, connected := whoami.GetIdOk()
	fmt.Println("Connected:", connected)

	queryType := goyeti.HOSTNAME
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

	observableResponse, reponse, err := client.ObservablesAPI.SearchApiV2ObservablesSearchPost(auth).ObservableSearchRequest(*observableSearchRequest).Execute()
	if err != nil {
		panic(err)
	}
	if reponse.StatusCode != 200 {
		panic(fmt.Errorf("Failed to search observables with status code %d", reponse.StatusCode))
	}
	obeservables := observableResponse.GetObservables()
	fmt.Printf("Total observables found %d for domain %s\n", len(observableResponse.Observables), domain)
	for _, observable := range obeservables {
		fmt.Printf("%s created at %s\n", observable.ASNOutput.Value, observable.ASNOutput.Created)
	}
	//Output: Connected: true
	//Total observables found 1 for domain email2.elibria.io
	//email2.elibria.io created at 2024-07-05 13:20:42.782787 +0000 UTC
}
