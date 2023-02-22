package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	openapiclient "github.com/ory/hydra-client-go"
	"net/http"
	"os"
)

type BasicAuthTransport2 struct {
	Username string
	Password string
}

func (t BasicAuthTransport2) RoundTrip(req *http.Request) (*http.Response, error) {
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

	grantType := "grantType_example"       // string |
	clientId := "82ca0ea2-c787-466d-98b6-2f72e0c7779d"         // string |  (optional)
	code := "code_example"                 // string |  (optional)
	redirectUri := ""   // string |  (optional)
	refreshToken := "refreshToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	configuration.Servers = []openapiclient.ServerConfiguration{
		{
			URL: "https://localhost:5444", // Public API URL
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	configuration.HTTPClient.Transport = BasicAuthTransport2{Username: "foo", Password: "bar"}
	resp, r, err := apiClient.PublicApi.Oauth2Token(context.Background()).GrantType(grantType).ClientId(clientId).Code(code).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.Oauth2Token``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2Token`: Oauth2TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicApi.Oauth2Token`: %v\n", resp)
}
