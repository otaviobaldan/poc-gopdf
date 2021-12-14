package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	go_pdf_generator "github.com/otaviobaldan/go-pdf-generator"
	"github.com/otaviobaldan/go-pdf-generator/config"
	"github.com/otaviobaldan/go-pdf-generator/constants"
	"io/ioutil"
	"time"
)

func main() {
	pdfConfig := config.NewPdfConfig(
		constants.OrientationPortrait,
		constants.UnitsMillimeters,
		constants.PaperSizeA4,
		20,
		20,
		40,
		true,
	)
	// margins in points
	//pdfConfig := config.NewPdfConfig(
	//	constants.OrientationPortrait,
	//	constants.UnitsPoints,
	//	constants.PaperSizeA4,
	//	85,
	//	56.7,
	//	85,
	//	true,
	//)
	footerCfg := &config.TextConfig{
		FontFamily: constants.FontBookmanOldStyle,
		Align:      constants.AlignCenter,
		Style:      "I",
		Size:       8,
	}
	titleCfg := &config.TextConfig{
		FontFamily: constants.FontBookmanOldStyle,
		Align:      constants.AlignCenter,
		Style:      "B",
		Size:       14,
	}
	subTitleCfg := &config.TextConfig{
		FontFamily: constants.FontBookmanOldStyle,
		Align:      constants.AlignLeft,
		Style:      "B",
		Size:       12,
	}
	textCfg := &config.TextConfig{
		FontFamily: constants.FontBookmanOldStyle,
		Align:      constants.AlignJustify,
		Style:      "",
		Size:       12,
	}
	pdf, err := go_pdf_generator.NewPdfGenerator(
		pdfConfig,
		nil,
		footerCfg,
		titleCfg,
		subTitleCfg,
		textCfg,
	)
	if err != nil {
		panic(err)
	}

	pdf.GenerateDefaultFooter("gerado pela plataforma mais contratos", true)
	pdf.GenerateTitle("LERO LERO DE TI")

	pdf.GenerateSubtitle("CHAPTER 1")
	// same line to test replace line break
	bytes, _ := ioutil.ReadFile("./text/parte1.txt")
	convertedString := string(bytes)

	pdf.GenerateText(convertedString)

	pdf.GenerateSubtitle("CHAPTER 2")
	bytes, _ = ioutil.ReadFile("./text/parte2.txt")
	pdf.GenerateText(string(bytes))

	pdf.GenerateSignature("Otavio Baldan")

	err = pdf.Pdf.OutputFileAndClose(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}
}

func main__() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	//pdf.MoveTo(20, 20)
	//pdf.LineTo(170, 20)
	////pdf.ArcTo(170, 40, 20, 20, 0, 90, 0)
	////pdf.CurveTo(190, 100, 105, 100)
	////pdf.CurveBezierCubicTo(20, 100, 105, 200, 20, 200)
	//pdf.ClosePath()
	////pdf.SetFillColor(200, 200, 200)
	//pdf.SetLineWidth(0.5)
	//pdf.DrawPath("DF")

	left, top, right, _ := pdf.GetMargins()
	width, _ := pdf.GetPageSize()

	bytes, _ := ioutil.ReadFile("./text/parte1.txt")
	pdf.MultiCell(0, constants.SizeTextHeight, string(bytes), "", "J", false)
	bytes, _ = ioutil.ReadFile("./text/parte2.txt")
	pdf.MultiCell(0, constants.SizeTextHeight, string(bytes), "", "J", false)
	currentY := pdf.GetY()

	fmt.Println("currentY ", currentY)
	fmt.Println("Width ", width)
	fmt.Println("left ", left)
	fmt.Println("top ", top)
	fmt.Println("right ", right)
	//space := 4.0

	//lineCenterY := 10 / 1.33
	//cell.Y += lineCenterY
	lineSize := float64(130)
	availableSpace := (width - left - right - lineSize) / 2
	signature := "assinatura aaaaa"

	pdf.Line(left+availableSpace, currentY+20, left+availableSpace+lineSize, 20+currentY)
	pdf.CellFormat(0, 50, signature, "", 1, "C", false, 0, "")

	err := pdf.OutputFileAndClose(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}
}
