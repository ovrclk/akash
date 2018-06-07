package state

import (
	"github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/types/base"
)

type AccountAdapter interface {
	Save(account *types.Account) error
	Get(base.Bytes) (*types.Account, error)
	KeyFor(base.Bytes) base.Bytes
}

type accountAdapter struct {
	db DB
}

func NewAccountAdapter(db DB) AccountAdapter {
	return &accountAdapter{db}
}

func (a *accountAdapter) Save(account *types.Account) error {
	key := a.KeyFor(account.Address)
	return saveObject(a.db, key, account)
}

func (a *accountAdapter) Get(address base.Bytes) (*types.Account, error) {

	acc := types.Account{}

	key := a.KeyFor(address)

	buf := a.db.Get(key)
	if buf == nil {
		return nil, nil
	}

	acc.Unmarshal(buf)

	return &acc, nil
}

func (a *accountAdapter) KeyFor(address base.Bytes) base.Bytes {
	return append([]byte(AccountPath), address...)
}
