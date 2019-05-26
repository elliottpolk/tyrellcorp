package {{ .Package | ToLower | Trim }}

import (
	"context"
	"testing"

	"{{ .Repository | ToLower | Trim }}/{{ .Package | ToLower | Trim }}/config"

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

	svr := &{{ .Name | Trim }}Server{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*{{ .Name | Trim  }}{
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "But I've never been to the moon! We need rest.",
			{{ .Name | Trim }}_2:     9,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			{{ .Name | Trim }}_2:     3,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			{{ .Name | Trim }}_2:     1000000,
		},
	}

	reqId := uuid.New().String()
	res, err := svr.Create(context.TODO(), &{{ .Name | Trim }}Request{
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

	svr := &{{ .Name | Trim }}Server{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*{{ .Name | Trim }}{
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "But I've never been to the moon! We need rest.",
			{{ .Name | Trim }}_2:     9,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			{{ .Name | Trim }}_2:     3,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			{{ .Name | Trim }}_2:     1000000,
		},
	}

	// warm the test dataset
	if err := Create(context.TODO(), items, client.Database(repo)); err != nil {
		t.Fatal(err)
	}

	reqId0 := uuid.New().String()
	res0, err := svr.Retrieve(context.TODO(), &{{ .Name | Trim }}Request{
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
	res1, err := svr.Retrieve(context.TODO(), &{{ .Name | Trim }}Request{
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

	svr := &{{ .Name | Trim }}Server{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*{{ .Name | Trim  }}{
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "But I've never been to the moon! We need rest.",
			{{ .Name | Trim }}_2:     9,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			{{ .Name | Trim }}_2:     3,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			{{ .Name | Trim }}_2:     1000000,
		},
	}

	db := client.Database(repo)

	// warm the test dataset
	if err := Create(context.TODO(), items, db); err != nil {
		t.Fatal(err)
	}

	// adjust the item
	item := items[0]

	item.{{ .Name | Trim }}_1 = "You are the last hope of the universe. Why not indeed! Aww, it's true. I've been hiding it for so long. Well, then good news! It's a suppository. OK, if everyone's finished being stupid."
	res0, err := RetrieveOne(context.TODO(), item.Id, db)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, item.{{ .Name | Trim }}_1, res0.{{ .Name | Trim }}_1)

	reqId0 := uuid.New().String()
	res1, err := svr.Update(context.TODO(), &{{ .Name | Trim }}Request{
		RequestId: reqId0,
		Payload:   []*{{ .Name | Trim }}{item},
	})
	assert.Empty(t, res1)
	assert.Error(t, err)
	assert.EqualValues(t, errors.Wrap(ErrInvalidUsername, "unable to update records").Error(), err.Error())

	reqId1 := uuid.New().String()
	res2, err := svr.Update(context.TODO(), &{{ .Name | Trim }}Request{
		RequestId: reqId1,
		Payload:   []*{{ .Name | Trim }}{item},
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

	svr := &{{ .Name | Trim }}Server{
		cmp:    &config.Composition{},
		client: client,
	}

	items := []*{{ .Name | Trim }}{
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "But I've never been to the moon! We need rest.",
			{{ .Name | Trim }}_2:     9,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "The spirit is willing, but the flesh is spongy and bruised. Incidentally, you have a dime up your nose. And until then, I can never die?",
			{{ .Name | Trim }}_2:     3,
		},
		&{{ .Name | Trim }}{
			RecordInfo: &RecordInfo{CreatedBy: "testing"},
			{{ .Name | Trim }}_1:     "I decline the title of Iron Cook and accept the lesser title of Zinc Saucier, which I just made up. Uhh… also, comes with double prize money. I am the man with no name, Zapp Brannigan! WINDMILLS DO NOT WORK THAT WAY! GOOD NIGHT!",
			{{ .Name | Trim }}_2:     1000000,
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
	res1, err := svr.Delete(context.TODO(), &{{ .Name | Trim }}Request{
		RequestId: reqId0,
		Payload:   []*{{ .Name | Trim }}{items[0]},
	})
	assert.NotEmpty(t, res1)
	assert.Equal(t, reqId0, res1.RequestId)
	assert.Error(t, err)
	assert.EqualValues(t, errors.Wrap(ErrInvalidUsername, "unable to delete records").Error(), err.Error())

	reqId1 := uuid.New().String()
	res2, err := svr.Delete(context.Background(), &{{ .Name | Trim }}Request{
		RequestId: reqId1,
		Payload:   []*{{ .Name | Trim }}{items[0]},
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
