package memmap

import (
	"github.com/metafates/application-design-test/internal/entity"
	"log/slog"
)

func New(logger *slog.Logger) *Repo {
	return &Repo{
		logger: logger,
		byID:   make(map[entity.OrderID]entity.Order, 100),
	}
}
