// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdinreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestCreateReceiver(t *testing.T) {
	cfg := createDefaultConfig().(*Config)

	mockLogsConsumer := consumertest.NewNop()
	lReceiver, err := newLogsReceiver(context.Background(), receivertest.NewNopSettings(), cfg, mockLogsConsumer)
	assert.Nil(t, err, "receiver creation failed")
	assert.NotNil(t, lReceiver, "receiver creation failed")
}
