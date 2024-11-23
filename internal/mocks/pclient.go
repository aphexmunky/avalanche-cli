// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/ava-labs/avalanchego/api"

	gas "github.com/ava-labs/avalanchego/vms/components/gas"

	ids "github.com/ava-labs/avalanchego/ids"

	mock "github.com/stretchr/testify/mock"

	platformvm "github.com/ava-labs/avalanchego/vms/platformvm"

	platformvmapi "github.com/ava-labs/avalanchego/vms/platformvm/api"

	rpc "github.com/ava-labs/avalanchego/utils/rpc"

	secp256k1 "github.com/ava-labs/avalanchego/utils/crypto/secp256k1"

	status "github.com/ava-labs/avalanchego/vms/platformvm/status"

	time "time"

	validators "github.com/ava-labs/avalanchego/snow/validators"
)

// PClient is an autogenerated mock type for the Client type
type PClient struct {
	mock.Mock
}

// ExportKey provides a mock function with given fields: ctx, user, address, options
func (_m *PClient) ExportKey(ctx context.Context, user api.UserPass, address ids.ShortID, options ...rpc.Option) (*secp256k1.PrivateKey, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportKey")
	}

	var r0 *secp256k1.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ids.ShortID, ...rpc.Option) (*secp256k1.PrivateKey, error)); ok {
		return rf(ctx, user, address, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ids.ShortID, ...rpc.Option) *secp256k1.PrivateKey); ok {
		r0 = rf(ctx, user, address, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secp256k1.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, ids.ShortID, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, address, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAtomicUTXOs provides a mock function with given fields: ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options
func (_m *PClient) GetAtomicUTXOs(ctx context.Context, addrs []ids.ShortID, sourceChain string, limit uint32, startAddress ids.ShortID, startUTXOID ids.ID, options ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs, sourceChain, limit, startAddress, startUTXOID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAtomicUTXOs")
	}

	var r0 [][]byte
	var r1 ids.ShortID
	var r2 ids.ID
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error)); ok {
		return rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) [][]byte); ok {
		r0 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ShortID); ok {
		r1 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ids.ShortID)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ID); ok {
		r2 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(ids.ID)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) error); ok {
		r3 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetBalance provides a mock function with given fields: ctx, addrs, options
func (_m *PClient) GetBalance(ctx context.Context, addrs []ids.ShortID, options ...rpc.Option) (*platformvm.GetBalanceResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 *platformvm.GetBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, ...rpc.Option) (*platformvm.GetBalanceResponse, error)); ok {
		return rf(ctx, addrs, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, ...rpc.Option) *platformvm.GetBalanceResponse); ok {
		r0 = rf(ctx, addrs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformvm.GetBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, ...rpc.Option) error); ok {
		r1 = rf(ctx, addrs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: ctx, blockID, options
func (_m *PClient) GetBlock(ctx context.Context, blockID ids.ID, options ...rpc.Option) ([]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBlock")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ([]byte, error)); ok {
		return rf(ctx, blockID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []byte); ok {
		r0 = rf(ctx, blockID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHeight provides a mock function with given fields: ctx, height, options
func (_m *PClient) GetBlockByHeight(ctx context.Context, height uint64, options ...rpc.Option) ([]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, height)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockByHeight")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, ...rpc.Option) ([]byte, error)); ok {
		return rf(ctx, height, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, ...rpc.Option) []byte); ok {
		r0 = rf(ctx, height, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, height, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchainStatus provides a mock function with given fields: ctx, blockchainID, options
func (_m *PClient) GetBlockchainStatus(ctx context.Context, blockchainID string, options ...rpc.Option) (status.BlockchainStatus, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockchainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockchainStatus")
	}

	var r0 status.BlockchainStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) (status.BlockchainStatus, error)); ok {
		return rf(ctx, blockchainID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) status.BlockchainStatus); ok {
		r0 = rf(ctx, blockchainID, options...)
	} else {
		r0 = ret.Get(0).(status.BlockchainStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockchainID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchains provides a mock function with given fields: ctx, options
func (_m *PClient) GetBlockchains(ctx context.Context, options ...rpc.Option) ([]platformvm.APIBlockchain, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockchains")
	}

	var r0 []platformvm.APIBlockchain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) ([]platformvm.APIBlockchain, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) []platformvm.APIBlockchain); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.APIBlockchain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentSupply provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetCurrentSupply(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentSupply")
	}

	var r0 uint64
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (uint64, uint64, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r2 = rf(ctx, subnetID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCurrentValidators provides a mock function with given fields: ctx, subnetID, nodeIDs, options
func (_m *PClient) GetCurrentValidators(ctx context.Context, subnetID ids.ID, nodeIDs []ids.NodeID, options ...rpc.Option) ([]platformvm.ClientPermissionlessValidator, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, nodeIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentValidators")
	}

	var r0 []platformvm.ClientPermissionlessValidator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) ([]platformvm.ClientPermissionlessValidator, error)); ok {
		return rf(ctx, subnetID, nodeIDs, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) []platformvm.ClientPermissionlessValidator); ok {
		r0 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.ClientPermissionlessValidator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeeConfig provides a mock function with given fields: ctx, options
func (_m *PClient) GetFeeConfig(ctx context.Context, options ...rpc.Option) (*gas.Config, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFeeConfig")
	}

	var r0 *gas.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) (*gas.Config, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) *gas.Config); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gas.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeeState provides a mock function with given fields: ctx, options
func (_m *PClient) GetFeeState(ctx context.Context, options ...rpc.Option) (gas.State, gas.Price, time.Time, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFeeState")
	}

	var r0 gas.State
	var r1 gas.Price
	var r2 time.Time
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) (gas.State, gas.Price, time.Time, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) gas.State); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(gas.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) gas.Price); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Get(1).(gas.Price)
	}

	if rf, ok := ret.Get(2).(func(context.Context, ...rpc.Option) time.Time); ok {
		r2 = rf(ctx, options...)
	} else {
		r2 = ret.Get(2).(time.Time)
	}

	if rf, ok := ret.Get(3).(func(context.Context, ...rpc.Option) error); ok {
		r3 = rf(ctx, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetHeight provides a mock function with given fields: ctx, options
func (_m *PClient) GetHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetHeight")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) (uint64, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1Validator provides a mock function with given fields: ctx, validationID, options
func (_m *PClient) GetL1Validator(ctx context.Context, validationID ids.ID, options ...rpc.Option) (platformvm.L1Validator, uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, validationID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetL1Validator")
	}

	var r0 platformvm.L1Validator
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (platformvm.L1Validator, uint64, error)); ok {
		return rf(ctx, validationID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) platformvm.L1Validator); ok {
		r0 = rf(ctx, validationID, options...)
	} else {
		r0 = ret.Get(0).(platformvm.L1Validator)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r1 = rf(ctx, validationID, options...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r2 = rf(ctx, validationID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMinStake provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetMinStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMinStake")
	}

	var r0 uint64
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (uint64, uint64, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r2 = rf(ctx, subnetID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProposedHeight provides a mock function with given fields: ctx, options
func (_m *PClient) GetProposedHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProposedHeight")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) (uint64, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRewardUTXOs provides a mock function with given fields: _a0, _a1, _a2
func (_m *PClient) GetRewardUTXOs(_a0 context.Context, _a1 *api.GetTxArgs, _a2 ...rpc.Option) ([][]byte, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRewardUTXOs")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetTxArgs, ...rpc.Option) ([][]byte, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetTxArgs, ...rpc.Option) [][]byte); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetTxArgs, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStake provides a mock function with given fields: ctx, addrs, validatorsOnly, options
func (_m *PClient) GetStake(ctx context.Context, addrs []ids.ShortID, validatorsOnly bool, options ...rpc.Option) (map[ids.ID]uint64, [][]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs, validatorsOnly)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetStake")
	}

	var r0 map[ids.ID]uint64
	var r1 [][]byte
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, bool, ...rpc.Option) (map[ids.ID]uint64, [][]byte, error)); ok {
		return rf(ctx, addrs, validatorsOnly, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, bool, ...rpc.Option) map[ids.ID]uint64); ok {
		r0 = rf(ctx, addrs, validatorsOnly, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ids.ID]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, bool, ...rpc.Option) [][]byte); ok {
		r1 = rf(ctx, addrs, validatorsOnly, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, bool, ...rpc.Option) error); ok {
		r2 = rf(ctx, addrs, validatorsOnly, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStakingAssetID provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetStakingAssetID(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetStakingAssetID")
	}

	var r0 ids.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (ids.ID, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubnet provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetSubnet(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (platformvm.GetSubnetClientResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSubnet")
	}

	var r0 platformvm.GetSubnetClientResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (platformvm.GetSubnetClientResponse, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) platformvm.GetSubnetClientResponse); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(platformvm.GetSubnetClientResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubnets provides a mock function with given fields: ctx, subnetIDs, options
func (_m *PClient) GetSubnets(ctx context.Context, subnetIDs []ids.ID, options ...rpc.Option) ([]platformvm.ClientSubnet, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSubnets")
	}

	var r0 []platformvm.ClientSubnet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ID, ...rpc.Option) ([]platformvm.ClientSubnet, error)); ok {
		return rf(ctx, subnetIDs, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ID, ...rpc.Option) []platformvm.ClientSubnet); ok {
		r0 = rf(ctx, subnetIDs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.ClientSubnet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetIDs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTimestamp provides a mock function with given fields: ctx, options
func (_m *PClient) GetTimestamp(ctx context.Context, options ...rpc.Option) (time.Time, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTimestamp")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) (time.Time, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) time.Time); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalStake provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetTotalStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalStake")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (uint64, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTx provides a mock function with given fields: ctx, txID, options
func (_m *PClient) GetTx(ctx context.Context, txID ids.ID, options ...rpc.Option) ([]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, txID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTx")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ([]byte, error)); ok {
		return rf(ctx, txID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []byte); ok {
		r0 = rf(ctx, txID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, txID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxStatus provides a mock function with given fields: ctx, txID, options
func (_m *PClient) GetTxStatus(ctx context.Context, txID ids.ID, options ...rpc.Option) (*platformvm.GetTxStatusResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, txID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTxStatus")
	}

	var r0 *platformvm.GetTxStatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (*platformvm.GetTxStatusResponse, error)); ok {
		return rf(ctx, txID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) *platformvm.GetTxStatusResponse); ok {
		r0 = rf(ctx, txID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformvm.GetTxStatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, txID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUTXOs provides a mock function with given fields: ctx, addrs, limit, startAddress, startUTXOID, options
func (_m *PClient) GetUTXOs(ctx context.Context, addrs []ids.ShortID, limit uint32, startAddress ids.ShortID, startUTXOID ids.ID, options ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs, limit, startAddress, startUTXOID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetUTXOs")
	}

	var r0 [][]byte
	var r1 ids.ShortID
	var r2 ids.ID
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error)); ok {
		return rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) [][]byte); ok {
		r0 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ShortID); ok {
		r1 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ids.ShortID)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ID); ok {
		r2 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(ids.ID)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) error); ok {
		r3 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetValidatorsAt provides a mock function with given fields: ctx, subnetID, height, options
func (_m *PClient) GetValidatorsAt(ctx context.Context, subnetID ids.ID, height platformvmapi.Height, options ...rpc.Option) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, height)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetValidatorsAt")
	}

	var r0 map[ids.NodeID]*validators.GetValidatorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, platformvmapi.Height, ...rpc.Option) (map[ids.NodeID]*validators.GetValidatorOutput, error)); ok {
		return rf(ctx, subnetID, height, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, platformvmapi.Height, ...rpc.Option) map[ids.NodeID]*validators.GetValidatorOutput); ok {
		r0 = rf(ctx, subnetID, height, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ids.NodeID]*validators.GetValidatorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, platformvmapi.Height, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, height, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IssueTx provides a mock function with given fields: ctx, tx, options
func (_m *PClient) IssueTx(ctx context.Context, tx []byte, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IssueTx")
	}

	var r0 ids.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, ...rpc.Option) (ids.ID, error)); ok {
		return rf(ctx, tx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, tx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, ...rpc.Option) error); ok {
		r1 = rf(ctx, tx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAddresses provides a mock function with given fields: ctx, user, options
func (_m *PClient) ListAddresses(ctx context.Context, user api.UserPass, options ...rpc.Option) ([]ids.ShortID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAddresses")
	}

	var r0 []ids.ShortID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ...rpc.Option) ([]ids.ShortID, error)); ok {
		return rf(ctx, user, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ...rpc.Option) []ids.ShortID); ok {
		r0 = rf(ctx, user, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.ShortID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SampleValidators provides a mock function with given fields: ctx, subnetID, sampleSize, options
func (_m *PClient) SampleValidators(ctx context.Context, subnetID ids.ID, sampleSize uint16, options ...rpc.Option) ([]ids.NodeID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, sampleSize)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SampleValidators")
	}

	var r0 []ids.NodeID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, uint16, ...rpc.Option) ([]ids.NodeID, error)); ok {
		return rf(ctx, subnetID, sampleSize, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, uint16, ...rpc.Option) []ids.NodeID); ok {
		r0 = rf(ctx, subnetID, sampleSize, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.NodeID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, uint16, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, sampleSize, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatedBy provides a mock function with given fields: ctx, blockchainID, options
func (_m *PClient) ValidatedBy(ctx context.Context, blockchainID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockchainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ValidatedBy")
	}

	var r0 ids.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) (ids.ID, error)); ok {
		return rf(ctx, blockchainID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, blockchainID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockchainID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validates provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) Validates(ctx context.Context, subnetID ids.ID, options ...rpc.Option) ([]ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Validates")
	}

	var r0 []ids.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ([]ids.ID, error)); ok {
		return rf(ctx, subnetID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []ids.ID); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPClient creates a new instance of PClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PClient {
	mock := &PClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
