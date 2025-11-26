package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func getJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("default_secret_key")
	}
	return []byte(secret)
}

func RegisterUser(ctx context.Context, input models.UserInput) error {
	// Check if user exists
	count, err := userCollection.CountDocuments(ctx, bson.M{"username": input.Username})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Check if this is the first user
	totalUsers, err := userCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	role := "user"
	if totalUsers == 0 {
		role = "admin"
	}

	user := models.User{
		ID:       uuid.New().String(),
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     role,
	}

	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func LoginUser(ctx context.Context, input models.UserInput) (string, error) {
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": input.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(getJwtSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func PromoteUser(ctx context.Context, username string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
