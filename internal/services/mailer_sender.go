package services

import (
	"context"
	"daily-quote/internal/domain"
	"github.com/mailersend/mailersend-go"
	"os"
)

func SendMail(ctx context.Context, quote domain.Quote) (string, error) {
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	subject := domain.SenderSubject
	from := getFromSender()
	recipients := getRecipients()
	variables := getVariables()
	personalization := getPersonalization(quote)

	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID(os.Getenv("EMAIL_TEMPLATE_ID"))
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		return "", err
	}

	return res.Header.Get("X-Message-Id"), nil
}

func getFromSender() mailersend.From {
	from := mailersend.From{}
	from.Name = domain.SenderName
	from.Email = domain.SenderEmail

	return from
}

func getRecipients() []mailersend.Recipient {
	var (
		recipients    []mailersend.Recipient
		manuRecipient mailersend.Recipient
		cataRecipient mailersend.Recipient
	)

	manuRecipient.Name = domain.RecipientManuName
	manuRecipient.Email = domain.RecipientManuEmail
	recipients = append(recipients, manuRecipient)

	cataRecipient.Name = domain.RecipientCataName
	cataRecipient.Email = domain.RecipientCataEmail
	recipients = append(recipients, cataRecipient)

	return recipients
}

func getVariables() []mailersend.Variables {
	return []mailersend.Variables{
		{
			Email: domain.RecipientManuEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "url",
					Value: domain.SenderUrlSite,
				},
			},
		},
		{
			Email: domain.RecipientCataEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "url",
					Value: domain.SenderUrlSite,
				},
			},
		},
	}
}

func getPersonalization(quote domain.Quote) []mailersend.Personalization {
	var (
		manuDailyQuotePersonalization mailersend.Personalization
		cataDailyQuotePersonalization mailersend.Personalization

		personalizations []mailersend.Personalization
	)

	manuDailyQuotePersonalization.Email = domain.RecipientManuEmail
	manuDailyQuotePersonalization.Data = map[string]interface{}{
		"name":         domain.RecipientManuName,
		"work":         quote.Work,
		"quote":        quote.Phrase,
		"author":       quote.Author,
		"date_created": quote.DateCreated,
	}
	personalizations = append(personalizations, manuDailyQuotePersonalization)

	cataDailyQuotePersonalization.Email = domain.RecipientCataEmail
	cataDailyQuotePersonalization.Data = map[string]interface{}{
		"name":         domain.RecipientCataName,
		"work":         quote.Work,
		"quote":        quote.Phrase,
		"author":       quote.Author,
		"date_created": quote.DateCreated,
	}
	personalizations = append(personalizations, cataDailyQuotePersonalization)

	return personalizations
}
