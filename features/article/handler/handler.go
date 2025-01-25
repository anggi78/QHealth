package handler

import (
	"net/http"
	"qhealth/domain"
	"qhealth/features/article"
	"qhealth/helpers"
	"qhealth/helpers/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	serv article.Service
}

func NewArticleHandler(serv article.Service) article.Handler {
	return &handler{serv: serv}
}

func (h *handler) CreateArticle(e echo.Context) error {
    articleReq := domain.ArticleReq{}
    if err := e.Bind(&articleReq); err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    date, err := helpers.ParsedDate(articleReq.Date)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    articleReq.Date = date

    fileImage, err := e.FormFile("image")
    if err != nil && err != http.ErrMissingFile {
        return helpers.CustomErr(e, "error handling image file: " + err.Error())
    }

    client := helpers.ConfigCloud()
    imageUrl := helpers.UploadFile(fileImage, client)
    articleReq.Image = imageUrl

    err = h.serv.CreateArticle(articleReq)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully create article", nil))
}

func (h *handler) GetAllArticle(e echo.Context) error {
	pageStr := e.QueryParam("page")
	title := e.QueryParam("title")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10

	_, userEmail, err := middleware.ExtractToken(e)
    if err != nil {
        return helpers.CustomErr(e, "invalid token")
    }

    user, err := h.serv.GetUserByEmail(userEmail)
    if err != nil {
        return helpers.CustomErr(e, "user not found")
    }

	articleList, total, err := h.serv.GetAllArticle(title, user.Id, page, pageSize)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	currentPage, allPages := helpers.CalculatePaginationValues(page, pageSize, total)
	nextPage := helpers.GetNextPage(currentPage, allPages)
	prevPage := helpers.GetPrevPage(currentPage)

	pagination := helpers.PaginationResponse{
		CurrentPage: currentPage,
		NextPage: nextPage,
		PrevPage: prevPage,
		AllPages: allPages,
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponsePage("successfully get all data", articleList, pagination))
}

func (h *handler) GetLatestArticle(e echo.Context) error {
	latestArticle, err := h.serv.GetLatestArticle()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully retrieved latest article", latestArticle))
}

func (h *handler) GetArticleById(e echo.Context) error {
	id := e.Param("id")

	article, err := h.serv.GetArticleById(id)
	if err != nil{
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, err.Error())
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get data", article))
}

func (h *handler) UpdateArticle(e echo.Context) error {
	id := e.Param("id")

	articleReq := domain.ArticleReq{}
	if err := e.Bind(&articleReq); err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	fileImage, err := e.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return helpers.CustomErr(e, "error handling image file: " + err.Error())
	}

	client := helpers.ConfigCloud()
	imageUrl := helpers.UploadFile(fileImage, client)
	articleReq.Image = imageUrl

	err = h.serv.UpdateArticle(id, articleReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data", nil))
}

func (h *handler) DeleteArticle(e echo.Context) error {
	id := e.Param("id")
	
	err := h.serv.DeleteArticle(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, err.Error())
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully deleted data", nil))
}