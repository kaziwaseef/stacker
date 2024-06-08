package util

import (
	"context"
	"log"
)

type logger struct {
	ctx context.Context
}

func Logger(ctx context.Context) *logger {
	return &logger{ctx: ctx}
}

func (l *logger) Log(value string) {
	if IsVerbose(l.ctx) {
		log.Println(value)
	}
}

func (l *logger) Fatal(value string) {
	if IsVerbose(l.ctx) {
		log.Fatalf(value)
	}
}
