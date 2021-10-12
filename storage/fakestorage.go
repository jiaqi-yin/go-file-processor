package storage

import (
	"time"

	"github.com/jiaqi-yin/go-file-processor/domain"
)

type fakeStorage struct{}

func (fs *fakeStorage) Save(item *domain.Item) {
	time.Sleep(1 * time.Millisecond)
}

func NewFakeStorage() Storage {
	return &fakeStorage{}
}
