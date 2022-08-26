// Code generated by MockGen. DO NOT EDIT.
// Source: vms/platformvm/txs/mempool/mempool.go

// Package mempool is a generated GoMock package.
package mempool

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockBlockTimer is a mock of BlockTimer interface.
type MockBlockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockBlockTimerMockRecorder
}

// MockBlockTimerMockRecorder is the mock recorder for MockBlockTimer.
type MockBlockTimerMockRecorder struct {
	mock *MockBlockTimer
}

// NewMockBlockTimer creates a new mock instance.
func NewMockBlockTimer(ctrl *gomock.Controller) *MockBlockTimer {
	mock := &MockBlockTimer{ctrl: ctrl}
	mock.recorder = &MockBlockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockTimer) EXPECT() *MockBlockTimerMockRecorder {
	return m.recorder
}

// ResetBlockTimer mocks base method.
func (m *MockBlockTimer) ResetBlockTimer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetBlockTimer")
}

// ResetBlockTimer indicates an expected call of ResetBlockTimer.
func (mr *MockBlockTimerMockRecorder) ResetBlockTimer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetBlockTimer", reflect.TypeOf((*MockBlockTimer)(nil).ResetBlockTimer))
}

// MockMempool is a mock of Mempool interface.
type MockMempool struct {
	ctrl     *gomock.Controller
	recorder *MockMempoolMockRecorder
}

// MockMempoolMockRecorder is the mock recorder for MockMempool.
type MockMempoolMockRecorder struct {
	mock *MockMempool
}

// NewMockMempool creates a new mock instance.
func NewMockMempool(ctrl *gomock.Controller) *MockMempool {
	mock := &MockMempool{ctrl: ctrl}
	mock.recorder = &MockMempoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMempool) EXPECT() *MockMempoolMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMempool) Add(tx *txs.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMempoolMockRecorder) Add(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMempool)(nil).Add), tx)
}

// DisableAdding mocks base method.
func (m *MockMempool) DisableAdding() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisableAdding")
}

// DisableAdding indicates an expected call of DisableAdding.
func (mr *MockMempoolMockRecorder) DisableAdding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAdding", reflect.TypeOf((*MockMempool)(nil).DisableAdding))
}

// EnableAdding mocks base method.
func (m *MockMempool) EnableAdding() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableAdding")
}

// EnableAdding indicates an expected call of EnableAdding.
func (mr *MockMempoolMockRecorder) EnableAdding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAdding", reflect.TypeOf((*MockMempool)(nil).EnableAdding))
}

// Get mocks base method.
func (m *MockMempool) Get(txID ids.ID) *txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", txID)
	ret0, _ := ret[0].(*txs.Tx)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockMempoolMockRecorder) Get(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMempool)(nil).Get), txID)
}

// GetDropReason mocks base method.
func (m *MockMempool) GetDropReason(txID ids.ID) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDropReason", txID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDropReason indicates an expected call of GetDropReason.
func (mr *MockMempoolMockRecorder) GetDropReason(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDropReason", reflect.TypeOf((*MockMempool)(nil).GetDropReason), txID)
}

// Has mocks base method.
func (m *MockMempool) Has(txID ids.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", txID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMempoolMockRecorder) Has(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMempool)(nil).Has), txID)
}

// HasDecisionTxs mocks base method.
func (m *MockMempool) HasDecisionTxs() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasDecisionTxs")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasDecisionTxs indicates an expected call of HasDecisionTxs.
func (mr *MockMempoolMockRecorder) HasDecisionTxs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasDecisionTxs", reflect.TypeOf((*MockMempool)(nil).HasDecisionTxs))
}

// HasProposalTx mocks base method.
func (m *MockMempool) HasProposalTx() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasProposalTx")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasProposalTx indicates an expected call of HasProposalTx.
func (mr *MockMempoolMockRecorder) HasProposalTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasProposalTx", reflect.TypeOf((*MockMempool)(nil).HasProposalTx))
}

// MarkDropped mocks base method.
func (m *MockMempool) MarkDropped(txID ids.ID, reason string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkDropped", txID, reason)
}

// MarkDropped indicates an expected call of MarkDropped.
func (mr *MockMempoolMockRecorder) MarkDropped(txID, reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDropped", reflect.TypeOf((*MockMempool)(nil).MarkDropped), txID, reason)
}

// PeekDecisionTxs mocks base method.
func (m *MockMempool) PeekDecisionTxs(maxTxsBytes int) []*txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekDecisionTxs", maxTxsBytes)
	ret0, _ := ret[0].([]*txs.Tx)
	return ret0
}

// PeekDecisionTxs indicates an expected call of PeekDecisionTxs.
func (mr *MockMempoolMockRecorder) PeekDecisionTxs(maxTxsBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekDecisionTxs", reflect.TypeOf((*MockMempool)(nil).PeekDecisionTxs), maxTxsBytes)
}

// PeekProposalTx mocks base method.
func (m *MockMempool) PeekProposalTx() *txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekProposalTx")
	ret0, _ := ret[0].(*txs.Tx)
	return ret0
}

// PeekProposalTx indicates an expected call of PeekProposalTx.
func (mr *MockMempoolMockRecorder) PeekProposalTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekProposalTx", reflect.TypeOf((*MockMempool)(nil).PeekProposalTx))
}

// PopDecisionTxs mocks base method.
func (m *MockMempool) PopDecisionTxs(maxTxsBytes int) []*txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopDecisionTxs", maxTxsBytes)
	ret0, _ := ret[0].([]*txs.Tx)
	return ret0
}

// PopDecisionTxs indicates an expected call of PopDecisionTxs.
func (mr *MockMempoolMockRecorder) PopDecisionTxs(maxTxsBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopDecisionTxs", reflect.TypeOf((*MockMempool)(nil).PopDecisionTxs), maxTxsBytes)
}

// PopProposalTx mocks base method.
func (m *MockMempool) PopProposalTx() *txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopProposalTx")
	ret0, _ := ret[0].(*txs.Tx)
	return ret0
}

// PopProposalTx indicates an expected call of PopProposalTx.
func (mr *MockMempoolMockRecorder) PopProposalTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopProposalTx", reflect.TypeOf((*MockMempool)(nil).PopProposalTx))
}

// Remove mocks base method.
func (m *MockMempool) Remove(txs []*txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", txs)
}

// Remove indicates an expected call of Remove.
func (mr *MockMempoolMockRecorder) Remove(txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMempool)(nil).Remove), txs)
}

// RemoveDecisionTxs mocks base method.
func (m *MockMempool) RemoveDecisionTxs(txs []*txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveDecisionTxs", txs)
}

// RemoveDecisionTxs indicates an expected call of RemoveDecisionTxs.
func (mr *MockMempoolMockRecorder) RemoveDecisionTxs(txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDecisionTxs", reflect.TypeOf((*MockMempool)(nil).RemoveDecisionTxs), txs)
}

// RemoveProposalTx mocks base method.
func (m *MockMempool) RemoveProposalTx(tx *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveProposalTx", tx)
}

// RemoveProposalTx indicates an expected call of RemoveProposalTx.
func (mr *MockMempoolMockRecorder) RemoveProposalTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProposalTx", reflect.TypeOf((*MockMempool)(nil).RemoveProposalTx), tx)
}