package main

import (
	"context"
	"log"
	"time"
	"translatorBot/cmd/api/app"
	"translatorBot/cmd/api/clients"
	"translatorBot/cmd/api/config"
	"translatorBot/cmd/api/domain"
	"translatorBot/cmd/api/providers"
)

func main() {
	config.InitializeConfigs()
	providers.InitializeProviders()
	clients.InitializeClients()
	healthCheckTranslator()
	app.StartDabot()
}

func healthCheckTranslator() {
	for {
		_, err := providers.TranslationProvider.Translate(context.Background(), domain.TranslationRequest{
			Quote:  "Are you up?",
			Source: "en",
			Target: "de",
		})
		if err == nil {
			log.Println("Ready to translate")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
