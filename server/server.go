package recordupdater

import (
	"context"

	pb "github.com/brotherlogic/recordupdater/proto"
)

type Server struct{}

func (s *Server) loadConfig(ctx context.Context) (*pb.Config, error) {
	return &pb.Config{}, nil
}

func (s *Server) saveConfig(ctx context.Context, config *pb.Config) error {
	return nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	config, err := s.loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	for _, update := range config.GetUpdates() {
		if update.InstanceId == req.GetInstanceId() {
			if req.GetUpdateTime() < update.GetNextUpdateTime() {
				update.NextUpdateTime = req.GetUpdateTime()
				update.UpdatePurpose = req.GetPurpose()
			}
			return &pb.UpdateResponse{}, s.saveConfig(ctx, config)
		}
	}

	config.Updates = append(config.Updates, &pb.Update{
		InstanceId:     req.GetInstanceId(),
		NextUpdateTime: req.GetUpdateTime(),
		UpdatePurpose:  req.GetPurpose(),
	})

	return &pb.UpdateResponse{}, s.saveConfig(ctx, config)
}
