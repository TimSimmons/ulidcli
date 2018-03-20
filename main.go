package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oklog/ulid"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("must specify exactly one argument, ex. ulid <id>")
		os.Exit(1)
	}

	id := os.Args[1]

	u, err := ulid.Parse(id)
	if err != nil {
		fmt.Printf("error parsing ULID: %s", err)
		os.Exit(1)
	}

	t := time.Unix(0, int64(u.Time())*int64(time.Millisecond))
	fmt.Println(t.String())
}
