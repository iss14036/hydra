package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	client "github.com/ory/hydra-client-go"
	"net/http"
)

/*
Some endpoints require basic authentication.
The following code example shows how to make an authenticated request to the Ory Hydra Admin API:
*/

type BasicAuthTransport struct {
	Username string
	Password string
}

func (t BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
			t.Username, t.Password)))))
	skipTLSClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	}
	return skipTLSClient.Transport.RoundTrip(req)
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		"daniel_client_id_2", "secret"))))

	config := client.NewConfiguration()
	config.Servers = []client.ServerConfiguration{
		{
			URL: "https://localhost:5445", // Admin API
		},
	}

	c := client.NewAPIClient(config)
	config.HTTPClient.Transport = BasicAuthTransport{Username: "foo", Password: "bar"}

	req := c.AdminApi.GetConsentRequest(context.Background()).ConsentChallenge("consentChallenge_example")
	fmt.Println(req.Execute())
}