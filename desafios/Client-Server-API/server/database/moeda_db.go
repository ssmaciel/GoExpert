package database

import (
	"context"
	"time"

	"github.com/ssmaciel/GoExpert/desafios/Client-Server-API/server/internal/entity"

	"gorm.io/gorm"
)

type Moeda struct {
	DB *gorm.DB
}

func NewMoeda(db *gorm.DB) *Moeda {
	return &Moeda{DB: db}
}

func (p *Moeda) Create(moeda *entity.Moeda) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10))
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Nanosecond*1))
	defer cancel()
	tx := p.DB.WithContext(ctx)
	err := tx.Create(moeda).Error
	if err != nil {
		ctx.Done()
		return err
	}
	return nil
}
