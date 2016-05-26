package account

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MGOAccount struct {
	*mgo.Database
}

const (
	accountC string = "accounts"
)

func (m *MGOAccount) Get(id bson.ObjectId) (*Account, error) {
	var acc Account
	err := m.C(accountC).FindId(id).One(&acc)
	return &acc, err
}

func (m *MGOAccount) Delete(id bson.ObjectId) error {
	return m.C(accountC).RemoveId(id)
}

func (m *MGOAccount) Insert(a *Account) error {
	return m.C(accountC).Insert(a)
}

func (m *MGOAccount) Update(id bson.ObjectId, a *Account) error {
	_, err := m.C(accountC).UpsertId(id, bson.M{"$set": a})
	return err
}

func (m *MGOAccount) AddProxy(id bson.ObjectId, p Proxy) error {
	_, err := m.C(accountC).UpsertId(
		id, bson.M{"$set": bson.M{"proxy": bson.M{
			"IP":   p.IP,
			"Port": p.Port,
		},
		},
		},
	)
	return err
}
