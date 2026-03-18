package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationAPI struct {
	Index []LocationEntry `json:"index"`
}

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

type FullArtistData struct {
	Artist
	Locations    []string            `json:"locations"` // Add tags if needed
	ConcertDates []string            `json:"dates"`
	Relations    map[string][]string `json:"relations"`
}
