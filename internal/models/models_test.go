package models

import (
	"encoding/json"
	"testing"
)

// Test Artist parsing

func TestArtistParsesFromJSON(t *testing.T) {
	raw := `{
        "id": 1,
        "name": "Queen",
        "image": "https://example.com/queen.jpg",
        "members": ["Freddie Mercury", "Brian May"],
        "creationDate": 1970,
        "firstAlbum": "1973-07-14"
    }`

	var artist Artist
	if err := json.Unmarshal([]byte(raw), &artist); err != nil {
		t.Fatalf("failed to parse artist: %v", err)
	}

	if artist.ID != 1 {
		t.Errorf("expected ID 1, got %d", artist.ID)
	}
	if artist.Name != "Queen" {
		t.Errorf("expected Name Queen, got %s", artist.Name)
	}
	if len(artist.Members) != 2 {
		t.Errorf("expected 2 members, got %d", len(artist.Members))
	}
	if artist.Members[0] != "Freddie Mercury" {
		t.Errorf("expected first member Freddie Mercury, got %s", artist.Members[0])
	}
}

// Test empty members

func TestArtistEmptyMembers(t *testing.T) {
	raw := `{
        "id": 2,
        "name": "Solo Act",
        "image": "",
        "members": [],
        "creationDate": 2000,
        "firstAlbum": "2001-01-01"
    }`

	var artist Artist
	if err := json.Unmarshal([]byte(raw), &artist); err != nil {
		t.Fatalf("failed to parse artist: %v", err)
	}

	if len(artist.Members) != 0 {
		t.Errorf("expected 0 members, got %d", len(artist.Members))
	}
}

// Test malformed JSON

func TestArtistMalformedJSON(t *testing.T) {
	raw := `{"id": "not-a-number"}`
	var artist Artist
	if err := json.Unmarshal([]byte(raw), &artist); err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

// Test Location parsing

func TestLocationEntryParses(t *testing.T) {
	raw := `{
        "index": [
            {"id": 1, "locations": ["New_York", "Los_Angeles"]},
            {"id": 2, "locations": ["London"]}
        ]
    }`

	var locAPI LocationAPI
	if err := json.Unmarshal([]byte(raw), &locAPI); err != nil {
		t.Fatalf("failed to parse LocationAPI: %v", err)
	}

	if len(locAPI.Index) != 2 {
		t.Errorf("expected 2 location entries, got %d", len(locAPI.Index))
	}
	if locAPI.Index[0].Locations[0] != "New_York" {
		t.Errorf("expected first location New_York, got %s", locAPI.Index[0].Locations[0])
	}
}

// Test Dates parsing

func TestDateEntryParses(t *testing.T) {
	raw := `{
        "index": [
            {"id": 1, "dates": ["2026-03-01", "2026-03-05"]},
            {"id": 2, "dates": ["2026-04-01"]}
        ]
    }`

	var datesAPI DatesAPI
	if err := json.Unmarshal([]byte(raw), &datesAPI); err != nil {
		t.Fatalf("failed to parse DatesAPI: %v", err)
	}

	if len(datesAPI.Index) != 2 {
		t.Errorf("expected 2 date entries, got %d", len(datesAPI.Index))
	}
	if datesAPI.Index[0].Dates[1] != "2026-03-05" {
		t.Errorf("expected second date 2026-03-05, got %s", datesAPI.Index[0].Dates[1])
	}
}

// Test Relations parsing

func TestRelationEntryParses(t *testing.T) {
	raw := `{
        "index": [
            {
                "id": 1,
                "datesLocations": {
                    "2026-03-01": ["New_York"],
                    "2026-03-05": ["Los_Angeles", "Chicago"]
                }
            }
        ]
    }`

	var relAPI RelationAPI
	if err := json.Unmarshal([]byte(raw), &relAPI); err != nil {
		t.Fatalf("failed to parse RelationAPI: %v", err)
	}

	if len(relAPI.Index) != 1 {
		t.Errorf("expected 1 relation entry, got %d", len(relAPI.Index))
	}

	if len(relAPI.Index[0].DatesLocations["2026-03-05"]) != 2 {
		t.Errorf("expected 2 locations on 2026-03-05, got %d", len(relAPI.Index[0].DatesLocations["2026-03-05"]))
	}
}

// Test FullArtistData composition

func TestFullArtistDataComposition(t *testing.T) {
	artist := Artist{
		ID:           1,
		Name:         "Queen",
		Members:      []string{"Freddie Mercury", "Brian May"},
		CreationDate: 1970,
		FirstAlbum:   "1973-07-14",
	}

	full := FullArtistData{
		Artist:       artist,
		Locations:    []string{"New_York", "Los_Angeles"},
		ConcertDates: []string{"2026-03-01", "2026-03-05"},
		Relations: map[string][]string{
			"2026-03-01": {"New_York"},
			"2026-03-05": {"Los_Angeles", "Chicago"},
		},
	}

	if full.Artist.ID != 1 || full.Locations[1] != "Los_Angeles" || len(full.Relations["2026-03-05"]) != 2 {
		t.Error("FullArtistData composition did not match expected values")
	}
}
