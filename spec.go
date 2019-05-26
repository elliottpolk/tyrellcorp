package tyrellcorp

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const repo string = "tyrellcorp_specs"

// Create ...
func Create(ctx context.Context, items []*Spec, db *mongo.Database) error {
	// swap to []interface{} because mongo needs it
	in := make([]interface{}, len(items))

	// validate / enrich required fields
	for i, item := range items {
		// the provider must specify at least the CreatedBy value
		if item.RecordInfo == nil {
			return ErrInvalidRecordInfo
		}

		// verify created_by is populated for an attempt at an audit
		if len(item.RecordInfo.CreatedBy) < 1 {
			return ErrInvalidCreatedBy
		}

		// ensure the Spec has an unique identifier
		if len(item.Id) < 1 {
			item.Id = uuid.New().String()
		}

		// ensure the created value is populated
		if item.RecordInfo.Created == nil || item.RecordInfo.Created.Seconds < 1 {
			item.RecordInfo.Created = &timestamp.Timestamp{Seconds: time.Now().Unix()}
		}

		in[i] = item
	}

	// write Spec to datastore
	if _, err := db.Collection(repo).InsertMany(ctx, in); err != nil {
		return err
	}

	// return the written element
	return nil
}

// RetrieveOne ...
func RetrieveOne(ctx context.Context, id string, db *mongo.Database) (*Spec, error) {
	res, err := Retrieve(ctx, bson.D{{"_id", id}}, db)
	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return nil, ErrNotFound
	}

	if len(res) > 1 {
		return nil, ErrMutlipleRecordsReturned
	}

	return res[0], nil
}

// TODO:
// - include limit on record retrieval
// - include pagination on record retrieval

// Retrieve ...
func Retrieve(ctx context.Context, filter bson.D, db *mongo.Database) ([]*Spec, error) {
	iter, err := db.Collection(repo).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer iter.Close(ctx)

	items := make([]*Spec, 0)
	for iter.Next(ctx) {
		item := &Spec{}
		if err := iter.Decode(&item); err != nil {
			return nil, errors.Wrapf(err, "unable to decode record")
		}
		items = append(items, item)
	}

	return items, nil
}

// Update ...
func Update(ctx context.Context, user string, spec *Spec, db *mongo.Database) error {
	// ensure the user provided a username in an attempt to audit
	if len(user) < 1 {
		return ErrInvalidUsername
	}

	if len(spec.Id) < 1 {
		return ErrInvalidId
	}

	prev, err := RetrieveOne(ctx, spec.Id, db)
	if err != nil {
		return errors.Wrapf(err, "unable to retrieve existing record for %s", spec.Id)
	}

	if prev == nil {
		return errors.Wrapf(err, "unable to retrieve existing record for %s", spec.Id)
	}

	if spec.History == nil {
		spec.History = make([]*Spec, len(prev.History))
	}
	copy(spec.History, prev.History)

	prev.History = nil
	spec.History = append(spec.History, prev)

	// TODO:
	// - check for previous item with id
	// - append to history field if found

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "update",
	}).Infof("attempting to update record with ID %s", spec.Id)

	spec.RecordInfo.Updated = &timestamp.Timestamp{Seconds: time.Now().Unix()}
	spec.RecordInfo.UpdatedBy = user

	res, err := db.Collection(repo).ReplaceOne(ctx, bson.D{{"_id", spec.Id}}, spec)
	if res != nil {
		return err
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "update",
	}).Infof("updated record %s", spec.Id)

	return nil
}

// Delete ...
func Delete(ctx context.Context, user string, items []*Spec, db *mongo.Database) error {
	// ensure the user provided a username in an attempt to audit
	if len(user) < 1 {
		return ErrInvalidUsername
	}

	ids := make([]string, len(items))
	for i, item := range items {
		if len(item.Id) < 1 {
			return ErrInvalidId
		}
		ids[i] = item.Id
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "delete",
	}).Infof("attempting to delete %d records", len(items))

	res, err := db.Collection(repo).DeleteMany(ctx, bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return errors.Wrapf(err, "deletion: expected %d - actually %d", len(ids), res.DeletedCount)
	}

	if want, got := int64(len(ids)), res.DeletedCount; got < want {
		return errors.Wrapf(ErrIncompleteAction, "deletion: expected %d - actually %d", want, got)
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "delete",
	}).Infof("deleted %d records", res.DeletedCount)

	return nil
}
