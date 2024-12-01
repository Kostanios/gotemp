package sqlservice

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func New(endpoint string, user string, password string, dbName string, maxOpenCon int, maxIdleCons int, maxLifetime time.Duration) (error, *gorm.DB) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, endpoint, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err, nil
	}

	sqlDb, err := db.DB()

	if err != nil {
		return err, nil
	}

	sqlDb.SetMaxOpenConns(maxOpenCon)
	sqlDb.SetMaxIdleConns(maxIdleCons)
	sqlDb.SetConnMaxLifetime(maxLifetime)

	return nil, db
}