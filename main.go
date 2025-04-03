package main

import (
	"embed"
	"emp-management/backend"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize database
	db, err := backend.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Create EmployeeService instance
	empService := backend.NewEmployeeService(db)

	// Create an instance of the App struct with DB
	app := NewApp(db)

	// Run Wails app
	err = wails.Run(&options.App{
		Title:  "Employee",
		Width:  1024,
		Height: 768,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,               // Bind App struct
			empService,    // ✅ Bind EmployeeService struct (fix function binding)
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
