// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdinreceiver

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

// This file implements factory for stdin receiver.

const (
	// The value of "type" key in configuration.
	typeStr        = "stdin"
	stabilityLevel = component.StabilityLevelDevelopment
)

type Config struct {
}

// NewFactory creates a factory for stdin receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		component.MustNewType(typeStr),
		createDefaultConfig,
		receiver.WithLogs(newLogsReceiver, stabilityLevel))
}

// CreateDefaultConfig creates the default configuration for stdin receiver.
func createDefaultConfig() component.Config {
	return &Config{}
}
