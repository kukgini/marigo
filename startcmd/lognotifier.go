package startcmd

import (
	"log"
)

type LogNotifier struct {
}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}

func (n *LogNotifier) Notify(topic string, message []byte) error {
	log.Printf("topic=%s, message=%s", topic, message)
	return nil
}
