# About
This is a simple study case project demonstrating how to create a desktop application using the Wails framework. The project integrates Vue.js for the frontend and Golang for the backend, using SQLite3 as the database. Wails supports HTML, CSS, JavaScript, and modern frontend frameworks.

# Technologies Used
* Vue.js – Frontend framework
* Golang – Backend language
* SQLite3 – Database
* Wails – Go-based framework for building desktop applications

# Installation & Setup
## Prerequisites

Ensure you have the following **installed**:
- Go (latest version)
- Node.js & npm (for Vue.js)
- Wails CLI (go install github.com/wailsapp/wails/v2/cmd/- wails@latest)
***
# Running the Project
## Live Development Mode
To run the project in live development mode with hot reload:
```sh 
wails dev 
```

## Building for Production
To build a redistributable production package:
```sh
wails build
```

This will generate an optimized binary suitable for distribution.

# Project Structure
```sh
├── frontend/   # Vue.js frontend
├── backend/    # Golang backend
├── wails.json  # Wails configuration
├── main.go     # Entry point for the application
└── README.md   # Project documentation
```
**License**
This project is open-source and available under the *MIT License.*