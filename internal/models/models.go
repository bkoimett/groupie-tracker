package models

// we are using the json:"-" becoz those fields are fetched from separate endpoints
// they were never part of /api/artists
type Artist struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Image        string              `json:"image"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"-"`
	ConcertDates []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
}

// Using named structs is used for better consistency across the models
// Locations can be used elsewhere

// The wrapper
type LocationAPI struct {
	Index []LocationEntry `json:"index"`
}

// The individual location entry
type LocationEntry struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesAPI struct {
	Index []DateEntry `json:"index"`
}
type DateEntry struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationAPI struct {
	Index []RelationEntry `json:"index"`
}
type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// FullArtsisData combines all information for display
// merge struct

type FullArtistData struct {
	Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}
