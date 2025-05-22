package main

import (
	"Bus-Backend/internal/Logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	Logging.InitLogger()

	dsn := "Database/bus_database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		Logging.Logs.Fatalf("Failed to connect to database: %v", err)
	}

	// Attach DB hook to global logger
	Logging.Logs.AddHook(Logging.NewGormHook(db))

	// Tesing Logs
	Logging.Logs.Info("Database connection established successfully")
	Logging.Logs.Warn("Logging initialized successfully")
	Logging.Logs.Error("GormHook attached to logger successfully")
	Logging.Logs.Debug("Debug level Logging enabled")
	Logging.Logs.Trace("Trace level Logging enabled")
	Logging.Logs.Fatal("Fatal level Logging enabled")

}
