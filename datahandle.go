package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func GetLinksSearch(SearchTerm string) []LinkType {

	var Results []LinkType

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
		// fmt.Println("Link Title: " + links.Links[i].Title)

		// Search Algorithim

		if fuzzy.Match(SearchTerm, links.Links[i].Title) ||
			fuzzy.Match(SearchTerm, links.Links[i].Tag) ||
			fuzzy.Match(SearchTerm, links.Links[i].URL) {
			Results = append(Results, links.Links[i])
			// fmt.Println("Link Title: " + links.Links[i].Title)
		}

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
	jsonFile, err := os.OpenFile("data.json", os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var links LinksType

	json.Unmarshal(byteValue, &links)

	links.Links = append(links.Links, link)

	stream, err := json.Marshal(links)

	fmt.Println(string(stream))

	jsonFile.Close()

	err1 := os.Truncate("data.json", 0)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Successfully Opened data.json")
	dataFile, err := os.OpenFile("data.json", os.O_RDWR, os.ModeAppend)
	byteStream, err := dataFile.WriteString(string(stream))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("wrote %d bytes\n", byteStream)
	dataFile.Close()

}
