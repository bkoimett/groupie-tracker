# 🎸 Groupie Tracker

A web application built with Go that displays information about bands and artists from an API.

## 🚀 Quick Start

### Prerequisites
- Go 1.16 or higher

### Run the Application

1. **Clone the repository**
```bash
git clone https://gitea.com/bkoimett/groupie-tracker.git
cd groupie-tracker
```

2. **Run the application**
```bash
go run cmd/main.go
```

3. **Open your browser**
```
http://localhost:8080
```

## 📁 Project Structure

```
groupie-tracker/
├── cmd/
│   ├── main.go              # Entry point
│   └── main_test.go
├── internal/
│   ├── handlers/             # HTTP handlers and API calls
│   ├── models/               # Data structures
│   ├── static/               # CSS and JavaScript files
│   │   ├── css/
│   │   └── js/
│   ├── templates/            # HTML templates
│   └── utils/                 # Error handling utilities
├── docs/                      # Documentation
├── go.mod
└── LICENSE
```

## ✨ Features

- Browse all artists in a grid layout
- View artist details (members, creation year, first album)
- See concert locations and dates
- Responsive design for mobile and desktop

## 🛠️ Built With

- Go standard library (`net/http`, `html/template`)
- HTML5, CSS3, JavaScript

## 👥 Team

- **Ivy Imoh** - API Integration
- **Favour Charles** - Server & Handlers  
- **Koimett Benjamin** - Frontend & UI

---

**Zone01 Kisumu - Module 1 Project**