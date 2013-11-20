package main

import (
	"errors"
	"fmt"
	"github.com/neilmiddleton/reddit"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal(errors.New("Subreddit must be provided"))
	}

	for _, subreddit := range os.Args[1:len(os.Args)] {
		items, err := reddit.Get(subreddit)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n", strings.ToUpper(subreddit))
		fmt.Println("=============")
		for _, item := range items {
			fmt.Println(item)
		}
	}
}
