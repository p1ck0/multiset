package log_test

import (
	"context"
	"testing"

	"github.com/p1ck0/multiset/log"
	"github.com/stretchr/testify/assert"
)

func TestLogCtx(t *testing.T) {
	logger := log.New(true)

	ctx := log.LoggerWithContext(context.Background(), logger)

	logger2 := log.LoggerFromContext(ctx)

	assert.Equal(t, logger, logger2)
}
