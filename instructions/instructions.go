package instructions

import "fmt"

func PrintHelpMenu() {
	fmt.Println(" ")
	fmt.Printf("             HELP MENU\n")
	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("             GOOGLE SEARCH COMMANDS\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-search    query used for api search\n")
	fmt.Printf("-site      filter results from this site\n")
	fmt.Printf("-sort      sort results in ascending or descending order using a or d\n")
	fmt.Printf("-exclude   exclude any results that has this phrase or word\n")
	fmt.Printf("-help      pull up help menu\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Printf("             AI COMMANDS\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-ai        triggers the ai call -requires content flag next\n")
	fmt.Printf("-content   question to ask the ai\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Printf("             BLOCKCHAIN COMMANDS\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-blockchain        triggers the blockchain call -required flag\n")
	fmt.Printf("-address           target address to query for\n")
	fmt.Printf("-chain             target chain to query on\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	
}

func PrintDescription(){
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Printf("             GoogleCli\n")
	fmt.Println(" ")
	fmt.Println("-------------------------------------")
	fmt.Printf("This is a command line interface for quick good searches in the terminal\n")
	fmt.Printf("Created with the purpose of freeing up time of having to load up a browser and then search\n")
	fmt.Printf("Results are clickable links to speed up your workflow\n")
	fmt.Printf("Questions can be asked directly to the AI through the command line\n")
	fmt.Println("-------------------------------------")
	fmt.Println(" ")
	fmt.Println("Thanks for using!")
	fmt.Println(" ")
}