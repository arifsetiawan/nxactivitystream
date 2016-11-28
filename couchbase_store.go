package nxactivitystream

import (
	"errors"

	"github.com/couchbase/gocb"
)

// CouchbaseStore is a representation of Couchbase Obj
type CouchbaseStore struct {
	Bucket *gocb.Bucket
}

// NewCouchbaseStore is a function to Initiate Couchbase store
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

// TopicFeeds feeds to a Topic or Users
func (c *CouchbaseStore) TopicFeeds(fType string, limit int, offset int, topicID string) ([]Activity, error) {

	extras := ""
	var params []interface{}
	params = append(params, topicID)
	params = append(params, limit)
	params = append(params, offset)

	if len(fType) > 0 {
		extras = "AND a.type = $4"
		params = append(params, fType)
	}

	q := gocb.NewN1qlQuery("SELECT a.* FROM nextflow a WHERE a.topic = $1 " + extras + " ORDER BY a.published DESC LIMIT $2 OFFSET $3")
	rows, err := c.Bucket.ExecuteN1qlQuery(q, params)

	if err != nil {
		return nil, err
	}

	var ac Activity
	var acs []Activity
	for i := 0; rows.Next(&ac); i++ {
		acs = append(acs, ac)
	}
	rows.Close()

	return acs, nil
}

// UserFeeds feeds to a User Stream
func (c *CouchbaseStore) UserFeeds(fType string, limit int, offset int, userID string) ([]Activity, error) {

	extras := ""
	var params []interface{}
	params = append(params, userID)
	params = append(params, limit)
	params = append(params, offset)

	if len(fType) > 0 {
		extras = "AND a.type = $4"
		params = append(params, fType)
	}

	q := gocb.NewN1qlQuery("SELECT a.* FROM nextflow a WHERE a.actor.id = $1 " + extras + " ORDER BY a.published DESC LIMIT $2 OFFSET $3")
	rows, err := c.Bucket.ExecuteN1qlQuery(q, params)

	if err != nil {
		return nil, err
	}

	var ac Activity
	var acs []Activity
	for i := 0; rows.Next(&ac); i++ {
		acs = append(acs, ac)
	}
	rows.Close()

	return acs, nil
}
