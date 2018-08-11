package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/OEP/ifttt-tools/internal/common"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ifttt-trigger [OPTIONS] EVENT [VALUE1 [VALUE2 [VALUE3]]]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	verbose := flag.Bool("v", false, "Switch on verbose mode")
	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 4 {
		flag.Usage()
		os.Exit(1)
	}

	if !*verbose {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	cfg := common.NewDefaultConfig()
	client := common.NewIFTTTClient(cfg)

	event := flag.Arg(0)
	args := flag.Args()[1:]

	err := client.Trigger(event, args...)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
