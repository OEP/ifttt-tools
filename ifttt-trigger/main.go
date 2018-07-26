package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OEP/ifttt-tools/internal/common"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ifttt-trigger EVENT [VALUE1 [VALUE2 [VALUE3]]]")
	}
	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 4 {
		flag.Usage()
		os.Exit(1)
	}

	cfg := common.NewDefaultConfig()
	client := common.NewIFTTTClient(cfg)

	event := flag.Arg(0)
	args := flag.Args()[1:]

	err := client.TriggerSlice(event, args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
