package usermanagement

import "github.com/golover/anonymousbot/internal/pkg/datastore"

func SetCurrentContact(userId, contactUserId string) {
	datastore.Set("contact_"+userId,contactUserId)
}

func GetCurrentContact(userId string) string {
	return datastore.Get("contact_"+userId)
}
