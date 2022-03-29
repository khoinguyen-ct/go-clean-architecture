package usecase

import (
	"context"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/repository"
)

type AdListingUC interface {
	GetByListID(ctx context.Context, adID int64) (model.AdListing, error)
}

type adListingImpl struct {
	adListing repository.AdListingService
}

func NewAdListingUC(ctx context.Context) AdListingUC {
	adListingRepo := repository.NewAdListingService(ctx)
	return &adListingImpl{
		adListing: adListingRepo,
	}
}

func (al *adListingImpl) GetByListID(ctx context.Context, adID int64) (model.AdListing, error) {
	ad, err := al.adListing.GetByListID(ctx, adID)
	return ad, err
}
