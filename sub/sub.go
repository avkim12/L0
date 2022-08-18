package sub

import (
	"encoding/json"
	"log"

	"github.com/avkim12/L0/cache"
	"github.com/avkim12/L0/models"
	"github.com/avkim12/L0/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "event-store"
	channelID = "channel"
)

func Subscribe(db *postgres.OrderDB, cache *cache.Cache) {

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}

	_, err = sc.Subscribe(channelID, func(m *stan.Msg) {

		var model models.Order

		err := json.Unmarshal(m.Data, &model)
		if err != nil {
			log.Fatal(err)
			return
		}

		v := validator.New()
		err = v.Struct(model)
		if err != nil {
			log.Fatal(err)
			return
		}

		order := postgres.Order{
			UID:   model.OrderUID,
			Model: m.Data,
		}

		err = db.CreateOrder(order)
		if err != nil {
			log.Fatal(err)
			return
		}

		cache.Set(model.OrderUID, m.Data)

	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}
}
