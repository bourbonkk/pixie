// Code generated by MockGen. DO NOT EDIT.
// Source: agent.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	pl_shared_k8s_metadatapb "pixielabs.ai/pixielabs/src/shared/k8s/metadatapb"
	controllers "pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers"
	metadatapb "pixielabs.ai/pixielabs/src/vizier/services/metadata/metadatapb"
	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
	pl_vizier_services_shared_agent "pixielabs.ai/pixielabs/src/vizier/services/shared/agentpb"
	reflect "reflect"
)

// MockAgentManager is a mock of AgentManager interface
type MockAgentManager struct {
	ctrl     *gomock.Controller
	recorder *MockAgentManagerMockRecorder
}

// MockAgentManagerMockRecorder is the mock recorder for MockAgentManager
type MockAgentManagerMockRecorder struct {
	mock *MockAgentManager
}

// NewMockAgentManager creates a new mock instance
func NewMockAgentManager(ctrl *gomock.Controller) *MockAgentManager {
	mock := &MockAgentManager{ctrl: ctrl}
	mock.recorder = &MockAgentManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentManager) EXPECT() *MockAgentManagerMockRecorder {
	return m.recorder
}

// RegisterAgent mocks base method
func (m *MockAgentManager) RegisterAgent(info *pl_vizier_services_shared_agent.Agent) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAgent", info)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAgent indicates an expected call of RegisterAgent
func (mr *MockAgentManagerMockRecorder) RegisterAgent(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAgent", reflect.TypeOf((*MockAgentManager)(nil).RegisterAgent), info)
}

// UpdateHeartbeat mocks base method
func (m *MockAgentManager) UpdateHeartbeat(agentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHeartbeat", agentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHeartbeat indicates an expected call of UpdateHeartbeat
func (mr *MockAgentManagerMockRecorder) UpdateHeartbeat(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHeartbeat", reflect.TypeOf((*MockAgentManager)(nil).UpdateHeartbeat), agentID)
}

// UpdateAgent mocks base method
func (m *MockAgentManager) UpdateAgent(info *pl_vizier_services_shared_agent.Agent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgent", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgent indicates an expected call of UpdateAgent
func (mr *MockAgentManagerMockRecorder) UpdateAgent(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgent", reflect.TypeOf((*MockAgentManager)(nil).UpdateAgent), info)
}

// DeleteAgent mocks base method
func (m *MockAgentManager) DeleteAgent(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent
func (mr *MockAgentManagerMockRecorder) DeleteAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockAgentManager)(nil).DeleteAgent), arg0)
}

