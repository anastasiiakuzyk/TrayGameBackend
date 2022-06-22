package game

import (
	"errors"
	"sync"
)

type Helpers struct {
	mp sync.Map
}

func NewHelpers() *Helpers {
	return &Helpers{sync.Map{}}
}

func (p *Helpers) Append(Helper *Helper) {
	p.mp.Store(Helper.UUID, Helper)
}
func (p *Helpers) Get(uuid string) (*Helper, error) {
	in, ok := p.mp.Load(uuid)
	if !ok {
		return nil, errors.New("Helper didn't exists")
	}

	return in.(*Helper), nil
}
func (p *Helpers) Remove(uuid string) {
	p.mp.Delete(uuid)
}

func (p *Helpers) Update(uuid string, Helper *Helper) {
	p.Remove(uuid)
	p.Append(Helper)
}

func (p *Helpers) ForEach(callback func(uuid string, Helper *Helper) bool) {
	p.mp.Range(func(key, value interface{}) bool {
		Helper := value.(*Helper)
		uuid := key.(string)
		return callback(uuid, Helper)
	})
}

func (p *Helpers) ToArray() []Helper {
	var Helpers []Helper = []Helper{}
	p.ForEach(func(uuid string, helper *Helper) bool {
		Helpers = append(Helpers, *helper)

		return true
	})

	return Helpers
}