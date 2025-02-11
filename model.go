package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"log"
)

type BedrockAgent struct {
	Client bedrockagentruntime.Client
}

func NewBedrock() *BedrockAgent {
	ctx := context.Background()
	var BedrockAgentRuntimeClient *bedrockagentruntime.Client

	// Load AWS Credentials
	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatal("Failed to load AWS credentials", err)
	}

	// create BedrockAgentClient
	BedrockAgentRuntimeClient = bedrockagentruntime.NewFromConfig(awsConfig)

	return &BedrockAgent{
		Client: *BedrockAgentRuntimeClient,
	}
}

func (bedrockAgent *BedrockAgent) RetrieveResponseFromKnowledgeBase(question string) string {
	// invoke bedrock agent runtime to retrieve and generate
	output, err := bedrockAgent.Client.RetrieveAndGenerate(
		context.TODO(),
		&bedrockagentruntime.RetrieveAndGenerateInput{
			Input: &types.RetrieveAndGenerateInput{
				Text: aws.String(question),
			},
			RetrieveAndGenerateConfiguration: &types.RetrieveAndGenerateConfiguration{
				Type: types.RetrieveAndGenerateTypeKnowledgeBase,
				KnowledgeBaseConfiguration: &types.KnowledgeBaseRetrieveAndGenerateConfiguration{
					KnowledgeBaseId: aws.String(KNOWLEDGE_BASE_ID),
					ModelArn:        aws.String(MODEL_ARN),
					RetrievalConfiguration: &types.KnowledgeBaseRetrievalConfiguration{
						VectorSearchConfiguration: &types.KnowledgeBaseVectorSearchConfiguration{
							NumberOfResults: aws.Int32(6),
						},
					},
				},
			},
		},
	)

	if err != nil {
		log.Fatal("RetrieveResponseFromKnowledgeBase::", err)
	}
	result := output.Output.Text
	return *result
}
