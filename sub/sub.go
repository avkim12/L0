package sub

import (
	"encoding/json"

	"github.com/avkim12/L0/model"
	"github.com/avkim12/L0/postgres"
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

	var model model.Order
	sub, err := sc.Subscribe(channelID, func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &model)
		if err != nil {
			panic(err)
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		panic(err)
	}

	postgres.CreateOrder(model)

	err = sub.Unsubscribe()
	if err != nil {
		panic(err)
	}
}