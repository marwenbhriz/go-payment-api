package models

type Company struct {
	Id                 int64  `gorm:"primaryKey" json:"id"`
	CorporateName      string `gorm:"type:varchar(300)" json:"corporate_name"`
	RepresentativeName string `gorm:"type:varchar(300)" json:"representative_name"`
	TelephoneNumber    string `gorm:"type:varchar(300)" json:"telephone_number"`
	PostCode           string `gorm:"type:varchar(300)" json:"post_code"`
	Address            string `gorm:"type:varchar(300)" json:"address"`
	User               User
}

type User struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	FullName     string `gorm:"type:varchar(300)" json:"full_name"`
	EmailAddress string `gorm:"type:varchar(300)" json:"email_address"`
	Password     string `gorm:"type:varchar(300)" json:"password"`
}

type Invoice struct {
	Id int64 `gorm:"primaryKey" json:"id"`
}
