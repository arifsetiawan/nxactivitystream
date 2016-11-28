package nxactivitystream

type ActivityStore interface {
	Create() error
	TopicFeeds() ([]Activity, error)
	UserFeeds() ([]Activity, error)
}
