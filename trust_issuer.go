package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
	client "github.com/ory/hydra-client-go"
)

func main() {
	skipTLSClient1 := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	}
	trustJwtGrantIssuerBody := *client.NewTrustJwtGrantIssuerBody(time.Now(), "https://jwt-idp.example.com", *client.NewJSONWebKey("RS256", "1603dfe0af8f4596", "RSA", "sig"), []string{"Scope_example"}) // TrustJwtGrantIssuerBody |  (optional)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "https://localhost:5445", // Admin API URL
		},
	}
	configuration.HTTPClient = skipTLSClient1
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminApi.TrustJwtGrantIssuer(context.Background()).TrustJwtGrantIssuerBody(trustJwtGrantIssuerBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.TrustJwtGrantIssuer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrustJwtGrantIssuer`: TrustedJwtGrantIssuer
	fmt.Fprintf(os.Stdout, "Response from `AdminApi.TrustJwtGrantIssuer`: %v\n", resp)
}