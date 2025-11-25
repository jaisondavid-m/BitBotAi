	package handlers

	import (
		"library/storage"
		"library/models"
		"strings"
		"library/genai_client"


		"context"
		"fmt"
		// "net/http"

		"github.com/gin-gonic/gin"
		genai "github.com/google/generative-ai-go/genai"
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

	func AskQuestions(c *gin.Context) {

	var body struct {
		Question string `json:"question"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Format"})
		return
	}

	// Load uploaded study material
	materials, err := storage.GetMaterial()
	if err != nil || len(materials) == 0 {
		c.JSON(400, gin.H{"error": "No study material is uploaded"})
		return
	}

	allText := strings.Join(materials, "\n\n")

	// ❗ Use global client (created only once)
	if genai_client.Client == nil {
		c.JSON(500, gin.H{"error": "AI Client not initialized"})
		return
	}

	ctx := context.Background()
	model := genai_client.Client.GenerativeModel("gemini-2.5-flash")

	prompt := fmt.Sprintf(`
		You are an Educational AI Assistant.

		RULES:
		1️⃣ Use the uploaded study material first.
		2️⃣ If not found, check trusted sources:
		   - BIT Sathy website info
		3️⃣ If still not found, search the internet.
		4️⃣ If nowhere found -> reply "Information not available."
		5️⃣ Non-academic questions must be rejected politely.

		======== STUDY MATERIAL ========
		%s

		======== QUESTION ========
		%s
	`, allText, body.Question)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		c.JSON(500, gin.H{"error": "Server Error while generating response"})
		return
	}

	answer := resp.Candidates[0].Content.Parts[0]

	c.JSON(200, gin.H{
		"answer": fmt.Sprintf("%v", answer),
	})
}