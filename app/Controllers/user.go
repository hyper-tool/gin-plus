package Controllers

import (
	"gin-plus/app/Models"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {

	user := c.DefaultQuery("user", "ben")
	age := c.GetInt("age")
	Models.Create(user, age)
	c.JSON(200, gin.H{"User": user})
}

func Excel(c *gin.Context) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(c.Writer)
}
