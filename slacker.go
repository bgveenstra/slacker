// package main // slacker

package slacker

import (
	"log";
	"os";
	"net/http";
	"net/url";
	"io/ioutil";
	"bytes";
	"fmt"
)

// func main() {

// 	// localCRUDBooks := "http://localhost:3000/books"
// 	// GetBooks(localCRUDBooks)
// 	// PostBook(localCRUDBooks)
// 	PostSlackMessage("test")

// }

func PostSlackMessage(message string){
	log.Println(message)

	slackTarget := os.Getenv("HOWLER_SLACK_WEBHOOK_URL")

	// create a jsonesque body for the post request
	reqStr := fmt.Sprintf("{\"text\":\"%s\"}", message)
	reqBody := bytes.NewBufferString(reqStr)
	response, err := http.Post(slackTarget, "application/json", reqBody)
	if err != nil {
		log.Fatal("post error", err)
	} 

	defer response.Body.Close()
	
	log.Println(response.Status)
	// if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("body read error", err)
		}
		// body is a byte slice
		log.Println("body", string(body[:]))
	// }


}

func GetBooks(targetUrl string) {

	response, err := http.Get(targetUrl)
	if err != nil {
		log.Fatal("get error", err)
	} 

	defer response.Body.Close()
	
	log.Println(response.Status)
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("body read error", err)
		}
		// body is a byte slice
		log.Println("body", string(body[:]))
	}
}

func PostBook(targetUrl string) {
	v := url.Values{}
	v.Set("title", "One Fish, Two Fish, Red Fish, Blue Fish")
	v.Set("author", "Dr. Seuss")  // image (url), releaseDate
	v.Set("image", "http://3.bp.blogspot.com/_KjHHNyQNQ-w/S8pY--rDoII/AAAAAAAAAEw/_Dez_1Kc5a8/s1600/one-fish-two-fish.jpg")
	log.Println(v)
	response, err := http.PostForm(targetUrl, v)
	if err != nil {
		log.Fatal("post error", err)
	} 

	defer response.Body.Close()
	
	log.Println(response.Status)
	// if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("body read error", err)
		}
		// body is a byte slice
		log.Println("body", string(body[:]))
	// }

}