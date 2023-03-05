package importer

import (
	"context"
	"encoding/json"
	"github.com/Hertucktor/archive-importer/config"
	"github.com/Hertucktor/archive-importer/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

func ImportCardsIntoDatabase(dbClient *mongo.Client, conf config.Config, page, delimiter int, logger *zap.SugaredLogger) error {
	logger.Infof("Start of import: %v", time.Now())
	var resp *http.Response
	var err error
	ctx := context.TODO()
	var multipleCards mongodb.MultipleCards

	for delimiter != 0 {

		// repeat API request if http status code != 200
		for i := 0; i < 3; i++ {
			resp, err = RequestAllCardsFromAPI(page)
			if err != nil {
				logger.Errorf("Error on attempt %d: %v\n", i+1, err)
				time.Sleep(2 * time.Second) // wait for 2 second before retrying
				continue
			}
			break // exit the loop if request was successful
		}
		logger.Infof("Request page number:%v", page)
		page++

		// take API response and fill data in mongodb.MultipleCards
		multipleCards, err = handleAPIResponse(resp, logger)
		if err != nil {
			logger.Fatal(err)
		}

		//If Card is in Database, update modified else insert card
		for _, card := range multipleCards.Cards {

			//TODO: increase performance:
			//instead of scraping the entire DB over and over again, create an additional Collection with unique
			//identifier info of each card
			found, err := FindCard(card.SetName, card.Number, dbClient, ctx, conf, logger)
			logger.Infof("Card was found in DB: %v", found)
			if err != nil {
				logger.Error(err)
			}
			if !found {
				err = InsertImportCard(card, dbClient, ctx, conf, logger)
				if err != nil {
					logger.Error(err)
				}
			} else {
				err = UpdateSingleCard(card.SetName, card.Number, dbClient, ctx, conf, logger)
				if err != nil {
					logger.Error(err)
				}
			}
		}
		delimiter = len(multipleCards.Cards)
	}
	logger.Infof("End of import: %v", time.Now().Unix())
	return nil
}

func handleAPIResponse(resp *http.Response, logger *zap.SugaredLogger) (mongodb.MultipleCards, error) {
	responseCardInfo := mongodb.MultipleCards{}
	//reads responseCardInfo body into []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return mongodb.MultipleCards{}, err
	}

	//parses responseCardInfo body []byte values into responseCardInfo
	if err = json.Unmarshal(body, &responseCardInfo); err != nil {
		logger.Error(err)
		return mongodb.MultipleCards{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(resp.Body)

	return responseCardInfo, nil
}
