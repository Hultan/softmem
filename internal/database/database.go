package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hultan/softteam/framework"
)

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

func (d *Database) GetAllNumbers() (map[string]NumberTable, error) {
	db, err := d.getDatabase()
	if err != nil {
		return nil, err
	}
	var numbers []NumberTable
	if result := db.Order("number asc").Find(&numbers); result.Error != nil {
		return nil, result.Error
	}

	var result = make(map[string]NumberTable)

	for _, v := range numbers {
		result[v.Number] = v
	}

	return result, nil
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
		fw := framework.NewFramework()
		server, err := fw.IO.ReadAllText("config/.credentials")
		if err != nil {
			panic(err)
		}
		// db, err := gorm.Open(databaseName, fmt.Sprintf(connectionString, server))
		db, err := gorm.Open(mysql.Open(fmt.Sprintf(connectionString, server)), &gorm.Config{})
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
	sqlDb, _ := d.db.DB()
	_ = sqlDb.Close()
	d.db = nil
	return
}
