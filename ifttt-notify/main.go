package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/OEP/ifttt-tools/internal/common"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ifttt-notify [OPTIONS] command...")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cfg := common.NewDefaultConfig()
	client := common.NewIFTTTClient(cfg)

	commandName := flag.Args()[0]
	commandArgs := flag.Args()[1:]

	cmd := exec.Command(commandName, commandArgs...)

	if err := cmd.Start(); err != nil {
		logTrigger(client, "ifttt-notify", commandName, "failure")
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		logTrigger(client, "ifttt-notify", commandName, "failure")
		log.Fatal(err)
	}

	logTrigger(client, "ifttt-notify", commandName, "success")
}

func logTrigger(client common.IFTTTClient, event string, values ...string) {
	flags := log.Flags()

	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	err := client.Trigger(event, values...)

	log.SetFlags(flags)
	log.SetOutput(ioutil.Discard)

	if err != nil {
		log.Println(err)
	}
}
