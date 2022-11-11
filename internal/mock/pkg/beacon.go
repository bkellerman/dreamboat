// Code generated by MockGen. DO NOT EDIT.
// Source: beacon.go

// Package mock_relay is a generated GoMock package.
package mock_relay

import (
	context "context"
	reflect "reflect"

	relay "github.com/blocknative/dreamboat/pkg"
	gomock "github.com/golang/mock/gomock"
)

// MockBeaconClient is a mock of BeaconClient interface.
type MockBeaconClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconClientMockRecorder
}

// MockBeaconClientMockRecorder is the mock recorder for MockBeaconClient.
type MockBeaconClientMockRecorder struct {
	mock *MockBeaconClient
}

// NewMockBeaconClient creates a new mock instance.
func NewMockBeaconClient(ctrl *gomock.Controller) *MockBeaconClient {
	mock := &MockBeaconClient{ctrl: ctrl}
	mock.recorder = &MockBeaconClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconClient) EXPECT() *MockBeaconClientMockRecorder {
	return m.recorder
}

// Endpoint mocks base method.
func (m *MockBeaconClient) Endpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockBeaconClientMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockBeaconClient)(nil).Endpoint))
}

// Genesis mocks base method.
func (m *MockBeaconClient) Genesis() (*relay.GenesisResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*relay.GenesisResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis.
func (mr *MockBeaconClientMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockBeaconClient)(nil).Genesis))
}

// GetProposerDuties mocks base method.
func (m *MockBeaconClient) GetProposerDuties(arg0 relay.Epoch) (*relay.RegisteredProposersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerDuties", arg0)
	ret0, _ := ret[0].(*relay.RegisteredProposersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposerDuties indicates an expected call of GetProposerDuties.
func (mr *MockBeaconClientMockRecorder) GetProposerDuties(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerDuties", reflect.TypeOf((*MockBeaconClient)(nil).GetProposerDuties), arg0)
}

// KnownValidators mocks base method.
func (m *MockBeaconClient) KnownValidators(arg0 relay.Slot) (relay.AllValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KnownValidators", arg0)
	ret0, _ := ret[0].(relay.AllValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KnownValidators indicates an expected call of KnownValidators.
func (mr *MockBeaconClientMockRecorder) KnownValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownValidators", reflect.TypeOf((*MockBeaconClient)(nil).KnownValidators), arg0)
}

// SubscribeToHeadEvents mocks base method.
func (m *MockBeaconClient) SubscribeToHeadEvents(ctx context.Context, slotC chan relay.HeadEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubscribeToHeadEvents", ctx, slotC)
}

// SubscribeToHeadEvents indicates an expected call of SubscribeToHeadEvents.
func (mr *MockBeaconClientMockRecorder) SubscribeToHeadEvents(ctx, slotC interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToHeadEvents", reflect.TypeOf((*MockBeaconClient)(nil).SubscribeToHeadEvents), ctx, slotC)
}

// SyncStatus mocks base method.
func (m *MockBeaconClient) SyncStatus() (*relay.SyncStatusPayloadData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus")
	ret0, _ := ret[0].(*relay.SyncStatusPayloadData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockBeaconClientMockRecorder) SyncStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockBeaconClient)(nil).SyncStatus))
}
