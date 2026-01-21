package main

import (
	"fmt"
	"sort"
)

type Journey struct {
	adj map[string][]string
}

func InitializeJourney() *Journey {
	return &Journey{
		adj: make(map[string][]string),
	}
}

func (j *Journey) AddDestination(from string, to string) {
	j.adj[from] = append(j.adj[from], to)
}

func (j *Journey) buildItinerary(airport string, itinerary *[]string) {
	airports := j.adj[airport]
	for len(j.adj[airport]) > 0 {
		currentAirport := airports[0]
		j.adj[airport] = j.adj[airport][1:]
		j.buildItinerary(currentAirport, itinerary)
	}
	*itinerary = append(*itinerary, airport)

}

func findItinerary(tickets [][]string) []string {
	journey := InitializeJourney()
	for _, ticket := range tickets {
		journey.AddDestination(ticket[0], ticket[1])
	}
	for k := range journey.adj {
		sort.Strings(journey.adj[k])
	}
	itinerary := []string{}
	journey.buildItinerary("JFK", &itinerary)
	for i, j := 0, len(itinerary)-1; i < j; i, j = i+1, j-1 {
		itinerary[i], itinerary[j] = itinerary[j], itinerary[i]
	}
	return itinerary
}

func main() {
	tickets := [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	}
	fmt.Printf("%v", findItinerary(tickets))
}
