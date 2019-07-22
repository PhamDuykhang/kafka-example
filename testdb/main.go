package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s","pdkhang","123456","localhost","27017","test")
	//fmt.Println(uri)
	//APname := "KRP"
	//client,err := mongo.NewClient(&options.ClientOptions{
	//	Auth:&options.Credential{
	//		AuthMechanism:auth.SCRAMSHA1,
	//		Username:"pdkhang",
	//		Password:"123456",
	//	},
	//	Hosts:[]string{"localhost:27017"},
	//	AppName:&APname,
	//
	//})

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://pdkhang:123456@localhost:27017/test?maxpoolsize=100"))
	if err != nil {
		logrus.Errorf("can't create client %v", err)
		return
	}

	err = client.Connect(context.TODO())
	if err != nil {
		logrus.Error(err)
		logrus.Errorf("can't connect to database %v", err)
		return
	}
	//NameONLY := true
	//databasename,err:=client.ListDatabaseNames(context.TODO(),nil,&options.ListDatabasesOptions{
	//	NameOnly:&NameONLY,
	//})
	//if err != nil{
	//	logrus.Errorf("can't fetch to database name %v",err)
	//	return
	//}
	//logrus.Info(databasename)
	data := bson.D{{
		"name",
		bson.D{{
			"$in",
			bson.A{"Alice", "Bob"},
		}},
	}}
	rs, err := client.Database("test").Collection("lo").InsertOne(context.TODO(), data)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Print(rs)

}
