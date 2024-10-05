package tcp

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/volvofixthis/pow-server/internal/core/ports"
	"github.com/volvofixthis/pow-server/internal/infra/config"
	"go.uber.org/zap"
)

// TCPConnWrapper wraps a TCP connection to implement Conn interface
type TCPConnWrapper struct {
	conn net.Conn
}

func (w *TCPConnWrapper) Read(p []byte) (int, error) {
	return w.conn.Read(p)
}

func (w *TCPConnWrapper) Write(p []byte) (int, error) {
	return w.conn.Write(p)
}

func (w *TCPConnWrapper) Close() error {
	return w.conn.Close()
}

type Worker struct {
	id          int
	jobQueue    <-chan net.Conn // Queue of connections to process
	ctx         context.Context
	wg          *sync.WaitGroup
	connAdapter ports.ConnAdapter
	c           *config.AppConfig
	log         *zap.Logger
}

func (w *Worker) Start() {
	for conn := range w.jobQueue {
		w.handleConnection(conn)
	}
}

// handleConnection processes the connection
func (w *Worker) handleConnection(conn net.Conn) {
	ctx, cancel := context.WithCancel(w.ctx)
	defer cancel()

	w.log.Info(
		"Trying to set connection deadlines",
		zap.Duration("readTimeout", w.c.ConReadTimeout),
		zap.Duration("writeTimeout", w.c.ConWriteTimeout),
	)
	if err := conn.SetReadDeadline(time.Now().Add(w.c.ConReadTimeout)); err != nil {
		w.log.Error("Unable to set read duration", zap.Error(err))
		return
	}
	if err := conn.SetWriteDeadline(time.Now().Add(w.c.ConWriteTimeout)); err != nil {
		w.log.Error("Unable to set read duration", zap.Error(err))
		return
	}

	w.log.Info("Received new connection")
	if err := w.connAdapter.Handle(ctx, &TCPConnWrapper{conn: conn}); err != nil {
		w.log.Error("Error when handling connection", zap.Error(err))
	}
}

func NewTCPServer(
	log *zap.Logger,
	config *config.AppConfig, connAdapter ports.ConnAdapter,
) *TCPServer {
	server := &TCPServer{
		address:        config.TCPAddress,
		jobQueue:       make(chan net.Conn, config.ConQueue), // Buffered channel to queue connections
		shutdownCh:     make(chan struct{}),                  // Shutdown channel
		workerPoolSize: config.WorkerPool,
		wg:             &sync.WaitGroup{},
		connAdapter:    connAdapter,
		c:              config,
		log:            log,
	}

	return server
}

type TCPServer struct {
	c              *config.AppConfig
	address        string
	jobQueue       chan net.Conn // Queue of incoming connections
	wg             *sync.WaitGroup
	shutdownCh     chan struct{} // Channel to signal server shutdown
	workerPoolSize int
	connAdapter    ports.ConnAdapter
	log            *zap.Logger
}

func StartServer(s *TCPServer) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.log.Info("Spawning workers", zap.Int("workerPoolSize", s.workerPoolSize))
	for i := 1; i <= s.workerPoolSize; i++ {
		worker := &Worker{
			id:          i,
			jobQueue:    s.jobQueue,
			ctx:         ctx,
			wg:          s.wg,
			connAdapter: s.connAdapter,
			c:           s.c,
			log:         s.log,
		}
		s.wg.Add(1)
		go worker.Start()
	}

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		s.log.Error("Failed to start TCP server", zap.Error(err))
		return
	}
	defer listener.Close()

	// Main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-s.shutdownCh:
				close(s.jobQueue) // Close the job queue to stop the workers
				return
			default:
				continue
			}
		}

		// Add connection to job queue for workers to pick up
		s.jobQueue <- conn
	}
}

func StopServer(s *TCPServer) error {
	s.shutdownCh <- struct{}{}
	return nil
}
