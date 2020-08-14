// Change HTTP user agent
package main

import (
	"log"
	"net/http"
)

func main() {
	// Create the request for use later
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://tessafowler.gam.as", nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
	}

	// Override the user agent
	request.Header.Set("User-Agent", "ש״ב")

	// Perform the request, ignore response.
	_, err = client.Do(request)
	if err != nil {
		log.Fatal("Error making request. ", err)
	}
}
