package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

func GetLinksSearch(SearchTerm string) []LinkType {
	
	Results := make([]LinkType, 10)

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var links LinksType

	json.Unmarshal(byteValue, &links)

	for i := 0; i < len(links.Links); i++ {
		fmt.Println("Link URL: " + links.Links[i].URL)
		fmt.Println("Link Title: " + links.Links[i].Title)
		fmt.Println("Link Tag: " + links.Links[i].Tag)
		fmt.Println("Link CreateTime: " + links.Links[i].CreateTime.String())

		// Search Algorithim
	}

	return Results
}

func GetLinksTag(SearchTerm string) []LinkType {
	
	var Results []LinkType

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var links LinksType

	json.Unmarshal(byteValue, &links)

	for i := 0; i < len(links.Links); i++ {
		// fmt.Println("Link URL: " + links.Links[i].URL)
		// fmt.Println("Link Title: " + links.Links[i].Title)
		// fmt.Println("Link Tag: " + links.Links[i].Tag)
		// fmt.Println("Link CreateTime: " + links.Links[i].CreateTime.String())

		if links.Links[i].Tag == SearchTerm {
			Results = append(Results, links.Links[i])
		}
	}

	for j := 0; j < len(Results); j++ {
		fmt.Println("Link URL: " + Results[j].URL)
	}

	return Results
}

func StoreLink(link LinkType) {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var links LinksType

	json.Unmarshal(byteValue, &links)
}