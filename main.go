package main

import (
	"flag"
	"fmt"
	"time"

	_ "github.com/jiaqi-yin/go-file-processor/config"
	fileprocessor "github.com/jiaqi-yin/go-file-processor/file_processor"
	"github.com/jiaqi-yin/go-file-processor/storage"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "/tmp/file.log", "read file path")
	flag.Parse()

	storage := storage.NewRedisService()

	start := time.Now()

	fileprocessor.SequentialFileProcessor(path, storage)

	elapse := time.Now().Sub(start)
	fmt.Println("Process time: ", elapse)
	fmt.Println("Exit")
}
