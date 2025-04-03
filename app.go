package main

import (
	"context"
	"database/sql"
	"emp-management/backend"
	// "log"
)

// App struct
type App struct {
	ctx         context.Context
	EmpService  *backend.EmployeeService
}

// NewApp creates a new App instance
func NewApp(db *sql.DB) *App {
	return &App{
		EmpService: backend.NewEmployeeService(db), // Initialize EmployeeService
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
