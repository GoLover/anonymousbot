package redis
type RedisDataStore struct {}
func InitDataStore() RedisDataStore{
	return RedisDataStore{}
}

func Get(key string) string {
	return ""
}

func Set(key,value string) {
	return
}
