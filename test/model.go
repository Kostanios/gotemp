package test

import "context"

type ContainerData struct {
	Url           string
	TerminateFunc func(ctx context.Context) error
}
