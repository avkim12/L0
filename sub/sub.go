package sub

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "event-store"
	channelID = "channel"
)

func Subscribe() {

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe(channelID, func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		// err := json.Unmarshal(m.Data, &order)
		if err != nil {
			panic(err)
		}
	}, stan.StartWithLastReceived())

	err = sub.Unsubscribe()
	if err != nil {
		panic(err)
	}
}
