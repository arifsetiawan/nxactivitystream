package nxactivitystream

// ActivityStore is
type ActivityStore interface {
	Create() error
	Remove() error
	Subscribe() error
	Unsubscribe() error
	TopicFeeds() ([]Activity, error)
	UserFeeds() ([]Activity, error)
	HomeFeeds() ([]Activity, error)
}
