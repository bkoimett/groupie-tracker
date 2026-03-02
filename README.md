# 🎸 Groupie Tracker

A full-stack web application that visualizes artist data from a RESTful API, displaying bands, their concert locations, and tour dates in an interactive and user-friendly interface.

## 📋 Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Implementation Plan](#implementation-plan)
- [Learning Outcomes](#learning-outcomes)
- [Team](#team)

## 🌟 Overview

Groupie Tracker is a web application that consumes a given API containing information about various bands and artists. The application processes and displays this data through multiple visualizations, making it easy for users to explore:

- Artist profiles with images and biographies
- Band members and formation details
- Upcoming and past concert locations
- Tour dates and schedules
- Relationships between artists, locations, and dates

## ✨ Features

### Core Features
- **Artist Gallery** - Browse all artists in a responsive card grid layout
- **Artist Details** - View comprehensive information about each artist:
  - Artist image and name
  - Creation year and first album date
  - Band members
  - Past and upcoming concert locations
  - Concert dates

### Client-Server Events (Choose One)
- **Search Functionality** - Search artists by name, members, or creation date
- **Advanced Filters** - Filter artists by:
  - Number of members
  - Creation date range
  - First album date range
  - Concert locations
- **Interactive Map** - Visualize concert locations on a map (bonus)
- **Date Timeline** - View tour dates on an interactive timeline

### Technical Features
- Clean, responsive design that works on desktop and mobile
- Server-side rendering with Go templates
- Graceful error handling (404, 500 error pages)
- No server crashes - robust error management
- RESTful API integration

## 🛠️ Tech Stack

### Backend
- **Language:** Go (Golang)
- **Packages:** Only standard library packages:
  - `net/http` - HTTP server and client
  - `html/template` - HTML templating
  - `encoding/json` - JSON parsing
  - `strings`, `strconv` - String manipulation
  - `sync` - Concurrency control

### Frontend
- **HTML5** - Semantic structure
- **CSS3** - Styling with Flexbox/Grid
- **Vanilla JavaScript** - DOM manipulation and event handling
- **No external CSS/JS frameworks allowed**

## 📁 Project Structure

```
groupie-tracker/
│
├── main.go                    # Entry point - server initialization
│
├── handlers/
│   ├── handlers.go            # Route handlers (home, artist, etc.)
│   └── api.go                 # API fetching logic
│
├── models/
│   └── models.go              # Data structures for API responses
│
├── templates/
│   ├── index.html             # Homepage with artist grid
│   ├── artist.html            # Individual artist details page
│   └── error.html             # Error page template
│
├── static/
│   ├── css/
│   │   └── style.css          # Main stylesheet
│   ├── js/
│   │   └── main.js            # Frontend interactions
│   └── images/                # Static images
│
├── go.mod                      # Go module file
├── go.sum                      # Go module checksums
└── README.md                   # Project documentation
```

## 🚀 Installation

### Prerequisites
- Go 1.16 or higher installed
- Git (optional)

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/bkoimett/groupie-tracker.git
cd groupie-tracker
```

2. **Initialize Go module** (if not already done)
```bash
go mod init groupie-tracker
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
1. Open your web browser
2. Navigate to `http://localhost:8080`
3. Browse the artist gallery
4. Click on any artist card to view details
5. Use search/filter features to find specific artists

### Example Interactions
- **View all artists** - Homepage displays all artists in a grid
- **Artist details** - Click any artist card for complete information
- **Search** - Type in search bar to filter artists
- **Filter** - Use dropdowns to filter by members, dates, etc.

## 🔌 API Reference

The application consumes the following API endpoints:

| Endpoint | Description | Data Includes |
|----------|-------------|---------------|
| `/artists` | Artist information | Name, image, year formed, first album, members |
| `/locations` | Concert locations | Past and upcoming venue locations |
| `/dates` | Concert dates | Past and upcoming tour dates |
| `/relation` | Relationship data | Links artists to dates and locations |

**Base URL:** `https://groupietrackers.herokuapp.com/api`

### Data Structure Example
```json
{
  "artists": [{
    "id": 1,
    "name": "Queen",
    "image": "https://example.com/queen.jpg",
    "members": ["Freddie Mercury", "Brian May"],
    "creationDate": 1970,
    "firstAlbum": "1973-07-13"
  }]
}
```

## 📅 Implementation Plan

### Phase 1: Foundation (Week 1)
- [ ] Set up Go module and project structure
- [ ] Create basic HTTP server
- [ ] Implement route handlers
- [ ] Serve static files (CSS, JS)
- [ ] Create base HTML templates

### Phase 2: API Integration (Week 2)
- [ ] Define Go structs for API data
- [ ] Fetch data from `/artists` endpoint
- [ ] Parse JSON responses
- [ ] Display artist cards on homepage
- [ ] Implement individual artist pages

### Phase 3: Enhanced Features (Week 3)
- [ ] Fetch and display concert locations
- [ ] Show tour dates on artist pages
- [ ] Style the application with CSS
- [ ] Make it responsive (mobile-friendly)

### Phase 4: Client-Server Events (Week 4)
- [ ] Implement search functionality
- [ ] Add filter options
- [ ] Create client-side interactions
- [ ] Add error handling pages

### Phase 5: Polish & Testing (Week 5)
- [ ] Write unit tests
- [ ] Optimize performance
- [ ] Add loading states
- [ ] Cross-browser testing
- [ ] Code cleanup and documentation

## 🎯 Learning Outcomes

By completing this project, you will learn:

### Backend Skills
- Building HTTP servers in Go
- Making API requests and parsing JSON
- Structuring a Go application
- Template rendering with `html/template`
- Error handling and graceful degradation
- Concurrency with goroutines

### Frontend Skills
- Semantic HTML5
- Responsive CSS with Flexbox/Grid
- DOM manipulation with vanilla JavaScript
- Event handling and user interactions
- Form processing and validation

### Software Engineering
- Project structure and organization
- Client-server architecture
- RESTful API consumption
- Data modeling and manipulation
- Testing and debugging

## 👥 Team

This project is developed by Mr. Koimett himselfu :

| Role | Responsibilities |
|------|------------------|
| **Backend Developer** | API integration, data processing, server logic |
| **Frontend Developer** | HTML/CSS, templates, user interface |
| **Full-Stack Developer** | Client-server events, testing, integration |

## 📝 License

This project is part of the curriculum at Zone01 Kisumu and is developed for educational purposes.

## 🙏 Acknowledgments

- Zone01 Kisumu for the project curriculum
- The Go community for excellent documentation
- Groupie Trackers API providers

---

**Ready to start building?** Check our [implementation guide](#-implementation-plan) and begin with Phase 1!