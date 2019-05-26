package tyrellcorp

import (
	"context"
	"time"

	"github.com/elliottpolk/tyrellcorp/config"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Server ...
type Server struct {
	cfg    *config.Composition
	client *mongo.Client
}

// NewServer ...
func NewServer(cfg *config.Composition, client *mongo.Client) SpecServiceServer {
	return &Server{
		cfg:    cfg,
		client: client,
	}
}

// Create ...
func (s *Server) Create(c context.Context, r *SpecRequest) (*Empty, error) {
	empty := &Empty{RequestId: r.RequestId}

	if s.client == nil {
		return empty, ErrNoValidMongoClient
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := s.client
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return empty, errors.Wrap(err, "unable verify connection to datastore")
	}

	log.Debug("generating replicants for supplied payload")
	for _, spec := range r.Payload {
		// ensure an existing replicant does not exist in the datastore
		prev, err := Retrieve(c, bson.D{{"name", spec.Name}}, client.Database(repo))
		if err != nil {
			return empty, errors.Wrapf(err, "unable to verify if replicant currently exists for %s", spec.Name)
		}

		if len(prev) > 0 {
			return empty, errors.Errorf("duplicate replicant with name %s", spec.Name)
		}

		// write new replicant to datastore with status "generating"
		spec.State = InProgress
		if err := Create(c, []*Spec{spec}, client.Database(repo)); err != nil {
			return empty, errors.Wrapf(err, "unable to write new replicant spec %s to datastore", spec.Name)
		}

		// generate replicant
		if err := GenerateReplicant(spec); err != nil {
			return empty, errors.Wrapf(err, "unable to generate the replicant %s", spec.Name)
		}
		log.Debugf("replicant %s generated", spec.Name)
	}

	// TODO:
	// - spin up service
	// - deploy to remote git hosting
	// - remove replicant dir

	return empty, nil
}

// Retrieve ...
func (s *Server) Retrieve(c context.Context, r *SpecRequest) (*SpecResponse, error) {
	if s.client == nil {
		return nil, ErrNoValidMongoClient
	}

	res := &SpecResponse{
		RequestId: r.RequestId,
	}

	client := s.client
	if err := client.UseSession(c, func(session mongo.SessionContext) error {
		defer session.EndSession(c)

		// retrieve 1 and return by ID if provided in request
		if id := r.Id; len(id) > 0 {
			item, err := RetrieveOne(c, id, client.Database(repo))
			if err != nil {
				return errors.Wrapf(err, "unable to retrieve record for id %s", id)
			}
			res.Payload = []*Spec{item}

			return nil
		}

		// TODO:
		// - handle a list of whitelisted request params for filtering
		items, err := Retrieve(c, bson.D{}, client.Database(repo))
		if err != nil {
			return errors.Wrap(err, "unable to retrieve records")
		}
		res.Payload = items

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

// Update ...
func (s *Server) Update(c context.Context, r *SpecRequest) (*Empty, error) {
	empty := &Empty{RequestId: r.RequestId}

	if len(r.Payload) < 1 {
		return empty, errors.New("invalid payload")
	}

	if s.client == nil {
		return empty, ErrNoValidMongoClient
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := s.client
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return empty, errors.Wrap(err, "unable verify connection to datastore")
	}

	if err := Update(c, r.Username, r.Payload[0], client.Database(repo)); err != nil {
		return empty, errors.Wrap(err, "unable update records")
	}

	return empty, nil
}

// Delete ...
func (s *Server) Delete(c context.Context, r *SpecRequest) (*Empty, error) {
	empty := &Empty{RequestId: r.RequestId}

	if s.client == nil {
		return empty, ErrNoValidMongoClient
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := s.client

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return empty, errors.Wrap(err, "unable verify connection to datastore")
	}

	// TODO:

	// if err := Delete(r.Payload, client.Database(repo)); err != nil {
	// 	return empty, err
	// }

	return empty, nil
}
