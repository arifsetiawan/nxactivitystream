package nxactivitystream

import (
	"errors"

	"github.com/couchbase/gocb"
)

type CouchbaseStore struct {
	Bucket *gocb.Bucket
}

func NewCouchbaseStore(hostName string, bucketName string, password string) (*CouchbaseStore, error) {
	cluster, err := gocb.Connect(hostName)
	if err != nil {
		return nil, errors.New("Can't connect to " + hostName)
	}

	bucket, err := cluster.OpenBucket(bucketName, password)
	if err != nil {
		return nil, errors.New("Invalid bucket password!")
	}

	return &CouchbaseStore{
		Bucket: bucket,
	}, nil
}

// Create an Activity
func (c *CouchbaseStore) Create(a Activity) error {

	if _, err := c.Bucket.Insert(a.ID, a, 0); err != nil {
		return err
	}

	return nil
}

// Feeds to a Topic or Users
func (c *CouchbaseStore) Feeds(key string) (Collection, error) {

}
