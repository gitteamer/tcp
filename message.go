package tcp

import "time"

const (
	maxMessageLength = 2048 // max length
	//heartbeat        = 0    // heartbeat package
	readTimeout  = time.Second * 3
	writeTimeout = time.Second * 3
)

/**
 * receive, send struct
 */
type Message struct {
	Type uint
	Data []byte
}
