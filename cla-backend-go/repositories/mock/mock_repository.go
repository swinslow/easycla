// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT
//

// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "github.com/communitybridge/easycla/cla-backend-go/gen/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddGithubRepository mocks base method
func (m *MockRepository) AddGithubRepository(ctx context.Context, externalProjectID, projectSFID string, input *models.GithubRepositoryInput) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGithubRepository", ctx, externalProjectID, projectSFID, input)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGithubRepository indicates an expected call of AddGithubRepository
func (mr *MockRepositoryMockRecorder) AddGithubRepository(ctx, externalProjectID, projectSFID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGithubRepository", reflect.TypeOf((*MockRepository)(nil).AddGithubRepository), ctx, externalProjectID, projectSFID, input)
}

// UpdateClaGroupID mocks base method
func (m *MockRepository) UpdateClaGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClaGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClaGroupID indicates an expected call of UpdateClaGroupID
func (mr *MockRepositoryMockRecorder) UpdateClaGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClaGroupID", reflect.TypeOf((*MockRepository)(nil).UpdateClaGroupID), ctx, repositoryID, claGroupID)
}

// EnableRepository mocks base method
func (m *MockRepository) EnableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepository indicates an expected call of EnableRepository
func (mr *MockRepositoryMockRecorder) EnableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableRepository", reflect.TypeOf((*MockRepository)(nil).EnableRepository), ctx, repositoryID)
}

// DisableRepository mocks base method
func (m *MockRepository) DisableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepository indicates an expected call of DisableRepository
func (mr *MockRepositoryMockRecorder) DisableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableRepository", reflect.TypeOf((*MockRepository)(nil).DisableRepository), ctx, repositoryID)
}

// DisableRepositoriesByProjectID mocks base method
func (m *MockRepository) DisableRepositoriesByProjectID(ctx context.Context, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableRepositoriesByProjectID", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepositoriesByProjectID indicates an expected call of DisableRepositoriesByProjectID
func (mr *MockRepositoryMockRecorder) DisableRepositoriesByProjectID(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableRepositoriesByProjectID", reflect.TypeOf((*MockRepository)(nil).DisableRepositoriesByProjectID), ctx, projectID)
}

// DisableRepositoriesOfGithubOrganization mocks base method
func (m *MockRepository) DisableRepositoriesOfGithubOrganization(ctx context.Context, externalProjectID, githubOrgName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableRepositoriesOfGithubOrganization", ctx, externalProjectID, githubOrgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepositoriesOfGithubOrganization indicates an expected call of DisableRepositoriesOfGithubOrganization
func (mr *MockRepositoryMockRecorder) DisableRepositoriesOfGithubOrganization(ctx, externalProjectID, githubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableRepositoriesOfGithubOrganization", reflect.TypeOf((*MockRepository)(nil).DisableRepositoriesOfGithubOrganization), ctx, externalProjectID, githubOrgName)
}

// GetRepository mocks base method
func (m *MockRepository) GetRepository(ctx context.Context, repositoryID string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository", ctx, repositoryID)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository
func (mr *MockRepositoryMockRecorder) GetRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockRepository)(nil).GetRepository), ctx, repositoryID)
}

// GetRepositoryByGithubID mocks base method
func (m *MockRepository) GetRepositoryByGithubID(ctx context.Context, externalID string, enabled bool) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByGithubID", ctx, externalID, enabled)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByGithubID indicates an expected call of GetRepositoryByGithubID
func (mr *MockRepositoryMockRecorder) GetRepositoryByGithubID(ctx, externalID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByGithubID", reflect.TypeOf((*MockRepository)(nil).GetRepositoryByGithubID), ctx, externalID, enabled)
}

// GetRepositoriesByCLAGroup mocks base method
func (m *MockRepository) GetRepositoriesByCLAGroup(ctx context.Context, claGroup string, enabled bool) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoriesByCLAGroup", ctx, claGroup, enabled)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByCLAGroup indicates an expected call of GetRepositoriesByCLAGroup
func (mr *MockRepositoryMockRecorder) GetRepositoriesByCLAGroup(ctx, claGroup, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoriesByCLAGroup", reflect.TypeOf((*MockRepository)(nil).GetRepositoriesByCLAGroup), ctx, claGroup, enabled)
}

// GetRepositoriesByOrganizationName mocks base method
func (m *MockRepository) GetRepositoriesByOrganizationName(ctx context.Context, gitHubOrgName string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoriesByOrganizationName", ctx, gitHubOrgName)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByOrganizationName indicates an expected call of GetRepositoriesByOrganizationName
func (mr *MockRepositoryMockRecorder) GetRepositoriesByOrganizationName(ctx, gitHubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoriesByOrganizationName", reflect.TypeOf((*MockRepository)(nil).GetRepositoriesByOrganizationName), ctx, gitHubOrgName)
}

// GetCLAGroupRepositoriesGroupByOrgs mocks base method
func (m *MockRepository) GetCLAGroupRepositoriesGroupByOrgs(ctx context.Context, projectID string, enabled bool) ([]*models.GithubRepositoriesGroupByOrgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCLAGroupRepositoriesGroupByOrgs", ctx, projectID, enabled)
	ret0, _ := ret[0].([]*models.GithubRepositoriesGroupByOrgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCLAGroupRepositoriesGroupByOrgs indicates an expected call of GetCLAGroupRepositoriesGroupByOrgs
func (mr *MockRepositoryMockRecorder) GetCLAGroupRepositoriesGroupByOrgs(ctx, projectID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCLAGroupRepositoriesGroupByOrgs", reflect.TypeOf((*MockRepository)(nil).GetCLAGroupRepositoriesGroupByOrgs), ctx, projectID, enabled)
}

// ListProjectRepositories mocks base method
func (m *MockRepository) ListProjectRepositories(ctx context.Context, externalProjectID, projectSFID string, enabled bool) (*models.ListGithubRepositories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectRepositories", ctx, externalProjectID, projectSFID, enabled)
	ret0, _ := ret[0].(*models.ListGithubRepositories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectRepositories indicates an expected call of ListProjectRepositories
func (mr *MockRepositoryMockRecorder) ListProjectRepositories(ctx, externalProjectID, projectSFID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectRepositories", reflect.TypeOf((*MockRepository)(nil).ListProjectRepositories), ctx, externalProjectID, projectSFID, enabled)
}
