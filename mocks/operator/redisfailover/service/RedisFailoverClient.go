// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha2 "github.com/spotahome/redis-operator/api/redisfailover/v1alpha2"

// RedisFailoverClient is an autogenerated mock type for the RedisFailoverClient type
type RedisFailoverClient struct {
	mock.Mock
}

// EnsureNotPresentRedisService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) EnsureNotPresentRedisService(rFailover *v1alpha2.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureRedisConfigMap provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureRedisConfigMap(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureRedisService provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureRedisService(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureRedisShutdownConfigMap provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureRedisShutdownConfigMap(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureRedisStatefulset provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureRedisStatefulset(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureSentinelConfigMap provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureSentinelConfigMap(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureSentinelDeployment provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureSentinelDeployment(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureSentinelService provides a mock function with given fields: rFailover, labels, ownerRefs
func (_m *RedisFailoverClient) EnsureSentinelService(rFailover *v1alpha2.RedisFailover, labels map[string]string, ownerRefs []v1.OwnerReference) error {
	ret := _m.Called(rFailover, labels, ownerRefs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover, map[string]string, []v1.OwnerReference) error); ok {
		r0 = rf(rFailover, labels, ownerRefs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
