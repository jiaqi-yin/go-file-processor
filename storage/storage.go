package storage

import (
	"github.com/jiaqi-yin/go-file-processor/domain"
)

type Storage interface {
	Save(item *domain.Item)
}
