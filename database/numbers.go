package database

type NumberTable struct {
	Number string `gorm:"column:number;primary_key"`
	Word   string `gorm:"column:word;size:20"`
}

type PEG struct {
	Number string
	Word   string
}

func (n *NumberTable) TableName() string {
	return "numbers"
}
