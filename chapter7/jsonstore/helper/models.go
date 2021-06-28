package helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type_JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates tables with plural names.
// Use this to supress it
func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitBD() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://gituser:passw0rd@localhost/mydb2?sslmode=disable")

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Shipment{}, &Package{})
	return db, nil
}
