package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	appAwsConfig "github.com/jiaqi-yin/go-file-processor/config"
	"github.com/jiaqi-yin/go-file-processor/domain"
)

type dynamodbService struct {
	Client *dynamodb.Client
	Table  string
}

func (ds *dynamodbService) Save(item *domain.Item) {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(ds.Table),
	}

	_, err = ds.Client.PutItem(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

func NewDynamodbService() Storage {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if appAwsConfig.Configurations.Aws.Endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           appAwsConfig.Configurations.Aws.Endpoint,
				SigningRegion: appAwsConfig.Configurations.Aws.Region,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(appAwsConfig.Configurations.Aws.Profile),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		panic(err)
	}

	return &dynamodbService{
		Client: dynamodb.NewFromConfig(cfg),
		Table:  appAwsConfig.Configurations.Aws.Dynamodb.Table,
	}
}
