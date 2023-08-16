package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

const (
	DatabasePath = "internal/repository/film.sqlite"
)

type IGetWhere interface {
	GetWhere(item interface{}, query string, args ...string)
}

type IMigrate interface {
	Migrate()
}

type ICreate interface {
	Create(item interface{}, conditions ...interface{})
}

type IDelete interface {
	Delete()
}

type IGetAll interface {
	GetAll(item interface{})
}
type Database struct {
	rwLock *sync.RWMutex
	db     *gorm.DB
}

func NewDatabase() *Database {
	database, err := gorm.Open(sqlite.Open(DatabasePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Database{
		db:     database,
		rwLock: &sync.RWMutex{},
	}
}

func (db Database) Create(item interface{}) {
	db.rwLock.Lock()
	defer db.rwLock.Unlock()
	db.db.Create(item)
}

func (db Database) Delete(item interface{}, conditions ...interface{}) {
	db.rwLock.Lock()
	defer db.rwLock.Unlock()
	db.db.Delete(item, conditions)
}

func (db Database) GetWhere(item interface{}, query string, args ...string) {
	db.rwLock.RLock()
	defer db.rwLock.RUnlock()
	db.db.Where(query, args).Find(item)
}

func (db Database) GetAll(item interface{}) {
	db.rwLock.RLock()
	defer db.rwLock.RUnlock()
	db.db.Find(item)

}
