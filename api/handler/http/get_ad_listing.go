package http

import (
	"github.com/labstack/echo/v4"
	"go-clean-architecture/internal/usecase"
	"net/http"
	"strconv"
)

func (sv *Server) GetAdListing(c echo.Context) (err error) {
	id, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)
	if err != nil {
		logger.Errorf("decode request basic err: %v", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get ad listing info",
		})
	}
	adListingUC := usecase.NewAdListingUC(sv.ctx)
	ad, err := adListingUC.GetByListID(sv.ctx, id)
	if err != nil {
		logger.Errorf("failed to get ad listing info err: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to get ad listing info",
		})
	}
	return c.JSON(http.StatusOK, ad)
}
