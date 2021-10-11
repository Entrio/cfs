package main

import (
	"context"
	gen "github.com/Entrio/cfs/internal/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	address     = "127.0.0.1:43567"
	defaultName = "UnknownEther"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to remote server")
	}
	defer conn.Close()

	client := gen.NewCFSPublicClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	r, err := client.GetServerInfo(ctx, &gen.ServerInfoRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to remote server")
	}

	log.Info().Str("name", r.Name).Msg("connected to server")

	_, err = client.CreateFarm(ctx, new(gen.Empty))
	checkErr(err, "farm info")
	log.Info().Msg("Created a farm")

}

func checkErr(err error, action string) {
	if err != nil {
		log.Fatal().Err(err).Msg(action)
	}
}
