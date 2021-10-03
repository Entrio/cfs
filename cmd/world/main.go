package main

import (
	"github.com/Entrio/cfs/internal/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

/*
This is the master server that controls the farming world
*/

var ()

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	t := time.NewTicker(time.Second / 30)
	w := models.NewWorld(models.WorldMeta{
		Name:        "Islands",
		Description: "A simple island map",
	})

	go func() {
		log.Debug().Msg("Starting game update loop")
		for {
			select {
			case <-t.C:
				if err := w.ProcessUpdate(); err != nil {
					log.Warn().Err(err).Msg("failed yp update world")
				}
			}
		}

	}()

	go test(w)
	<-w.Quit()
	log.Info().Msg("Program exit")
}

func test(w *models.World) {
	t := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-w.Quit():
			break
		case <-t.C:
			w.AddFarm(models.Wood)
		}
	}
}
