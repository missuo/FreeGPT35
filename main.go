package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type OpenAIResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []OpenAIChoice `json:"choices"`
}

type OpenAIChoice struct {
	Index        int         `json:"index"`
	Delta        OpenAIDelta `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason *string     `json:"finish_reason"`
}

type OpenAIDelta struct {
	Content string `json:"content"`
}

type DuckDuckGoResponse struct {
	Role    string `json:"role"`
	Message string `json:"message"`
	Created int64  `json:"created"`
	ID      string `json:"id"`
	Action  string `json:"action"`
	Model   string `json:"model"`
}

func chatWithDuckDuckGo(c *gin.Context, messages []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}) {
	userAgent := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0"
	headers := map[string]string{
		"User-Agent":      userAgent,
		"Accept":          "text/event-stream",
		"Accept-Language": "de,en-US;q=0.7,en;q=0.3",
		"Accept-Encoding": "gzip, deflate, br",
		"Referer":         "https://duckduckgo.com/",
		"Content-Type":    "application/json",
		"Origin":          "https://duckduckgo.com",
		"Connection":      "keep-alive",
		"Cookie":          "dcm=1",
		"Sec-Fetch-Dest":  "empty",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Site":  "same-origin",
		"Pragma":          "no-cache",
		"TE":              "trailers",
	}

	statusURL := "https://duckduckgo.com/duckchat/v1/status"
	chatURL := "https://duckduckgo.com/duckchat/v1/chat"

	// get vqd_4
	req, err := http.NewRequest("GET", statusURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("x-vqd-accept", "1")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	vqd4 := resp.Header.Get("x-vqd-4")

	payload := map[string]interface{}{
		"model":    "gpt-3.5-turbo-0125",
		"messages": messages,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err = http.NewRequest("POST", chatURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("x-vqd-4", vqd4)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	flusher, _ := c.Writer.(http.Flusher)

	var response OpenAIResponse
	response.Choices = make([]OpenAIChoice, 1)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if bytes.HasPrefix(line, []byte("data: ")) {
			chunk := line[6:]

			if bytes.HasPrefix(chunk, []byte("[DONE]")) {
				// send stop
				c.Data(http.StatusOK, "text/plain", []byte("data: {\"id\":\"chatcmpl-9HOzx2PhnYCLPxQ3Dpa2OKoqR2lgl\",\"object\":\"chat.completion\",\"created\":1713934697,\"model\":\"gpt-3.5-turbo-0125\",\"choices\":[{\"index\":0,\"delta\":{\"content\":\"\"},\"logprobs\":null,\"finish_reason\":\"stop\"}]}\n\n"))
				flusher.Flush()

				// send done
				c.Data(http.StatusOK, "text/plain", []byte("data: [DONE]"))
				flusher.Flush()
				return
			}

			var data DuckDuckGoResponse
			err = json.Unmarshal(chunk, &data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			response.ID = data.ID
			response.Object = "chat.completion"
			response.Created = data.Created
			response.Model = data.Model
			response.Choices[0].Delta.Content = data.Message

			responseBytes, err := json.Marshal(response)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Data(http.StatusOK, "text/plain", append(append([]byte("data: "), responseBytes...), []byte("\n\n")...))
			flusher.Flush()

			response.Choices[0].Delta.Content = ""
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/v1/chat/completions", func(c *gin.Context) {
		var req OpenAIRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// only support user role
		for i := range req.Messages {
			if req.Messages[i].Role == "system" {
				req.Messages[i].Role = "user"
			}
		}
		// set model to gpt-3.5-turbo-0125
		req.Model = "gpt-3.5-turbo-0125"

		chatWithDuckDuckGo(c, req.Messages)
	})

	r.Run(":8080")
}
