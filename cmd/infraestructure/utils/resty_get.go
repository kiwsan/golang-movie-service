package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func RestyGet() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://reqres.in/api/users?page=2")

	// Explore response object
	fmt.Println("-----------------------------------------------------")
	fmt.Println("https://reqres.in/api/users?page=2")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
