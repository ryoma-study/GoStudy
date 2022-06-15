package test

import (
	"fmt"
	"testing"
	"time"
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

// Connect creates a connection.
func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

func TestOptions(t *testing.T) {
	connection, _ := NewConnect("127.0.0.1:3000")
	fmt.Printf("%v\n", connection)

	connection, _ = NewConnect("127.0.0.1:3000", WithCaching(true), WithTimeout(20))
	fmt.Printf("%v\n", connection)
}

//const (
//	defaultTimeout = 10
//	defaultCaching = false
//)
//
//type Connection struct {
//	addr    string
//	cache   bool
//	timeout time.Duration
//}
//
//type ConnectionOptions struct {
//	Caching bool
//	Timeout time.Duration
//}
//
//func NewDefaultOptions() *ConnectionOptions {
//	return &ConnectionOptions{
//		Caching: defaultCaching,
//		Timeout: defaultTimeout,
//	}
//}
//
//// NewConnect creates a connection with options.
//func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error) {
//	return &Connection{
//		addr:    addr,
//		cache:   opts.Caching,
//		timeout: opts.Timeout,
//	}, nil
//}
//
//func TestOptions(t *testing.T) {
//	defaultOptions := NewDefaultOptions()
//	connection, _ := NewConnect("127.0.0.1:3000", defaultOptions)
//	fmt.Printf("%v\n", connection)
//
//	options := &ConnectionOptions{Caching: true, Timeout: 20}
//	connection, _ = NewConnect("127.0.0.1:3000", options)
//	fmt.Printf("%v\n", connection)
//}

//const (
//	defaultTimeout = 10
//	defaultCaching = false
//)
//
//type Connection struct {
//	addr    string
//	cache   bool
//	timeout time.Duration
//}
//
//// NewConnect creates a connection.
//func NewConnect(addr string) (*Connection, error) {
//	return &Connection{
//		addr:    addr,
//		cache:   defaultCaching,
//		timeout: defaultTimeout,
//	}, nil
//}
//
//// NewConnectWithOptions creates a connection with options.
//func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
//	return &Connection{
//		addr:    addr,
//		cache:   cache,
//		timeout: timeout,
//	}, nil
//}
//
//func TestOptions(t *testing.T) {
//	connection, _ := NewConnect("127.0.0.1:3000")
//	fmt.Printf("%v\n", connection)
//
//	connection, _ = NewConnectWithOptions("127.0.0.1:3000", true, 20)
//	fmt.Printf("%v\n", connection)
//}
