package pow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volvofixthis/pow-server/internal/core/models"
	"github.com/volvofixthis/pow-server/internal/core/utils"
)

// Helper function to create a mock PowTask
func mockPowTask(hash []byte) *models.PowTask {
	return &models.PowTask{
		Hash: hash,
		// You can add more fields for a PowTask if necessary for testing.
	}
}

func TestPowRepositoryCreateAndGet(t *testing.T) {
	repo := NewPowRepository()
	ctx := context.TODO()

	// Mock task data
	hash, _ := utils.GenerateSalt(utils.OutputLength)
	task := mockPowTask(hash)

	// Create task in the repository
	err := repo.Create(ctx, task)
	assert.NoError(t, err, "expected no error when creating a task")

	// Get task from the repository
	retrievedTask, err := repo.Get(ctx, hash)
	assert.NoError(t, err, "expected no error when retrieving a task")
	assert.Equal(t, task, retrievedTask, "expected the retrieved task to be the same as the created task")

	// Ensure the task is deleted after retrieval
	_, err = repo.Get(ctx, hash)
	assert.ErrorIs(t, err, ErrNotFound, "expected task not to be found after retrieval")
}

func TestPowRepositoryGetNonExistentTask(t *testing.T) {
	repo := NewPowRepository()
	ctx := context.TODO()

	// Attempt to get a task that doesn't exist
	nonExistentHash, _ := utils.GenerateSalt(utils.OutputLength)
	_, err := repo.Get(ctx, nonExistentHash)
	assert.ErrorIs(t, err, ErrNotFound, "expected error for non-existent task")
}

func TestPowRepositoryDelete(t *testing.T) {
	repo := NewPowRepository()
	ctx := context.TODO()

	// Mock task data
	hash, _ := utils.GenerateSalt(utils.OutputLength)
	task := mockPowTask(hash)

	// Create task in the repository
	err := repo.Create(ctx, task)
	assert.NoError(t, err, "expected no error when creating a task")

	// Delete task from the repository
	err = repo.Delete(ctx, task)
	assert.NoError(t, err, "expected no error when deleting a task")

	// Ensure task cannot be retrieved after deletion
	_, err = repo.Get(ctx, hash)
	assert.ErrorIs(t, err, ErrNotFound, "expected task not to be found after deletion")
}

func TestPowRepositoryItems(t *testing.T) {
	repo := NewPowRepository()
	ctx := context.TODO()

	// Mock multiple tasks
	hash1, _ := utils.GenerateSalt(utils.OutputLength)
	task1 := mockPowTask(hash1)
	hash2, _ := utils.GenerateSalt(utils.OutputLength)
	task2 := mockPowTask(hash2)

	// Create tasks in the repository
	err := repo.Create(ctx, task1)
	assert.NoError(t, err, "expected no error when creating a task")
	err = repo.Create(ctx, task2)
	assert.NoError(t, err, "expected no error when creating a task")

	// Retrieve all items from the repository
	var tasks []*models.PowTask
	repo.Items(ctx)(func(task *models.PowTask) bool {
		tasks = append(tasks, task)
		return true
	})

	// Ensure that all tasks were retrieved
	assert.Len(t, tasks, 2, "expected two tasks in the repository")
	assert.Contains(t, tasks, task1, "expected task1 to be present")
	assert.Contains(t, tasks, task2, "expected task2 to be present")
}
