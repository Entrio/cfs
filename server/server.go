package server

import (
	"context"
	"github.com/Entrio/cfs/internal/models"
	gen "github.com/Entrio/cfs/internal/proto"
)

type server struct {
	gen.UnimplementedCFSPublicServer
	world *models.World
}

func NewPublicServer(w *models.World) gen.CFSPublicServer {
	return &server{
		world: w,
	}
}

func (s server) GetServerInfo(ctx context.Context, request *gen.ServerInfoRequest) (*gen.ServerInfoResponse, error) {
	return &gen.ServerInfoResponse{Name: s.world.GetName()}, nil
}

func (s server) CreateFarm(ctx context.Context, request *gen.Empty) (*gen.Empty, error) {
	return new(gen.Empty), s.world.AddFarm(models.Wood)
}

func (s server) GetFarms(context.Context, *gen.Empty) (*gen.GetFarmsResponse, error) {
	farms := make([]*gen.Farm, 0)
	for _, f := range s.world.GetFarmManager().GetFarms() {
		farms = append(farms, &gen.Farm{
			Id:    f.GetID().String(),
			Type:  models.ResourceTypeString[f.GetResourceType()],
			Cells: int32(f.GetCellsCount()),
		})
	}
	return &gen.GetFarmsResponse{Farms: farms}, nil
}
