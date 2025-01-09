package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"

	"log"
)

var BedrockClient *bedrockruntime.Client

// bedrock agent runtime client
var BedrockAgentRuntimeClient *bedrockagentruntime.Client

func init() {

	// load aws credentials from profile demo using config
	awsCfg1, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ap-south-1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// create bedrock runtime client
	BedrockClient = bedrockruntime.NewFromConfig(awsCfg1)
	// create bedrock agent runtime client
	BedrockAgentRuntimeClient = bedrockagentruntime.NewFromConfig(awsCfg1)

}

func main() {
	//gobedrock.HandleRetrieveAndGenerate(w, r, BedrockAgentRuntimeClient)
}
