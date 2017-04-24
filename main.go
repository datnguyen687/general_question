package main

import (
	"os"
	"fmt"
	"general_web_api_query/UserConfig"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Printf("Usage: %s [user_id]\nExample: %s 1\n", args[0], args[0])
		os.Exit(0)
	}

	user := new(UserConfig.User)
	for arg := range args {
		user.PopulateData(strconv.Itoa(arg))
	}
	user.PrintDetail()
}
