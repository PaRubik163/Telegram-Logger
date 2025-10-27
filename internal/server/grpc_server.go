package server

import (
	"context"
	"fmt"
	"net"
	"teleglogger/internal/bot"
	pb "teleglogger/pkg/api/logger"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct{
	pb.UnimplementedTelegramLoggerServer
	bot *bot.Bot
	s *grpc.Server
}

func NewServer(bot *bot.Bot) (*Server) {
	return &Server{
		bot: bot,
		s: grpc.NewServer(),
	}
}

func (s *Server) Run(addr string) (error) {
	listner, err := net.Listen("tcp", ":"+addr)
	if err != nil{
		return fmt.Errorf("failed to listen server %v", err)
	}

	pb.RegisterTelegramLoggerServer(s.s, &Server{})

	logrus.Info("Start listen server")
	if err := s.s.Serve(listner); err != nil{
		return fmt.Errorf("failed to serve server %v", err)
	}

	return nil
}

func (s *Server) SendLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	if err := s.bot.Send(req.Topic, req.Level, req.Text); err != nil{
		return &pb.LogResponse{Ok: false}, err
	}

	return &pb.LogResponse{Ok: true}, nil
}