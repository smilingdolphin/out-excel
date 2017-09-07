package controllers

import (
	"fmt"

	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fpay/gopress"
)

type Summary struct {
	Desc          string  `json:"desc"`
	Name          string  `json:"name"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	ReceiveAmount float32 `json:"receive_amount"`
	ReceiveCount  int     `json:"receive_count"`
	RealAmount    float32 `json:"real_amount"`
	RealCount     int     `json:"real_count"`
	RefundAmount  float32 `json:"refund_amount"`
	RefundCount   int     `json:"refund_count"`
	CouponAmount  float32 `json:"coupon_amount"`
	CouponCount   int     `json:"coupon_count"`
	Commission    float32 `json:"commission"`
	SettleAmount  float32 `json:"settle_amount"`
}

type Data []interface{}

type Sheet struct {
	Sheet     string   `json:"sheet"`
	Summaries Summary  `json:"summaries"`
	Fields    []string `json:"fields"`
	Datas     []Data   `json:"datas"`
}

type JsonData struct {
	Merchant string  `json:"merchant"`
	Excel    []Sheet `json:"excel"`
}

type Row []string

type Table struct {
	Name string `json:"name"`
	Rows []Row  `json:"rows"`
}

type TableData struct {
	Excel []Table `json:"excel"`
}

// ExcelController
type ExcelController struct {
	// Uncomment this line if you want to use services in the app
	// app *gopress.App
}

// NewExcelController returns excel controller instance.
func NewExcelController() *ExcelController {
	return new(ExcelController)
}

// RegisterRoutes registes routes to app
// It is used to implements gopress.Controller.
func (c *ExcelController) RegisterRoutes(app *gopress.App) {
	// Uncomment this line if you want to use services in the app
	// c.app = app

	app.POST("/excel/export", c.ExportExcelExAction)
}

// SampleGetAction Action
// Parameter gopress.Context is just alias of echo.Context
func (c *ExcelController) ExportExcelAction(ctx gopress.Context) error {
	excel := new(JsonData)
	ctx.Bind(excel)

	xlsx := excelize.NewFile()
	startRow := 0
	startCol := 0
	sheetName := "Sheet1"
	for k, sheet := range excel.Excel {
		startRow = 1
		startCol = 65
		if k >= 1 {
			xlsx.NewSheet(k+1, sheet.Sheet)
			sheetName = fmt.Sprintf("Sheet%d", k+1)
		}

		// 如果有summaries数据
		if sheet.Summaries.Desc != "" {
			xlsx.SetCellValue(sheetName, "A1", sheet.Summaries.Desc)
			xlsx.SetCellValue(sheetName, "B1", sheet.Summaries.Name)
			xlsx.SetCellValue(sheetName, "A2", "起始日期：")
			xlsx.SetCellValue(sheetName, "B2", sheet.Summaries.StartDate)
			xlsx.SetCellValue(sheetName, "C2", "起始日期：")
			xlsx.SetCellValue(sheetName, "D2", sheet.Summaries.EndDate)
			xlsx.SetCellValue(sheetName, "A3", "应收总计：")
			xlsx.SetCellValue(sheetName, "B3", fmt.Sprintf("共%.2f元", sheet.Summaries.ReceiveAmount))
			xlsx.SetCellValue(sheetName, "C3", fmt.Sprintf("共%d笔", sheet.Summaries.ReceiveCount))
			xlsx.SetCellValue(sheetName, "A4", "实收总计：")
			xlsx.SetCellValue(sheetName, "B4", fmt.Sprintf("共%.2f元", sheet.Summaries.RealAmount))
			xlsx.SetCellValue(sheetName, "C4", fmt.Sprintf("共%d笔", sheet.Summaries.RealCount))
			xlsx.SetCellValue(sheetName, "A5", "退款总计：")
			xlsx.SetCellValue(sheetName, "B5", fmt.Sprintf("共%.2f元", sheet.Summaries.RefundAmount))
			xlsx.SetCellValue(sheetName, "C5", fmt.Sprintf("共%d笔", sheet.Summaries.RefundCount))
			xlsx.SetCellValue(sheetName, "A6", "商家优惠总计：")
			xlsx.SetCellValue(sheetName, "B6", fmt.Sprintf("共%.2f元", sheet.Summaries.CouponAmount))
			xlsx.SetCellValue(sheetName, "C6", fmt.Sprintf("共%d笔", sheet.Summaries.CouponCount))
			xlsx.SetCellValue(sheetName, "A7", "手续费总计：")
			xlsx.SetCellValue(sheetName, "B7", fmt.Sprintf("共%.2f元", sheet.Summaries.Commission))
			xlsx.SetCellValue(sheetName, "A8", "结算总计：")
			xlsx.SetCellValue(sheetName, "B8", fmt.Sprintf("共%.2f元", sheet.Summaries.SettleAmount))
			xlsx.SetCellValue(sheetName, "C8", "结算总计=应收总计-退款总计-商家优惠总计-手续费总计")
			startRow = 10
		}

		// Fields
		for _, f := range sheet.Fields {
			xlsx.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(startCol), startRow), f)
			startCol++
		}

		// Datas
		for _, d := range sheet.Datas {
			startCol = 65
			startRow++
			for _, dd := range d {
				xlsx.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(startCol), startRow), dd)
				startCol++
			}
		}

		xlsx.SetSheetName(sheetName, sheet.Sheet)
	}

	xlsx.SetActiveSheet(1)

	filename := "./" + excel.Merchant + ".xlsx"
	err := xlsx.SaveAs(filename)

	if err != nil {
		fmt.Println(err)
	}

	return ctx.File(filename)
}

func (c *ExcelController) ExportExcelExAction(ctx gopress.Context) error {
	excel := new(TableData)
	ctx.Bind(excel)

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	startCol := 0
	for k, tbl := range excel.Excel {
		if k >= 1 {
			xlsx.NewSheet(k+1, tbl.Name)
			sheetName = fmt.Sprintf("Sheet%d", k+1)
		}
		for l, r := range tbl.Rows {
			startCol = 65
			if len(r) > 0 {
				for _, v := range r {
					xlsx.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(startCol), l+1), v)
					startCol++
				}
			}
		}

		xlsx.SetSheetName(sheetName, tbl.Name)
	}

	xlsx.SetActiveSheet(1)

	filename := fmt.Sprintf("/tmp/%s.xlsx", time.Now().Format("20060102150504.000"))
	err := xlsx.SaveAs(filename)

	if err != nil {
		fmt.Println(err)
	}

	ctx.File(filename)
	return os.Remove(filename)
}
