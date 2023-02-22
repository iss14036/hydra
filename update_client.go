package main

import (
	"context"
	"crypto/tls"
	"fmt"
	client "github.com/ory/hydra-client-go"
	"net/http"
	"os"
)

func main() {
	skipTLSClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	}
	id := "daniel_client_id_6"
	oAuth2Client := *client.NewOAuth2Client()// string | The id of the OAuth 2.0 Client.
	oAuth2Client.SetGrantTypes([]string{"client_credentials"})
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "https://localhost:5445", // Public API URL
		},
	}

	configuration.HTTPClient = skipTLSClient
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminApi.UpdateOAuth2Client(context.Background(), id).OAuth2Client(oAuth2Client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.UpdateOAuth2Client``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOAuth2Client`: OAuth2Client
	fmt.Fprintf(os.Stdout, "Response from `AdminApi.UpdateOAuth2Client`: %v\n", resp)
}
