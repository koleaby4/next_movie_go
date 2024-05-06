package plexapi

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string, headers map[string]string, data map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln("error constructing request", err)
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)

	if err != nil {
		return "", errors.New("error sending request: " + err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("error reading body:" + err.Error())
	}
	return string(body), nil

}

func GetUserToken(appToken string) string {
	url := "http://hp-server2:32400/security/token?type=delegation&scope=all"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("X-Plex-Token", appToken)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func Do(userToken string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://plex.tv/api/v2/user", nil)

	if err != nil {
		log.Fatalln("failed contacting plex.tv", err)
	}

	req.Header.Add("accept", "application/json")

	query := req.URL.Query()
	query.Add("X-Plex-Product", "next_movie_go")
	query.Add("X-Plex-Client-Identifier", "4649f6f7-7bec-41bb-aa2b-ba6067c506e")
	query.Add("X-Plex-Token", userToken)

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln("error doing request to Plex" + err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("error reading body" + err.Error())
	}
	log.Println(string(body))

	/*
		$ curl -X GET https://plex.tv/api/v2/user \
		  -H 'accept: application/json' \
		  -d 'X-Plex-Product=My Cool Plex App' \
		  -d 'X-Plex-Client-Identifier=<clientIdentifier>' \
		  -d 'X-Plex-Token=<userToken>'
	*/
}
