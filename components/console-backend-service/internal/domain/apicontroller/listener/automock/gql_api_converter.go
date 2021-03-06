// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1alpha2 "github.com/kyma-project/kyma/components/api-controller/pkg/apis/gateway.kyma-project.io/v1alpha2"

// gqlApiConverter is an autogenerated mock type for the gqlApiConverter type
type gqlApiConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlApiConverter) ToGQL(in *v1alpha2.Api) *gqlschema.API {
	ret := _m.Called(in)

	var r0 *gqlschema.API
	if rf, ok := ret.Get(0).(func(*v1alpha2.Api) *gqlschema.API); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.API)
		}
	}

	return r0
}
