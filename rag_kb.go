package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"net/http"
)

func HandleRetrieveAndGenerate(w http.ResponseWriter, r *http.Request, client *bedrockagentruntime.Client) {

	// parse user messages
	type Content struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}

	type Message struct {
		Role    string    `json:"role"`
		Content []Content `json:"content"`
	}

	var request struct {
		Messages []Message `json:"messages"`
	}

	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		fmt.Println(error)
	}

	messages := request.Messages

	// pop the last message as user question
	userQuestion := messages[len(messages)-1].Content[0].Text

	// invoke bedrock agent runtime to retrieve and generate
	output, error := client.RetrieveAndGenerate(
		context.TODO(),
		&bedrockagentruntime.RetrieveAndGenerateInput{
			Input: &types.RetrieveAndGenerateInput{
				Text: aws.String(userQuestion),
			},
			RetrieveAndGenerateConfiguration: &types.RetrieveAndGenerateConfiguration{
				Type: types.RetrieveAndGenerateTypeKnowledgeBase,
				KnowledgeBaseConfiguration: &types.KnowledgeBaseRetrieveAndGenerateConfiguration{
					KnowledgeBaseId: aws.String(KNOWLEDGE_BASE_ID),
					ModelArn:        aws.String(KNOWLEDGE_BASE_MODEL_ID),
					RetrievalConfiguration: &types.KnowledgeBaseRetrievalConfiguration{
						VectorSearchConfiguration: &types.KnowledgeBaseVectorSearchConfiguration{
							NumberOfResults: aws.Int32(KNOWLEDGE_BASE_NUMBER_OF_RESULT),
						},
					},
				},
			},
		},
	)

	if error != nil {
		fmt.Println(error)
	}

	// write output to client
	json.NewEncoder(w).Encode(output)

}
