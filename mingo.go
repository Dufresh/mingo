package mingo

import (
	"time"
)

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

// Conn represents a connection to server.
type Conn interface {
	// MarkIdleTime mark time becoming idle.
	MarkIdleTime()

	// GetIdleTime get time becoming idle.
	GetIdleTime() time.Time

	// Close closes the connection.
	Close() error

	// Err returns a non-nil value when the connection is not usable.
	Err() error

	// Post a command to the server and returns the received response.
	Post(request string, args ...interface{}) (response interface{}, err error)

	// Send writes the command to the client's output buffer.
	Send(command string, args ...interface{}) error

	// Flush flushes the output buffer to the server.
	Flush() error

	// Receive receives a single reply from the server
	Receive() (reply interface{}, err error)
}

// Scanner is implemented by types which want to control how their value is
// interpreted when read from server.
type Scanner interface {
	// An error should be returned if the value cannot be stored without
	// loss of information.
	RedisScan(src interface{}) error
}
