package db

import (
	"fmt"
	"money-manager/core/models"
	"money-manager/internal/constant"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEnv struct {
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
	dsn := getDNS(getDBEnv())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		print(dsn)
		panic("failed to connect to the database")
	}

	fmt.Println("Database connection established!")
	return db
}

func getDNS(config DBEnv) string {
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

func getDBEnv() DBEnv {
	dbConfig := DBEnv{
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

func AutoMigrations(db *gorm.DB) {
	// Migrate the schema
	// TODO: use migrations
	db.AutoMigrate(&models.Management{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.MoneyAccount{})
	db.AutoMigrate(&models.Money{})
	db.AutoMigrate(&models.Location{})
}

func DropTables(db *gorm.DB) {
	db.Migrator().DropTable(&models.Management{})
	db.Migrator().DropTable(&models.Account{})
	db.Migrator().DropTable(&models.MoneyAccount{})
	db.Migrator().DropTable(&models.Money{})
	db.Migrator().DropTable(&models.Location{})
}
