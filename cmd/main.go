package main

import (
	"github.com/make-money-fast/empty/config"
	"github.com/make-money-fast/empty/internal/model"
	"github.com/make-money-fast/empty/internal/router"
)

func main() {
	config.Load()
	model.Load()
	router.Start()
}
