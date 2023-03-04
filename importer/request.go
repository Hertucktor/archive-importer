package importer

import (
	"encoding/json"
	"github.com/Hertucktor/archive-importer/mongodb"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

/*
RequestAllCardsFromAPI receives a response with type *http.Response from
the mtg api containing 100 cards.
Returning the response and an error
*/
func RequestAllCardsFromAPI(page int, logger *zap.SugaredLogger) (mongodb.MultipleCards, error) {
	var response mongodb.MultipleCards
	//GET request to URL with page param
	apiUrl := "https://api.magicthegathering.io/v1/cards?page=" + strconv.Itoa(page)
	apiResponse := makeRequest(apiUrl, logger)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(apiResponse.Body)
	//TODO: repeat on !200 until response or system abort
	if apiResponse.StatusCode != 200 {
		//sleepTime := "2"
		logger.Errorf("Http status code:%v", apiResponse.StatusCode)
		//logger.Info("error in API request, sleep for: %v and repeat API request", sleepTime)
		//time.Sleep(2 * time.Second)
	}
	//reads response body into []byte
	body, err := io.ReadAll(apiResponse.Body)
	if err != nil {
		logger.Error(err)
	}

	//parses response body []byte values into response
	if err = json.Unmarshal(body, &response); err != nil {
		logger.Error(err)
		return response, err
	}

	return response, err
}

func makeRequest(apiUrl string, logger *zap.SugaredLogger) *http.Response {
	logger.Info(apiUrl)
	resp, err := http.Get(apiUrl)
	if err != nil {
		logger.Fatalf("problem with http GET request\n%v", err)
	}
	return resp
}
