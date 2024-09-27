package handler

import (
	"net/http"
	articleview "qhealth/features/article_view"
	"qhealth/helpers"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
)

type handler struct {
	serv articleview.Service
}

func NewViewHandler(serv articleview.Service) articleview.Handler {
	return &handler{serv: serv}
}

// func (h *handler) GetAllView(e echo.Context) error {
// 	viewList, err := h.serv.GetAllView()
// 	if err != nil {
// 		return helpers.CustomErr(e, err.Error())
// 	}

// 	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", viewList))
// }

func (h *handler) GetAllView(e echo.Context) error {
    viewList, err := h.serv.GetAllView()
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", viewList))
}

func (h *handler) GetArticleTrackView(e echo.Context) error {
    articleId := e.Param("id")

    userId, _, err := middleware.ExtractToken(e)  
    if err != nil {
        return helpers.CustomErr(e, "invalid token")
    }

    article, err := h.serv.GetArticleTrackView(userId, articleId)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", article))
}