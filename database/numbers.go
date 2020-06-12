package database

type NumberTable struct {
	Number string `gorm:"column:number;primary_key"`
	Word   string `gorm:"column:word;size:20"`
}

func (n *NumberTable) TableName() string {
	return "numbers"
}
