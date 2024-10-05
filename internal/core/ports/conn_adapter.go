package ports

import "context"

type Conn interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}

type ConnAdapter interface {
	Handle(ctx context.Context, conn Conn) error
}
