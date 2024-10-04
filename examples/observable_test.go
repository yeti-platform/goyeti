// Description: This is a simple example of how to use the goyeti package to:
// - authenticate with the Yeti platform.
// - do arbitrary queries in the Yeti platform.
// - search for observables in the Yeti platform.
// Note that you need to setup the YETI_URL and YETI_API_KEY environment variables to run this example.
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yeti-platform/goyeti"
)

func ExampleObservable() {
	domain := "email2.elibria.io"
	client := goyeti.NewYetiClient(
		os.Getenv("YETI_URL"),
		443,
		os.Getenv("YETI_API_KEY"),
	)
	err := client.Init()
	if err != nil {
		panic(err)
	}
	whoami, err := client.Query("api/v2/auth/me", http.MethodGet, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected: %t\n", whoami["enabled"].(bool))
	observables, err := client.ObservablesSearch(domain, "hostname")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total observables found %d for domain %s\n", observables.Total, domain)
	err = client.Close()
	if err != nil {
		panic(err)
	}
	//Output: Connected: true
	//Total observables found 1 for domain email2.elibria.io
}
