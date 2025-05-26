package memstorage

import (
	"testing"

	"github.com/Nikita213-hub/testTaskGo/models"
	"github.com/stretchr/testify/assert"
)

func TestAddQuote(t *testing.T) {
	tests := []struct {
		name    string
		quote   string
		author  string
		wantErr bool
		checkID bool
	}{
		{
			name:    "valid quote",
			quote:   "Test quote",
			author:  "Test author",
			wantErr: false,
			checkID: true,
		},
		{
			name:    "empty quote",
			quote:   "",
			author:  "Test author",
			wantErr: false,
			checkID: true,
		},
		{
			name:    "empty author",
			quote:   "Test quote",
			author:  "",
			wantErr: false,
			checkID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewMemStorage()
			quote, err := storage.AddQuote(tt.quote, tt.author)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.checkID {
					assert.Equal(t, &models.QuoteData{Quote: tt.quote, Author: tt.author}, quote)
					assert.Equal(t, tt.quote, quote.Quote)
					assert.Equal(t, tt.author, quote.Author)
				}
			}
		})
	}
}

func TestGetAllQuotes(t *testing.T) {
	storage := NewMemStorage()

	storage.AddQuote("Quote 1", "Author 1")
	storage.AddQuote("Quote 2", "Author 1")
	storage.AddQuote("Quote 3", "Author 2")

	tests := []struct {
		name    string
		filter  string
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all quotes",
			filter:  "",
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "filter by author 1",
			filter:  "Author 1",
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "filter by author 2",
			filter:  "Author 2",
			wantLen: 1,
			wantErr: false,
		},
		{
			name:    "filter by non-existent author",
			filter:  "Author 3",
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quotes, err := storage.GetAllQuotes(tt.filter)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantLen, len(quotes))
			}
		})
	}
}

func TestGetRandomQuote(t *testing.T) {
	storage := NewMemStorage()

	t.Run("empty storage", func(t *testing.T) {
		quote, err := storage.GetRandomQuote()
		assert.Error(t, err)
		assert.Nil(t, quote)
	})

	storage.AddQuote("Test quote", "Test author")
	t.Run("storage with one quote", func(t *testing.T) {
		quote, err := storage.GetRandomQuote()
		assert.NoError(t, err)
		assert.NotNil(t, quote)
		assert.Equal(t, "Test quote", quote.Quote)
		assert.Equal(t, "Test author", quote.Author)
	})
}

func TestDeleteQuote(t *testing.T) {
	storage := NewMemStorage()

	quote, _ := storage.AddQuote("Test quote", "Test author")
	_ = quote
	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "delete existing quote",
			id:      0,
			wantErr: false,
		},
		{
			name:    "delete non-existent quote",
			id:      -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.DeleteQuoteById(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				quote, _ := storage.GetQuoteById(tt.id)
				assert.Nil(t, quote)
			}
		})
	}
}

func TestGetQuoteById(t *testing.T) {
	storage := NewMemStorage()

	quote, _ := storage.AddQuote("Test quote", "Test author")
	_ = quote
	tests := []struct {
		name    string
		id      int
		want    bool
		wantErr bool
	}{
		{
			name:    "get existing quote",
			id:      0,
			want:    true,
			wantErr: false,
		},
		{
			name:    "get non-existent quote",
			id:      999,
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote, err := storage.GetQuoteById(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.want {
					assert.NotNil(t, quote)
					assert.Equal(t, "Test quote", quote.Quote)
					assert.Equal(t, "Test author", quote.Author)
				} else {
					assert.Nil(t, quote)
				}
			}
		})
	}
}
