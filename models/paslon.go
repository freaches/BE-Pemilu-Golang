package models

type Paslon struct {
	ID            uint
	Name          string
	VisionMission string `gorm:"column:visionMission"`
	Image         string
	Partai        []PartaiPaslon
}

type PaslonVote struct {
	ID   uint
	Name string
}

type PaslonPartai struct {
	ID   uint
	Name string
}

func (PaslonVote) TableName() string {
	return "paslons"
}

func (PaslonPartai) TableName() string {
	return "paslons"
}
