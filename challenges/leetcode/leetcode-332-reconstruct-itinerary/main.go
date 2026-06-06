package main

import (
	"slices"
	"strings"
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
	if airports, ok := j.adj[airport]; !ok || len(airports) == 0 {
		*itinerary = append(*itinerary, airport)
		return
	}
	currentAirport := j.adj[airport][0]
	j.adj[airport] = j.adj[airport][1:]
	j.buildItinerary(currentAirport, itinerary)
}

func findItinerary(tickets [][]string) []string {
	journey := InitializeJourney()
	for _, ticket := range tickets {
		journey.AddDestination(ticket[0], ticket[1])
	}
	for k := range journey.adj {
		slices.SortFunc(journey.adj[k], func(i, j string) int { return -strings.Compare(i, j) })
	}
	itinerary := []string{}
	journey.buildItinerary("JFK", &itinerary)
	for i, j := 0, len(itinerary)-1; i < j; i, j = i+1, j-1 {
		itinerary[i], itinerary[j] = itinerary[j], itinerary[i]
	}
	return itinerary
}
