package domain

func ReqToArticle(a ArticleReq, userId string) Articles {
	return Articles{
		Writer:  a.Writer,
		Title:   a.Title,
		Content: a.Content,
		Date:    a.Date,
		Image:   a.Image,
		IdUser:  userId,
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
		User: Users{
			Name:     a.User.Name,
			Email:    a.User.Email,
			Phone:    a.User.Phone,
			Address:  a.User.Address,
			Image:    a.User.Image,
			Birth:    a.User.Birth,
			JK:       a.User.JK,
			Nik:      a.User.Nik,
			ImageKtp: a.User.ImageKtp,
		},
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
