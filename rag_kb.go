package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
)

func HandleRetrieveAndGenerate(client *bedrockagentruntime.Client) {

	// parse user messages
	type Content struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}

	type Message struct {
		Role    string    `json:"role"`
		Content []Content `json:"content"`
	}

	//var request struct {
	//	Messages []Message `json:"messages"`
	//}

	//error := json.NewDecoder(r.Body).Decode(&request)
	//
	//if error != nil {
	//	fmt.Println(error)
	//}
	//

	// pop the last message as user question
	userQuestion := "What time is it?"

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
					KnowledgeBaseId: aws.String("BDKIVJYOVN"),
					ModelArn:        aws.String("anthropic.claude-3-haiku-20240307-v1:0"),
					RetrievalConfiguration: &types.KnowledgeBaseRetrievalConfiguration{
						VectorSearchConfiguration: &types.KnowledgeBaseVectorSearchConfiguration{
							NumberOfResults: aws.Int32(6),
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
	//json.NewEncoder(w).Encode(output)
	result := output.Output.Text

	fmt.Printf("%+v", *result)

}
