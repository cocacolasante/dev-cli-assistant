package main

import "os"


func InitExample() {
	os.Setenv("SEARCH_API_KEY", "")
	os.Setenv("OPEN_AI_KEY", "")
	os.Setenv("ETH_MAINNET_URL","")
	os.Setenv("POLYGON_MAINNET_URL", "")
	os.Setenv("ABRITRUM_MAINNET_URL", "")
	os.Setenv("BASE_MAINNET_URL", "")
}
