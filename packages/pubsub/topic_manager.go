package pubsub

import (
	"errors"
	"fmt"
	"log"
)
import "sync"

type topics struct {
	topicsMap  *sync.Map
}

func (p *PubSub) CreateTopic(topicName string) (bool, error) {
	_, ok := p.topics.topicsMap.Load(topicName)

	if ok {
		fmt.Println("topic Already Exists", topicName)
		return false, errors.New("topic already exists")
	} else {
		p.topics.topicsMap.Store(topicName, &topic{topicId: topicName, subscriptions: *new(sync.Map)})
		return true, nil
	}
}

func (p *PubSub) DeleteTopic(TopicID string) (bool, error) {
	log.Printf("Deleting topic: %q", TopicID)
	topicVar, ok := p.topics.topicsMap.Load(TopicID)
	if ok {
		log.Printf("Deleting all subscriptions for topic: %q", TopicID)
		// first removing all subscriptions of that Topic from subscriptionTopicsMap
		topicVar.(*topic).subscriptions.Range(func(key, _ interface{}) bool {
			// deleting all subscription before deleting topic
			_, err := p.DeleteSubscription(key.(string))
			print(err)
			return true
		})
		log.Printf("All subscriptions for topic: %q deleted", TopicID)
		//then deleting topic from topicsMap
		p.topics.topicsMap.Delete(TopicID)

		log.Printf("DeleteTopic()-> topic: %q deleted \n", TopicID)
		return true, nil
	} else {
		log.Printf("DeleteTopic()-> TopicID %q don't exist \n", TopicID)
		return false, errors.New("TopicID don't exist")
	}

}
