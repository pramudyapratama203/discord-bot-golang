package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func IsToxic(text string, apiKey string) (bool, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return false, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")

	model.SetTemperature(0.1)

	prompt := fmt.Sprintf(`Analisis apakah kalimat ini kasar, toksik, atau mengandung hinaan: "%s". 
	Jawab HANYA dengan satu kata: YA atau TIDAK.`, text)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return false, err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		result := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
		// Cek jawaban
		if strings.Contains(strings.ToUpper(result), "YA") {
			return true, nil
		}
	}
	return false, nil
}
