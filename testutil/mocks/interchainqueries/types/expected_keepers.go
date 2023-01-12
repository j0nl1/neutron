// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../x/interchainqueries/types/expected_keepers.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockContractManagerKeeper is a mock of ContractManagerKeeper interface.
type MockContractManagerKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagerKeeperMockRecorder
}

// MockContractManagerKeeperMockRecorder is the mock recorder for MockContractManagerKeeper.
type MockContractManagerKeeperMockRecorder struct {
	mock *MockContractManagerKeeper
}

// NewMockContractManagerKeeper creates a new mock instance.
func NewMockContractManagerKeeper(ctrl *gomock.Controller) *MockContractManagerKeeper {
	mock := &MockContractManagerKeeper{ctrl: ctrl}
	mock.recorder = &MockContractManagerKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagerKeeper) EXPECT() *MockContractManagerKeeperMockRecorder {
	return m.recorder
}

// HasContractInfo mocks base method.
func (m *MockContractManagerKeeper) HasContractInfo(ctx types.Context, contractAddress types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasContractInfo", ctx, contractAddress)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasContractInfo indicates an expected call of HasContractInfo.
func (mr *MockContractManagerKeeperMockRecorder) HasContractInfo(ctx, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasContractInfo", reflect.TypeOf((*MockContractManagerKeeper)(nil).HasContractInfo), ctx, contractAddress)
}

// SudoError mocks base method.
func (m *MockContractManagerKeeper) SudoError(ctx types.Context, senderAddress types.AccAddress, request types1.Packet, details string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoError", ctx, senderAddress, request, details)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoError indicates an expected call of SudoError.
func (mr *MockContractManagerKeeperMockRecorder) SudoError(ctx, senderAddress, request, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoError", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoError), ctx, senderAddress, request, details)
}

// SudoKVQueryResult mocks base method.
func (m *MockContractManagerKeeper) SudoKVQueryResult(ctx types.Context, contractAddress types.AccAddress, queryID uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoKVQueryResult", ctx, contractAddress, queryID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoKVQueryResult indicates an expected call of SudoKVQueryResult.
func (mr *MockContractManagerKeeperMockRecorder) SudoKVQueryResult(ctx, contractAddress, queryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoKVQueryResult", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoKVQueryResult), ctx, contractAddress, queryID)
}

// SudoResponse mocks base method.
func (m *MockContractManagerKeeper) SudoResponse(ctx types.Context, senderAddress types.AccAddress, request types1.Packet, msg []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoResponse", ctx, senderAddress, request, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoResponse indicates an expected call of SudoResponse.
func (mr *MockContractManagerKeeperMockRecorder) SudoResponse(ctx, senderAddress, request, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoResponse", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoResponse), ctx, senderAddress, request, msg)
}

// SudoTimeout mocks base method.
func (m *MockContractManagerKeeper) SudoTimeout(ctx types.Context, senderAddress types.AccAddress, request types1.Packet) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoTimeout", ctx, senderAddress, request)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoTimeout indicates an expected call of SudoTimeout.
func (mr *MockContractManagerKeeperMockRecorder) SudoTimeout(ctx, senderAddress, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoTimeout", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoTimeout), ctx, senderAddress, request)
}

// SudoTxQueryResult mocks base method.
func (m *MockContractManagerKeeper) SudoTxQueryResult(ctx types.Context, contractAddress types.AccAddress, queryID uint64, height int64, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoTxQueryResult", ctx, contractAddress, queryID, height, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoTxQueryResult indicates an expected call of SudoTxQueryResult.
func (mr *MockContractManagerKeeperMockRecorder) SudoTxQueryResult(ctx, contractAddress, queryID, height, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoTxQueryResult", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoTxQueryResult), ctx, contractAddress, queryID, height, data)
}
