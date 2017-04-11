package utils

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"fmt"
	"github.com/containous/traefik/log"
	"golang.org/x/net/context"
	"os"
)

var (
	topicUpdateCorporation *pubsub.Topic
	topicUpdateCharacter   *pubsub.Topic
	topicUpdateAlliance    *pubsub.Topic
)

func InitPubSub() (*pubsub.Client, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT1"))
	if err != nil {
		log.Fatal(err)
	}

	topicUpdateCorporation, _ = client.CreateTopic(ctx, "UPDATE_CORPORATION")
	createSubscription(client, ctx, topicUpdateCorporation, "update_corporation")

	topicUpdateAlliance, _ = client.CreateTopic(ctx, "UPDATE_ALLIANCE")
	createSubscription(client, ctx, topicUpdateAlliance, "update_alliance")

	topicUpdateCharacter, _ = client.CreateTopic(ctx, "UPDATE_CHARACTER")
	createSubscription(client, ctx, topicUpdateCharacter, "update_character")

	return client, nil
}

func createSubscription(client *pubsub.Client, ctx context.Context, topic *pubsub.Topic, route string) {
	url := fmt.Sprintf("http://localhost:8000/api/pubsub/%v/", route)

	client.CreateSubscription(
		ctx,
		"MAIN",
		topic,
		0,
		&pubsub.PushConfig{Endpoint: url},
	)
}

func UpdateCorporation(id int) {
	msg, _ := json.Marshal(id)

	publishMessage(msg, topicUpdateCorporation)
}

func UpdateAlliance(id int) {
	msg, _ := json.Marshal(id)

	publishMessage(msg, topicUpdateAlliance)
}

func UpdateCharacter(id int) {
	msg, _ := json.Marshal(id)

	publishMessage(msg, topicUpdateCharacter)
}

func publishMessage(msg []byte, topic *pubsub.Topic) error {
	ctx := context.Background()

	r := topic.Publish(ctx, &pubsub.Message{Data: msg})

	_, err := r.Get(ctx)
	if err != nil {
		log.Printf("Error while publishing to Google PubSub: %v", err)
		return err
	}

	return nil
}
