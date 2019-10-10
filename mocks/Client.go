// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"
import xf "github.com/motain/xfetch-go"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Read provides a mock function with given fields: key, fetchable
func (_m *Client) Read(key string, fetchable xf.Fetchable) (float64, float64, error) {
	ret := _m.Called(key, fetchable)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string, xf.Fetchable) float64); ok {
		r0 = rf(key, fetchable)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(string, xf.Fetchable) float64); ok {
		r1 = rf(key, fetchable)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, xf.Fetchable) error); ok {
		r2 = rf(key, fetchable)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: key, ttl, delta, fetchable
func (_m *Client) Update(key string, ttl time.Duration, delta float64, fetchable xf.Fetchable) error {
	ret := _m.Called(key, ttl, delta, fetchable)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Duration, float64, xf.Fetchable) error); ok {
		r0 = rf(key, ttl, delta, fetchable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
