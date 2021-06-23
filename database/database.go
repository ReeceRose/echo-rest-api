package database

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

var (
	DBConn *gorm.DB
)
