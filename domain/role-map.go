package domain

func ReqToRole(r RoleReq) Role {
	return Role{
		Name: r.Name,
	}
}

func RoleToResp(r Role) RoleResp {
	return RoleResp{
		Id:   r.Id,
		Name: r.Name,
	}
}

func ListRoleToResp(r []Role) []RoleResp {
	result := []RoleResp{}
	for _, v := range r {
		data := RoleToResp(v)
		result = append(result, data)
	}
	return result
}