# 🎸 Groupie Tracker

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)](LICENSE)
[![Testing](https://img.shields.io/badge/Test-Driven_Development-green?style=for-the-badge)](#-test-driven-development)
[![Status](https://img.shields.io/badge/status-in_development-yellow?style=for-the-badge)](#-3-day-sprint-plan)

A full-stack web application built with Go that visualizes artist data from a RESTful API. This project was developed using **Test-Driven Development (TDD)** methodology by a team of three developers in a 3-day sprint.

![Groupie Tracker Demo](https://via.placeholder.com/800x400?text=Groupie+Tracker+Screenshot)

## 📋 Table of Contents
- [Overview](#-overview)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [3-Day Sprint Plan](#-3-day-sprint-plan)
- [Test-Driven Development](#-test-driven-development)
- [Team Roles](#-team-roles)
- [Installation](#-installation)
- [Usage](#-usage)
- [Testing](#-testing)
- [Daily Checklist](#-daily-checklist)
- [API Reference](#-api-reference)
- [Contributing](#-contributing)
- [Team](#-team)
- [License](#-license)

## 🌟 Overview

Groupie Tracker consumes an API containing information about various bands and artists, processing and displaying the data through multiple visualizations. Users can explore:

- **Artist profiles** with images and biographies
- **Band members** and formation details
- **Upcoming and past concert locations**
- **Tour dates** and schedules
- **Relationships** between artists, locations, and dates

Built with clean architecture and 100% Test-Driven Development, this project demonstrates robust software engineering practices in a rapid 3-day development cycle.

## ✨ Features

### Core Features
- **Artist Gallery** - Browse all artists in a responsive card grid layout
- **Artist Details** - View comprehensive information about each artist:
  - Artist image and name
  - Creation year and first album date
  - Band members list
  - Past and upcoming concert locations
  - Concert dates
  - Tour schedule (dates mapped to locations)

### Technical Features
- **Test-Driven Development** - 100% test coverage approach
- **Concurrent API Calls** - Goroutines for parallel data fetching
- **Responsive Design** - Mobile-first CSS with Flexbox/Grid
- **Graceful Error Handling** - Custom error pages and fallbacks
- **Template Caching** - Optimized template rendering
- **RESTful API Integration** - Consumes multiple API endpoints

## 🛠️ Tech Stack

### Backend
- **Language:** Go (Golang)
- **Testing:** Go's built-in `testing` package
- **HTTP:** Standard `net/http` package
- **Templates:** `html/template` with custom functions
- **Concurrency:** Goroutines and channels

### Frontend
- **HTML5** - Semantic structure
- **CSS3** - Styling with Flexbox/Grid
- **Vanilla JavaScript** - DOM manipulation (minimal)

## 📁 Project Structure

```text
groupie-tracker/
│
├── cmd/
│   ├── main.go                    # Entry point - server initialization
│   └── main_test.go               # Integration tests
│
├── internal/
│   ├── handlers/
│   │   ├── handlers.go            # HTTP route handlers
│   │   ├── handlers_test.go       # Handler tests (TDD)
│   │   ├── api.go                 # API client functions
│   │   └── api_test.go            # API client tests (TDD)
│   │
│   ├── models/
│   │   ├── models.go              # Data structures
│   │   └── models_test.go         # Model tests (TDD)
│   │
│   ├── templates/
│   │   ├── index.html             # Homepage with artist grid
│   │   ├── artist.html            # Individual artist details
│   │   ├── error.html             # Error page template
│   │   └── template_test.go       # Template function tests (TDD)
│   │
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css          # Main stylesheet
│   │   └── js/
│   │       └── main.js            # Frontend interactions
│   │
│   └── testutils/
│       └── testutils.go           # Shared test helpers
│
├── .gitignore                      # Git ignore rules
├── go.mod                          # Go module file
├── go.sum                          # Go module checksums
├── LICENSE                         # MIT License
└── README.md                       # Project documentation

```

## 🗓️ 3-Day Sprint Plan

### **Day 1: Foundation & Core Models** 🚀

#### Morning Session (3 hours) - Project Setup
| Time | Activity | Team Member | Tests to Write |
|------|----------|-------------|----------------|
| 9:00 - 9:30 | Project initialization | All | - |
| 9:30 - 10:30 | Set up test utilities | All (Pair) | `testutils_test.go` |
| 10:30 - 12:00 | Create project structure | All | - |

#### Afternoon Session (4 hours) - TDD Implementation
| Time | Member 1 (API & Data) | Member 2 (Handlers) | Member 3 (Templates) |
|------|----------------------|---------------------|----------------------|
| 1:00 - 2:30 | **🔴 Write failing tests:** | **🔴 Write failing tests:** | **🔴 Write failing tests:** |
| | `models_test.go` - Artist unmarshaling | `handlers_test.go` - HomeHandler | `template_test.go` - FormatLocation |
| | `models_test.go` - Validation | `handlers_test.go` - ErrorHandler | `template_test.go` - FormatDate |
| 2:30 - 4:00 | **✅ Make tests pass:** | **✅ Make tests pass:** | **✅ Make tests pass:** |
| | Implement Artist struct | Implement basic handlers | Implement formatting functions |
| | Add JSON tags | Add error handling | Test with sample data |
| 4:00 - 5:00 | **🔄 Refactor & Integrate** | **🔄 Refactor & Integrate** | **🔄 Refactor & Integrate** |
| | Clean model code | Improve error messages | Optimize functions |
| | Add comments | Add request validation | Add edge cases |

#### Day 1 Deliverables
- ✅ Project structure complete
- ✅ All model tests passing
- ✅ Basic handler tests passing
- ✅ Template function tests passing
- ✅ Server runs without errors

### **Day 2: API Integration & Server Logic** 🔌

#### Morning Session (4 hours) - API Client
| Time | Member 1 (API & Data) | Member 2 (Handlers) | Member 3 (Templates) |
|------|----------------------|---------------------|----------------------|
| 9:00 - 10:30 | **🔴 Write failing tests:** | **🔴 Write failing tests:** | **🔴 Write failing tests:** |
| | `api_test.go` - FetchArtists | `handlers_test.go` - ArtistHandler | `template_test.go` - Template rendering |
| | `api_test.go` - Error cases | `handlers_test.go` - Invalid IDs | `template_test.go` - HTML structure |
| 10:30 - 12:00 | **✅ Make tests pass:** | **✅ Make tests pass:** | **✅ Make tests pass:** |
| | Implement API client | Implement ArtistHandler | Create index.html template |
| | Add error handling | Add ID validation | Add artist cards structure |

#### Afternoon Session (4 hours) - Concurrent Data Fetching
| Time | Member 1 (API & Data) | Member 2 (Handlers) | Member 3 (Templates) |
|------|----------------------|---------------------|----------------------|
| 1:00 - 2:30 | **🔴 Write failing tests:** | **🔴 Write failing tests:** | **🔴 Write failing tests:** |
| | `api_test.go` - Locations | `handlers_test.go` - Concurrent calls | `template_test.go` - Artist details |
| | `api_test.go` - Dates | `handlers_test.go` - Data aggregation | `css_test.go` - Responsive classes |
| 2:30 - 4:00 | **✅ Make tests pass:** | **✅ Make tests pass:** | **✅ Make tests pass:** |
| | Implement Locations fetch | Add goroutines for API calls | Create artist.html template |
| | Implement Dates fetch | Add channels for data | Add CSS styling |
| 4:00 - 5:00 | **Integration Testing** | **Integration Testing** | **Integration Testing** |
| | All run `main_test.go` together | Fix integration issues | Verify UI rendering |

#### Day 2 Deliverables
- ✅ API client fully functional
- ✅ Artist detail page working
- ✅ Concurrent data fetching implemented
- ✅ Templates render with real data
- ✅ CSS styling complete

### **Day 3: Polish & Client-Server Events** ✨

#### Morning Session (4 hours) - Search/Filter Feature
| Time | Member 1 (API & Data) | Member 2 (Handlers) | Member 3 (Templates) |
|------|----------------------|---------------------|----------------------|
| 9:00 - 10:30 | **🔴 Write failing tests:** | **🔴 Write failing tests:** | **🔴 Write failing tests:** |
| | Search data filtering | Search handler tests | Search UI tests |
| | Filter logic tests | Query parameter handling | Filter form rendering |
| 10:30 - 12:00 | **✅ Make tests pass:** | **✅ Make tests pass:** | **✅ Make tests pass:** |
| | Implement search logic | Add search endpoint | Add search form to UI |
| | Implement filter functions | Handle search requests | Add JavaScript for filtering |

#### Afternoon Session (4 hours) - Final Polish
| Time | Member 1 (API & Data) | Member 2 (Handlers) | Member 3 (Templates) |
|------|----------------------|---------------------|----------------------|
| 1:00 - 2:30 | **Edge Case Tests:** | **Performance Tests:** | **UI Polish Tests:** |
| | Empty API responses | Load testing | Mobile responsiveness |
| | Rate limiting | Connection timeouts | Browser compatibility |
| 2:30 - 4:00 | **Implement fixes** | **Implement fixes** | **Implement fixes** |
| | Add caching | Improve error pages | Add loading states |
| | Retry logic | Add request logging | Add animations |
| 4:00 - 5:00 | **Final Integration & Demo** | **Final Integration & Demo** | **Final Integration & Demo** |
| | Run full test suite | Deploy locally | Present working app |

#### Day 3 Deliverables
- ✅ Search/filter functionality working
- ✅ All tests passing (100% coverage)
- ✅ Responsive design complete
- ✅ Error handling graceful
- ✅ Final demo ready

## 🧪 Test-Driven Development

This project was built using **Test-Driven Development (TDD)** methodology:

### TDD Workflow
1. **🔴 RED** - Write a failing test
2. **✅ GREEN** - Write minimal code to pass the test
3. **🔄 REFACTOR** - Clean up code while keeping tests green

### Test Coverage
- **Unit Tests** - Individual components (models, handlers, API client)
- **Integration Tests** - End-to-end functionality
- **Template Tests** - HTML rendering and formatting functions
- **Edge Cases** - Error handling, invalid inputs, empty responses

### Running Tests
```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test -v ./handlers
go test -v ./models
go test -v ./templates
```

## 👥 Team Roles

### Team Member 1 (Ivy Imoh): Backend (API & Data)
- **Focus:** Data fetching, models, API integration
- **Tests:** API client tests, model validation
- **Daily Tasks:**
  - Day 1: Model structs and validation
  - Day 2: API client implementation
  - Day 3: Search/filter logic

### Team Member 2 (Favour Charles): Backend (Handlers & Server)
- **Focus:** HTTP handlers, routing, server logic
- **Tests:** Handler tests, error handling
- **Daily Tasks:**
  - Day 1: Basic handlers and routing
  - Day 2: Artist handler with concurrency
  - Day 3: Search endpoints and performance

### Team Member 3 (koimett Benjamin (lead)): Frontend (Templates & UI)
- **Focus:** HTML templates, CSS, template functions
- **Tests:** Template rendering, CSS responsiveness
- **Daily Tasks:**
  - Day 1: Template functions and base templates
  - Day 2: Artist details template and CSS
  - Day 3: Search UI and final polish

## 🚀 Installation

### Prerequisites
- Go 1.16 or higher
- Git

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/groupie-tracker.git
cd groupie-tracker
```

2. **Install dependencies**
```bash
go mod download
```

3. **Run the application**
```bash
go run main.go
```

4. **Open your browser**
```
http://localhost:8080
```

## 📖 Usage

### Starting the Server
```bash
# Default port 8080
go run main.go

# Or specify a custom port
PORT=3000 go run main.go
```

### Accessing the Application
1. Open your web browser to `http://localhost:8080`
2. Browse the artist gallery
3. Click on any artist card to view details
4. Explore concert locations, dates, and tour schedules
5. Use search/filter to find specific artists

### Example Routes
| Route | Description |
|-------|-------------|
| `/` | Homepage - Artist gallery |
| `/artist?id=1` | Details for artist with ID 1 |
| `/search?q=queen` | Search for artists (Day 3) |
| Any invalid route | Custom 404 error page |

## 🧪 Testing

### Run All Tests
```bash
go test -v ./...
```

### Run Tests with Coverage Report
```bash
go test -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Specific Package
```bash
# Test handlers package
go test -v ./handlers

# Test models package
go test -v ./models

# Test templates package
go test -v ./templates
```

### Integration Tests
```bash
go test -v -run Integration
```

## ✅ Daily Checklist

### Day 1 Checklist
- [ ] Project structure created
- [ ] Test utilities implemented
- [ ] Model tests written (RED)
- [ ] Model code implemented (GREEN)
- [ ] Handler tests written (RED)
- [ ] Basic handlers implemented (GREEN)
- [ ] Template function tests written (RED)
- [ ] Format functions implemented (GREEN)
- [ ] Server runs without errors

### Day 2 Checklist
- [ ] API client tests written (RED)
- [ ] API client implemented (GREEN)
- [ ] Artist handler tests for concurrency (RED)
- [ ] Concurrent fetching implemented (GREEN)
- [ ] Artist detail template tests (RED)
- [ ] Artist detail template created (GREEN)
- [ ] CSS tests written (RED)
- [ ] CSS styling implemented (GREEN)
- [ ] Integration tests passing

### Day 3 Checklist
- [ ] Search/filter tests written (RED)
- [ ] Search functionality implemented (GREEN)
- [ ] Edge case tests written
- [ ] Edge cases handled
- [ ] Performance tests passing
- [ ] Mobile responsive verified
- [ ] All tests passing (100% coverage)
- [ ] Final demo ready

## 🔌 API Reference

The application consumes the Groupie Trackers API:

**Base URL:** `https://groupietrackers.herokuapp.com/api`

| Endpoint | Description | Data Includes |
|----------|-------------|---------------|
| `/artists` | Artist information | Name, image, members, creation date, first album |
| `/locations/{id}` | Concert locations | Past and upcoming venues |
| `/dates/{id}` | Concert dates | Past and upcoming tour dates |
| `/relation/{id}` | Relationship data | Dates mapped to locations |

### Data Structure Example
```json
{
  "id": 1,
  "name": "Queen",
  "image": "https://example.com/queen.jpg",
  "members": ["Freddie Mercury", "Brian May"],
  "creationDate": 1970,
  "firstAlbum": "1973-07-13"
}
```

## 🤝 Contributing

This project was developed for educational purposes. If you'd like to contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests first (TDD approach)
4. Implement your changes
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## 👨‍💻 Team

### Development Team

| Name | Role | Responsibilities |
|------|------|------------------|
| [Team Member 1 - Ivy Imooh] | Backend API Developer | Data models, API integration, tests |
| [Team Member 2 - Favour Charles] | Backend Server Developer | HTTP handlers, routing, server logic |
| [Team Member 3 - Koimett Benjamin] | Frontend Developer | HTML templates, CSS, UI tests |

### Mentors
- Zone01 Kisumu Staff

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Zone01 Kisumu** - For the project curriculum and guidance
- **The Go Community** - For excellent documentation and best practices
- **Groupie Trackers API** - For providing the data

## 🎯 Learning Outcomes

Through this project, we learned:

- **Test-Driven Development** methodology in practice
- **Go concurrency** with goroutines and channels
- **HTTP server** implementation in Go
- **API integration** and JSON parsing
- **HTML templating** with custom functions
- **Responsive CSS** design principles
- **Team collaboration** with Git
- **Error handling** and graceful degradation
- **Rapid prototyping** in a 3-day sprint
- **Agile development** with daily standups

---

## 🏁 Quick Start for New Teams

```bash
# Day 1 Morning
git clone <repo>
cd groupie-tracker
go mod init groupie-tracker
mkdir handlers models templates static testutils
touch main.go main_test.go

# Assign roles and start Day 1 tasks
echo "Let's go! 🚀"
```

**Made with ❤️ by the Groupie Tracker Team**  
*Zone01 Kisumu - Module 1 Project*