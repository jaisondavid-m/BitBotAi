	package handlers

	import (
		"library/storage"
		"library/models"
		"strings"


		"context"
		"fmt"
		"net/http"
		"os"

		"github.com/gin-gonic/gin"
		genai "github.com/google/generative-ai-go/genai"
		"google.golang.org/api/option"
	)


	func UploadText(c *gin.Context){

		var body models.Body

		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		if err := storage.SaveMaterial(body.Text); err != nil {
			c.JSON(500, gin.H{"error": "DB Error: Could not store"})
			return
		}
		
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
		// text, err := storage.GetMaterial()
		// if err != nil || text == "" {
		// 	c.JSON(400, gin.H{"error": "No study material is uploaded"})
		// 	return
		// }
		materials, err := storage.GetMaterial()
		if err != nil || len(materials) == 0 {
			c.JSON(400, gin.H{"error": "No study material is uploaded"})
			return
		}

		// ⬅️ JOIN all rows into a single text
		allText := strings.Join(materials, "\n\n")
		apikey:=os.Getenv("GEMINI_API_KEY")

		ctx := context.Background()
		client,err :=genai.NewClient(ctx , option.WithAPIKey(apikey))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"server error"})
			return
		}

		model := client.GenerativeModel("gemini-2.5-flash")

	prompt := fmt.Sprintf(`
		You are an educational AI assistant.

		RULES:
		1. You must ALWAYS prioritize using the uploaded  material as the source of truth.
		2. If the user question is academic (school/college subjects like science, math, history, computer science, biology, chemistry, geography, etc.) AND the study material does not contain the answer, THEN you are allowed to check the following trusted websites:
		- https://www.bitsathy.ac.in/
		- Information about Bannari Amman Institute of Technology
		If the answer is still not found, you may search the internet to provide a correct answer.
		3. If the question is NOT related to education or academics (examples: gaming, money, hacking, cheating, entertainment, personal questions, illegal questions), you MUST reply:
		"❌ This question is not related to education. I can answer only academic questions."

		MATERIAL:
		%s

		USER QUESTION:
		%s

		Remember:
		- Always check if the question is educational.	
		- If educational → answer using material; if missing → check the trusted websites; if still missing → respond with "information not available".
		- If not educational → refuse.
	`, allText, body.Question)





		resp,err:= model.GenerateContent(ctx,genai.Text(prompt))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Server Error"})
			return
		}

		answer := resp.Candidates[0].Content.Parts[0]

		c.JSON(200,gin.H{"answer":fmt.Sprintf("%v",answer)})

	}