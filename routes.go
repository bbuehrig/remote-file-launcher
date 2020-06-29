package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// RoutesStruct defines all configured routes
type RoutesStruct struct {
	Routes []RouteStruct `json:"routes"`
}

// RouteStruct defines one single route. The given path will be routed to one of the given URLs for opening the given file.
type RouteStruct struct {
	Name     string `json:"name"`
	FilePath string `json:"filepath"`
	URLPath  string `json:"urlpath"`
}

func getRouteConfig(file string) (RouteStruct, error) {
	for _, route := range routes.Routes {
		if route.URLPath == file {
			return route, nil
		}
	}
	return RouteStruct{}, errors.New("Route for " + file + " not found!")
}

// parseJSONConfig parses the json-routes config file and builds the routes for echo.
func parseJSONConfig() error {
	// parse config file
	data, err := ioutil.ReadFile("./routes.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &routes)
	return nil
}
