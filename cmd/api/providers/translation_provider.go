package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"translatorBot/cmd/api/domain"
	"translatorBot/cmd/api/utils/apierrors"
)

const (
	LibreTranslateURL = "http://translator:5000/translate"

	ContentTypeHeaderKey       = "Content-Type"
	ApplicationJsonHeaderValue = "application/json"
)

type TranslationProviderInterface interface {
	Translate(ctx context.Context, request domain.TranslationRequest) (*domain.TranslationResponse, apierrors.ApiError)
}

type translationProvider struct {
	client *resty.Client
}

func (p *translationProvider) Translate(ctx context.Context, request domain.TranslationRequest) (*domain.TranslationResponse, apierrors.ApiError) {
	response, err :=
		p.client.R().SetHeader(ContentTypeHeaderKey, ApplicationJsonHeaderValue).SetBody(request).Post(LibreTranslateURL)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}

	var translationResp domain.TranslationResponse
	if unmarshalErr := json.Unmarshal(response.Body(), &translationResp); unmarshalErr != nil {
		fmt.Println(unmarshalErr.Error())
		return nil, apierrors.NewInternalServerApiError(unmarshalErr.Error(), unmarshalErr)
	}
	return &translationResp, nil
}
