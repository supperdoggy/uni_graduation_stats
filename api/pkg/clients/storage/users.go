package storage

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *mongodb) CreateUser(ctx context.Context, email, fullname string, password []byte) (*db.User, error) {
	u := db.User{
		ID:        uuid.New().String(),
		Email:     email,
		FullName:  fullname,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}
	_, err := c.users.InsertOne(ctx, u)
	return &u, err
}

func (c *mongodb) DeleteUser(ctx context.Context, id string) error {
	_, err := c.users.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *mongodb) UpdateUser(ctx context.Context, id, email string, password []byte) error {
	_, err := c.users.UpdateByID(ctx, id, bson.M{"$set": bson.M{"email": email, "password": password, "edited_at": time.Now().Unix()}})
	return err
}

func (c *mongodb) GetUser(ctx context.Context, id string) (*db.User, error) {
	resp := c.users.FindOne(ctx, bson.M{"_id": id})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var user db.User
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil

}

func (c *mongodb) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	resp := c.users.FindOne(ctx, bson.M{"email": email})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var user db.User
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *mongodb) NewEmailCode(ctx context.Context, email, code string) error {

	// check if email occupied
	user, err := c.GetUserByEmail(ctx, email)
	if err != nil && err != mongo.ErrNoDocuments || user != nil {
		return errors.New("email occupied")
	}

	emailcode := db.EmailCode{
		Email:  email,
		Code:   code,
		Expire: time.Now().Add(time.Minute * 10),
	}

	c.emailCodeCache.mut.Lock()
	c.emailCodeCache.m[email] = emailcode
	c.emailCodeCache.mut.Unlock()

	return nil
}

func (c *mongodb) CheckEmailCode(ctx context.Context, email, code string) (bool, error) {
	c.emailCodeCache.mut.Lock()
	defer c.emailCodeCache.mut.Unlock()

	emailcode, ok := c.emailCodeCache.m[email]
	if !ok {
		return false, errors.New("no such email")
	}

	if emailcode.Code != code {
		return false, errors.New("wrong code")
	}

	if emailcode.Expire.Before(time.Now()) {
		return false, errors.New("expired")
	}

	delete(c.emailCodeCache.m, email)

	return true, nil
}
