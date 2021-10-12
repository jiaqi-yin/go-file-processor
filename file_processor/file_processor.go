package fileprocessor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jiaqi-yin/go-file-processor/domain"
	"github.com/jiaqi-yin/go-file-processor/storage"
)

type Reader interface {
	Read(chan string)
}

type Writer interface {
	Write(chan *domain.Item, chan bool)
}

type FileReader struct {
	Path string
}

type StorageWriter struct {
	Storage storage.Storage
}

type FileProcessor struct {
	ReadChan        chan string
	WriteChan       chan *domain.Item
	ExitProcessChan chan bool
	ExitWriteChan   chan bool
	Read            Reader
	Write           Writer
}

func (r *FileReader) Read(readChan chan string) {
	file, err := os.Open(r.Path)
	defer file.Close()

	if err != nil {
		panic(fmt.Sprintf("failed openning file: %s", err.Error()))
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		readChan <- scanner.Text()
	}

	close(readChan)
}

func (fp *FileProcessor) Process() {
	for v := range fp.ReadChan {
		data := strings.Split(v, ",")
		item := &domain.Item{
			UUID:      data[0],
			Firstname: data[1],
			Lastname:  data[2],
			Created:   time.Now(),
		}
		fp.WriteChan <- item
	}

	fp.ExitProcessChan <- true
}

func (w *StorageWriter) Write(writeChan chan *domain.Item, exitWriteChan chan bool) {
	for v := range writeChan {
		w.Storage.Save(v)
	}

	exitWriteChan <- true
}
