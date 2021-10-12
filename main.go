package main

import (
	"flag"
	"fmt"
	"time"

	_ "github.com/jiaqi-yin/go-file-processor/config"
	"github.com/jiaqi-yin/go-file-processor/domain"
	fileprocessor "github.com/jiaqi-yin/go-file-processor/file_processor"
	"github.com/jiaqi-yin/go-file-processor/storage"
)

func main() {
	var numRoutine int = 100 //runtime.NumCPU()
	fmt.Println("Go routines: ", numRoutine)

	var path string
	flag.StringVar(&path, "path", "/tmp/file.log", "read file path")
	flag.Parse()

	reader := &fileprocessor.FileReader{
		Path: path,
	}
	writer := &fileprocessor.StorageWriter{
		// Storage: storage.NewDynamodbService(),
		// Storage: storage.NewFakeStorage(),
		Storage: storage.NewRedisService(),
	}

	fp := &fileprocessor.FileProcessor{
		ReadChan:        make(chan string, 1000),
		WriteChan:       make(chan *domain.Item, 1000),
		ExitProcessChan: make(chan bool, numRoutine),
		ExitWriteChan:   make(chan bool, numRoutine),
		Read:            reader,
		Write:           writer,
	}

	start := time.Now()

	go fp.Read.Read(fp.ReadChan)

	for i := 0; i < numRoutine; i++ {
		go fp.Process()
	}

	go func(exitProcessChan chan bool, writeChan chan *domain.Item) {
		for i := 0; i < numRoutine; i++ {
			<-exitProcessChan
		}
		close(writeChan)
	}(fp.ExitProcessChan, fp.WriteChan)

	for i := 0; i < numRoutine; i++ {
		go fp.Write.Write(fp.WriteChan, fp.ExitWriteChan)
	}

	for i := 0; i < numRoutine; i++ {
		<-fp.ExitWriteChan
	}

	elapse := time.Now().Sub(start)
	fmt.Println("Process time: ", elapse)
	fmt.Println("Exit")
}
