package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/kumaresan1983/todoserver/pkg/models"
)

func CreateTodo(c *gin.Context) {
	var todoInput models.ToDo

	if err := c.ShouldBindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming you have a middleware that sets the current user in the context
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := currentUser.(models.Users)

	todo := models.ToDo{
		Title:     todoInput.Title,
		Content:   todoInput.Content,
		AuthorID:  user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if result := initializers.DB.Create(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	// Assuming you have a middleware that sets the current user in the context
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := currentUser.(models.Users)

	// Extract query parameters
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	sortBy := c.Query("sort_by")
	sortOrder := c.Query("sort_order")

	// Default values if not provided in the query
	page := 1
	limit := 10

	// Convert strings to integers
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = l
		}
	}

	// Query database with pagination and sorting
	var todos []models.ToDo
	query := initializers.DB.Where("author_id = ?", user.ID)

	// Sorting
	if sortBy != "" {
		order := sortBy
		if sortOrder == "desc" {
			order += " DESC"
		}
		query = query.Order(order)
	}

	// Pagination
	query = query.Offset((page - 1) * limit).Limit(limit)

	result := query.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	// Create a slice for simplified todos
	simplifiedTodos := make([]models.TodoSimplified, len(todos))
	for i, todo := range todos {
		simplifiedTodos[i] = todo.Simplified()
	}

	c.JSON(http.StatusOK, simplifiedTodos)
}

func GetTodoByID(c *gin.Context) {
	todoID := c.Param("id")

	var todo models.ToDo
	result := initializers.DB.First(&todo, "id = ?", todoID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodoByID deletes a todo by ID
func DeleteTodoByID(c *gin.Context) {
	// Assuming you have a middleware that sets the current user in the context
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := currentUser.(models.Users)

	todoID := c.Param("id")

	// Check if the todo exists
	var todo models.ToDo
	result := initializers.DB.First(&todo, "id = ? AND author_id = ?", todoID, user.ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Delete the todo
	if result := initializers.DB.Delete(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Todo deleted successfully"})
}

// CompleteTodoByID completes a todo by updating the Completed and CompletedAt fields
func CompleteTodoByID(c *gin.Context) {
	// Assuming you have a middleware that sets the current user in the context
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := currentUser.(models.Users)

	todoID := c.Param("id")

	// Check if the todo exists
	var todo models.ToDo
	result := initializers.DB.First(&todo, "id = ? AND author_id = ?", todoID, user.ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Bind request JSON to struct to get Completed flag
	var completeRequest struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&completeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the todo with the Completed flag and CompletedAt
	todo.Completed = completeRequest.Completed
	if todo.Completed {
		todo.CompletedAt = time.Now()
	} else {
		todo.CompletedAt = time.Time{} // Set to zero value if not completed
	}

	// Save the updated todo
	if result := initializers.DB.Save(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Todo completed successfully", "todo": todo.Simplified()})
}
