package models

type Partai struct {
	ID            uint
	Name          string
	PartyLeader   string `gorm:"column:partyLeader"`
	VisionMission string `gorm:"column:visionMission"`
	Address       string
	Image         string
	PaslonID      uint `gorm:"column:paslonId"`
	Paslon        PaslonPartai
}

type PartaiPaslon struct {
	ID       uint
	Name     string
	PaslonID uint `gorm:"column:paslonId"`
}

func (PartaiPaslon) TableName() string {
	return "partais"
}
