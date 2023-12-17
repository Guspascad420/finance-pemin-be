package controllers

import (
	"finance-be/database"
	"finance-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBalances(c *gin.Context) {
	var balances []models.Balance

	record := database.Db.Find(&balances)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": balances})
}

func UpdateBalance(c *gin.Context) {
	var balance models.Balance
	if jsonErr := c.ShouldBindJSON(&balance); jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonErr.Error()})
		c.Abort()
		return
	}

	update := database.Db.Save(&balance)
	if update.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": update.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "balance successfully updated"})
}

func DeleteBalance(c *gin.Context) {
	id := c.Param("id")
	deletion := database.Db.Delete(&models.Balance{}, id)

	if deletion.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": deletion.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "balance successfully deleted"})
}

func CreateNewBalance(c *gin.Context) {
	var balance models.Balance
	if jsonErr := c.ShouldBindJSON(&balance); jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonErr.Error()})
		c.Abort()
		return
	}

	record := database.Db.Create(&balance)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "successfully created new balance"})
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	record := database.Db.Find(&transactions)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": transactions})
}

func CreateNewTransaction(c *gin.Context) {
	var transaction models.Transaction
	if jsonErr := c.ShouldBindJSON(&transaction); jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": jsonErr.Error()})
		c.Abort()
		return
	}

	record := database.Db.Create(&transaction)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "successfully created new transaction"})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	deletion := database.Db.Delete(&models.Transaction{}, id)

	if deletion.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": deletion.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "transaction successfully deleted"})
}
