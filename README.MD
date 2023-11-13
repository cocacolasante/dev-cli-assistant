# Developer Cli Assistant

Thank you for checking out this repo!

This is a command line interface designed to help speed up developers workflow

You are able to google search, ask questions to chatgpt and query the blockchain and smart contracts

# Command list

	-------------------------------------
	             GOOGLE SEARCH COMMANDS
	-------------------------------------
	
	-search    query used for api search
	-site      filter results from this site
	-sort      sort results in ascending or descending order using a or d
	-exclude   exclude any results that has this phrase or word
	-help      pull up help menu
	-------------------------------------
	
	             AI COMMANDS
	-------------------------------------
	
	-ai        triggers the ai call -requires content flag next
	-content   question to ask the ai
	-------------------------------------
	
	             BLOCKCHAIN COMMANDS
	-------------------------------------
	
	-blockchain        triggers the blockchain call -required flag
	-address           target address to query for
	-chain             target chain to query on
	-contract          target contract to query 
	-erc20             flag to query ERC20 tokens
	-erc721            flag to query ERC721 tokens
	-token             flag for nft token id
	-------------------------------------

# Example
    Calling AI -stream response
        godev -ai -content "how to spell cake in italian?" -stream
    
    Google Search
        godev -search "how to write an essay" 
    
    Blockchain Query
        godev -blockchain -chain polygon -address 0x7F6C7eb6e7Ae8491212bd3E11498F2A6c5c9979B -erc20 -contract 0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619 

# How To Use
    Step 1: update the .init.example.go file with the correct api keys (will compile and add env variables at build time)
    Step 2: run command - go build ./cmd/godev
    Step 3: move binary into bin folder - sudo mv godev /usr/local/bin/    
    Step 4: access commands using godev in command line
    Step 5: call godev -search "star wars" should return correct results
