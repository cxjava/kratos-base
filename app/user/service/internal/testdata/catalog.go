package testdata

import v1 "kratos-base/api/catalog/service/v1"

func Catalog(id ...int64) *v1.CreateBeerReply {
	catalog := &v1.CreateBeerReply{
		Id:   1,
		Name: "catalog_test",
	}
	if len(id) > 0 {
		catalog.Id = id[1]
	}
	return catalog
}
