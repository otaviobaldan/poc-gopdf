package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"path/filepath"
	"time"
)

type Document struct {
	pdf        gofpdf.Pdf
	translator func(string) string
}

func NewDocument() *Document {
	pdf := gofpdf.New("P", "mm", "A4", "")
	return &Document{
		pdf:        pdf,
		translator: pdf.UnicodeTranslatorFromDescriptor(""),
	}
}

func main() {
	doc := NewDocument()

	mainTitle := "Lero Lero de TI"
	doc.pdf.SetMargins(20, 40, 20)
	doc.pdf.SetTitle(mainTitle, false)
	doc.pdf.SetAuthor("obaldan", true)

	doc.buildFooter()
	printChapter := func(chapNum int, titleStr, fileStr string, addPage bool) {
		if addPage {
			doc.pdf.AddPage()
			doc.buildTitle(mainTitle)
		}
		doc.buildAreaTitle(titleStr)
		doc.buildClause(chapNum)
		doc.buildBody(fileStr)
	}
	printChapter(1, "COMPUTER SCIENCE", TextFile("parte1.txt"), true)
	printChapter(2, "BIG DATA", TextFile("parte2.txt"), false)
	printChapter(3, "NETWORK", TextFile("parte3.txt"), false)

	err := doc.pdf.OutputFileAndClose(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}
}

func (d *Document) buildBody(fileStr string) {
	// Read text file
	txtStr, err := ioutil.ReadFile(fileStr)
	if err != nil {
		d.pdf.SetError(err)
	}
	d.pdf.SetFont("Arial", "", 12)
	// Output justified text
	d.pdf.MultiCell(0, 5, d.translator(string(txtStr)), "", "", false)
	// Line break
	d.pdf.Ln(-1)
}

func (d *Document) buildClause(clauseNumber int) {
	clause := fmt.Sprintf("Clause %d:", clauseNumber)
	// Arial 12
	d.pdf.SetFont("Arial", "B", 12)
	// Title
	d.pdf.CellFormat(0, 6, d.translator(clause),
		"", 1, "L", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildAreaTitle(titleStr string) {
	// 	// Arial 12
	d.pdf.SetFont("Arial", "B", 12)
	// Title
	d.pdf.CellFormat(0, 6, d.translator(titleStr),
		"", 1, "L", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildFooter() {
	d.pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		d.pdf.SetY(-15)
		// Arial italic 8
		d.pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		d.pdf.SetTextColor(128, 128, 128)
		// Page number
		d.pdf.CellFormat(0, 10, d.translator(fmt.Sprintf("Pág. %d", d.pdf.PageNo())),
			"", 0, "R", false, 0, "")
	})
}

func (d *Document) buildTitle(titleStr string) {
	// 	// Arial 12
	d.pdf.SetFont("Arial", "B", 14)
	// Title
	d.pdf.CellFormat(0, 9, d.translator(titleStr),
		"", 1, "C", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildHeader(titleStr string) {
	d.pdf.SetHeaderFunc(func() {
		// Arial bold 15
		d.pdf.SetFont("Arial", "B", 14)
		// Calculate width of title and position
		wd := d.pdf.GetStringWidth(titleStr) + 6
		d.pdf.SetX((210 - wd) / 2)
		// Title
		d.pdf.CellFormat(wd, 9, d.translator(titleStr), "0", 0, "C", false, 0, "")
		// Line break
		d.pdf.Ln(10)
	})
}

func TextFile(fileStr string) string {
	return filepath.Join("./text", fileStr)
}
