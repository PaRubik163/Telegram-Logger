package server

import (
	"context"
	"fmt"
	"net"
	"encoding/json"
	"github.com/PaRubik163/Telegram-Logger/internal/bot"
	pb "github.com/PaRubik163/Telegram-Logger/pkg/api/logger"

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

	pb.RegisterTelegramLoggerServer(s.s, s)

	logrus.Info("Start listen server")
	if err := s.s.Serve(listner); err != nil{
		return fmt.Errorf("failed to serve server %v", err)
	}

	return nil
}

func (s *Server) SendLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	if req.Data == nil{
		return &pb.LogResponse{Ok: false}, fmt.Errorf("missing Data filed in request")
	}

	dataBytes, err := json.MarshalIndent(req.Data, "", "  ")
	if err != nil {
		return &pb.LogResponse{Ok: false}, fmt.Errorf("failed to marshal log data: %v", err)
	}
	go func (topic, level, text string){
		if err := s.bot.Send(topic, level, text); err != nil{
			 fmt.Printf("Failed to send message %v ", err)
		}
	}(req.Topic, req.Level, string(dataBytes))

	return &pb.LogResponse{Ok: true}, nil
}