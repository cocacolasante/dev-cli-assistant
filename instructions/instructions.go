package instructions

import "fmt"

func PrintHelpMenu() {
	fmt.Printf("HELP MENU\n")
	fmt.Printf("-search    query used for api search\n")
	fmt.Printf("-site      filter results from this site\n")
	fmt.Printf("-sort      sort results in ascending or descending order using a or d\n")
	fmt.Printf("-help      pull up help menu\n")
}