package service

import (
	"go_concurrent_server/db"
	"go_concurrent_server/resource"
)

func ServiceClient(input resource.Request) (resource.Response, error) {
	result, err := db.ServiceClient(input)
	if err != nil {
		return result, err

	}

	return result, nil

}
