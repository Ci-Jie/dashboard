package pod

import (
	"log"

	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	"github.com/kubernetes/dashboard/src/app/backend/resource/event"
	client "k8s.io/client-go/kubernetes"
)

// GetEventsForPod gets events that are associated with this pod.
func GetEventsForPod(client client.Interface, dsQuery *dataselect.DataSelectQuery, namespace,
	podName string) (*common.EventList, error) {
	
	podEvents, err := event.GetPodEvents(client, namespace, podName)
	if err != nil {
		return nil, err
	}

	if !event.IsTypeFilled(podEvents) {
		podEvents = event.FillEventsType(podEvents)
	}

	events := event.CreateEventList(podEvents, dsQuery)

	log.Printf("Found %d events related to %s pod in %s namespace", len(events.Events), podName,
		namespace)

	return &events, nil
}
