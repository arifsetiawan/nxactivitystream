package nxactivitystream

import "time"

type Object struct {
}

type Link struct {
}

// Activity objects are specializations of the base Object type that provide information about actions that have either already occurred, are in the process of occurring, or may occur in the future.
type Activity struct {
	ID        string      `json:"id"`
	Topic     string      `json:"topic,omitempty"`
	Type      string      `json:"type,omitempty"`
	Name      string      `json:"name,omitempty"`
	Actor     *BaseObject `json:"actor,omitempty"`
	Target    *BaseObject `json:"target,omitempty"`
	Object    *BaseObject `json:"object,omitempty"`
	Published time.Time   `json:"published, omitempty"`
}

type IntransitiveActivity struct {
}

type Collection struct {
}

type OrderedCollection struct {
}

type CollectionPage struct {
}

type OrderedCollectionPage struct {
}

// BaseObject is represent of general object
type BaseObject struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url,omitempty"`
	ObjectType  string                 `json:"object_type,omitempty"`
	DisplayName string                 `json:"display_name,omitempty"`
	Content     string                 `json:"content,omitempty"`
	MetaData    map[string]interface{} `json:"meta_data,omitempty"`
}

// Subscription is
type Subscription struct {
	ID     string      `json:"id"`
	Actor  *BaseObject `json:"actor,omitempty"`
	Topics []string    `json:"topics,omitempty"`
}
