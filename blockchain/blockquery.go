package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	IERC721 "github.com/cocacolasante/googlecli/goerc721"
	IERC20 "github.com/cocacolasante/googlecli/goierc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BcRequest struct {
	Address string
	Chain string
	Contract string
	TokenId string
}

func NewBcRequest(address string, chain string, contract string, tokenId string) *BcRequest{
	return &BcRequest{
		Address: address,
		Chain: chain,
		Contract: contract,
		TokenId: tokenId,
	}
}

func(bc *BcRequest) GetEthBalance(){
	client := getClient(bc.Chain)
	if bc.Address == ""{
		log.Fatal("No address provided")
		return
	}
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
	contract := common.HexToAddress(bc.Contract)
	instance, err := IERC20.NewIERC20(contract, client)
	if err != nil {
	  log.Fatal(err)
	}
	
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
	  log.Fatal(err)
	}
	

	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("Account %s\n", account)
	fmt.Printf("ERC20 Contract %s\n", contract)
	fmt.Printf("Address balance: %d\n", balance)
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
}

func(bc *BcRequest) GetNFTBalanceOf(){
	client := getClient(bc.Chain)
	account := common.HexToAddress(bc.Address)
	contract := common.HexToAddress(bc.Contract)
	instance, err := IERC721.NewIERC721(contract, client)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("Account %s\n", account)
	fmt.Printf("ERC721 Contract %s\n", contract)
	fmt.Printf("Address balance Of: %d\n", balance)
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
}

func(bc *BcRequest) GetNFTOwnerOf(){
	client := getClient(bc.Chain)
	
	contract := common.HexToAddress(bc.Contract)
	instance, err := IERC721.NewIERC721(contract, client)
	if err != nil {
		log.Fatal(err)
		return
	}
	var token big.Int
	_, success := token.SetString(bc.TokenId, 10)
	if !success {
		log.Fatal("Error converting string to big.Int")
		return
	}
	ownerOf, err := instance.OwnerOf(&bind.CallOpts{}, &token)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("Token Id %s\n", &token)
	fmt.Printf("ERC721 Contract %s\n", contract)
	fmt.Printf("Address Owner Of: %s\n", ownerOf)
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
}




func getClient(chain string) *ethclient.Client {
	var client *ethclient.Client

	switch chain {
	case "polygon":
		fmt.Printf("Polygon\n")

		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("POLYGON_MAINNET_URL")
		conn, err := ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}
		
		client = conn
	case "arbitrum":
		fmt.Printf("Arbitrum\n")

		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("ABRITRUM_MAINNET_URL")
		conn, err := ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}
		
		client = conn
	case "base":
		fmt.Printf("Base\n")


		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("BASE_MAINNET_URL")
		conn, err := ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}
		
		client = conn
	default:

		// default case is eth mainnet
		
		fmt.Printf("Getting Blockchain Client \n")

		ethURL := os.Getenv("ETH_MAINNET_URL")
		conn, err := ethclient.Dial(ethURL)
		if err != nil {
			log.Fatal(err)

		}
		
		client = conn

	}

	return client

}