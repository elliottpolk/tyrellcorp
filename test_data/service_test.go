package super

import (
	"context"
	"testing"

	"github.com/elliottpolk/super/config"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO:
// - convert to a real service test
// - include a docker-compose.yml file which:
//   - starts up a mongodb instance
// 	 - runs the tests in a golang:latest container

func TestServiceCreate(t *testing.T) {
	client, clean := dbclient(t)
	defer clean(repo)

	svr := &DopeServer{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*Dope{
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "But I've never been to the moon! We need rest.",
			Dope_2:     9,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			Dope_2:     3,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			Dope_2:     1000000,
		},
	}

	reqId := uuid.New().String()
	res, err := svr.Create(context.TODO(), &DopeRequest{
		Payload:   items,
		RequestId: reqId,
	})
	if err != nil {
		t.Fatal(err)
	}

	// should be an "Empty" with the original request ID
	assert.NotEmpty(t, res)
	assert.Equal(t, reqId, res.RequestId)

	// confirm item
	got, err := Retrieve(context.TODO(), bson.D{}, client.Database(repo))
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, got, len(items))
}

func TestServiceRetrieve(t *testing.T) {
	client, clean := dbclient(t)
	defer clean(repo)

	svr := &DopeServer{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*Dope{
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "But I've never been to the moon! We need rest.",
			Dope_2:     9,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			Dope_2:     3,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			Dope_2:     1000000,
		},
	}

	// warm the test dataset
	if err := Create(context.TODO(), items, client.Database(repo)); err != nil {
		t.Fatal(err)
	}

	reqId0 := uuid.New().String()
	res0, err := svr.Retrieve(context.TODO(), &DopeRequest{
		RequestId: reqId0,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, res0)
	assert.NotNil(t, res0.Payload)
	assert.Equal(t, reqId0, res0.RequestId)
	assert.Len(t, res0.Payload, len(items))

	item := items[0]
	reqId1 := uuid.New().String()
	res1, err := svr.Retrieve(context.TODO(), &DopeRequest{
		Id:        item.Id,
		RequestId: reqId1,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, res1)
	assert.NotNil(t, res1.Payload)
	assert.Equal(t, reqId1, res1.RequestId)
	assert.Len(t, res1.Payload, 1)
}

func TestServiceUpdate(t *testing.T) {
	client, clean := dbclient(t)
	defer clean(repo)

	svr := &DopeServer{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*Dope{
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "But I've never been to the moon! We need rest.",
			Dope_2:     9,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			Dope_2:     3,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			Dope_2:     1000000,
		},
	}

	db := client.Database(repo)

	// warm the test dataset
	if err := Create(context.TODO(), items, db); err != nil {
		t.Fatal(err)
	}

	// adjust the item
	item := items[0]

	item.Dope_1 = "You are the last hope of the universe. Why not indeed! Aww, it's true. I've been hiding it for so long. Well, then good news! It's a suppository. OK, if everyone's finished being stupid."
	res0, err := RetrieveOne(context.TODO(), item.Id, db)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, item.Dope_1, res0.Dope_1)

	reqId0 := uuid.New().String()
	res1, err := svr.Update(context.TODO(), &DopeRequest{
		RequestId: reqId0,
		Payload:   []*Dope{item},
	})
	assert.Empty(t, res1)
	assert.Error(t, err)
	assert.EqualValues(t, errors.Wrap(ErrInvalidUsername, "unable to update records").Error(), err.Error())

	reqId1 := uuid.New().String()
	res2, err := svr.Update(context.TODO(), &DopeRequest{
		RequestId: reqId1,
		Payload:   []*Dope{item},
		Username:  "fake_user",
	})
	if err != nil {
		t.Fatal(err)
	}

	if !assert.NotNil(t, res2) {
		t.FailNow()
	}
	if !assert.NotNil(t, res2.Payload) {
		t.FailNow()
	}

	assert.Equal(t, reqId1, res2.RequestId)
	assert.Len(t, res2.Payload, 1)

	res3, err := RetrieveOne(context.TODO(), item.Id, db)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.NotNil(t, res3) {
		t.FailNow()
	}
	assert.EqualValues(t, item, res3)
}

func TestServiceDelete(t *testing.T) {
	client, clean := dbclient(t)
	defer clean(repo)

	svr := &DopeServer{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*Dope{
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "But I've never been to the moon! We need rest.",
			Dope_2:     9,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			Dope_2:     3,
		},
		&Dope{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			Dope_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			Dope_2:     1000000,
		},
	}

	db := client.Database(repo)

	// warm the test dataset
	if err := Create(context.TODO(), items, db); err != nil {
		t.Fatal(err)
	}

	// verify
	res0, err := Retrieve(context.TODO(), bson.D{}, db)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, res0, len(items))

	reqId0 := uuid.New().String()
	res1, err := svr.Delete(context.TODO(), &DopeRequest{
		RequestId: reqId0,
		Payload:   []*Dope{items[0]},
	})
	assert.NotEmpty(t, res1)
	assert.Equal(t, reqId0, res1.RequestId)
	assert.Error(t, err)
	assert.EqualValues(t, errors.Wrap(ErrInvalidUsername, "unable to delete records").Error(), err.Error())

	reqId1 := uuid.New().String()
	res2, err := svr.Delete(context.Background(), &DopeRequest{
		RequestId: reqId1,
		Payload:   []*Dope{items[0]},
		Username:  "fake_user",
	})
	if err != nil {
		t.Fatal(err)
	}

	if !assert.NotNil(t, res2) {
		t.FailNow()
	}
	assert.Equal(t, reqId1, res2.RequestId)

	res3, err := Retrieve(context.TODO(), bson.D{}, db)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.NotNil(t, res3) {
		t.FailNow()
	}
	assert.Len(t, res3, len(items)-1)
}
