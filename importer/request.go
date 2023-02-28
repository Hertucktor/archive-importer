package importer

import (
	"encoding/json"
	"errors"
	"github.com/Hertucktor/archive-importer/mongodb"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

/*
RequestAllCards receives a response with type *http.Response from
the mtg api containing 100 cards.
Returning the response and an error
*/
func RequestAllCards(page int, logger *zap.SugaredLogger) (mongodb.MultipleCards, error) {
	var response mongodb.MultipleCards
	var resp *http.Response
	var err error
	var body []byte
	//GET request to URL with page param
	if resp, err = http.Get("https://api.magicthegathering.io/v1/cards?page=" + strconv.Itoa(page)); err != nil {
		logger.Error("Error: problem with http GET request\n")
		return response, err
	}

	logger.Infof("HTTP GET REQUEST TO https://api.magicthegathering.io/v1/cards?page=\n%v", page)

	defer func() {
		if err = resp.Body.Close(); err != nil {
			logger.Fatalf("Fatal: couldn't close response body\n")
		}
	}()
	//checks if there is a http status code other than 200 in the response
	if resp.StatusCode != 200 {
		err = errors.New("http statuscode != 200")
		logger.Errorf("Http status code:\n%v, error:%v", resp.StatusCode, err)
		return response, err
	}
	//reads response body into []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		logger.Error(err)
		return response, err
	}
	//parses response body []byte values into response
	if err = json.Unmarshal(body, &response); err != nil {
		logger.Error(err)
		return response, err
	}

	return response, err
}
