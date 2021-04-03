package usermanagement

import (
	"github.com/golover/anonymousbot/internal/pkg/datastore"
	"github.com/golover/anonymousbot/internal/pkg/utils"
)

func AddUser(userId string) string {
	uniqueId := LookupUserByUserID(userId)
	if  uniqueId != "" {
		return uniqueId
	}
	uniqueId = utils.RandStringRunes(8)
	datastore.Set("uid_"+userId,uniqueId)
	datastore.Set("unq_"+uniqueId,userId)
	return uniqueId
}