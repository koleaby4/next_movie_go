package plexapi

import (
	"io"
	"log"
	"net/http"
)

func GetUserToken(appToken string) string {
	url := "http://10.10.10.47:32400/security/token"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln("error getting Plex user token", err)
	}
	req.Header.Add("X-Plex-Token", appToken)

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln("error getting Plex user token", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("error getting Plex user token", err)
	}
	return string(body)
}
