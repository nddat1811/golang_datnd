package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nddat1811/simple_project_golang/entity"
	//migrate "github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type sQLConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Dbname    string `json:"dbname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

// func MigrationDB(db *gorm.DB) error {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return err
// 	}

// 	migrations := &migrate.FileMigrationSource{
// 		Dir: MigrationPath,
// 	}

// 	n, err := migrate.Exec(sqlDB, DriverName, migrations, migrate.Up)

// 	if err != nil {
// 		return err
// 	}

// 	n, err = migrate.Exec(sqlDB, DriverName, migrations, migrate.Up)

// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Applied migraions! ", n)

// 	return nil
// }

// func loadDBConfigs(filepath string) *sQLConfig {
// 	configFile, err := os.Open(filepath)

// 	defer configFile.Close()

// 	if err != nil {
// 		return nil
// 	}

// 	jsonParser := json.NewDecoder(configFile)
// 	config := sQLConfig{}
// 	jsonParser.Decode(&config)

// 	return &config
// }
//SetupDatabaseConnection is creating a new connection to DB
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3308)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})
	fmt.Println("Connect success")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
