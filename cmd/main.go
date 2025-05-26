package main

import (
	"context"

	"github.com/Nikita213-hub/testTaskGo/handlers"
	"github.com/Nikita213-hub/testTaskGo/helpers"
	"github.com/Nikita213-hub/testTaskGo/server"
	memstorage "github.com/Nikita213-hub/testTaskGo/storage/memStorage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel // i will be able to use that func to stop the server if i face any critical error
	storage := memstorage.NewMemStorage()
	quotesHandlers := handlers.NewQuotesHandlers(storage)
	addr, port := helpers.GetHostAddr()
	server := server.NewServer(addr, ":"+port)
	err := server.Run(ctx, quotesHandlers)
	if err != nil {
		panic(err)
	}
}
