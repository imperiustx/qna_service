package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database is database connector
type Database struct {
	db                *mongo.Database
	session           *mongo.Client
	writeSessionsLock sync.Mutex
}

//NewDefaultDatabase creates new database connector with default parameters
func NewDefaultDatabase() (*Database, error) {
	username := env("DB_USERNAME", "")
	password := env("DB_PASSWORD", "")
	host := env("DB_HOST", "")
	port := env("DB_PORT", "")
	dbname := env("DB_NAME", "")

	return NewDatabase(username, password, host, port, dbname)
}

//NewDatabase creates new database connector
func NewDatabase(username, password, host, port, dbname string) (*Database, error) {
	db, session, err := connect(username, password, host, port, dbname)
	if err != nil {
		return nil, err
	}

	return &Database{db: db, session: session}, nil
}

func connect(username, password, host, port, dbname string) (db *mongo.Database, client *mongo.Client, err error) {
	var connectOnce sync.Once
	connectOnce.Do(func() {
		db, client, err = connectToMongo(username, password, host, port, dbname)
	})
	return db, client, err
}

func connectToMongo(username, password, host, port, dbname string) (db *mongo.Database, client *mongo.Client, err error) {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, dbname)

	clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return db, client, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return db, client, err
	}

	db = client.Database(dbname)

	return db, client, nil
}

// Lock protects the writing database.
func (ds *Database) Lock() {
	ds.writeSessionsLock.Lock()
}

// Unlock protects the writing database.
func (ds *Database) Unlock() {
	ds.writeSessionsLock.Unlock()
}

// Create executes an insert command to insert a single document into the collection.
func (ds *Database) Create(collection string, document interface{}) (interface{}, error) {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := c.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// CreateMany executes an insert command to insert multiple documents into the collection. If write errors occur
func (ds *Database) CreateMany(collection string, documents []interface{}) ([]interface{}, error) {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := c.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

// Find executes a find command and returns a Cursor over the matching documents in the collection.
func (ds *Database) Find(collection string, filter interface{}) (*mongo.Cursor, error) {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return c.Find(ctx, filter)
}

// FindOne executes a find command and returns a Jmap for one document in the collection.
func (ds *Database) FindOne(collection string, filter interface{}) (interface{}, error) {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r := c.FindOne(ctx, filter)
	var result interface{}
	err := r.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll executes a find command and returns a list  matching documents in the collection.
func (ds *Database) FindAll(collection string, filter interface{}) ([]interface{}, error) {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	results := make([]interface{}, 0)

	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result interface{}
		if err := cursor.Decode(result); err == nil {
			results = append(results, result)
		}
	}
	return results, nil
}

// Update executes an update command to update at most one document in the collection.
func (ds *Database) Update(collection string, filter, update interface{}) error {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete executes a delete command to delete documents from the collection.
func (ds *Database) Delete(collection string, filter interface{}) error {
	c := ds.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := c.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// Aggregate executes an aggregate command against the collection and returns a cursor over the resulting documents.
func (ds *Database) Aggregate(collection string, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c := ds.db.Collection(collection)

	cursor, err := c.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return cursor, err
}
