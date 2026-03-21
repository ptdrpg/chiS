package handler

import (
	"os"
	"os/exec"
	"path/filepath"
)

func AddAllDependancies(directory string, db string) {
	var ressources = []string{"gorm.io/gorm", "github.com/go-chi/chi/v5", "github.com/joho/godotenv", filepath.Join("gorm.io/driver/", db), "github.com/spf13/cobra", "github.com/go-chi/cors", "github.com/golang-jwt/jwt", "github.com/google/uuid", "golang.org/x/crypto/bcrypt", "github.com/golang-jwt/jwt/v5"}
	for _, dep := range ressources {
		installDep := exec.Command("go", "get", dep)
		installDep.Dir = directory
		installDep.Stderr = os.Stderr
		err := installDep.Run()
		if err != nil {
			ErrorHandler(err)
			return
		}
	}
}
