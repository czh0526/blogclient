package main

import (
	"blog/x/blog/types"
	"context"
	"fmt"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"log"
)

func main() {
	cosmosClient, err := cosmosclient.New(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	accountName := "alice"
	address, err := cosmosClient.Address(accountName)
	if err != nil {
		log.Fatal(err)
	}

	msg := &types.MsgCreatePost{
		Creator: address.String(),
		Title:   "Hello!",
		Body:    "This is the first post",
	}

	txResp, err := cosmosClient.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MsgCreatePost: \n\n")
	fmt.Println(txResp)

	queryClient := types.NewQueryClient(cosmosClient.Context())
	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\nAll Posts:\n\n")
	fmt.Println(queryResp)
}
