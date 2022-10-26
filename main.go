package main

import (
	"io"
	"log"
	"nhlapi/nhlapi"
	"os"
	"sync"
	"time"
)

func main() {
	// benchmarking time

	now := time.Now()

	// OPen, create, append if exists

	rosterFile, err := os.OpenFile("roster.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error Opening File roster.txt: %v", err)
	}

	defer rosterFile.Close()

	// Use multiple write points

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlapi.GetAllTeams()

	if err != nil {
		log.Fatalf("Error getting all teams %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(len(teams)) // wait group counter

	results := make(chan []nhlapi.Roster) // channel will hold slice of nhlApi Rosters (multiple Slices)

	for _, team := range teams {
		log.Println("------------------------------------")
		log.Printf("Name %s", team.Name)
		log.Println("------------------------------------")
		
		go func (team nhlapi.Team)  {
			roster, err := nhlapi.GetRoster(team.ID)
			if err != nil {
				log.Fatalf("error getting roster %v", err)
			}

			results <- roster
			wg.Done()
		}(team)

	}

	go func ()  {
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v", time.Since(now).String())
}

func display(results chan []nhlapi.Roster) {
	for r := range results {  // no index defined ?
		for _, ros := range r {
			log.Println("-------------------------------------")
			log.Printf("ID: %d\n", ros.Person.ID)
			log.Printf("Name: %s\n", ros.Person.FullName)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.JerseyNumber)
			log.Println("-------------------------------------")
		}
	}
}