package account

import "gopkg.in/mgo.v2/bson"

type AccountManager interface {
	Get(id bson.ObjectId) (*Account, error)
	Delete(id bson.ObjectId) error
	Insert(Account) error
	Update(id bson.ObjectId, a Account) error
	AddProxy(id bson.ObjectId, p Proxy) error
}

type ProxyManager interface {
	Get(id bson.ObjectId) (*Proxy, error)
	Delete(id bson.ObjectId) error
	Insert(Proxy) error
	Update(id bson.ObjectId, a *Account) error
}
