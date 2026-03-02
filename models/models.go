package models

// Artist represents the structure from the /artists endpoint
type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    // We'll add these fields to store related data
    Locations    []string `json:"-"`
    ConcertDates []string `json:"-"`
    Relations    map[string][]string `json:"-"`
}

// LocationsAPI represents the response from /locations/{id}
type LocationsAPI struct {
    Index []struct {
        ID        int      `json:"id"`
        Locations []string `json:"locations"`
    } `json:"index"`
}

// DatesAPI represents the response from /dates/{id}
type DatesAPI struct {
    Index []struct {
        ID    int      `json:"id"`
        Dates []string `json:"dates"`
    } `json:"index"`
}

// RelationAPI represents the response from /relation/{id}
type RelationAPI struct {
    Index []struct {
        ID             int                 `json:"id"`
        DatesLocations map[string][]string `json:"datesLocations"`
    } `json:"index"`
}

// FullArtistData combines all information for display
type FullArtistData struct {
    Artist
    Locations    []string
    Dates        []string
    Relations    map[string][]string
}