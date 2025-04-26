package main

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/Mckee404/microcms_search_api/repository"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// クエリパラメータで受け取った検索語を使ってデータを取得
	q, _ := url.QueryUnescape(event.QueryStringParameters["q"])

	result, err := repository.SearchBlog(q)
	if err != nil {
		// エラーが発生した場合は、エラーメッセージをJSONで返す
		errorResponse := map[string]string{"error": "Error searching blog"}
		errorJSON, _ := json.Marshal(errorResponse)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(errorJSON),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// 成功した場合、結果をJSONで返す
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		errorResponse := map[string]string{"error": "Internal Server Error"}
		errorJSON, _ := json.Marshal(errorResponse)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(errorJSON),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// 成功時のレスポンス
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonBytes),
	}, nil
}

func main() {
	lambda.Start(handler)
}
