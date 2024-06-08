package util

import (
	"context"
	"log"

	"github.com/kaziwaseef/stacker/internal/types"
)

type logger struct {
	ctx context.Context
}

func Logger(ctx context.Context) *logger {
	return &logger{ctx: ctx}
}

func (l *logger) Log(value string) {
	v := l.ctx.Value(types.Verbose)
	if val, ok := v.(bool); ok && val {
		log.Println(value)
	}
}

func (l *logger) Fatal(value string) {
	v := l.ctx.Value("verbose")
	if val, ok := v.(bool); ok && val {
		log.Fatalf(value)
	}
}
