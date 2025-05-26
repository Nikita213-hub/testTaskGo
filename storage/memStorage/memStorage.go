package memstorage

import (
	"errors"
	"math/rand/v2"
	"sync"

	"github.com/Nikita213-hub/testTaskGo/models"
	"github.com/Nikita213-hub/testTaskGo/storage"
)

type Storage struct {
	counter int
	Quotes  map[int]*models.QuoteData
	mx      sync.RWMutex
}

func NewMemStorage() storage.Storage {
	return &Storage{
		counter: 0,
		Quotes:  make(map[int]*models.QuoteData),
	}
}

func (strg *Storage) AddQuote(quote, author string) (*models.QuoteData, error) {
	strg.mx.Lock()
	defer strg.mx.Unlock()

	id := strg.counter
	strg.Quotes[int(id)] = &models.QuoteData{Quote: quote, Author: author}
	strg.counter++
	return strg.Quotes[int(id)], nil
}

func (strg *Storage) DeleteQuoteById(id int) error {
	strg.mx.Lock()
	defer strg.mx.Unlock()

	if _, isExisting := strg.Quotes[id]; !isExisting {
		return errors.New("there is no such quote to delete")
	}
	delete(strg.Quotes, id)
	strg.counter--
	return nil
}

func (strg *Storage) GetAllQuotes(filter string) (map[int]*models.QuoteData, error) {
	strg.mx.RLock()
	defer strg.mx.RUnlock()

	if filter == "" {
		return strg.Quotes, nil
	}

	filteredQuotes := make(map[int]*models.QuoteData)
	for k, v := range strg.Quotes {
		if filter == v.Author {
			filteredQuotes[k] = v
		}
	}
	return filteredQuotes, nil
}

func (strg *Storage) GetRandomQuote() (*models.QuoteData, error) {
	strg.mx.RLock()
	defer strg.mx.RUnlock()

	if strg.counter == 0 {
		return nil, errors.New("there are no quotes in the storage")
	}
	randId := rand.IntN(strg.counter)
	return strg.Quotes[randId], nil
}

func (strg *Storage) GetQuoteById(id int) (*models.QuoteData, error) {
	strg.mx.RLock()
	defer strg.mx.RUnlock()
	quote, ok := strg.Quotes[id]
	if !ok {
		return nil, errors.New("there is no such quote")
	}
	return quote, nil
}
