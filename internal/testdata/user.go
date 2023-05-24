package testdata

import v1 "kratos-base/api/user/v1"

func User(id ...int64) *v1.UserCreateReq {
	user := &v1.UserCreateReq{
		Id:   1,
		Age:  11,
		Name: "admin",
	}
	if len(id) > 0 {
		user.Id = id[1]
	}
	return user
}