// GetActiveAgents mocks base method
func (m *MockAgentManager) GetActiveAgents() ([]*pl_vizier_services_shared_agent.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveAgents")
	ret0, _ := ret[0].([]*pl_vizier_services_shared_agent.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveAgents indicates an expected call of GetActiveAgents
func (mr *MockAgentManagerMockRecorder) GetActiveAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveAgents", reflect.TypeOf((*MockAgentManager)(nil).GetActiveAgents))
}

// AddToFrontOfAgentQueue mocks base method
func (m *MockAgentManager) AddToFrontOfAgentQueue(arg0 string, arg1 *pl_shared_k8s_metadatapb.ResourceUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFrontOfAgentQueue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToFrontOfAgentQueue indicates an expected call of AddToFrontOfAgentQueue
func (mr *MockAgentManagerMockRecorder) AddToFrontOfAgentQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFrontOfAgentQueue", reflect.TypeOf((*MockAgentManager)(nil).AddToFrontOfAgentQueue), arg0, arg1)
}

// GetFromAgentQueue mocks base method
func (m *MockAgentManager) GetFromAgentQueue(arg0 string) ([]*pl_shared_k8s_metadatapb.ResourceUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromAgentQueue", arg0)
	ret0, _ := ret[0].([]*pl_shared_k8s_metadatapb.ResourceUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromAgentQueue indicates an expected call of GetFromAgentQueue
func (mr *MockAgentManagerMockRecorder) GetFromAgentQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromAgentQueue", reflect.TypeOf((*MockAgentManager)(nil).GetFromAgentQueue), arg0)
}

// GetMetadataUpdates mocks base method
func (m *MockAgentManager) GetMetadataUpdates(hostname *controllers.HostnameIPPair) ([]*pl_shared_k8s_metadatapb.ResourceUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataUpdates", hostname)
	ret0, _ := ret[0].([]*pl_shared_k8s_metadatapb.ResourceUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataUpdates indicates an expected call of GetMetadataUpdates
func (mr *MockAgentManagerMockRecorder) GetMetadataUpdates(hostname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataUpdates", reflect.TypeOf((*MockAgentManager)(nil).GetMetadataUpdates), hostname)
}

// AddUpdatesToAgentQueue mocks base method
func (m *MockAgentManager) AddUpdatesToAgentQueue(arg0 string, arg1 []*pl_shared_k8s_metadatapb.ResourceUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUpdatesToAgentQueue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUpdatesToAgentQueue indicates an expected call of AddUpdatesToAgentQueue
func (mr *MockAgentManagerMockRecorder) AddUpdatesToAgentQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpdatesToAgentQueue", reflect.TypeOf((*MockAgentManager)(nil).AddUpdatesToAgentQueue), arg0, arg1)
}

// ApplyAgentUpdate mocks base method
func (m *MockAgentManager) ApplyAgentUpdate(update *controllers.AgentUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyAgentUpdate", update)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyAgentUpdate indicates an expected call of ApplyAgentUpdate
func (mr *MockAgentManagerMockRecorder) ApplyAgentUpdate(update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyAgentUpdate", reflect.TypeOf((*MockAgentManager)(nil).ApplyAgentUpdate), update)
}

// HandleUpdate mocks base method
func (m *MockAgentManager) HandleUpdate(arg0 *controllers.UpdateMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleUpdate", arg0)
}

// HandleUpdate indicates an expected call of HandleUpdate
func (mr *MockAgentManagerMockRecorder) HandleUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockAgentManager)(nil).HandleUpdate), arg0)
}

// NewAgentUpdateCursor mocks base method
func (m *MockAgentManager) NewAgentUpdateCursor() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAgentUpdateCursor")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// NewAgentUpdateCursor indicates an expected call of NewAgentUpdateCursor
func (mr *MockAgentManagerMockRecorder) NewAgentUpdateCursor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAgentUpdateCursor", reflect.TypeOf((*MockAgentManager)(nil).NewAgentUpdateCursor))
}

// DeleteAgentUpdateCursor mocks base method
func (m *MockAgentManager) DeleteAgentUpdateCursor(cursorID uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteAgentUpdateCursor", cursorID)
}

// DeleteAgentUpdateCursor indicates an expected call of DeleteAgentUpdateCursor
func (mr *MockAgentManagerMockRecorder) DeleteAgentUpdateCursor(cursorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentUpdateCursor", reflect.TypeOf((*MockAgentManager)(nil).DeleteAgentUpdateCursor), cursorID)
}

// GetAgentUpdates mocks base method
func (m *MockAgentManager) GetAgentUpdates(cursorID uuid.UUID) ([]*metadatapb.AgentUpdate, *storepb.ComputedSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentUpdates", cursorID)
	ret0, _ := ret[0].([]*metadatapb.AgentUpdate)
	ret1, _ := ret[1].(*storepb.ComputedSchema)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAgentUpdates indicates an expected call of GetAgentUpdates
func (mr *MockAgentManagerMockRecorder) GetAgentUpdates(cursorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentUpdates", reflect.TypeOf((*MockAgentManager)(nil).GetAgentUpdates), cursorID)
}

// UpdateConfig mocks base method
func (m *MockAgentManager) UpdateConfig(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockAgentManagerMockRecorder) UpdateConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockAgentManager)(nil).UpdateConfig), arg0, arg1, arg2, arg3)
}
