package main

import (
	"fmt"
	"os"

	"github.com/zenoss/go-auth0/auth0"
)

func main() {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	if auth0Domain == "" {
		fmt.Println("Error: AUTH0_DOMAIN must be defined")
		return
	}
	authzAPIURL := os.Getenv("AUTH0_AUTHORIZATION_API_URL")
	if authzAPIURL == "" {
		fmt.Println("Error: AUTH0_AUTHORIZATION_API_URL must be defined")
		return
	}
	authzAPIAudience := os.Getenv("AUTH0_AUTHORIZATION_API_AUDIENCE")
	if authzAPIAudience == "" {
		fmt.Println("Error: AUTH0_AUTHORIZATION_API_AUDIENCE must be defined")
		return
	}
	authzClientID := os.Getenv("AUTH0_AUTHORIZATION_CLIENT_ID")
	if authzClientID == "" {
		fmt.Println("Error: AUTH0_AUTHORIZATION_CLIENT_ID must be defined")
		return
	}
	authzClientSecret := os.Getenv("AUTH0_AUTHORIZATION_CLIENT_SECRET")
	if authzClientSecret == "" {
		fmt.Println("Error: AUTH0_AUTHORIZATION_CLIENT_SECRET must be defined")
		return
	}
	fmt.Println("Creating authz client")
	authorization := auth0.AuthzClientFromCredentials(
		auth0Domain,
		auth0.API{
			URL:          authzAPIURL,
			Audience:     []string{authzAPIAudience},
			ClientID:     authzClientID,
			ClientSecret: authzClientSecret,
		},
	)
	fmt.Println("Creating auth0 config")
	cfg, err := NewAuth0Config("auth0_config.yaml")
	if err != nil {
		fmt.Printf("WHAT THE ERR: %e", err)
	}

	fmt.Printf("%+v\n", cfg)
	fmt.Printf("Authz config is\n%+v\n\n", cfg)
	fmt.Println("Creating Roles")
	results := cfg.Authorization.CreateRoles(authorization.Roles)
	for _, result := range results {
		fmt.Println("Request data:")
		fmt.Printf("\t%+v\n", result.Request)
		fmt.Println("Result:")
		fmt.Printf("\t%+v\n", result.Result)
		fmt.Println()
	}

	fmt.Println("Creating Permissions")
	results = cfg.Authorization.CreatePermissions(authorization.Permissions)
	for _, result := range results {
		fmt.Println("Request data:")
		fmt.Printf("\t%+v\n", result.Request)
		fmt.Println("Result:")
		fmt.Printf("\t%+v\n", result.Result)
		fmt.Println()
	}
}
