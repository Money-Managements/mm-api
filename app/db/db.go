package db

import (
	"fmt"
	"money-manager/app/db/model"
	"money-manager/internal/constant"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	DBName    string
	Charset   string
	ParseTime bool
	Loc       string
}

func ConnectDB() *gorm.DB {
	dsn := getDNS(getDBConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Database connection established!")

	// Migrate the schema TODO: use migrations
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Location{})
	db.AutoMigrate(&model.Money{})
	db.AutoMigrate(&model.MoneyAccount{})
	db.AutoMigrate(&model.MoneyTransaction{})

	return db
}

func getDNS(config DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Charset,
		config.ParseTime,
		config.Loc,
	)

	return dsn
}

func getDBConfig() DBConfig {
	dbConfig := DBConfig{
		Username:  os.Getenv(constant.DBUser),
		Password:  os.Getenv(constant.DBPassword),
		Host:      os.Getenv(constant.DBHost),
		Port:      os.Getenv(constant.DBPort),
		DBName:    os.Getenv(constant.DBName),
		Charset:   "utf8mb4",
		ParseTime: true,
		Loc:       "UTC",
	}

	return dbConfig
}
