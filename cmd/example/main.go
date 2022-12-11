package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lichess "github.com/mikeb26/lichess-go-sdk/openapi"
)

// Visit https://lichess.org/account/oauth/token to generate one
const LichessApiToken = "<REPLACE_WITH_YOUR_API_TOKEN>"

var NullUser = lichess.User{}

func main() {
	config := lichess.NewConfiguration()
	client := lichess.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), lichess.ContextAccessToken,
		LichessApiToken)

	userMap := make(map[string]lichess.User)
	err := getUsers(userMap)
	if err != nil {
		log.Fatalf("Failed to get users: %v\n", err)
	}

	err = getUserInfos(ctx, client, userMap)
	if err != nil {
		log.Fatalf("Failed to get user data: %v\n", err)
	}

	printUserInfos(userMap)
}

func getUsers(userMap map[string]lichess.User) error {
	userMap["VincentKeymer2004"] = NullUser
	userMap["RebeccaHarris"] = NullUser // Naroditsky
	userMap["AnishGiri"] = NullUser
	userMap["alireza2003"] = NullUser

	return nil
}

func getUserInfos(ctx context.Context, client *lichess.APIClient,
	userMap map[string]lichess.User) error {

	// https://lichess.org/api#tag/Users/operation/apiUsers
	const MaxUsersPerRequest = 300

	userList := ""
	count := 0
	var err error
	for key, _ := range userMap {
		if userList == "" {
			userList = key
		} else {
			userList = userList + "," + key
		}
		count++

		if count == MaxUsersPerRequest {
			err = getOneUserInfosBatch(ctx, client, userList, userMap)
			if err != nil {
				return err
			}
			userList = ""
			count = 0
		}
	}

	// last batch
	if userList != "" {
		err = getOneUserInfosBatch(ctx, client, userList, userMap)
	}

	return err
}

func getOneUserInfosBatch(ctx context.Context, client *lichess.APIClient,
	userList string, userMap map[string]lichess.User) error {

	var usersReply []lichess.User
	var httpResp *http.Response
	var err error

	for {
		usersReply, httpResp, err =
			client.UsersApi.ApiUsers(ctx).Body(lichess.PtrString(userList)).Execute()
		if err != nil {
			return fmt.Errorf("Client failure: %w", err)
		}
		if httpResp.StatusCode == 429 {
			// https://lichess.org/page/api-tips says wait a minute
			fmt.Fprintf(os.Stderr, "429 recv; sleeping 1min...")
			time.Sleep(1 * time.Minute)
			continue
		} else if httpResp.StatusCode != 200 {
			return fmt.Errorf("HTTP failure: %v", httpResp)
		} else {
			break
		}
	}

	for _, u := range usersReply {
		username := u.GetUsername().(string)
		_, ok := userMap[username]
		if !ok {
			return fmt.Errorf("returned username %v was not already in map", u.GetUsername())
		}

		userMap[username] = u
	}

	return nil
}

func printUserInfos(userMap map[string]lichess.User) {
	for key, u := range userMap {
		if u == NullUser {
			log.Fatalf("missing user entry for %v", key)
		}
		fmt.Printf("-------------------------------------\n")
		fmt.Printf("id:%v\n", u.GetId())
		fmt.Printf("  username:%v\n", u.GetUsername())
		if u.HasPerfs() {
			fmt.Printf("  perfs:\n")
			perfs := u.GetPerfs()
			if perfs.Blitz != nil {
				fmt.Printf("    blitz:\n")
				fmt.Printf("      rating:%v\n", perfs.Blitz.GetRating())
				fmt.Printf("      games:%v\n", perfs.Blitz.GetGames())
			}
			if perfs.Rapid != nil {
				fmt.Printf("    rapid:\n")
				fmt.Printf("      rating:%v\n", perfs.Rapid.GetRating())
				fmt.Printf("      games:%v\n", perfs.Rapid.GetGames())
			}
			if perfs.Classical != nil {
				fmt.Printf("    classical:\n")
				fmt.Printf("      rating:%v\n", perfs.Classical.GetRating())
				fmt.Printf("      games:%v\n", perfs.Classical.GetGames())
			}
		}
		fmt.Printf("  createdAt:%v\n", lichessTime2GoTime(u.GetCreatedAt()))
		fmt.Printf("  disabled:%v\n", u.GetDisabled())
		fmt.Printf("  tosViolation:%v\n", u.GetTosViolation())
		fmt.Printf("  profile:\n")
		prof := u.GetProfile()
		fmt.Printf("    firstname:%v\n", prof.GetFirstName())
		fmt.Printf("    lastname:%v\n", prof.GetLastName())
		fmt.Printf("    country:%v\n", prof.GetCountry())
		fmt.Printf("    location:%v\n", prof.GetLocation())
		fmt.Printf("    links:%v\n", prof.GetLinks())
		fmt.Printf("    bio:%v\n", prof.GetBio())
		fmt.Printf("  seenAt:%v\n", lichessTime2GoTime(u.GetSeenAt()))
		fmt.Printf("  verified:%v\n", u.GetVerified())
		fmt.Printf("  playTime:%v\n", lichessDuration2String(u.GetPlayTime()))
		fmt.Printf("  title:%v\n", u.GetTitle())
		fmt.Printf("  patron:%v\n", u.GetPatron())
	}
}

func lichessTime2GoTime(timeIn interface{}) time.Time {
	timeValFloatMsecs := timeIn.(float64)
	timeValIntNsecs := int64(timeValFloatMsecs * 1000000)
	return time.Unix(0, timeValIntNsecs)
}

func lichessDuration2String(durIn interface{}) string {
	durValPlayTime := durIn.(lichess.PlayTime)
	durValFloatSecs := durValPlayTime.GetTotal().(float64)
	durValIntSecs := int64(durValFloatSecs)
	durVal := time.Duration(durValIntSecs) * time.Second

	seconds := uint64(durVal.Seconds()) % 60
	minutes := uint64(durVal.Minutes()) % 60
	hours := uint64(durVal.Hours()) % 24
	days := uint64(durVal.Hours()) / 24

	ret := fmt.Sprintf("%v seconds", seconds)
	if minutes > 0 {
		ret = fmt.Sprintf("%v minutes, %v", minutes, ret)
	}
	if hours > 0 {
		ret = fmt.Sprintf("%v hours, %v", hours, ret)
	}
	if days > 0 {
		ret = fmt.Sprintf("%v days, %v", days, ret)
	}

	return ret
}
