package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ptdrpg/chiS/handler"
)

func ConfigDB(directory string, db string) {
	sql := fmt.Sprintf(`
package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/%s"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connexion() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%%s user=%%s password=%%s dbname=%%s port=%%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, errors := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errors != nil {
		log.Fatalf("failed to connect to the database: %%v", errors)
	}

	db.AutoMigrate()

	DB = db
}
	`, db)

	sqlLite := fmt.Sprintf(`
package app

import (
	"log"
	"%s/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func COnnexion() {
	db, errors := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if errors != nil {
		log.Fatalf("failed to connect to the database: %%v", errors)
	}

	db.AutoMigrate()
	DB = db
}

	`, directory)

	appPath := filepath.Join(directory, "app", "db.go")
	app, err := os.Create(appPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	if db != "sqlite" {
		variableEnv := `DB_HOST=localhost
	DB_USER=dbusername
	DB_PASSWORD=dbpassword
	DB_NAME=dbname
	DB_PORT=5432
	SECRET_KEY=secretkey`
		envPath := filepath.Join(directory, ".env")
		env, err := os.Create(envPath)
		if err != nil {
			handler.ErrorHandler(err)
		}
		env.Write([]byte(variableEnv))
		env.Close()
		app.Write([]byte(sql))
	} else {
		app.Write([]byte(sqlLite))
	}
	app.Close()
}
