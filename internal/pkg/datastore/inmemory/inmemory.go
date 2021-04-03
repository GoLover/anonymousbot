package inmemory

type InMemoryDataStore struct {
}
var inMemoryObject map[string]string
var initializedimds InMemoryDataStore
func InitDataStore() InMemoryDataStore {
	inMemoryObject = make(map[string]string)
	return initializedimds
}

func (imds InMemoryDataStore) Get(key string) string {
	return inMemoryObject[key]
}

func (imds InMemoryDataStore) Set(key,value string) {
	inMemoryObject[key] = value
}