package eventsubscriptions

import (
	"encoding/json"
	"fmt"
)

var _ EventSubscriptionDestination = WebHookEventSubscriptionDestination{}

type WebHookEventSubscriptionDestination struct {
	Properties *WebHookEventSubscriptionDestinationProperties `json:"properties,omitempty"`

	// Fields inherited from EventSubscriptionDestination
}

var _ json.Marshaler = WebHookEventSubscriptionDestination{}

func (s WebHookEventSubscriptionDestination) MarshalJSON() ([]byte, error) {
	type wrapper WebHookEventSubscriptionDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebHookEventSubscriptionDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebHookEventSubscriptionDestination: %+v", err)
	}
	decoded["endpointType"] = "WebHook"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebHookEventSubscriptionDestination: %+v", err)
	}

	return encoded, nil
}