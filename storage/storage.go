package storage

import "github.com/Nikita213-hub/testTaskGo/models"

type Storage interface {
	AddQuote(quoute string, author string) (*models.QuoteData, error)
	DeleteQuoteById(id int) error
	GetAllQuotes(filter string) (map[int]*models.QuoteData, error)
	GetRandomQuote() (*models.QuoteData, error)
	GetQuoteById(id int) (*models.QuoteData, error)
}
