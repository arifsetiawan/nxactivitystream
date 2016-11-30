package nxactivitystream

type ActivityStore interface {
	Create() error
	Remove() error
	TopicFeeds() ([]Activity, error)
	UserFeeds() ([]Activity, error)
}
