package repository

import "daily-quote/internal/domain"

func GetDailyQuote() (domain.Quote, error) {
	resultQuote, err := ClientDB.Query(
		"SELECT author, work, phrase, date_created FROM `quotes`.`quotes` ORDER BY RAND() LIMIT 1;")
	if err != nil {
		return domain.Quote{}, err
	}

	var quote domain.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return domain.Quote{}, err
		}
	}

	return quote, nil
}
