package queue_adapter

import (
	logging "trip_scheduler/logger"

	"github.com/go-stomp/stomp"
)

type Logger interface {
	Log(string)
	LogError(string)
}

type QueueAdapter struct {
	logger  Logger
	address string
	conn    *stomp.Conn
}

func (queueAdapter *QueueAdapter) Init(logger logging.Logger) {
	queueAdapter.logger = logger
	queueAdapter.address = "localhost:61613"
}

func (queueAdapter *QueueAdapter) Connect() {
	conn, err := stomp.Dial("tcp", queueAdapter.address)
	if err != nil {
		queueAdapter.logger.LogError("Failed dialing to " + queueAdapter.address)
	}

	queueAdapter.conn = conn
}

func (queueAdapter *QueueAdapter) SendString(message string, dest string) {
	err := queueAdapter.conn.Send(
		dest,            // destination
		"text/plain",    // content-type
		[]byte(message), // body
		stomp.SendOpt.Receipt,
		stomp.SendOpt.Header("expires", "2049-12-31 23:59:59"))

	if err != nil {
		queueAdapter.logger.LogError("Failed sending " + message + " to " + dest)
	}
}

func (queueAdapter *QueueAdapter) Disconnect() {
	queueAdapter.conn.Disconnect()
}
