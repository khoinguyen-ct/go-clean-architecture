package repository

import (
	"context"
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/pkg/client"
)

type AdListingResponse struct {
	Ad AdInfo `json:"ad"`
}

type AdInfo struct {
	AdID       int64  `json:"ad_id"`
	ListID     int64  `json:"list_id"`
	CategoryID int64  `json:"category"`
	Body       string `json:"body"`
	Subject    string `json:"subject"`
}

type AdListingService interface {
	GetByAdID(ctx context.Context, adID int64) (model.AdListing, error)
	GetByListID(ctx context.Context, listID int64) (model.AdListing, error)
}

type adListingServiceImpl struct {
	httpClient *client.HTTPClient
}

func (a *adListingServiceImpl) GetByAdID(ctx context.Context, adID int64) (model.AdListing, error) {
	return model.AdListing{}, nil
}

func (a *adListingServiceImpl) GetByListID(ctx context.Context, listID int64) (ad model.AdListing, err error) {
	var adResp AdListingResponse
	url := fmt.Sprintf("%v/%v",config.GetConfiguration().AdListingDomain, listID)
	err = a.httpClient.SendHTTPRequest(ctx, "GET", url, nil, &adResp)
	ad = model.AdListing{
		AdID:       adResp.Ad.AdID,
		ListID:     adResp.Ad.ListID,
		CategoryID: adResp.Ad.CategoryID,
		Body:       adResp.Ad.Body,
		Subject:    adResp.Ad.Subject,
	}
	return ad, err
}

func NewAdListingService(ctx context.Context) AdListingService {
	httpCli := client.NewHttpClient("ad-listing")
	return &adListingServiceImpl{
		httpClient: httpCli,
	}
}
