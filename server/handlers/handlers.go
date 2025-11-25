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
					You are an Educational AI Assistant.

					Your job is to answer ONLY academic questions correctly using the following rules:

					============================
					📘 RULE ORDER (MUST follow):
					============================

					1️⃣ **Check the uploaded study material first.**  
					- If the answer is found in the material, reply using ONLY that information.

					2️⃣ **If the answer is NOT found in the material:**
					- Check ONLY the trusted academic source:
						- https://www.bitsathy.ac.in/
						- Information about Bannari Amman Institute of Technology

					3️⃣ **If still not found AND the question is educational:**
					- You may search the internet and answer accurately.

					4️⃣ **If still not found from all sources:**
					- Reply: "Information not available."

					5️⃣ **If the question is NOT related to academics (science, maths, history, geography, computer science, biology, chemistry, college syllabus, etc.):**
					- Reply: "❌ This question is not related to education. I can answer only academic questions."

					============================
					📘 MATERIAL (Use this first)
					============================
					%s

					============================
					📘 USER QUESTION
					============================
					%s

					Remember:
					- STRICTLY follow the order: Material → Trusted Sites → Internet → Not Available
					- NEVER skip directly to "information not available".
					- ONLY search internet for academic questions.
					- ALWAYS refuse non-academic questions.
					`, allText, body.Question)






		resp,err:= model.GenerateContent(ctx,genai.Text(prompt))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Server Error"})
			return
		}

		answer := resp.Candidates[0].Content.Parts[0]

		c.JSON(200,gin.H{"answer":fmt.Sprintf("%v",answer)})

	}