package economy_of_tests

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

type ClassifiedAd struct {
	Title string
	Body  string
	Price int
}

type ClassifiedAdRow struct {
	gorm.Model
	ClassifiedAd
}

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&ClassifiedAdRow{})
	if err != nil {
		panic(err)
	}

	return &Database{DB: db}
}

func (d *Database) Create(c ClassifiedAd) (id uint, err error) {
	row := ClassifiedAdRow{ClassifiedAd: c}
	tx := d.DB.Create(&row)
	return row.ID, tx.Error
}

func (d *Database) Get(id uint) (c *ClassifiedAd, err error) {
	row := &ClassifiedAdRow{}
	tx := d.DB.Find(row, id)
	return &row.ClassifiedAd, tx.Error
}

func (d *Database) Update(id uint, c ClassifiedAd) error {
	row := ClassifiedAdRow{Model: gorm.Model{ID: id}, ClassifiedAd: c}
	tx := d.DB.Save(&row)
	return tx.Error
}

func (d *Database) List() (ads []ClassifiedAd, err error) {
	var rows []ClassifiedAdRow
	tx := d.DB.Find(&rows)
	for _, row := range rows {
		ads = append(ads, row.ClassifiedAd)
	}
	return ads, tx.Error
}
