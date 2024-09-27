package service

import (
	"errors"
	"qhealth/domain"
	articleview "qhealth/features/article_view"
)

type service struct {
	repo articleview.Repository
}

func NewViewService(repo articleview.Repository) articleview.Service {
	return &service{repo: repo}
}

func (s *service) GetAllView() ([]domain.ArticleViewResp, error) {
    views, err := s.repo.GetAllView()
    if err != nil {
        return nil, err
    }

    var result []domain.ArticleViewResp

    userMap := make(map[string]*domain.ArticleViewResp)

    for _, view := range views {
        if _, exists := userMap[view.IdUser]; !exists {
            userMap[view.IdUser] = &domain.ArticleViewResp{
				Id: view.Id,
                IdUser: view.IdUser,
                User: domain.UserResp{
                    Name:     view.User.Name,
                    Email:    view.User.Email,
                    Phone:    view.User.Phone,
                    Address:  view.User.Address,
                    Image:    view.User.Image,
                    Birth:    view.User.Birth,
                    JK:       view.User.JK,
                    Nik:      view.User.Nik,
                    ImageKtp: view.User.ImageKtp,
                },
                Article: []domain.ArticleResp{},
            }
        }

        userMap[view.IdUser].Article = append(userMap[view.IdUser].Article, domain.ArticleResp{
            Id:      view.Article.Id,
            Writer:  view.Article.Writer,
            Title:   view.Article.Title,
            Content: view.Article.Content,
            Date:    view.Article.Date,
            Image:   view.Article.Image,
        })
    }

    for _, userView := range userMap {
        result = append(result, *userView)
    }

    return result, nil
}


func (s *service) GetArticleTrackView(userId, articleId string) (domain.Articles, error) {
	article, err := s.repo.GetArticleById(articleId)
	if err != nil {
		return domain.Articles{}, errors.New("article not found")
	}

	articleView := domain.ArticleView{
		IdUser: userId,
		IdArticle: articleId,
	}

	if err := s.repo.CreateArticleView(articleView); err != nil {
		return domain.Articles{}, errors.New("failed to track article view")
	}

	return article, nil
}
