package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

type LaptopRequest struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}

type Recommendation struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func getLaptopRecommendation(c echo.Context) error {
	var reqData LaptopRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	client := openai.NewClient(os.Getenv("TOKEN_OPENAI"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that recommends laptops.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Recommend a laptop with a budget of %d for %s use.", reqData.Budget, reqData.Purpose),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	recommendation := resp.Choices[0].Message.Content

	response := Recommendation{
		Status: "success",
		Data:   recommendation,
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.POST("/recommendation", getLaptopRecommendation)

	e.Start(":8000")
}
