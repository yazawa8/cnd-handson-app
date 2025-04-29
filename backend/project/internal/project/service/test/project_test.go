package test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/repository/test"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	// Set environment variable to skip DB connection during test execution
	os.Setenv("DB_SKIP_CONNECT", "true")
}

// Helper function to create test project data
func createTestProject() *model.Project {
	id := uuid.New()
	ownerID := uuid.New()
	now := time.Now()
	return &model.Project{
		ID:          id,
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     ownerID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Helper function to create test project with specific owner ID
func createTestProjectWithOwner(ownerID uuid.UUID) *model.Project {
	id := uuid.New()
	now := time.Now()
	return &model.Project{
		ID:          id,
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     ownerID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// TestCreateProject tests the CreateProject function
func TestCreateProject(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)

	// Create service instance
	projectService := service.NewProjectService(mockRepo)

	// Test cases
	testCases := []struct {
		name          string
		projectName   string
		description   string
		ownerID       uuid.UUID
		mockBehavior  func(ownerID uuid.UUID)
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:        "Success: Project creation succeeds",
			projectName: "Test Project",
			description: "This is a test project",
			ownerID:     uuid.New(),
			mockBehavior: func(ownerID uuid.UUID) {
				// Return project with the same owner ID as passed parameter
				mockRepo.On("CreateProject", mock.AnythingOfType("*model.Project")).Return(
					createTestProjectWithOwner(ownerID), nil)
			},
			expectedError: false,
		},
		{
			name:        "Error: Empty project name",
			projectName: "",
			description: "Description",
			ownerID:     uuid.New(),
			mockBehavior: func(ownerID uuid.UUID) {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:        "Error: Empty owner ID",
			projectName: "Project Name",
			description: "Description",
			ownerID:     uuid.Nil,
			mockBehavior: func(ownerID uuid.UUID) {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:        "Error: Repository error",
			projectName: "Test Project",
			description: "This is a test project",
			ownerID:     uuid.New(),
			mockBehavior: func(ownerID uuid.UUID) {
				mockRepo.On("CreateProject", mock.AnythingOfType("*model.Project")).Return(nil, errors.New("DB error"))
			},
			expectedError: true,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior(tc.ownerID)

			// Execute function under test
			result, err := projectService.CreateProject(tc.projectName, tc.description, tc.ownerID)
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				if tc.errorCode != 0 {
					statusErr, ok := status.FromError(err)
					assert.True(t, ok, "Should be a gRPC status error")
					assert.Equal(t, tc.errorCode, statusErr.Code(), "Error code should match")
				}
				assert.Nil(t, result, "Result should be nil")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, result, "Result should not be nil")
				assert.Equal(t, tc.projectName, result.Name, "Project name should match")
				assert.Equal(t, tc.description, result.Description, "Description should match")
				assert.Equal(t, tc.ownerID, result.OwnerID, "Owner ID should match")
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetProject tests the GetProject function
func TestGetProject(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)

	// Create service instance
	projectService := service.NewProjectService(mockRepo)

	// Test data
	testProject := createTestProject()
	validID := testProject.ID
	invalidID := uuid.New()

	// Test cases
	testCases := []struct {
		name          string
		projectID     uuid.UUID
		mockBehavior  func()
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:      "Success: Project retrieval succeeds",
			projectID: validID,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
			},
			expectedError: false,
		},
		{
			name:      "Error: Empty project ID",
			projectID: uuid.Nil,
			mockBehavior: func() {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:      "Error: Non-existent project",
			projectID: invalidID,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", invalidID).Return(nil, errors.New("not found"))
			},
			expectedError: true,
			errorCode:     codes.NotFound,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior()

			// Execute function under test
			result, err := projectService.GetProject(tc.projectID)
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				if tc.errorCode != 0 {
					statusErr, ok := status.FromError(err)
					assert.True(t, ok, "Should be a gRPC status error")
					assert.Equal(t, tc.errorCode, statusErr.Code(), "Error code should match")
				}
				assert.Nil(t, result, "Result should be nil")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, result, "Result should not be nil")
				assert.Equal(t, testProject.ID, result.ID, "Project ID should match")
				assert.Equal(t, testProject.Name, result.Name, "Project name should match")
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestListProjects tests the ListProjects function
func TestListProjects(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)
	
	// Create service instance
	projectService := service.NewProjectService(mockRepo)
	
	// Test data
	testProjects := []*model.Project{createTestProject(), createTestProject()}
	
	// Test cases
	testCases := []struct {
		name          string
		mockBehavior  func()
		expectedError bool
		expectedCount int
	}{
		{
			name: "Success: Project list retrieval succeeds",
			mockBehavior: func() {
				mockRepo.On("ListProjects").Return(testProjects, nil)
			},
			expectedError: false,
			expectedCount: 2,
		},
		{
			name: "Error: Repository error",
			mockBehavior: func() {
				mockRepo.On("ListProjects").Return(nil, errors.New("DB error"))
			},
			expectedError: true,
			expectedCount: 0,
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior()
			
			// Execute function under test
			results, err := projectService.ListProjects()
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				assert.Nil(t, results, "Results should be nil")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, results, "Results should not be nil")
				assert.Len(t, results, tc.expectedCount, "Result count should match")
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestListProjectsByOwner tests the ListProjectsByOwner function
func TestListProjectsByOwner(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)
	
	// Create service instance
	projectService := service.NewProjectService(mockRepo)
	
	// Test data
	testProjects := []*model.Project{createTestProject(), createTestProject()}
	ownerID := uuid.New()
	
	// Test cases
	testCases := []struct {
		name          string
		ownerID       uuid.UUID
		mockBehavior  func()
		expectedError bool
		errorCode     codes.Code
		expectedCount int
	}{
		{
			name:    "Success: Owner's project list retrieval succeeds",
			ownerID: ownerID,
			mockBehavior: func() {
				mockRepo.On("GetProjectsByOwnerID", ownerID).Return(testProjects, nil)
			},
			expectedError: false,
			expectedCount: 2,
		},
		{
			name:    "Error: Empty owner ID",
			ownerID: uuid.Nil,
			mockBehavior: func() {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
			expectedCount: 0,
		},
		{
			name:    "Error: Repository error",
			ownerID: ownerID,
			mockBehavior: func() {
				mockRepo.On("GetProjectsByOwnerID", ownerID).Return(nil, errors.New("DB error"))
			},
			expectedError: true,
			expectedCount: 0,
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior()
			
			// Execute function under test
			results, err := projectService.ListProjectsByOwner(tc.ownerID)
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				if tc.errorCode != 0 {
					statusErr, ok := status.FromError(err)
					assert.True(t, ok, "Should be a gRPC status error")
					assert.Equal(t, tc.errorCode, statusErr.Code(), "Error code should match")
				}
				assert.Nil(t, results, "Results should be nil")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, results, "Results should not be nil")
				assert.Len(t, results, tc.expectedCount, "Result count should match")
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestUpdateProject tests the UpdateProject function
func TestUpdateProject(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)
	
	// Create service instance
	projectService := service.NewProjectService(mockRepo)
	
	// Test data
	testProject := createTestProject()
	validID := testProject.ID
	newName := "Updated Project Name"
	newDesc := "Updated description"
	
	// Test cases
	testCases := []struct {
		name          string
		projectID     uuid.UUID
		updateName    *string
		updateDesc    *string
		mockBehavior  func()
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:       "Success: Name and description update succeeds",
			projectID:  validID,
			updateName: &newName,
			updateDesc: &newDesc,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
				mockRepo.On("UpdateProject", mock.AnythingOfType("*model.Project")).Return(
					&model.Project{
						ID:          validID,
						Name:        newName,
						Description: newDesc,
						OwnerID:     testProject.OwnerID,
						CreatedAt:   testProject.CreatedAt,
						UpdatedAt:   time.Now(),
					}, nil)
			},
			expectedError: false,
		},
		{
			name:       "Success: Only name update succeeds",
			projectID:  validID,
			updateName: &newName,
			updateDesc: nil,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
				mockRepo.On("UpdateProject", mock.AnythingOfType("*model.Project")).Return(
					&model.Project{
						ID:          validID,
						Name:        newName,
						Description: testProject.Description,
						OwnerID:     testProject.OwnerID,
						CreatedAt:   testProject.CreatedAt,
						UpdatedAt:   time.Now(),
					}, nil)
			},
			expectedError: false,
		},
		{
			name:       "Error: Empty project ID",
			projectID:  uuid.Nil,
			updateName: &newName,
			updateDesc: &newDesc,
			mockBehavior: func() {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:       "Error: Update to empty name",
			projectID:  validID,
			updateName: func() *string { s := ""; return &s }(),
			updateDesc: &newDesc,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:       "Error: Non-existent project",
			projectID:  uuid.New(),
			updateName: &newName,
			updateDesc: &newDesc,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("not found"))
			},
			expectedError: true,
			errorCode:     codes.NotFound,
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior()
			
			// Execute function under test
			result, err := projectService.UpdateProject(tc.projectID, tc.updateName, tc.updateDesc)
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				if tc.errorCode != 0 {
					statusErr, ok := status.FromError(err)
					assert.True(t, ok, "Should be a gRPC status error")
					assert.Equal(t, tc.errorCode, statusErr.Code(), "Error code should match")
				}
				assert.Nil(t, result, "Result should be nil")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, result, "Result should not be nil")
				if tc.updateName != nil {
					assert.Equal(t, *tc.updateName, result.Name, "Updated name should match")
				}
				if tc.updateDesc != nil {
					assert.Equal(t, *tc.updateDesc, result.Description, "Updated description should match")
				}
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestDeleteProject tests the DeleteProject function
func TestDeleteProject(t *testing.T) {
	// Create mock repository
	mockRepo := new(test.MockProjectRepository)
	
	// Create service instance
	projectService := service.NewProjectService(mockRepo)
	
	// Test data
	testProject := createTestProject()
	validID := testProject.ID
	
	// Test cases
	testCases := []struct {
		name          string
		projectID     uuid.UUID
		mockBehavior  func()
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:      "Success: Project deletion succeeds",
			projectID: validID,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
				mockRepo.On("DeleteProject", validID).Return(nil)
			},
			expectedError: false,
		},
		{
			name:      "Error: Empty project ID",
			projectID: uuid.Nil,
			mockBehavior: func() {
				// Mock should not be called
			},
			expectedError: true,
			errorCode:     codes.InvalidArgument,
		},
		{
			name:      "Error: Non-existent project",
			projectID: uuid.New(),
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("not found"))
			},
			expectedError: true,
			errorCode:     codes.NotFound,
		},
		{
			name:      "Error: Deletion error",
			projectID: validID,
			mockBehavior: func() {
				mockRepo.On("GetProjectByID", validID).Return(testProject, nil)
				mockRepo.On("DeleteProject", validID).Return(errors.New("delete error"))
			},
			expectedError: true,
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil
			tc.mockBehavior()
			
			// Execute function under test
			err := projectService.DeleteProject(tc.projectID)
			
			// Verify expected results
			if tc.expectedError {
				assert.Error(t, err, "Should return an error")
				if tc.errorCode != 0 {
					statusErr, ok := status.FromError(err)
					assert.True(t, ok, "Should be a gRPC status error")
					assert.Equal(t, tc.errorCode, statusErr.Code(), "Error code should match")
				}
			} else {
				assert.NoError(t, err, "Should not return an error")
			}
			
			// Verify mock calls
			mockRepo.AssertExpectations(t)
		})
	}
}