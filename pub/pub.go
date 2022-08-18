package pub

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-client"
	channelID = "channel"
)

func Publish() {

	uid := uuid.New()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	msg := fmt.Sprintf(validModel, uid.String())

	err = sc.Publish(channelID, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
