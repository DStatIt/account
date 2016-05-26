package account

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MGOProxy struct {
	*mgo.Database
}

const (
	proxyC string = "proxys"
)

func (m *MGOProxy) Get(id bson.ObjectId) (*Proxy, error) {
	var prox Proxy
	err := m.C(proxyC).FindId(id).One(&prox)
	return &prox, err
}

func (m *MGOProxy) Delete(id bson.ObjectId) error {
	return m.C(proxyC).RemoveId(id)
}

func (m *MGOProxy) Insert(p *Proxy) error {
	return m.C(proxyC).Insert(bson.M{"$set": p})
}

func (m *MGOProxy) Update(id bson.ObjectId, p *Proxy) error {
	_, err := m.C(proxyC).UpsertId(id, p)
	return err
}
