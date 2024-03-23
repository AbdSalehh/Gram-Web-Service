package controllers

import (
	"MyGram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentGetAll(c *gin.Context) {
    comments, err := services.GetAllComments()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": "Oops! Something went wrong.",
        })
        return
    }

    c.JSON(http.StatusOK, comments)
}

func CommentCreate(c *gin.Context) {
    userID := getUserIDFromContext(c)
    comment, err := services.CreateComment(c, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func CommentUpdate(c *gin.Context) {
    userID := getUserIDFromContext(c)
    comment, err := services.UpdateComment(c, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func CommentDelete(c *gin.Context) {
    userID := getUserIDFromContext(c)
    if err := services.DeleteComment(c, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": "Oops! Something went wrong.",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Your comment has been successfully deleted",
    })
}
