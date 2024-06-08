package util

import (
	"context"

	"github.com/kaziwaseef/stacker/internal/types"
)

func IsVerbose(ctx context.Context) bool {
	v := ctx.Value(types.Verbose)
	val, ok := v.(bool)
	return ok && val
}
