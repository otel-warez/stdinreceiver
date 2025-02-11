// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdinreceiver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestConsumeLine(t *testing.T) {
	sink := new(consumertest.LogsSink)
	config := createDefaultConfig()
	r := stdinReceiver{logsConsumer: sink, config: config.(*Config), done: make(chan bool, 1)}
	r.consumeLine(context.Background(), "foo")
	lds := sink.AllLogs()
	assert.Equal(t, 1, len(lds))
	log := lds[0].ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Body().Str()
	assert.Equal(t, "foo", log)
}

func TestReceiveLines(t *testing.T) {
	read, write, err := os.Pipe()
	assert.NoError(t, err)
	stdin = read
	sink := new(consumertest.LogsSink)
	config := createDefaultConfig()
	settings := receivertest.NewNopSettings()
	o, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             settings.ID,
		Transport:              "",
		ReceiverCreateSettings: settings})
	require.NoError(t, err)
	r := stdinReceiver{logsConsumer: sink, config: config.(*Config), obsrecv: o, done: make(chan bool, 1)}
	err = r.Start(context.Background(), componenttest.NewNopHost())
	assert.NoError(t, err)
	write.WriteString("foo\nbar\nfoobar\n")
	write.WriteString("foo\r\nbar\nfoobar\n")
	time.Sleep(time.Second * 1)
	lds := sink.AllLogs()
	assert.Equal(t, 6, len(lds))

	read.Chmod(000)
	write.WriteString("foo\nbar\nfoobar\n")
	time.Sleep(time.Second * 1)
	write.Close() // close stdin early.

	err = r.Shutdown(context.Background())
	assert.NoError(t, err)
}
