	package handlers

	import (
		"library/storage"

		"context"
		"fmt"
		"net/http"
		"os"

		"github.com/gin-gonic/gin"
		genai "github.com/google/generative-ai-go/genai"
		"google.golang.org/api/option"
	)


	func UploadText(c *gin.Context){
		var body struct{
			Text string `json:"text"`
		}

		if err := c.BindJSON(&body); err !=nil{
			c.JSON(400,gin.H{"error":"Invalid Format"})
			return
		}

		storage.SavedText=body.Text
		c.JSON(200,gin.H{"message":"material Uploaded Successfully"})
	}

	func AskQuestions(c *gin.Context){
		
		var body struct{
			Question  string `json:"question"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid Format"})
			return
		}
		if storage.SavedText==""{
			c.JSON(http.StatusBadRequest,gin.H{"error":"No study Material is Uploaded"})
			return
		}
		apikey:=os.Getenv("GEMINI_API_KEY")
		fmt.Println("API KEY:", apikey)


		ctx := context.Background()
		client,err :=genai.NewClient(ctx , option.WithAPIKey(apikey))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"server error"})
			return
		}

		model := client.GenerativeModel("gemini-2.5-flash")

		prompt := fmt.Sprintf(`
							You are an educational assistant.
							Answer ONLY using the following study material:
							%s
							student Question:%s

							if the answer is not in the material , respond:
							"😔we have not trained this model to answer this Questions"
							`,storage.SavedText,body.Question)

		resp,err:= model.GenerateContent(ctx,genai.Text(prompt))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Server Error"})
			return
		}

		answer := resp.Candidates[0].Content.Parts[0]

		c.JSON(200,gin.H{"answer":fmt.Sprintf("%v",answer)})

	}