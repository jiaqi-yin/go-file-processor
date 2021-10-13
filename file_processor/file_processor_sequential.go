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

func SequentialFileProcessor(path string, storage storage.Storage) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(fmt.Sprintf("failed openning file: %s", err.Error()))
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		item := &domain.Item{
			UUID:      data[0],
			Firstname: data[1],
			Lastname:  data[2],
			Created:   time.Now(),
		}
		storage.Save(item)
	}
}
