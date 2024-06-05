package main

import (
	"EPL/match_service/match_service/proto"
	"context"
	"database/sql"
	"log"
	"net"
	"time"

	"EPL/match_service/internal/entity"
	"EPL/match_service/internal/interfaces/repository"
	"EPL/match_service/internal/usecases"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type matchServiceServer struct {
	proto.UnimplementedMatchServiceServer `json:"proto.UnimplementedMatchServiceServer"`
	createUseCase                         *usecases.CreateMatch    `json:"createUseCase,omitempty"`
	getUseCase                            *usecases.FindMatchByID  `json:"getUseCase,omitempty"`
	updateUseCase                         *usecases.UpdateMatch    `json:"updateUseCase,omitempty"`
	deleteUseCase                         *usecases.DeleteMatch    `json:"deleteUseCase,omitempty"`
	listUseCase                           *usecases.FindAllMatches `json:"listUseCase,omitempty"`
}

func (s *matchServiceServer) CreateMatch(ctx context.Context, req *proto.CreateMatchRequest) (*proto.CreateMatchResponse, error) {
	match := entity.NewMatch{
		HomeTeam:  req.HomeTeam,
		AwayTeam:  req.AwayTeam,
		Date:      req.Date,
		Status:    req.Status,
		HomeScore: int(req.HomeScore),
		AwayScore: int(req.AwayScore),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdMatch, err := s.createUseCase.Execute(match)
	if err != nil {
		return nil, err
	}

	return &proto.CreateMatchResponse{
		Id:        int32(createdMatch.ID),
		HomeTeam:  createdMatch.HomeTeam,
		AwayTeam:  createdMatch.AwayTeam,
		Date:      createdMatch.Date,
		Status:    createdMatch.Status,
		HomeScore: int32(createdMatch.HomeScore),
		AwayScore: int32(createdMatch.AwayScore),
		CreatedAt: timestamppb.New(createdMatch.CreatedAt),
	}, nil
}

func (s *matchServiceServer) GetMatch(ctx context.Context, req *proto.GetMatchRequest) (*proto.GetMatchResponse, error) {
	match, err := s.getUseCase.Execute(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.GetMatchResponse{
		Id:        int32(match.ID),
		HomeTeam:  match.HomeTeam,
		AwayTeam:  match.AwayTeam,
		Date:      match.Date,
		Status:    match.Status,
		HomeScore: int32(match.HomeScore),
		AwayScore: int32(match.AwayScore),
		CreatedAt: timestamppb.New(match.CreatedAt),
		UpdatedAt: timestamppb.New(match.UpdatedAt),
	}, nil
}

func (s *matchServiceServer) UpdateMatch(ctx context.Context, req *proto.UpdateMatchRequest) (*proto.UpdateMatchResponse, error) {
	match := entity.NewMatch{
		ID:        int(req.Id),
		HomeTeam:  req.HomeTeam,
		AwayTeam:  req.AwayTeam,
		Date:      req.Date,
		Status:    req.Status,
		HomeScore: int(req.HomeScore),
		AwayScore: int(req.AwayScore),
		UpdatedAt: time.Now(),
	}

	updatedMatch, err := s.updateUseCase.Execute(match)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateMatchResponse{
		Id:        int32(updatedMatch.ID),
		HomeTeam:  updatedMatch.HomeTeam,
		AwayTeam:  updatedMatch.AwayTeam,
		Date:      updatedMatch.Date,
		Status:    updatedMatch.Status,
		HomeScore: int32(updatedMatch.HomeScore),
		AwayScore: int32(updatedMatch.AwayScore),
		CreatedAt: timestamppb.New(updatedMatch.CreatedAt),
		UpdatedAt: timestamppb.New(updatedMatch.UpdatedAt),
	}, nil
}

func (s *matchServiceServer) DeleteMatch(ctx context.Context, req *proto.DeleteMatchRequest) (*proto.DeleteMatchResponse, error) {
	id := int(req.Id)

	if err := s.deleteUseCase.Execute(id); err != nil {
		return nil, err
	}

	return &proto.DeleteMatchResponse{}, nil
}

func (s *matchServiceServer) ListMatches(ctx context.Context, req *proto.ListMatchesRequest) (*proto.ListMatchesResponse, error) {
	matches, err := s.listUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var matchResponses []*proto.Match
	for _, match := range matches {
		matchResponses = append(matchResponses, &proto.Match{
			Id:        int32(match.ID),
			HomeTeam:  match.HomeTeam,
			AwayTeam:  match.AwayTeam,
			Date:      match.Date,
			Status:    match.Status,
			HomeScore: int32(match.HomeScore),
			AwayScore: int32(match.AwayScore),
			CreatedAt: timestamppb.New(match.CreatedAt),
			UpdatedAt: timestamppb.New(match.UpdatedAt),
		})
	}

	return &proto.ListMatchesResponse{Matches: matchResponses}, nil
}

func main() {
	// Set up the database connection
	db, err := sql.Open("postgres", "postgres://postgres:AidosK05@localhost:5433/epl_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewMatchRepository(db)

	createUseCase := &usecases.CreateMatch{Repository: repo}
	getUseCase := &usecases.FindMatchByID{Repository: repo}
	updateUseCase := &usecases.UpdateMatch{Repository: repo}
	deleteUseCase := &usecases.DeleteMatch{Repository: repo}
	listUseCase := &usecases.FindAllMatches{Repository: repo}

	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMatchServiceServer(grpcServer, &matchServiceServer{
		createUseCase: createUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
		listUseCase:   listUseCase,
	})

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
