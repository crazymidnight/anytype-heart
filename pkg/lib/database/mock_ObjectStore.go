// Code generated by mockery v2.34.1. DO NOT EDIT.

package database

import mock "github.com/stretchr/testify/mock"

// MockObjectStore is an autogenerated mock type for the ObjectStore type
type MockObjectStore struct {
	mock.Mock
}

type MockObjectStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObjectStore) EXPECT() *MockObjectStore_Expecter {
	return &MockObjectStore_Expecter{mock: &_m.Mock}
}

// Query provides a mock function with given fields: q
func (_m *MockObjectStore) Query(q Query) ([]Record, int, error) {
	ret := _m.Called(q)

	var r0 []Record
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(Query) ([]Record, int, error)); ok {
		return rf(q)
	}
	if rf, ok := ret.Get(0).(func(Query) []Record); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Record)
		}
	}

	if rf, ok := ret.Get(1).(func(Query) int); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(Query) error); ok {
		r2 = rf(q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockObjectStore_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockObjectStore_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - q Query
func (_e *MockObjectStore_Expecter) Query(q interface{}) *MockObjectStore_Query_Call {
	return &MockObjectStore_Query_Call{Call: _e.mock.On("Query", q)}
}

func (_c *MockObjectStore_Query_Call) Run(run func(q Query)) *MockObjectStore_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Query))
	})
	return _c
}

func (_c *MockObjectStore_Query_Call) Return(records []Record, total int, err error) *MockObjectStore_Query_Call {
	_c.Call.Return(records, total, err)
	return _c
}

func (_c *MockObjectStore_Query_Call) RunAndReturn(run func(Query) ([]Record, int, error)) *MockObjectStore_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRaw provides a mock function with given fields: filters, limit, offset
func (_m *MockObjectStore) QueryRaw(filters *Filters, limit int, offset int) ([]Record, error) {
	ret := _m.Called(filters, limit, offset)

	var r0 []Record
	var r1 error
	if rf, ok := ret.Get(0).(func(*Filters, int, int) ([]Record, error)); ok {
		return rf(filters, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(*Filters, int, int) []Record); ok {
		r0 = rf(filters, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Record)
		}
	}

	if rf, ok := ret.Get(1).(func(*Filters, int, int) error); ok {
		r1 = rf(filters, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockObjectStore_QueryRaw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRaw'
type MockObjectStore_QueryRaw_Call struct {
	*mock.Call
}

// QueryRaw is a helper method to define mock.On call
//   - filters *Filters
//   - limit int
//   - offset int
func (_e *MockObjectStore_Expecter) QueryRaw(filters interface{}, limit interface{}, offset interface{}) *MockObjectStore_QueryRaw_Call {
	return &MockObjectStore_QueryRaw_Call{Call: _e.mock.On("QueryRaw", filters, limit, offset)}
}

func (_c *MockObjectStore_QueryRaw_Call) Run(run func(filters *Filters, limit int, offset int)) *MockObjectStore_QueryRaw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Filters), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockObjectStore_QueryRaw_Call) Return(_a0 []Record, _a1 error) *MockObjectStore_QueryRaw_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockObjectStore_QueryRaw_Call) RunAndReturn(run func(*Filters, int, int) ([]Record, error)) *MockObjectStore_QueryRaw_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObjectStore creates a new instance of MockObjectStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObjectStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObjectStore {
	mock := &MockObjectStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
