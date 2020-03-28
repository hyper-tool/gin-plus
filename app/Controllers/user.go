package Controllers

import (
	"fmt"
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
	for i := 1; i < 30000; i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i), 100)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i), 100)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(c.Writer)
}
