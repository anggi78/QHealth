package domain

func ReqToArticle(a ArticleReq) Articles {
	return Articles{
		Writer:  a.Writer,
		Title:   a.Title,
		Content: a.Content,
		Date:    a.Date,
		Image:   a.Image,
	}
}

func ArticleToResp(a Articles) ArticleResp {
	return ArticleResp{
		Id:      a.Id,
		Writer:  a.Writer,
		Title:   a.Title,
		Content: a.Content,
		Date:    a.Date,
		Image:   a.Image,
		Status:  a.Status,
		User:    a.User,
	}
}

func ListArticleToResp(a []Articles) []ArticleResp {
	result := []ArticleResp{}
	for _, v := range a {
		data := ArticleToResp(v)
		result = append(result, data)
	}
	return result
}