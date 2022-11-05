package domain

type TranslationRequest struct {
	Quote  string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
	Format string `json:"format"`
	ApiKey string `json:"api_key"`
}

type TranslationResponse struct {
	TranslatedText string `json:"translatedText"`
}
