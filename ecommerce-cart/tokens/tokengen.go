package tokens

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iamvibs/golang-projects/ecommerce-cart/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.StandardClaims
}

var userData *mongo.Collection = database.UserData(database.Client, "Users")
var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email, fname, lname, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: fname,
		Last_Name:  lname,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is already expired"
		return
	}

	return claims, msg
}

func UpdateAllTokens(signedToken, signedRefreshToken, userID string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateobj primitive.D

	updateobj = append(updateobj, bson.E{Key: "token", Value: signedToken})
	updateobj = append(updateobj, bson.E{Key: "refresh_token", Value: signedRefreshToken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append(updateobj, bson.E{Key: "updated_at", Value: updated_at})

	upsert := true
	filter := bson.M{"user_id": userID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userData.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateobj}}, &opt)
	if err != nil {
		log.Panic(err)
		return
	}
}
