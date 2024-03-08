package models

type Vote struct {
	ID       uint
	UserID   uint `gorm:"column:userId"`
	User     UserPaslon
	PaslonID uint `gorm:"column:paslonId"`
	Paslon   PaslonVote
}
