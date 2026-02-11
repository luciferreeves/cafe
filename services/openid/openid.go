package openid

import (
	"cafe/config"
	"cafe/types"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"slices"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	Provider     *oidc.Provider
	OAuth2Config *oauth2.Config
	Verifier     *oidc.IDTokenVerifier
)

func init() {
	if config.OpenID.DiscoveryURL == "" {
		log.Fatal("OPENID_DISCOVERY_URL not configured. OpenID authentication is required.")
	}

	ctx := context.Background()
	var err error

	Provider, err = oidc.NewProvider(ctx, config.OpenID.DiscoveryURL)
	if err != nil {
		log.Fatalf("Failed to initialize OpenID provider: %v", err)
	}

	OAuth2Config = &oauth2.Config{
		ClientID:     config.OpenID.ClientID,
		ClientSecret: config.OpenID.ClientSecret,
		RedirectURL:  config.OpenID.CallbackURL,
		Endpoint:     Provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile", "groups"},
	}

	Verifier = Provider.Verifier(&oidc.Config{
		ClientID: config.OpenID.ClientID,
	})

	log.Println("OpenID Connect provider initialized successfully")
}

func GenerateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GetAuthURL(state string) string {
	return OAuth2Config.AuthCodeURL(state)
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return OAuth2Config.Exchange(ctx, code)
}

func VerifyIDToken(ctx context.Context, rawIDToken string) (*oidc.IDToken, error) {
	return Verifier.Verify(ctx, rawIDToken)
}

func GetUserInfo(ctx context.Context, token *oauth2.Token, idToken *oidc.IDToken) (*types.UserInfo, error) {
	var userInfo types.UserInfo
	if err := idToken.Claims(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse ID token claims: %v", err)
	}

	userInfoEndpoint, err := Provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
	if err != nil {
		log.Printf("Warning: Failed to fetch additional user info from userinfo endpoint: %v", err)
		return &userInfo, nil
	}

	var additionalClaims types.UserInfo
	if err := userInfoEndpoint.Claims(&additionalClaims); err != nil {
		log.Printf("Warning: Failed to parse userinfo endpoint claims into UserInfo: %v", err)
		return &userInfo, nil
	}

	if len(additionalClaims.Groups) > 0 {
		userInfo.Groups = additionalClaims.Groups
	}

	return &userInfo, nil
}

func IsAdmin(userInfo *types.UserInfo) bool {
	return slices.Contains(userInfo.Groups, "administrator")
}
