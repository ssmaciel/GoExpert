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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 10))
	tx := p.DB.WithContext(ctx)
	defer cancel()

	return tx.Create(moeda).Error
}