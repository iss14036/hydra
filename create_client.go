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
	clientName := "daniel_client_id_6"
	oAuth2Client := *client.NewOAuth2Client() // OAuth2Client |
	oAuth2Client.SetClientId("daniel_client_id_6")
	oAuth2Client.SetClientName(clientName)
	oAuth2Client.GrantTypes = []string{"authorization_code", "client_credentials"}
	oAuth2Client.SetClientSecret("daniel_client_id_6")
	oAuth2Client.SetScope("email tracking profile")
	oAuth2Client.SetMetadata(
		map[string]interface{}{
			"account_id":  1,
			"customer_id": 1,
		},
	)
	oAuth2Client.SetJwks(
		map[string]interface{}{
			"kty": "RSA",
			"e":   "AQAB",
			"kid": "d8e91f55-67e0-4e56-a066-6a5f0c2efdf7",
			"n":   "nzyis1ZjfNB0bBgKFMSvvkTtwlvBsaJq7S5wA-kzeVOVpVWwkWdVha4s38XM_pa_yr47av7-z3VTmvDRyAHcaT92whREFpLv9cj5lTeJSibyr_Mrm_YtjCZVWgaOYIhwrXwKLqPr_11inWsAkfIytvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0e-lf4s4OxQawWD79J9_5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWbV6L11BWkpzGXSW4Hv43qa-GSYOD2QU68Mb59oSk2OB-BtOLpJofmbGEGgvmwyCI9Mw",
		},
	)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "https://localhost:5445", // Admin API URL
		},
	}
	configuration.HTTPClient = skipTLSClient
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminApi.CreateOAuth2Client(context.Background()).OAuth2Client(oAuth2Client).Execute()
	if err != nil {
		switch r.StatusCode {
		case http.StatusConflict:
			fmt.Fprintf(os.Stderr, "Conflict when creating oAuth2Client: %v\n", err)
		default:
			fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.CreateOAuth2Client``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
	}
	// response from `CreateOAuth2Client`: OAuth2Client
	fmt.Fprintf(os.Stdout, "Created client with name %s\n", resp.GetClientName())

	limit := int64(20)
	offset := int64(0)
	clients, r, err := apiClient.AdminApi.ListOAuth2Clients(context.Background()).Limit(limit).Offset(offset).ClientName(clientName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ListOAuth2Clients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "We have %d clients\n", len(clients))
	fmt.Fprintf(os.Stdout, "First client name: %s\n", clients[0].GetClientName())

}

////docker-compose -f quickstart.yml up --build