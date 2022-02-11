package storage

import (
	"errors"
	"github.com/OpenCal-FYDP/PreferenceManagement/rpc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	NotFound = errors.New("key not found")
)

const tableName = "UserPreferences"

type Storage struct {
	client *dynamodb.DynamoDB
}

func New() *Storage {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(sess)

	return &Storage{client}
}

func (s *Storage) GetUserProfile(req *rpc.GetUserProfileReq, res *rpc.GetUserProfileRes) error {
	result, err := s.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userID": {
				S: aws.String(req.UserID),
			},
		},
	})
	if err != nil {
		return err
	}

	if result.Item == nil {
		return NotFound
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, &res)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) SetUserProfile(req *rpc.SetUserProfileReq, res *rpc.SetUserProfileRes) error {
	av, err := dynamodbattribute.MarshalMap(req)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = s.client.PutItem(input)
	return err
}
