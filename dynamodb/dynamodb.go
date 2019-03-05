package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

type CoffeeItem struct {
	ClickType string `json:"clickType"`
	Timestamp int    `json:"timestamp,omitempty"`
	Username  string `json:"username,omitempty"`
}

func GetUsers() (map[string]int, error) {

	var usermap = make(map[string]int)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	},
	)
	if err != nil {
		return usermap, err
	}
	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("coffee"),
		ExpressionAttributeNames: map[string]*string{
			"#username": aws.String("username"),
		},
		FilterExpression: aws.String("attribute_exists(#username)"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		return usermap, err
	}

	fmt.Printf("Scan Results: %v\n", result)

	coffeeItem := CoffeeItem{}

	for _, item := range result.Items {
		fmt.Printf("Item: %v\n", item)
		err = dynamodbattribute.UnmarshalMap(item, &coffeeItem)
		if err != nil {
			return usermap, err
		}
		count := usermap[coffeeItem.Username]
		count++
		usermap[coffeeItem.Username] = count
	}

	fmt.Printf("Map: %v\n", usermap)
	return usermap, err
}

func UpdateLastCoffee(username string) error {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	},
	)

	if err != nil {
		return err
	}
	fmt.Printf("Session: %v\n", sess)

	svc := dynamodb.New(sess)
	fmt.Printf("Service: %v\n", svc)

	result1, err := GetLastTime(svc, "SINGLE")
	result2, err := GetLastTime(svc, "DOUBLE")

	if err != nil {
		return err
	}

	coffeeResult := result1
	if result2.Timestamp > result1.Timestamp {
		coffeeResult = result2
	}

	fmt.Printf("Results: %v %v %v\n", result1, result2, coffeeResult)

	UpdateItemInTable(svc, coffeeResult, username)

	return err
}

func UpdateItemInTable(svc *dynamodb.DynamoDB, coffee CoffeeItem, username string) {
	updateInput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":user": {
				S: aws.String(username),
			},
			":time": {
				N: aws.String(strconv.Itoa(coffee.Timestamp)),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#timestamp": aws.String("timestamp"),
		},
		ConditionExpression: aws.String("#timestamp = :time"),
		TableName:           aws.String("coffee"),
		Key: map[string]*dynamodb.AttributeValue{
			"clickType": {
				S: aws.String(coffee.ClickType),
			},
			"timestamp": {
				N: aws.String(strconv.Itoa(coffee.Timestamp)),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set username = :user"),
	}
	fmt.Printf("Input: %v\n", updateInput)

	result, err := svc.UpdateItem(updateInput)
	fmt.Printf("Update: %v  Error: %v \n", result, err)
}

func GetLastTime(svc *dynamodb.DynamoDB, key string) (CoffeeItem, error) {
	item := CoffeeItem{}
	params := &dynamodb.QueryInput{
		KeyConditions: map[string]*dynamodb.Condition{
			"clickType": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(key),
					},
				},
			},
		},
		Limit:            aws.Int64(1),
		ScanIndexForward: aws.Bool(false),
		TableName:        aws.String("coffee"),
	}
	result, err := svc.Query(params)
	if err != nil {
		return item, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	return item, err
}
