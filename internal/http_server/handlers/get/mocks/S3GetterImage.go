// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3GetterImage is an autogenerated mock type for the S3GetterImage type
type S3GetterImage struct {
	mock.Mock
}

// GetObject provides a mock function with given fields: ctx, params, optFns
func (_m *S3GetterImage) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 *s3.GetObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) *s3.GetObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewS3GetterImage creates a new instance of S3GetterImage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewS3GetterImage(t interface {
	mock.TestingT
	Cleanup(func())
}) *S3GetterImage {
	mock := &S3GetterImage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}