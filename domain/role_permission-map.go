package domain

func RolePermissionToResp(v RolePermissions) RolePermissionResp {
	return RolePermissionResp{
		Id:        v.Id,
		CanCreate: v.CanCreate,
		CanRead:   v.CanRead,
		CanEdit:   v.CanEdit,
		CanDelete: v.CanDelete,
		IdRole:    v.IdRole,
	}
}

func ListRolePermissionToResp(v []RolePermissions) []RolePermissionResp {
	result := []RolePermissionResp{}
	for _, value := range v{
		data := RolePermissionToResp(value)
		result = append(result, data)
	}
	return result
}

func RolePermissionRespToRolePermission(resp RolePermissionResp) RolePermissions {
	return RolePermissions{
		Id:        resp.Id,
		CanCreate: resp.CanCreate,
		CanRead:   resp.CanRead,
		CanEdit:   resp.CanEdit,
		CanDelete: resp.CanDelete,
		IdRole:    resp.IdRole,
	}
}
