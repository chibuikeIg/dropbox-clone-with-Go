package repository

import (
	"context"
	"errors"
	"filemetadata-service/internals/app/config"
	"filemetadata-service/internals/core/ports"
	"flag"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/fatih/color"
)

type DynamoDbRepository struct {
	Client    *dynamodb.Client
	Ctx       context.Context
	TableName string
	// db query express clause
	queryKey map[string]types.AttributeValue
}

func NewDynamoDbRepository() *DynamoDbRepository {

	ctx := context.TODO()
	cfg := config.NewAwsCredentials()
	client := dynamodb.NewFromConfig(aws.Config{
		Region:      os.Getenv("AWS_REGION"),
		Credentials: cfg,
	})

	/// setup commands for database migration
	flagSet := flag.NewFlagSet("create", flag.ExitOnError)
	tablename := flagSet.String("table_name", "", "use this to set the table name")
	primary_key := flagSet.String("primary_key", "", "use this to set the primary key for table")
	range_key := flagSet.String("range_key", "", "use this to set the range key")
	migration_type := flagSet.String("mt", "new", "use this to specify if you want to create a new table or update existing one. Set to `update` if updating")

	if len(os.Args) >= 2 {
		flagSet.Parse(os.Args[2:])
	}

	tname := *tablename
	pk := *primary_key
	rk := *range_key
	mt := *migration_type

	if tname != "" && pk != "" && rk != "" {

		if mt == "" || mt == "new" {

			_, err := RunMigration(client, tname, pk, rk)
			if err != nil {
				log.Fatalf("unable to run dynamodb migration, %v", err)
			}

			color.Green("Table: %s was created successfully", tname)

		} else if mt != "" || mt == "update" {
			_, err := CreateSecondaryIndex(client, tname, pk, rk)
			if err != nil {
				log.Fatalf("unable to run dynamodb migration, %v", err)
			}

			color.Green("Table: %s secondary index was created successfully", tname)
		}

	}

	return &DynamoDbRepository{Client: client, Ctx: ctx}
}

func (db *DynamoDbRepository) Table(tname string) ports.FileMetaDataDBRepository {
	db.TableName = tname
	return db
}

func (db *DynamoDbRepository) Find(condition []string, dataModel any) any {

	// loop through conditions
	// and convert to aws attribute value
	keyEx := expression.Key(condition[0]).Equal(expression.Value(condition[1]))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
		return nil
	}

	/// Query aws dynamoDB to get item
	response, err := db.Client.Query(db.Ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(db.TableName),
		IndexName:                 aws.String(condition[0] + "Index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		log.Printf("Couldn't query result. Here's why: %v\n", err)
		return nil
	}

	if response.Count == 0 {
		return nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &dataModel)

	if err != nil {
		log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
		return nil
	}

	return nil
}

func (db *DynamoDbRepository) Get() (any, error) {

	return nil, nil
}

func (db *DynamoDbRepository) First() (any, error) {

	var data any

	response, err := db.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: db.queryKey, TableName: aws.String(db.TableName),
	})

	if err != nil {
		log.Printf("Couldn't retreive %s record. Here's why: %v\n", db.TableName, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &data)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	return data, err
}

func (db *DynamoDbRepository) Create(data any) (any, error) {

	item, err := attributevalue.MarshalMap(data)

	if err != nil {
		panic(item)
	}

	_, err = db.Client.PutItem(db.Ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName), Item: item,
	})

	return data, err
}

func (db *DynamoDbRepository) Update(data map[string]string) (any, error) {

	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	var err error
	var response *dynamodb.UpdateItemOutput

	var attributeMap any

	// build expression
	update := expression.Set(expression.Name(keys[0]), expression.Value(data[keys[0]]))
	for k := range keys {
		if k == 0 {
			continue
		}
		update.Set(expression.Name(keys[k]), expression.Value(data[keys[k]]))
	}
	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	}

	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = db.Client.UpdateItem(db.Ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(db.TableName),
			Key:                       db.queryKey,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update record. Here's why: %v\n", err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}

	return attributeMap, err
}

func (db *DynamoDbRepository) Where(data [][]any) ports.FileMetaDataDBRepository {

	if len(data) <= 1 {
		log.Printf("Couldn't build expression for update. Here's why: inappropriate clause provided")
	}

	keys := make(map[string]types.AttributeValue)

	for k := range data {

		marshalledVal, err := attributevalue.Marshal(data[k][len(data[k])-1].(string))

		if err != nil {
			panic(err)
		}

		keys[data[k][0].(string)] = marshalledVal
	}

	db.queryKey = keys

	return db
}

func (db *DynamoDbRepository) Delete() error {
	_, err := db.Client.DeleteItem(db.Ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(db.TableName), Key: db.queryKey,
	})

	if err != nil {
		log.Printf("Couldn't delete record from the table. Here's why: %v\n", err)
		return err
	}

	return nil
}

func RunMigration(dbClient *dynamodb.Client, tableName string, pk string, rk string) (string, error) {

	var tableDesc *types.TableDescription

	// check if table already exists
	if exists, err := TableExists(dbClient, tableName); err == nil && exists {
		return tableName, nil
	}

	// create table
	table, err := dbClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{ // set primary/index keys
			AttributeName: aws.String(pk),
			AttributeType: types.ScalarAttributeTypeS,
		}, { // set primary/index keys
			AttributeName: aws.String(rk),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String(pk),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String(rk),
			KeyType:       types.KeyTypeRange,
		}},
		TableName: aws.String(tableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})

	if err != nil {
		log.Printf("Couldn't create table %v. Here's why: %v\n", tableName, err)
	} else {

		// wait for table to finish creating before returning
		waiter := dynamodb.NewTableExistsWaiter(dbClient)

		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName)}, 5*time.Minute)

		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}

		tableDesc = table.TableDescription

		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
	}

	return string(tableDesc.TableStatus), err
}

func CreateSecondaryIndex(dbClient *dynamodb.Client, tableName string, pk string, rk string) (*types.TableDescription, error) {

	// update table and setup secondary index
	table, err := dbClient.UpdateTable(context.TODO(), &dynamodb.UpdateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{ // set primary/index keys
			AttributeName: aws.String(rk),
			AttributeType: types.ScalarAttributeTypeS,
		}, { // set primary/index keys
			AttributeName: aws.String(pk),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		GlobalSecondaryIndexUpdates: []types.GlobalSecondaryIndexUpdate{
			{
				Create: &types.CreateGlobalSecondaryIndexAction{

					IndexName: aws.String(pk + "Index"),
					KeySchema: []types.KeySchemaElement{{
						AttributeName: aws.String(pk),
						KeyType:       types.KeyTypeHash,
					}, {
						AttributeName: aws.String(rk),
						KeyType:       types.KeyTypeRange,
					}},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(10),
						WriteCapacityUnits: aws.Int64(10),
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
				},
			},
		},

		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	// wait for table to finish creating before returning
	waiter := dynamodb.NewTableExistsWaiter(dbClient)

	err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName)}, 5*time.Minute)

	if err != nil {
		return nil, err
	}

	tableDesc := table.TableDescription

	return tableDesc, nil
}

func TableExists(client *dynamodb.Client, tableName string) (bool, error) {
	exists := true
	_, err := client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", tableName, err)
		}
		exists = false
	}
	return exists, err
}

func DeleteTable(client *dynamodb.Client, tableName string) error {
	_, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName)})
	if err != nil {
		log.Printf("Couldn't delete table %v. Here's why: %v\n", tableName, err)
	}
	return err
}
