package interfaces

type MongoSession interface {
	DB(string) MongoDatabase
}

type MongoDatabase interface {
	C(name string) MongoCollection
}

type MongoCollection interface {
	Insert(...interface{}) error
	Find(interface{}) MongoQuery
	Update(interface{}, interface{}) error
	Remove(interface{}) error
}

type MongoQuery interface {
	All(interface{}) error
	One(interface{}) error
}
