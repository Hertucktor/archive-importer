package importer

import (
	"context"
	"github.com/Hertucktor/archive-importer/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

func ImportCardsIntoDatabase(dbClient *mongo.Client, conf config.Config, page, delimiter int, logger *zap.SugaredLogger) error {
	ctx := context.TODO()

	logger.Infof("Start of import: %v", time.Now())
	for delimiter != 0 {
		requestAllCards, err := RequestAllCardsFromAPI(page, logger)
		if err != nil {
			return err
		}
		//If Card is in Database, update modified else insert card
		for _, card := range requestAllCards.Cards {
			//TODO: instead of scraping the entire DB over and over again, create an additional Collection with unique
			//identifier info of each card
			found, err := FindCard(card.SetName, card.Number, dbClient, ctx, conf)
			if err != nil {
				return err
			}
			if !found {
				if err = InsertImportCard(card, dbClient, ctx, conf, logger); err != nil {
					return err
				}
			} else {
				if err = UpdateSingleCard(card, card.SetName, card.Number, dbClient, ctx, conf, logger); err != nil {
					return err
				}
			}
		}

		logger.Infof("Request page number:%v", page)

		increadePageImpression(page)

		delimiter = len(requestAllCards.Cards)
	}
	logger.Infof("End of import: %v", time.Now().Unix())
	return nil
}

func increadePageImpression(page int) (increasedPage int) {
	increasedPage = page + 1
	return
}
