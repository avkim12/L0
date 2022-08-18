package pub

import (
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "event-store"
	channelID = "channel"
)

func Publish() {

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	err = sc.Publish(channelID, []byte(validModel))
	if err != nil {
		panic(err)
	}
}
