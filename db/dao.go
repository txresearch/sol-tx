package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Dao struct {
	db *gorm.DB
}

type Config struct {
	User     string
	Password string
	Url      string
	Scheme   string
	Port     uint
	Debug    bool
}

func New(cfg *Config) *Dao {
	log := logger.Default
	if cfg.Debug {
		log = log.LogMode(logger.Info)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Url, cfg.User, cfg.Password, cfg.Scheme, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: log})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Block{}, &Transaction{}, &Trade{}, &Token{}, &UserOwner{}, &Pool{})
	if err != nil {
		panic(err)
	}
	s := &Dao{
		db: db,
	}
	return s
}

func (d *Dao) SaveBlock(b *Block) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "height"},
		},
		DoNothing: true,
	}).Save(b).Error
}

func (d *Dao) SaveTransaction(t *Transaction) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "hash"},
		},
		DoNothing: true,
	}).Save(t).Error
}

func (d *Dao) SaveTrade(t *Trade) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "tx_hash"},
			{Name: "tx_seq"},
		},
		DoNothing: true,
	}).Save(t).Error
}

func (d *Dao) SaveToken(t *Token) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "hash"},
		},
		DoNothing: true,
	}).Save(t).Error
}

func (d *Dao) SaveUserOwner(u *UserOwner) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user"},
		},
		DoNothing: true,
	}).Save(u).Error
}

func (d *Dao) SavePool(p *Pool) error {
	return d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "hash"},
		},
		DoNothing: true,
	}).Save(p).Error
}
