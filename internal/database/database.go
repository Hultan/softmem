package database

import (
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

type Database struct {
	db *gorm.DB
}

func (d *Database) UpdateStatistics(number NumberTable) error {
	db, err := d.getDatabase()
	if err != nil {
		return err
	}
	if result := db.Save(number); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database) GetAllNumbers() ([]NumberTable, error) {
	db, err := d.getDatabase()
	if err != nil {
		return nil, err
	}
	var numbers []NumberTable
	if result := db.Order("number asc").Find(&numbers); result.Error != nil {
		return nil, result.Error
	}

	return numbers, nil
}

func (d *Database) GetNumber(number int) (*NumberTable, error) {
	db, err := d.getDatabase()
	if err != nil {
		return nil, err
	}
	result := NumberTable{}
	if result := db.Where("number=?", number).First(&result); result.Error != nil {
		return nil, err
	}

	return &result, nil
}

func (d *Database) getDatabase() (*gorm.DB, error) {
	if d.db == nil {
	db, err := gorm.Open(databaseName, connectionString)
		if err != nil {
			return nil, err
		}
		d.db = db
	}
	return d.db, nil
}

func (d *Database) CloseDatabase() {
	if d.db == nil {
		return
	}
	d.db.Close()
	d.db = nil
	return
}
