// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	entity "github.com/alkamalp/crm-golang/entity"
	mock "github.com/stretchr/testify/mock"
)

// ActorInterfaceRepo is an autogenerated mock type for the ActorInterfaceRepo type
type ActorInterfaceRepo struct {
	mock.Mock
}

// CreateActor provides a mock function with given fields: actor
func (_m *ActorInterfaceRepo) CreateActor(actor *entity.Actor) (*entity.Actor, error) {
	ret := _m.Called(actor)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) (*entity.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor) *entity.Actor); ok {
		r0 = rf(actor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteActor provides a mock function with given fields: username
func (_m *ActorInterfaceRepo) DeleteActor(username string) (interface{}, error) {
	ret := _m.Called(username)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActorById provides a mock function with given fields: id
func (_m *ActorInterfaceRepo) GetActorById(id uint) (entity.Actor, error) {
	ret := _m.Called(id)

	var r0 entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Actor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Actor)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginActor provides a mock function with given fields: actor
func (_m *ActorInterfaceRepo) LoginActor(actor *entity.Actor) (*entity.Actor, error) {
	ret := _m.Called(actor)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) (*entity.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor) *entity.Actor); ok {
		r0 = rf(actor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateActor provides a mock function with given fields: actor, id
func (_m *ActorInterfaceRepo) UpdateActor(actor *entity.Actor, id uint) (*entity.Actor, error) {
	ret := _m.Called(actor, id)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, uint) (*entity.Actor, error)); ok {
		return rf(actor, id)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor, uint) *entity.Actor); ok {
		r0 = rf(actor, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor, uint) error); ok {
		r1 = rf(actor, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewActorInterfaceRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewActorInterfaceRepo creates a new instance of ActorInterfaceRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActorInterfaceRepo(t mockConstructorTestingTNewActorInterfaceRepo) *ActorInterfaceRepo {
	mock := &ActorInterfaceRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
