package stub_test

import (
	"testing"

	"github.com/rlafferty/twg/suite"
	"github.com/rlafferty/twg/suite/stub"
	"github.com/rlafferty/twg/suite/suitetest"
)

var _ suite.UserStore = &stub.UserStore{}

func TestUserStore(t *testing.T) {
	us := &stub.UserStore{}
	suitetest.UserStore(t, us, nil, nil)
}

func TestUserStore_withStruct(t *testing.T) {
	us := &stub.UserStore{}
	tests := suitetest.UserStoreSuite{
		UserStore: us,
	}
	tests.All(t)
}
