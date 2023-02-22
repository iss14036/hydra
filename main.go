package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	openapiclient "github.com/ory/hydra-client-go"
)

func main() {
	skipTLSClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	}
	oAuth2Client := *openapiclient.NewOAuth2Client() // OAuth2Client |

	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = skipTLSClient
	configuration.Servers = []openapiclient.ServerConfiguration{
		{
			URL: "https://localhost:5445", // Admin API URL
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminApi.CreateOAuth2Client(context.Background()).OAuth2Client(oAuth2Client).Execute()
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.CreateOAuth2Client``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOAuth2Client`: OAuth2Client
	fmt.Fprintf(os.Stdout, "Response from `AdminApi.CreateOAuth2Client`: %v\n", resp)
}
