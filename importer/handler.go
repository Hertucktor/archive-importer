package importer

import (
	"github.com/Hertucktor/archive-importer/config"
	"go.uber.org/zap"
	"time"
)

func ImportCardsIntoDatabase(conf config.Config, page, delimiter int, logger *zap.SugaredLogger) error {
	client, ctx, err := ImporterClient(conf.DBUser, conf.DBPass, conf.DBPort, conf.DBName)
	if err != nil {
		return err
	}

	logger.Infof("Start of import: %v", time.Now().Unix())
	for delimiter != 0 {
		requestAllCards, err := RequestAllCards(page, logger)
		if err != nil {
			return err
		}

		for _, card := range requestAllCards.Cards {
			//If Card is in Database, update modified else insert card
			found, err := FindCard(card.SetName, card.Number, client, ctx, conf)
			if err != nil {
				return err
			}
			if found != true {
				if err = InsertImportCard(card, client, ctx, conf, logger); err != nil {
					return err
				}
			} else {
				if err = UpdateSingleCard(card, card.SetName, card.Number, client, ctx, conf, logger); err != nil {
					return err
				}
			}
		}

		logPageImpression(page, logger)

		increadePageImpression(page)

		delimiter = len(requestAllCards.Cards)
	}
	logger.Infof("End of import: %v", time.Now().Unix())
	return err
}

func logPageImpression(page int, logger *zap.SugaredLogger) {
	switch page {
	case 100:
		logger.Infof("Reached page: %v, at time: %v", page, time.Now().Unix())
	case 200:
		logger.Infof("Reached page: %v, at time: %v", page, time.Now().Unix())
	case 300:
		logger.Infof("Reached page: %v, at time: %v", page, time.Now().Unix())
	case 400:
		logger.Infof("Reached page: %v, at time: %v", page, time.Now().Unix())
	case 500:
		logger.Infof("Reached page: %v, at time: %v", page, time.Now().Unix())
	default:
		logger.Infof("Request page number:%v one page = 100 cards", page)
	}
}

func increadePageImpression(page int) (increasedPage int) {
	increasedPage = page + 1
	return
}
