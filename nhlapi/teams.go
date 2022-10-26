package nhlapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		TimeZone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue,omitempty"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division        struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		NameShort    string `json:"nameShort"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId"`
		TeamName    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	ShortName       string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}

// Note api response returns 
// {
// 	"copyright" : "NHL and the NHL Shield are registered 
// 	"teams" : [ 
// 	{
// 	  "id" : 1,
// 	  ...
//  to pick out only the teams struct

type nhlTTeamsResponse struct {
	Teams []Team `json:"teams"`
}

// Returns Slice of Team or an error
func GetAllTeams() ([]Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams", baseUrl))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// close after error/ api call

	var response nhlTTeamsResponse
	err = json.NewDecoder(res.Body).Decode(&response) // decode to response
	if err != nil {
		return nil, err
	}

	return response.Teams, nil
}