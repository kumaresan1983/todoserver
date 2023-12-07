package initializers

import (
	"fmt"
	"os"

	"github.com/kumaresan1983/todoserver/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const projectDirName = "todoserver" // change to relevant project name

// ConnectDB initializes the database connection
func ConnectDB() error {

	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// currentWorkDirectory, _ := os.Getwd()
	// rootPath := projectName.Find([]byte(currentWorkDirectory))
	if os.Getenv("GIN_MODE") == "test" {
		// Use an in-memory SQLite database for testing
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			logrus.WithError(err).Error("Failed to connect to the Database")
			return fmt.Errorf("failed to connect to the Database: %v", err)
		}

		logrus.Info("Running Migrations...")
		err = DB.AutoMigrate(&models.Users{}, &models.ToDo{})
		if err != nil {
			logrus.WithError(err).Error("Failed to run migrations")
			return fmt.Errorf("failed to run migrations: %v", err)
		}

		logrus.Info("Connected Successfully to the Database 🚀")
		return nil

	} else {
		// Load database configuration from the environment file
		config, err := LoadConfig(".")

		// Construct DSN from the loaded configuration
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DatabaseUserName, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.WithError(err).Error("Failed to connect to the Database")
			return fmt.Errorf("failed to connect to the Database: %v", err)
		}

		logrus.Info("Running Migrations...")
		err = DB.AutoMigrate(&models.Users{}, &models.ToDo{})
		if err != nil {
			logrus.WithError(err).Error("Failed to run migrations")
			return fmt.Errorf("failed to run migrations: %v", err)
		}

		logrus.Info("Connected Successfully to the Database 🚀")
		return nil

	}

	// logrus.Info("Running Migrations...")
	// err = DB.AutoMigrate(&models.Users{}, &models.ToDo{})
	// if err != nil {
	// 	logrus.WithError(err).Error("Failed to run migrations")
	// 	return fmt.Errorf("failed to run migrations: %v", err)
	// }

	// logrus.Info("Connected Successfully to the Database 🚀")
	// return nil
}
