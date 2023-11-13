package blockchain

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type BcRequest struct {
	Address string
	Chain string
	Contract string
}

func NewBcRequest(address string, chain string, contract string) *BcRequest{
	return &BcRequest{
		Address: address,
		Chain: chain,
		Contract: contract,
	}
}

func(bc *BcRequest) GetEthBalance(){
	client := getClient(bc.Chain)
	account := common.HexToAddress(bc.Address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("Balance: %s\n", balance) // 25893180161173005034
	fmt.Printf("Address: %s\n",bc.Address) // 25893180161173005034
	fmt.Printf("Chain: %s\n", bc.Chain ) // 25893180161173005034
	fmt.Println("-------------------------------------")
	fmt.Println(" ")

}

func(bc *BcRequest) GetTokenBalanceOfAddress() {
	client := getClient(bc.Chain)
	account := common.HexToAddress(bc.Address)


	fmt.Printf("client %s\n", client)
	fmt.Printf("Account %s\n", account)
	 
}



func getClient(chain string) *ethclient.Client {
	var client *ethclient.Client

	switch chain {
	case "polygon":
		fmt.Printf("Polygon\n")

		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("POLYGON_MAINNET_URL")
		client, err = ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}
	case "arbitrum":
		fmt.Printf("Arbitrum\n")

		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("ABRITRUM_MAINNET_URL")
		client, err = ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}

	default:
		// default case is eth mainnet
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("ETH_MAINNET_URL")
		client, err = ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}

	}

	return client

}