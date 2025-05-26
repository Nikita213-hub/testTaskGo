package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nikita213-hub/testTaskGo/models"
	"github.com/Nikita213-hub/testTaskGo/server"
	"github.com/Nikita213-hub/testTaskGo/storage"
)

type QuotesHandlers struct {
	storage storage.Storage
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewQuotesHandlers(strg storage.Storage) server.Router {
	return &QuotesHandlers{
		storage: strg,
	}
}

func (qh *QuotesHandlers) Routes() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/quotes",
			Handler: qh.AddQuote,
		},
		{
			Method:  "GET",
			Path:    "/quotes",
			Handler: qh.GetAllQuotes,
		},
		{
			Method:  "GET",
			Path:    "/quotes/random",
			Handler: qh.GetRandomQuote,
		},
		{
			Method:  "DELETE",
			Path:    "/quotes/{id}",
			Handler: qh.DeleteQuote,
		},
	}
}

func (qh *QuotesHandlers) AddQuote(w http.ResponseWriter, req *http.Request) {
	var quoteData models.QuoteData
	if err := json.NewDecoder(req.Body).Decode(&quoteData); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid request body format")
		return
	}

	if quoteData.Quote == "" || quoteData.Author == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Quote and author cannot be empty")
		return
	}

	quote, err := qh.storage.AddQuote(quoteData.Quote, quoteData.Author)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to add quote: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(quote); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to encode response: "+err.Error())
		return
	}
}

func (qh *QuotesHandlers) GetAllQuotes(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	var quotes map[int]*models.QuoteData
	var err error

	if len(params) > 0 {
		if !params.Has("author") {
			sendErrorResponse(w, http.StatusBadRequest, "Invalid query parameter")
			return
		}
		quotes, err = qh.storage.GetAllQuotes(params.Get("author"))
	} else {
		quotes, err = qh.storage.GetAllQuotes("")
	}

	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to get quotes: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to encode response: "+err.Error())
		return
	}
}

func (qh *QuotesHandlers) GetRandomQuote(w http.ResponseWriter, req *http.Request) {
	quote, err := qh.storage.GetRandomQuote()
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, "No quotes available in the storage: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quote); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to encode response: "+err.Error())
		return
	}
}

func (qh *QuotesHandlers) DeleteQuote(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	if len(id) == 0 {
		sendErrorResponse(w, http.StatusBadRequest, "Quote Id is required")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid quote Id format")
		return
	}

	_, err = qh.storage.GetQuoteById(idInt)
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, "Quote not found: "+err.Error())
		return
	}

	err = qh.storage.DeleteQuoteById(idInt)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete quote: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Quote successfully deleted"})
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
