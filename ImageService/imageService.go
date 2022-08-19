package imageservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ImageUrls struct {
	Full    string `json:"full"`
	Regular string `json:"regular"`
}

type randomImage struct {
	Id        string    `json:"id"`
	ImageUrls ImageUrls `json:"urls"`
}

func GetRandomUnsplashImage() string {
	response, error := http.Get("https://api.unsplash.com/photos/random/?client_id=XXXX&count=1")

	if error != nil {
		log.Fatalln(error)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Fatalln(error)
	}

	urls := make([]randomImage, 0)
	jsonError := json.Unmarshal([]byte(body), &urls)

	if jsonError != nil {
		log.Fatalln(jsonError)
	}

	return urls[0].ImageUrls.Regular
}
