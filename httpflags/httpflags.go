package httpflags

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


type HTTPReq struct{
	Url string
	ApiKey string 
	Method string


}

func NewHtpReq(url string, apikey string, method string) *HTTPReq{
	return &HTTPReq{
		Url: url,
		ApiKey: apikey,
		Method: method,
	}
}

func(hq *HTTPReq) MakeRequest(){
	request, err :=  http.NewRequest(hq.Method, hq.Url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Response Body:", string(body))

}