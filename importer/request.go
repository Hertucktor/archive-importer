package importer

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
RequestAllCardsFromAPI receives a response with type *http.Response from
the mtg api containing 100 cards.
Returning the response and an error
*/
func RequestAllCardsFromAPI(page int) (*http.Response, error) {
	//GET request to URL with page param
	apiUrl := "https://api.magicthegathering.io/v1/cards?page=" + strconv.Itoa(page)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}
	return resp, err
}
