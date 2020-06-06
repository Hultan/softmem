package main

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

type database struct {
	db *gorm.DB
}

func (d *database) GetAllNumbers() ([]NumberTable, error) {
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

func (d *database) GetNumber(number int) (*NumberTable, error) {
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

func (d *database) getDatabase() (*gorm.DB, error) {
	if d.db == nil {
	db, err := gorm.Open("mysql", "per:KnaskimGjwQ6M!@tcp(192.168.1.3:3306)/softmemo?parseTime=True")
		if err != nil {
			return nil, err
		}
		d.db = db
	}
	return d.db, nil
}

func (d *database) CloseDatabase() {
	if d.db == nil {
		return
	}
	d.db.Close()
	d.db = nil
	return
}
