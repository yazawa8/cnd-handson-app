package test

import (
	"os"
	"testing"
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/repository"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/pkg/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var testRepo repository.ProjectRepository

func init() {
	// Set environment variable to skip DB connection during test execution
	os.Setenv("DB_SKIP_CONNECT", "true")
}

// TestMain handles setup and teardown for all tests
func TestMain(m *testing.M) {
	// Setup in-memory SQLite database
	var err error
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize test database: " + err.Error())
	}

	// Migrate tables
	err = testDB.AutoMigrate(&model.Project{})
	if err != nil {
		panic("Failed to migrate tables: " + err.Error())
	}

	// Create test repository
	// Explicitly set DB
	db.SetDB(testDB)
	testRepo = repository.NewProjectRepository()

	// Run tests
	code := m.Run()

	// Cleanup after tests
	// Clear environment variables (if needed)
	// os.Unsetenv("DB_SKIP_CONNECT")
	
	// Exit
	os.Exit(code)
}

// createTestProject creates test project data
func createTestProject() *model.Project {
	return &model.Project{
		ID:          uuid.New(),
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// TestCreateProject tests CreateProject function
func TestCreateProject(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Test project
	project := createTestProject()

	// Execute test
	result, err := testRepo.CreateProject(project)

	// Verification
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, project.ID, result.ID, "ID should match")
	assert.Equal(t, project.Name, result.Name, "Name should match")
	assert.Equal(t, project.Description, result.Description, "Description should match")

	// Verify data is saved in database
	var savedProject model.Project
	err = testDB.Where("id = ?", project.ID).First(&savedProject).Error
	assert.NoError(t, err, "Should be able to retrieve the saved project")
	assert.Equal(t, project.Name, savedProject.Name, "Saved name should match")
}

// TestGetProjectByID tests GetProjectByID function
func TestGetProjectByID(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Test project
	project := createTestProject()
	testDB.Create(project)

	// Test with existing ID
	result, err := testRepo.GetProjectByID(project.ID)
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, project.ID, result.ID, "ID should match")

	// Test with non-existent ID
	nonExistentID := uuid.New()
	result, err = testRepo.GetProjectByID(nonExistentID)
	assert.Error(t, err, "Should return an error")
	assert.Nil(t, result, "Result should be nil for non-existent ID")
}

// TestGetProjectsByOwnerID tests GetProjectsByOwnerID function
func TestGetProjectsByOwnerID(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Test owner ID
	ownerID := uuid.New()

	// Create two projects with the same owner ID
	project1 := createTestProject()
	project1.OwnerID = ownerID
	testDB.Create(project1)

	project2 := createTestProject()
	project2.OwnerID = ownerID
	testDB.Create(project2)

	// Create one project with a different owner ID
	project3 := createTestProject()
	testDB.Create(project3)

	// Test retrieving projects by owner ID
	results, err := testRepo.GetProjectsByOwnerID(ownerID)
	assert.NoError(t, err, "Should not return an error")
	assert.Len(t, results, 2, "Should return 2 results")

	// Test with non-existent owner ID
	nonExistentOwnerID := uuid.New()
	results, err = testRepo.GetProjectsByOwnerID(nonExistentOwnerID)
	assert.NoError(t, err, "Should not return an error (empty results)")
	assert.Empty(t, results, "Results should be empty")
}

// TestListProjects tests ListProjects function
func TestListProjects(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Create 3 test projects
	testDB.Create(createTestProject())
	testDB.Create(createTestProject())
	testDB.Create(createTestProject())

	// Test retrieving project list
	results, err := testRepo.ListProjects()
	assert.NoError(t, err, "Should not return an error")
	assert.Len(t, results, 3, "Should return 3 results")
}

// TestUpdateProject tests UpdateProject function
func TestUpdateProject(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Test project
	project := createTestProject()
	testDB.Create(project)

	// Update project
	project.Name = "Updated Project Name"
	project.Description = "Updated description"

	// Execute update
	result, err := testRepo.UpdateProject(project)
	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, "Updated Project Name", result.Name, "Name should be updated")
	assert.Equal(t, "Updated description", result.Description, "Description should be updated")

	// Verify in database
	var updatedProject model.Project
	err = testDB.Where("id = ?", project.ID).First(&updatedProject).Error
	assert.NoError(t, err, "Should be able to retrieve the updated project")
	assert.Equal(t, "Updated Project Name", updatedProject.Name, "Name in database should be updated")
	assert.Equal(t, "Updated description", updatedProject.Description, "Description in database should be updated")
}

// TestDeleteProject tests DeleteProject function
func TestDeleteProject(t *testing.T) {
	// Clean data before test
	testDB.Where("1 = 1").Delete(&model.Project{})

	// Test project
	project := createTestProject()
	testDB.Create(project)

	// Execute delete
	err := testRepo.DeleteProject(project.ID)
	assert.NoError(t, err, "Should not return an error")

	// Verify deletion
	var deletedProject model.Project
	err = testDB.Where("id = ?", project.ID).First(&deletedProject).Error
	assert.Error(t, err, "Should not be able to retrieve the deleted project")

	// Test deleting non-existent ID
	nonExistentID := uuid.New()
	err = testRepo.DeleteProject(nonExistentID)
	assert.NoError(t, err, "Deleting non-existent ID should not return an error")
}