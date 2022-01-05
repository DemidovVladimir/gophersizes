package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dghubble/oauth1"
)

func viewHashTagHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		getPage, _ := template.ParseFiles("hash_tags.html")
		getPage.Execute(w, r)
	case "POST":
		var country string
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		country = r.FormValue("country")
		tags := getHashTags(country)
		postPage, _ := template.ParseFiles("hash_tags.html")
		postPage.Execute(w, tags)
	default:
		fmt.Fprintf(w, "Unable to get result.")
	}
}

func getHashTags(country string) []string {
	var listHashTags []string

	consumerKey := "lEmUUbDRNLHNV0wXZxleTXGPG"
	consumerSecret := "EBxb0ovNBb4A2MuhbdDdeYQOAvJG2pVmK0MLx9uuhkTfh4cLqf"
	accessToken := "816628106559651842-qvfryyocRKw2GwCguTHfs5uN3VQ7ZjU"
	accessSecret := "9dvBbhuoyvBm49XUjxE92lMn4iQpLwFVjWFwzVSMkAORi"

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/1.1/trends/place.json?id=" + country

	resp, _ := httpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var JSON = strings.TrimLeft(string(body), "[")
	JSON = strings.TrimRight(JSON, "]")

	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)
	trends := info["trends"].([]interface{})

	for _, element := range trends {
		if trendList, ok := element.(map[string]interface{}); ok {
			for key, value := range trendList {
				// Filter hashtags started with #
				if strings.Contains(key, "name") && strings.Contains(value.(string), "#") {
					listHashTags = append(listHashTags, value.(string))
				}
			}
		}
	}
	return listHashTags
}

func main() {
	port := flag.Int("port", 3000, "Run server on port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", viewHashTagHandler)

	fmt.Printf("Start running server on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
