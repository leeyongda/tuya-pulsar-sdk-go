package pulsarx

import (
	"fmt"
	"strings"
)

const (
	env_test = "event-test"
	env_prod = "event"
)

func getTenant(topic string) string {
	topic = strings.TrimPrefix(topic, "persistent://")
	end := strings.Index(topic, "/")
	return topic[:end]
}

func GetSubscriptionName(topic string) string {
	return getTenant(topic) + "-sub"
}

func GetTopicForAccessID(accessID string) string {
	topic := fmt.Sprintf("persistent://%s/out/%s", accessID, env_prod)
	return topic
}

func GetTopicTestForAccessID(accessID string) string {
	topic := fmt.Sprintf("persistent://%s/out/%s", accessID, env_test)
	return topic
}
