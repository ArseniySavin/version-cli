package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"version-cli/internal/cli"
)

func main() {

	if err := cli.New("ArseniySavin",
		fmt.Sprintf("1.0.0-Alpha%s+%d",
			"2",
			time.Now().UnixNano())).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
