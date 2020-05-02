package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	exampleSugarLog()

	exampleLog()
}

func exampleSugarLog() {
	sugar := zap.NewExample().Sugar()
	defer func() {
		_ = sugar.Sync()
	}()

	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

func exampleLog() {
	logger := zap.NewExample()
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}
