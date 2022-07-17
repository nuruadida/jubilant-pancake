package application

import (
	"Nurul-trial/config/env"
	"Nurul-trial/router"
	"github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		return
	}
}

func StartApp() {

	address := env.Config.Environment + ":" + env.Config.AppPort
	confi := gin.Config{
		ListenAddr: address,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is running at >>>> " + address)
		},
	}
	gin.Run(confi)
}
