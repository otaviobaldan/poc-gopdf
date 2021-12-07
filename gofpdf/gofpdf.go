package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Document struct {
	pdf        *gofpdf.Fpdf
	translator func(string) string
}

func NewDocument() *Document {
	pdf := gofpdf.New("P", "pt", "A4", "")
	return &Document{
		pdf:        pdf,
		translator: pdf.UnicodeTranslatorFromDescriptor(""),
	}
}

const margin = 40
const lorem = "Lorem ipsum dolor sit amet consectetur adipiscing elit tellus rutrum suspendisse aliquet dignissim sem, faucibus tempor erat quisque vehicula sociosqu magna in praesent cursus habitant felis. Turpis ornare taciti habitant posuere inceptos vivamus viverra vulputate tempus convallis, molestie condimentum iaculis pellentesque leo ante est sed placerat curabitur vitae, suscipit mattis cursus fringilla lobortis litora sollicitudin justo nunc. Dignissim facilisi lectus natoque fermentum risus etiam integer mi iaculis nam ornare, augue porttitor blandit aliquet elementum sagittis faucibus habitasse et vitae."

func main_test() {
	var pdf = gofpdf.New("P", "pt", "A4", "")
	pdf.AddUTF8Font("Bookman", "", "./font/bookman-old-style.ttf")

	pdf.SetMargins(85, 85, 56.7)

	pdf.SetAutoPageBreak(true, margin)

	pdf.AddPage()

	txtStr, err := ioutil.ReadFile("./text/parte1.txt")
	if err != nil {
		pdf.SetError(err)
	}

	pdf.SetFont("Bookman", "", 12)
	pdf.MultiCell(0.0, 18.0, string(txtStr), "", "J", false)

	err = pdf.OutputFileAndClose("./pdfs/test2.pdf")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

}

func main() {
	doc := NewDocument()

	mainTitle := "Lero Lero de TI"

	doc.pdf.SetMargins(85, 85, 56.7)

	doc.pdf.SetAutoPageBreak(true, margin)
	doc.pdf.AddPage()

	doc.pdf.SetTitle(mainTitle, true)
	doc.pdf.SetAuthor("obaldan", true)

	doc.registerFonts()

	doc.buildFooter()
	printChapter := func(chapNum int, titleStr, fileStr string, addPage bool) {
		if addPage {
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
	d.pdf.SetFont("Bookman", "", 12)
	// Output justified text
	d.pdf.MultiCell(0.0, 18.0, string(txtStr), "", "J", false)
	// Line break
	d.pdf.Ln(-1)
}

func (d *Document) buildClause(clauseNumber int) {
	clause := fmt.Sprintf("Clause %d:", clauseNumber)
	// Bookman 12
	d.pdf.SetFont("Bookman", "B", 12)
	// Title
	d.pdf.CellFormat(0, 17, clause,
		"", 1, "L", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildAreaTitle(titleStr string) {
	// 	// Bookman 12
	d.pdf.SetFont("Bookman", "B", 12)
	// Title
	d.pdf.CellFormat(0, 17, titleStr,
		"", 1, "L", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildFooter() {
	d.pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		d.pdf.SetY(-42.5)
		// Bookman italic 8
		d.pdf.SetFont("Bookman", "I", 8)

		d.pdf.SetTextColor(224, 224, 224)
		// Page number
		d.pdf.CellFormat(0, 28.34, "Contrato emitido pela plataforma maiscontratos.com em 06/12/2021",
			"", 0, "C", false, 0, "")

		// Text color in gray
		d.pdf.SetTextColor(128, 128, 128)
		// Page number
		d.pdf.CellFormat(0, 28.34, fmt.Sprintf("Pág. %d", d.pdf.PageNo()),
			"", 0, "R", false, 0, "")

	})
}

func (d *Document) buildTitle(titleStr string) {
	// 	// Bookman 12
	d.pdf.SetFont("Bookman", "B", 14)
	// Title
	d.pdf.CellFormat(0, 25, titleStr,
		"", 1, "C", false, 0, "")
	// Line break
	d.pdf.Ln(4)
}

func (d *Document) buildHeader(titleStr string) {
	d.pdf.SetHeaderFunc(func() {
		// Bookman bold 15
		d.pdf.SetFont("Bookman", "B", 14)
		// Calculate width of title and position
		wd := d.pdf.GetStringWidth(titleStr) + 6
		d.pdf.SetX((210 - wd) / 2)
		// Title
		d.pdf.CellFormat(wd, 9, titleStr, "0", 0, "C", false, 0, "")
		// Line break
		d.pdf.Ln(10)
	})
}

func (d *Document) registerFonts() {
	//jsonFileBytes, _ := ioutil.ReadFile("./font/bookman-old-style.json")
	//zFileBytes, _ := ioutil.ReadFile("./font/bookman-old-style.z")
	//d.pdf.AddFontFromBytes("Bookman", "", jsonFileBytes, zFileBytes)
	//
	//jsonFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-bold.json")
	//zFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-bold.z")
	//d.pdf.AddFontFromBytes("Bookman", "B", jsonFileBytes, zFileBytes)
	//
	//jsonFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-italic.json")
	//zFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-italic.z")
	//d.pdf.AddFontFromBytes("Bookman", "I", jsonFileBytes, zFileBytes)
	//
	//jsonFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-bold-italic.json")
	//zFileBytes, _ = ioutil.ReadFile("./font/bookman-old-style-bold-italic.z")
	//d.pdf.AddFontFromBytes("Bookman", "BI", jsonFileBytes, zFileBytes)

	d.pdf.AddUTF8Font("Bookman", "", "./font/bookman-old-style.ttf")
	d.pdf.AddUTF8Font("Bookman", "B", "./font/bookman-old-style-bold.ttf")
	d.pdf.AddUTF8Font("Bookman", "BI", "./font/bookman-old-style-bold-italic.ttf")
	d.pdf.AddUTF8Font("Bookman", "I", "./font/bookman-old-style-italic.ttf")
}

func TextFile(fileStr string) string {
	return filepath.Join("./text", fileStr)
}
