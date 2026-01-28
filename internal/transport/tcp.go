package transport

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strings"

	"github.com/FelipeFelipeRenan/govoid/internal/engine"
)

type Server struct {
	listenAddr string
	store      *engine.StringStore
	logger     *slog.Logger
}

func New(addr string, store *engine.StringStore, logger *slog.Logger) *Server {
	return &Server{
		listenAddr: addr,
		store:      store,
		logger:     logger,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("falha ao iniciar TCP listener: %w", err)
	}

	defer listener.Close()

	s.logger.Info("GoVoid listening", "addr", s.listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("Erro no accept", "error", err)
			continue
		}

		go s.handleConn(conn)
	}

}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				s.logger.Error("erro de leitura", "remote_addr", conn.RemoteAddr(), "error", err)
			}
			return
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		response := s.executeCommand(line)

		_, _ = conn.Write([]byte(response + "\n"))
	}
}

func (s *Server) executeCommand(input string) string {
	parts := strings.Fields(input)
	if len(parts) == 0{
		return "ERR empty command"
	}

	cmd := strings.ToUpper(parts[0])

	switch cmd {
	case "SET":
		if len(parts) < 3{
			return "ERR usage: SET key value"
		}
		s.store.Set(parts[1], parts[2])
		return "OK"
	
	case "GET":
		if len(parts) < 2{
			return "ERR usage: GET key"
		}

		val, ok := s.store.Get(parts[1])
		if !ok {
			return "(nil)"
		}
		return val
	default:
		return "ERR unknown command"
	}
}
