package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	homebox "github.com/jeffresc/homebox-sdk-go/client"
)

func main() {
	// Get configuration from environment variables
	baseURL := os.Getenv("HOMEBOX_URL")
	if baseURL == "" {
		log.Fatal("HOMEBOX_URL environment variable is required")
	}

	username := os.Getenv("HOMEBOX_USERNAME")
	password := os.Getenv("HOMEBOX_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("HOMEBOX_USERNAME and HOMEBOX_PASSWORD environment variables are required")
	}

	// Create HTTP client
	httpClient := &http.Client{}

	// Create unauthenticated client for login
	c, err := homebox.NewClientWithResponses(baseURL, homebox.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Login to get authentication token
	fmt.Println("Logging in...")
	loginResp, err := c.PostV1UsersLoginWithResponse(ctx, nil, homebox.PostV1UsersLoginJSONRequestBody{
		Username: &username,
		Password: &password,
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	if loginResp.JSON200 == nil {
		log.Fatalf("Login failed: unexpected response status %d", loginResp.StatusCode())
	}

	token := loginResp.JSON200.Token
	fmt.Printf("Successfully logged in! Token: %s...\n", (*token)[:20])

	// Create authenticated client
	authClient, err := homebox.NewClientWithResponses(
		baseURL,
		homebox.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", *token)
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create authenticated client: %v", err)
	}

	// Example 1: Get all items
	fmt.Println("\nFetching items...")
	itemsResp, err := authClient.GetV1ItemsWithResponse(ctx, &homebox.GetV1ItemsParams{})
	if err != nil {
		log.Fatalf("Failed to get items: %v", err)
	}

	if itemsResp.JSON200 != nil && itemsResp.JSON200.Items != nil {
		fmt.Printf("Found %d items:\n", len(*itemsResp.JSON200.Items))
		for _, item := range *itemsResp.JSON200.Items {
			fmt.Printf("  - %s (ID: %s)\n", *item.Name, *item.Id)
		}
	} else {
		fmt.Println("No items found or unable to parse response")
	}

	// Example 2: Get all locations
	fmt.Println("\nFetching locations...")
	locationsResp, err := authClient.GetV1LocationsWithResponse(ctx, &homebox.GetV1LocationsParams{})
	if err != nil {
		log.Fatalf("Failed to get locations: %v", err)
	}

	if locationsResp.JSON200 != nil {
		fmt.Printf("Found %d locations:\n", len(*locationsResp.JSON200))
		for _, location := range *locationsResp.JSON200 {
			fmt.Printf("  - %s (ID: %s)\n", *location.Name, *location.Id)
		}
	} else {
		fmt.Println("No locations found or unable to parse response")
	}

	// Example 3: Get all labels
	fmt.Println("\nFetching labels...")
	labelsResp, err := authClient.GetV1LabelsWithResponse(ctx)
	if err != nil {
		log.Fatalf("Failed to get labels: %v", err)
	}

	if labelsResp.JSON200 != nil {
		fmt.Printf("Found %d labels:\n", len(*labelsResp.JSON200))
		for _, label := range *labelsResp.JSON200 {
			fmt.Printf("  - %s (ID: %s)\n", *label.Name, *label.Id)
		}
	} else {
		fmt.Println("No labels found or unable to parse response")
	}

	fmt.Println("\nExample completed successfully!")
}
