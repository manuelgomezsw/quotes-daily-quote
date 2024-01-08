package services

import (
	"context"
	"daily-quote/internal/domain"
	"daily-quote/internal/repository"
)

func SendDailyQuote(ctx context.Context) (string, error) {
	dailyQuote, err := repository.GetDailyQuote()
	if err != nil {
		return "", err
	}

	completeDataDailyQuote(&dailyQuote)

	confirmationID, err := SendMail(ctx, dailyQuote)
	if err != nil {
		return "", err
	}

	return confirmationID, nil
}

func completeDataDailyQuote(quote *domain.Quote) {
	if quote.Author == "" {
		quote.Author = domain.Desconocido
	}
}
