package unit1

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

type ClassifiedAd struct {
	gorm.Model
	Title string
	Body  string
	Price int
}

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	db, _ := gorm.Open(sqlite.Open("file::memory:"))
	//db, _ := gorm.Open(sqlite.Open("ads.sqlite"))
	_ = db.AutoMigrate(&ClassifiedAd{})
	return &Database{DB: db}
}

func (d *Database) Create(c ClassifiedAd) (id uint, err error) {
	tx := d.DB.Create(&c)
	return c.ID, tx.Error
}

func (d *Database) Get(id uint) (*ClassifiedAd, error) {
	ad := &ClassifiedAd{}
	tx := d.DB.Find(ad, id)
	return ad, tx.Error
}

func (d *Database) Update(id uint, c ClassifiedAd) error {
	ad := ClassifiedAd{Model: gorm.Model{ID: id}, Title: c.Title, Body: c.Body, Price: c.Price}
	tx := d.DB.Save(&ad)
	return tx.Error
}

func (d *Database) List() (ads []ClassifiedAd, err error) {
	tx := d.DB.Find(&ads)
	return ads, tx.Error
}
