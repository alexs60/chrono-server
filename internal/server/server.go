package server

import (
	"bufio"
	"chrono/internal/store" // Import our new store package
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

// Server holds the dependencies for our TCP server, primarily the data store.
type Server struct {
	store *store.Store
}

// NewServer creates a new Server instance.
func NewServer(s *store.Store) *Server {
	return &Server{store: s}
}

// Start listens on the given address and begins accepting connections.
func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		// Handle each connection in a new goroutine.
		go s.handleConnection(conn)
	}
}

// handleConnection manages a single client connection from start to finish.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Read error:", err)
			}
			return
		}

		parts := strings.Fields(strings.TrimSpace(cmdString))
		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])
		args := parts[1:]

		switch command {
		case "SET":
			if len(args) < 3 {
				conn.Write([]byte("ERROR: SET requires a key, a value, and a TTL in seconds\n"))
				continue
			}
			key := args[0]
			value := args[1]
			ttl, err := strconv.Atoi(args[2])
			if err != nil {
				conn.Write([]byte("ERROR: TTL must be an integer\n"))
				continue
			}
			// Use the store's Set method.
			s.store.Set(key, value, ttl)
			conn.Write([]byte("OK\n"))

		case "GET":
			if len(args) < 1 {
				conn.Write([]byte("ERROR: GET requires a key\n"))
				continue
			}
			key := args[0]
			// Use the store's Get method.
			value, ok := s.store.Get(key)
			if !ok {
				conn.Write([]byte("(nil)\n"))
			} else {
				conn.Write([]byte(value + "\n"))
			}

		default:
			conn.Write([]byte("ERROR: Unknown command '" + command + "'\n"))
		}
	}
}
