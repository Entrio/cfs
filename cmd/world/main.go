package main

import (
	"github.com/Entrio/cfs/internal/models"
	gen "github.com/Entrio/cfs/internal/proto"
	"github.com/Entrio/cfs/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
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

	{
		go func() {
			log.Info().Msg("Starting GRPC server")
			lis, err := net.Listen("tcp", "0.0.0.0:43567")
			if err != nil {
				log.Fatal().Err(err).Msg("creating GRPC listener handle")
			}
			s := grpc.NewServer()
			gen.RegisterCFSPublicServer(s, server.NewPublicServer(w))
			log.Info().Msg("GRPC server started")
			if err = s.Serve(lis); err != nil {
				log.Fatal().Err(err).Msg("GRPC server listen")
			}
		}()
	}

	//go test(w)
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
