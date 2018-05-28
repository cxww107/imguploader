package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	mongoUserSecret = "MONGO_USERNAME"
	mongoPassSecret = "MONGO_PASSWORD"
)

// Mongoer is interface to Mongo client
type Mongoer interface {
}

// Mongo is a struct to mongodb operations
type Mongo struct {
	DBName string
	*mgo.Session
}

// NewMongo returns mongodb handler with custom functions
func NewMongo(addrs string, database string, ssl bool, secretPath string) (*Mongo, error) {
	m := &Mongo{DBName: database}

	if ssl {
		di, err := getSSLDialInfo(addrs, database)
		if err != nil {
			return nil, fmt.Errorf("not able reach mongo db at: %s \n %v", addrs, err)
		}

		if m.Session, err = mgo.DialWithInfo(di); err != nil {
			return nil, fmt.Errorf("not able reach mongo db at: %s \n %v", addrs, err)
		}
	} else {
		u, err := getFromSecret(mongoUserSecret, secretPath)
		if err != nil {
			return nil, fmt.Errorf("error reading docker secret %s, %v", mongoUserSecret, err)
		}

		pass, err := getFromSecret(mongoPassSecret, secretPath)
		if err != nil {
			return nil, fmt.Errorf("error reading docker secret %s, %v", mongoUserSecret, err)
		}

		log.Printf("Dial to mongo: %v %v %v", addrs, u, pass)
		m.Session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    strings.Split(addrs, ","),
			Username: u,
			Password: pass})

		if err != nil {
			return nil, fmt.Errorf("not able reach mongo db at: %s \n %v", addrs, err)
		}
	}

	return m, nil
}

// InsertOneFile writes a struct to mongodb
func (m *Mongo) InsertOneFile(ctx context.Context, collection string, doc []byte) error {
	sess := m.Session.Copy()
	defer sess.Close()

	f, err := sess.DB(m.DBName).GridFS(collection).Create("")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(doc)

	return err
}

// FindIDs is method
func (m *Mongo) FindIDs(ctx context.Context, collection string, pipeline []bson.M) []bson.ObjectId {
	sess := m.Session.Copy()
	defer sess.Close()

	c := sess.DB(m.DBName).C(collection)

	iter := c.Pipe(pipeline).Iter()
	defer iter.Close()

	type ID struct {
		ID bson.ObjectId `bson:"_id,omitempty"`
	}
	var r ID
	res := make([]bson.ObjectId, 0)
	for iter.Next(&r) {
		res = append(res, r.ID)
	}

	return res
}

// Get key from docker secret
func getFromSecret(secret, secretPath string) (string, error) {
	p := filepath.Join(secretPath, secret)
	d, err := ioutil.ReadFile(p)
	if err != nil {
		return "", fmt.Errorf("failed to read docker secret at %s, %v", secretPath, err)
	}

	return string(d), nil
}
