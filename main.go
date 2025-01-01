package main

import (
	"net/http" // Package for HTTP constants and utilities

	"github.com/gin-gonic/gin" // Gin framework for building REST APIs
)

// User struct defines the structure of the user object
type User struct {
	ID    int    `json:"id"`    // User ID
	Name  string `json:"name"`  // User name
	Email string `json:"email"` // User email
	Phone string `json:"phone"` // User phone number
}

func main() {
	InitDatabase()   // Initialize the database connection
	defer DB.Close() // Ensure the database connection is closed when the program exits

	router := gin.Default() // Create a new Gin router instance

	// GET /users - Return all users
	router.GET("/users", func(c *gin.Context) {
		var users []User                                                // Slice to store all users
		rows, _ := DB.Query("SELECT id, name, email, phone FROM users") // Execute SQL query to fetch all users
		defer rows.Close()                                              // Ensure the rows object is closed after use

		// Iterate through the rows and map data to the User struct
		for rows.Next() {
			var user User
			rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone) // Map row data to the User struct
			users = append(users, user)                               // Append user to the users slice
		}

		c.JSON(http.StatusOK, users) // Return the list of users as a JSON response
	})

	// GET /users/:id - Return a user by ID
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")                                                             // Get the user ID from the URL parameter
		var user User                                                                   // Create a User object to store the result
		row := DB.QueryRow("SELECT id, name, email, phone FROM users WHERE id = ?", id) // Execute SQL query
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)                 // Map row data to the User struct
		if err != nil {                                                                 // Handle error if no user is found
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user) // Return the user data as a JSON response
	})

	// POST /users - Create a new user
	router.POST("/users", func(c *gin.Context) {
		var user User     // Create a User object to store the request data
		c.BindJSON(&user) // Bind the JSON request body to the User object

		// Execute SQL query to insert the new user into the database
		_, err := DB.Exec("INSERT INTO users (name, email, phone) VALUES (?, ?, ?)", user.Name, user.Email, user.Phone)
		if err != nil { // Handle error if the insertion fails
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created"}) // Return success message
	})

	// PUT /users/:id - Update a user
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id") // Get the user ID from the URL parameter
		var user User       // Create a User object to store the request data
		c.BindJSON(&user)   // Bind the JSON request body to the User object

		// Execute SQL query to update the user in the database
		_, err := DB.Exec("UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?", user.Name, user.Email, user.Phone, id)
		if err != nil { // Handle error if the update fails
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated"}) // Return success message
	})

	// DELETE /users/:id - Delete a user
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id") // Get the user ID from the URL parameter

		// Execute SQL query to delete the user from the database
		_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil { // Handle error if the deletion fails
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) // Return success message
	})

	router.Run(":8080") // Start the Gin server on port 8080
}
