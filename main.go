package main

import (
	"context"
	"log"
	"os"

	"github.com/sigstore/cosign/v2/pkg/providers"
	_ "github.com/sigstore/cosign/v2/pkg/providers/github"
	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Invalid args")
	}

	tokenURL := os.Args[1]
	clientID := os.Args[2]

	ctx := context.Background()
	if !providers.Enabled(ctx) {
		log.Fatalf("incorrect environment")
	}
	token, err := providers.Provide(ctx, "guac")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if token == "" {
		log.Fatalf("empty token")
	}
	log.Printf("ID Token aquired")

	var conf oauth2.Config
	conf.Endpoint.TokenURL = tokenURL
	options := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer"),
		oauth2.SetAuthURLParam("assertion", token),
		oauth2.SetAuthURLParam("scope", "openid"),
		oauth2.SetAuthURLParam("client_id", clientID),
	}
	tok, err := conf.Exchange(ctx, "", options...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if tok.AccessToken == "" {
		log.Fatalf("empty token")
	}
	log.Printf("access Token aquired")

	// var scopes []string
	// var audience string
	// var ctx context.Context
	// var conf oauth2.Config
	// var tok oauth2.Token
	// //... fill in scopes, audience, ctx, conf, and get original access token in tok ...
	// options := []oauth2.AuthCodeOption{
	// 	oauth2.SetAuthURLParam("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange"),
	// 	oauth2.SetAuthURLParam("subject_token", tok.AccessToken),
	// 	oauth2.SetAuthURLParam("subject_token_type", "urn:ietf:params:oauth:token-type:access_token"),
	// }
	// if len(scopes) > 0 {
	// 	options = append(options, oauth2.SetAuthURLParam("scope", strings.Join(scopes, " ")))
	// }
	// if audiences != "" {
	// 	options = append(options, oauth2.SetAuthURLParam("audience", audience))
	// }
	// tok, err = conf.Exchange(ctx, "", options...)

}
