package main

import (
	"flag"

	"http_tool/internal/client"
)

func main() {
	parallel := flag.Int("parallel", 10, "an int")

	flag.Parse()
	if *parallel < 0 {
		panic("parallel value must be greater than zero")
	}

	opts := client.Options{
		Parallel: *parallel,
		URLS:     flag.Args(),
	}
	httpCustomClient := client.New(opts)

	err := httpCustomClient.Request()
	if err != nil {
		panic(err)
	}
}
