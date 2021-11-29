package eventsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderTopicTypeId{}

// ProviderTopicTypeId is a struct representing the Resource ID for a Provider Topic Type
type ProviderTopicTypeId struct {
	SubscriptionId string
	TopicTypeName  string
}

// NewProviderTopicTypeID returns a new ProviderTopicTypeId struct
func NewProviderTopicTypeID(subscriptionId string, topicTypeName string) ProviderTopicTypeId {
	return ProviderTopicTypeId{
		SubscriptionId: subscriptionId,
		TopicTypeName:  topicTypeName,
	}
}

// ParseProviderTopicTypeID parses 'input' into a ProviderTopicTypeId
func ParseProviderTopicTypeID(input string) (*ProviderTopicTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderTopicTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderTopicTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.TopicTypeName, ok = parsed.Parsed["topicTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'topicTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviderTopicTypeIDInsensitively parses 'input' case-insensitively into a ProviderTopicTypeId
// note: this method should only be used for API response data and not user input
func ParseProviderTopicTypeIDInsensitively(input string) (*ProviderTopicTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderTopicTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderTopicTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.TopicTypeName, ok = parsed.Parsed["topicTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'topicTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviderTopicTypeID checks that 'input' can be parsed as a Provider Topic Type ID
func ValidateProviderTopicTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderTopicTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Topic Type ID
func (id ProviderTopicTypeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.EventGrid/topicTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.TopicTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Topic Type ID
func (id ProviderTopicTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftEventGrid", "Microsoft.EventGrid", "Microsoft.EventGrid"),
		resourceids.StaticSegment("staticTopicTypes", "topicTypes", "topicTypes"),
		resourceids.UserSpecifiedSegment("topicTypeName", "topicTypeValue"),
	}
}

// String returns a human-readable description of this Provider Topic Type ID
func (id ProviderTopicTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Topic Type Name: %q", id.TopicTypeName),
	}
	return fmt.Sprintf("Provider Topic Type (%s)", strings.Join(components, "\n"))
}