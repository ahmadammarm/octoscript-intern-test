package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/database"
	"test/model"
)



// @Summary Get all posts
// @Description Get all social media post tasks
// @Tags posts
// @Produce json
// @Success 200 {array} model.Post
// @Router /posts [get]
func GetPosts(c *gin.Context) {     // Handler untuk mengambil semua data post dari database
	var posts []model.Post
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}


// @Summary Create a new post
// @Description Create a new social media post task
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.Post true "Post object"
// @Success 200 {object} model.Post
// @Router /posts [post]
func CreatePost(c *gin.Context) {   // Handler untuk membuat post baru ke database
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}



// @Summary Get a post
// @Description Get a social media post task by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {  // Handler untuk mengambil post berdasarkan ID
	var post model.Post
	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}



// @Summary Update a post
// @Description Update a social media post task
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body model.Post true "Post object"
// @Success 200 {object} model.Post
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {   // Handler untuk mengupdate post berdasarkan ID
	var post model.Post
	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var input model.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&post).Updates(input)
	c.JSON(http.StatusOK, post)
}



// @Summary Delete a post
// @Description Delete a social media post task
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} string
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {   // Handler untuk menghapus post berdasarkan ID
	var post model.Post
	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
