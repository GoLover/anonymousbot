package usermanagement

import "github.com/golover/anonymousbot/internal/pkg/datastore"

func LookupUserByUniqueId(uniqueId string) string {
	return datastore.Get("unq_"+uniqueId)
}

func LookupUserByUserID(userId string) string {
	return datastore.Get("uid_"+userId)
}