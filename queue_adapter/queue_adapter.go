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
		queueAdapter.logger.LogError("Failad dialing to " + queueAdapter.address)
		return
	}

	queueAdapter.conn = conn
}

func (queueAdapter *QueueAdapter) SendString(message string, dest string) {
	err := queueAdapter.conn.Send(
		dest,            // destination
		"text/plain",    // content-type
		[]byte(message)) // body

	if err != nil {
		queueAdapter.logger.LogError("Failed sending " + message + " to " + dest)
		return
	}
}

func (queueAdapter *QueueAdapter) Disconnect() {
	queueAdapter.conn.Disconnect()
	queueAdapter.logger.Log("Connection disconnected from " + queueAdapter.address)
}
