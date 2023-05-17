package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/p1ck0/multiset/log"
	"github.com/p1ck0/multiset/multireq"
	"github.com/p1ck0/multiset/parser"
)

func main() {
	var file string
	var async bool
	flag.StringVar(&file, "f", "", "file to read")
	flag.BoolVar(&async, "a", false, "async mode")
	flag.Parse()

	logger := log.New()
	if file == "" || file[len(file)-5:] != ".json" {
		logger.Error("no file specified")
		return
	}

	content, err := os.ReadFile(file)
	if err != nil {
		logger.Error("could not read file")
		return
	}

	logger.Info("START PARSE")
	requests, err := parser.Parse(content)
	if err != nil {
		logger.Error("could not parse file")
		return
	}
	logger.Info("PARSE DONE")

	ctx, stop := signal.NotifyContext(log.LoggerWithContext(context.Background(), logger), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if async {
		logger.Info("START SEND ASYNC MULTI REQ")
		multireq.SendMultiReqAsync(ctx, requests)
	} else {
		logger.Info("START SEND MULTI REQ")
		multireq.SendMultiReq(ctx, requests)
	}
}
