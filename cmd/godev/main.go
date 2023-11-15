package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/cocacolasante/googlecli/blockchain"
	"github.com/cocacolasante/googlecli/httpflags"
	"github.com/cocacolasante/googlecli/instructions"
	"github.com/cocacolasante/googlecli/openaiapi"
	searchQueries "github.com/cocacolasante/googlecli/search"
)

const VERSION = "1.0.1"

func main() {

	Init()
	//  GOOGLE API FLAGS
	var phrase string
	var excludeterm string
	var helpFlag bool 
	var descriptionFlag bool
	var versionFlag bool

	flag.StringVar(&phrase, "search", "", "the search string")
	targetSite := flag.String("site", "", "Search results for a specific site")
	sort := flag.String("sort", "", "Sort result by date")
	flag.StringVar(&excludeterm, "exclude", "", "terms to exclude")
	flag.BoolVar(&helpFlag, "help", false, "print usage instructions")
	flag.BoolVar(&descriptionFlag, "description", false, "prints cli description")
	flag.BoolVar(&versionFlag, "v", false, "prints cli version")

	// OPEN AI FLAGS
	var usingAi bool
	var content string
	var stream bool
	flag.BoolVar(&usingAi, "ai", false, "Use this call to get ai input")
	flag.StringVar(&content, "content", "", "content to ask ai")
	flag.BoolVar(&stream,"stream", false, "change ai response type to stream answer")

	// blockchain flags
	var callingBc bool
	var isERC20 bool
	var isERC721 bool
	var address string
	var chain string
	var contract string
	var nftTokenId string
	flag.BoolVar(&callingBc, "blockchain", false, "flag to call the blockchain")
	flag.BoolVar(&isERC20, "erc20", false, "flag to call erc20 contract")
	flag.BoolVar(&isERC721, "erc721", false, "flag to call erc20 contract")
	flag.StringVar(&address, "address", "", "target address")
	flag.StringVar(&chain, "chain", "", "target chain")
	flag.StringVar(&contract, "contract", "", "target contract")
	flag.StringVar(&nftTokenId, "token", "", "nft token number to check")
	
	// http flags
	var isHttp bool
	var urlString string
	var methodType string
	flag.BoolVar(&isHttp, "http", false, "flag to make http request")
	flag.StringVar(&urlString, "url", "", "target url")
	flag.StringVar(&methodType, "method", "GET", "method type")
	





	flag.Parse()

	if helpFlag {
		instructions.PrintHelpMenu()
		return
	}

	if descriptionFlag {
		instructions.PrintDescription()
		return
	}
	if versionFlag {
		fmt.Printf("Current Version: %s\n", VERSION)
		return
	}

	if usingAi {
		
		aiReq := openaiapi.NewAiRequest(content)
		if aiReq.Content == ""{
			log.Fatal("empty ai request")
			return 
		}
		if stream {
			aiReq.NewStreamCall()
		}
		
		aiReq.ApiCall()
		return
	}

	if callingBc {
		bcReq := blockchain.NewBcRequest(address, chain, contract, nftTokenId)
		if isERC20 {
			bcReq.GetTokenBalanceOfAddress()
			return
		} else if isERC721 {
			if nftTokenId != ""{
				bcReq.GetNFTOwnerOf()
				return
			}
			bcReq.GetNFTBalanceOf()
			return
		}else {
			
			bcReq.GetEthBalance()
			return 

		}
	}

	if isHttp {		
		newReq := httpflags.NewHtpReq(urlString, "", strings.ToUpper(methodType))
		newReq.MakeRequest()
		return
	}
	
	if phrase == ""{
		log.Fatal("no search phrase")
		return
	}
	search := url.QueryEscape(phrase)
	exclude := url.QueryEscape(excludeterm)
	
	fmt.Printf("Search string is: %s\n", string(search))



	queryStruct := searchQueries.NewQuery(search, *targetSite, *sort, exclude)
	searchQuery := queryStruct.NewURL()
	fmt.Println(searchQuery)

	searchQueries.GetGoogleResponse(searchQuery)

	

	
}	

