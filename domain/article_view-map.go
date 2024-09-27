package domain

func ArticleViewToResp(v ArticleView) ArticleViewResp {
	return ArticleViewResp{
		Id:        v.Id,
		IdUser:    v.IdUser,
		ViewedAt:  v.ViewedAt,
		User: UserResp{
			Name:     v.User.Name,
			Email:    v.User.Email,
			Phone:    v.User.Phone,
			Address:  v.User.Address,
			Image:    v.User.Image,
			Birth:    v.User.Birth,
			JK:       v.User.JK,
			Nik:      v.User.Nik,
			ImageKtp: v.User.ImageKtp,
		},
		Article: []ArticleResp{},
	}
}

func ListArticleViewToResp(a []ArticleView) []ArticleViewResp {
	result := []ArticleViewResp{}
	for _, v := range a {
		data := ArticleViewToResp(v)
		result = append(result, data)
	}
	return result
}
