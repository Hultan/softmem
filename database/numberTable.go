package database

type NumberTable struct {
	Number     string `gorm:"column:number;primary_key"`
	Word       string `gorm:"column:word;size:20"`
	Correct    int    `gorm:"column:correct;"`
	HasChanged bool   `gorm:"-"` // ignore this field
}

func (n *NumberTable) TableName() string {
	return "numbers"
}
