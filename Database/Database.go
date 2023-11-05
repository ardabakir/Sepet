package Database

type Database interface {
	GetItem(key string, keyName string, castTo interface{}) error
	PutItem(item interface{}) error
}
