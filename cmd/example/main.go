package main

import (
	"context"
	"fmt"
	"log"

	lichess "github.com/mikeb26/lichess-go-sdk/openapi"
)

// Visit https://lichess.org/account/oauth/token to generate one
const LichessApiToken = "<REPLACE_WITH_YOUR_API_TOKEN>"

func main() {
	config := lichess.NewConfiguration()
	client := lichess.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), lichess.ContextAccessToken,
		LichessApiToken)

	acctReply, httpResp, err := client.AccountApi.AccountMe(ctx).Execute()
	if err != nil {
		log.Fatalf("Client failure: %v", err)
	}
	if httpResp.StatusCode != 200 {
		log.Fatalf("HTTP failure: %v", httpResp)
	}

	fmt.Printf("Account Info:\n")

	// unfortunately the auto-generated code is missing output types
	// and instead seems to return interface{} for everything; parse by hand
	// for now
	myMap, ok := acctReply.(map[string]interface{})
	if !ok {
		fmt.Printf("Account Info:\n  resp:%v\n  type:%T\n", acctReply,
			acctReply)
		return
	}

	for key, val := range myMap {
		switch key {
		case "perfs":
			perfMap, ok := val.(map[string]interface{})
			if !ok {
				// just skip for now
				continue
			}
			fmt.Printf("  perfs:\n")
			for pkey, pval := range perfMap {
				switch pkey {
				case "classical":
					fallthrough
				case "rapid":
					fallthrough
				case "blitz":
					fallthrough
				case "bullet":
					pvalMap, ok := pval.(map[string]interface{})
					if !ok {
						// just skip for now
						continue
					}
					fmt.Printf("    %v\n", pkey)
					for pvalKey, pvalVal := range pvalMap {
						fmt.Printf("      %v:%v\n", pvalKey, pvalVal)
					}

				default:
					// just skip for now
				}
			}
		case "playTime":
			fallthrough
		case "seenAt":
			fallthrough
		case "createdAt":
			fallthrough
		case "count":
			fallthrough
		case "profile":
			// just skip for now
		default:
			fmt.Printf("  %v:%v\n", key, val)
		}
	}
}
