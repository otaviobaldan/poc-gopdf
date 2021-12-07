package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"image"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"time"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	//err := RunPdfReport("unidoc-report.pdf")
	//if err != nil {
	//	panic(err)
	//}

	//err := CreateContract(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	err := CreateContract("./pdfs/unidoc-test.pdf")
	if err != nil {
		panic(err)
	}
}

func CreateContract(outputPath string) error {
	bookmanFontStd, err := model.NewPdfFontFromTTFFile("./font/bookman-old-style.ttf")
	if err != nil {
		return err
	}

	bookmanFontBold, err := model.NewPdfFontFromTTFFile("./font/bookman-old-style-bold.ttf")
	if err != nil {
		return err
	}

	c := creator.New()
	c.SetPageMargins(30, 20, 30, 20)
	drawFooter(c, bookmanFontStd)

	// title
	p := c.NewParagraph("Lero Lero de TI")
	p.SetFont(bookmanFontBold)
	p.SetFontSize(14)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetColor(creator.ColorRGBFrom8bit(0, 0, 0))
	c.Draw(p)

	ch := c.NewChapter("Clausula 1")

	chapterFont := bookmanFontBold
	chapterFontColor := creator.ColorRGBFrom8bit(0, 0, 0)
	chapterFontSize := 12.0

	normalFont := bookmanFontStd
	normalFontColor := creator.ColorRGBFrom8bit(0, 0, 0)
	normalFontSize := 12.0

	ch.GetHeading().SetFont(chapterFont)
	ch.GetHeading().SetFontSize(chapterFontSize)
	ch.GetHeading().SetColor(chapterFontColor)

	c.Draw(ch)

	// Paragraphs.
	sc := ch.NewSubchapter("COMPUTER SCIENCE")
	//sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	txtStr, err := ioutil.ReadFile(TextFile("parte1.txt"))
	p = c.NewParagraph(string(txtStr))
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	//p.SetMargins(0, 0, 5, 0)
	p.SetTextAlignment(creator.TextAlignmentJustify)
	sc.Add(p)

	c.Draw(sc)

	//chapter := drawTitle(c, bookmanFontBold, "Clause 1")
	//subChapter := drawArea(chapter, bookmanFontBold, "COMPUTER SCIENCE")
	//drawChapterBody(c, subChapter, bookmanFontStd, "parte1.txt")

	err = c.WriteToFile(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func TextFile(fileStr string) string {
	return filepath.Join("./text", fileStr)
}

func drawChapterBody(c *creator.Creator, sc *creator.Chapter, bookmanFontStd *model.PdfFont, filePath string) {
	text := TextFile(filePath)
	p := c.NewParagraph(text)
	p.SetFont(bookmanFontStd)
	p.SetFontSize(12)
	p.SetColor(creator.ColorRGBFrom8bit(0, 0, 0))
	p.SetTextAlignment(creator.TextAlignmentJustify)

	sc.Add(p)
}

func drawArea(ch *creator.Chapter, bookmanFontBold *model.PdfFont, area string) *creator.Chapter {
	// Paragraphs.
	sc := ch.NewSubchapter(area)
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(bookmanFontBold)
	sc.GetHeading().SetFontSize(12)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(0, 0, 0))

	ch.Add(sc)

	return sc
}

func drawTitle(c *creator.Creator, bookmanFontBold *model.PdfFont, title string) *creator.Chapter {
	ch := c.NewChapter(title)

	chapterFont := bookmanFontBold
	chapterFontColor := creator.ColorRGBFrom8bit(0, 0, 0)
	chapterFontSize := 14.0

	ch.GetHeading().SetFont(chapterFont)
	ch.GetHeading().SetFontSize(chapterFontSize)
	ch.GetHeading().SetColor(chapterFontColor)

	return ch
}

func drawFooter(c *creator.Creator, bookmanFontStd *model.PdfFont) {
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := c.NewParagraph("Contrato emitido pela plataforma maiscontratos.com em 06/12/2021")
		p.SetFont(bookmanFontStd)
		p.SetFontSize(8)
		//p.SetPos(50, 20)
		p.SetTextAlignment(creator.TextAlignmentCenter)
		p.SetColor(creator.ColorRGBFrom8bit(224, 224, 224))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = c.NewParagraph(strPage)
		p.SetFont(bookmanFontStd)
		p.SetFontSize(8)
		p.SetTextAlignment(creator.TextAlignmentRight)
		//p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(128, 128, 128))
		block.Draw(p)
	})
}

func RunPdfReport(outputPath string) error {
	robotoFontRegular, err := model.NewPdfFontFromTTFFile("./font/bookman-old-style.ttf")
	if err != nil {
		return err
	}

	robotoFontPro, err := model.NewPdfFontFromTTFFile("./font/bookman-old-style-bold.ttf")
	if err != nil {
		return err
	}

	c := creator.New()
	c.SetPageMargins(30, 20, 30, 20)

	// Generate the table of contents.
	c.AddTOC = true
	toc := c.TOC()
	hstyle := c.NewTextStyle()
	hstyle.Color = creator.ColorRGBFromArithmetic(0.2, 0.2, 0.2)
	hstyle.FontSize = 28
	toc.SetHeading("Table of Contents", hstyle)
	lstyle := c.NewTextStyle()
	lstyle.FontSize = 14
	toc.SetLineStyle(lstyle)

	//logoImg, err := c.NewImageFromFile("./unidoc-logo.png")
	//if err != nil {
	//	return err
	//}
	//
	//logoImg.ScaleToHeight(25)
	//logoImg.SetPos(58, 20)

	DoDocumentControl(c, robotoFontRegular, robotoFontPro)

	DoFeatureOverview(c, robotoFontRegular, robotoFontPro)

	// Setup a front page (always placed first).
	c.CreateFrontPage(func(args creator.FrontpageFunctionArgs) {
		DoFirstPage(c, robotoFontRegular, robotoFontPro)
	})

	// Draw a header on each page.
	//c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
	//	// Draw the header on a block. The block size is the size of the page's top margins.
	//	block.Draw(logoImg)
	//})

	// Draw footer on each page.
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := c.NewParagraph("unidoc.io")
		p.SetFont(robotoFontRegular)
		p.SetFontSize(8)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = c.NewParagraph(strPage)
		p.SetFont(robotoFontRegular)
		p.SetFontSize(8)
		p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	err = c.WriteToFile(outputPath)
	if err != nil {
		return err
	}

	return nil
}

// Generates the front page.
func DoFirstPage(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	p := c.NewParagraph("UniDoc")
	p.SetFont(helvetica)
	p.SetFontSize(48)
	p.SetMargins(85, 0, 150, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	p = c.NewParagraph("Example Report")
	p.SetFont(helveticaBold)
	p.SetFontSize(30)
	p.SetMargins(85, 0, 0, 0)
	p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	c.Draw(p)

	t := time.Now().UTC()
	dateStr := t.Format("1 Jan, 2006 15:04")

	p = c.NewParagraph(dateStr)
	p.SetFont(helveticaBold)
	p.SetFontSize(12)
	p.SetMargins(90, 0, 5, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)
}

// Document control page.
func DoDocumentControl(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	ch := c.NewChapter("Document control")
	ch.SetMargins(0, 0, 40, 0)
	ch.GetHeading().SetFont(fontRegular)
	ch.GetHeading().SetFontSize(18)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	sc := ch.NewSubchapter("Issuer details")
	sc.GetHeading().SetFont(fontRegular)
	sc.GetHeading().SetFontSize(18)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	issuerTable := c.NewTable(2)
	issuerTable.SetMargins(0, 0, 30, 0)

	pColor := creator.ColorRGBFrom8bit(72, 86, 95)
	bgColor := creator.ColorRGBFrom8bit(56, 68, 67)

	p := c.NewParagraph("Issuer")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell := issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("UniDoc")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Address")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("Klapparstig 16, 101 Reykjavik, Iceland")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Email")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBackgroundColor(bgColor)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("sales@unidoc.io")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Web")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("unidoc.io")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Author")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("UniDoc report generator")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	sc.Add(issuerTable)

	// 1.2 - Document history
	sc = ch.NewSubchapter("Document History")
	sc.SetMargins(0, 0, 5, 0)
	sc.GetHeading().SetFont(fontRegular)
	sc.GetHeading().SetFontSize(18)
	sc.GetHeading().SetColor(pColor)

	histTable := c.NewTable(3)
	histTable.SetMargins(0, 0, 30, 50)

	histCols := []string{"Date Issued", "UniDoc Version", "Type/Change"}
	for _, histCol := range histCols {
		p = c.NewParagraph(histCol)
		p.SetFont(fontBold)
		p.SetFontSize(10)
		p.SetColor(creator.ColorWhite)
		cell = histTable.NewCell()
		cell.SetBackgroundColor(bgColor)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		cell.SetContent(p)
	}

	dateStr := common.ReleasedAt.Format("1 Jan, 2006 15:04")

	histVals := []string{dateStr, common.Version, "First issue"}
	for _, histVal := range histVals {
		p = c.NewParagraph(histVal)
		p.SetFont(fontRegular)
		p.SetFontSize(10)
		p.SetColor(pColor)
		cell = histTable.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		cell.SetContent(p)
	}

	sc.Add(histTable)

	err := c.Draw(ch)
	if err != nil {
		panic(err)
	}
}

// Chapter giving an overview of features.
// TODO: Add code snippets and show more styles and options.
func DoFeatureOverview(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	// Ensure that the chapter starts on a new page.
	c.NewPage()

	ch := c.NewChapter("Feature overview")

	chapterFont := fontRegular
	chapterFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	chapterFontSize := 18.0

	normalFont := fontRegular
	normalFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	normalFontSize := 10.0

	bgColor := creator.ColorRGBFrom8bit(56, 68, 67)

	ch.GetHeading().SetFont(chapterFont)
	ch.GetHeading().SetFontSize(chapterFontSize)
	ch.GetHeading().SetColor(chapterFontColor)

	p := c.NewParagraph("This chapter demonstrates a few of the features of UniDoc that can be used for report generation.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	ch.Add(p)

	// Paragraphs.
	sc := ch.NewSubchapter("Paragraphs")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Paragraphs are used to represent text, as little as a single character, a word or " +
		"multiple words forming multiple sentences. UniDoc handles automatically wrapping those across lines and pages, making " +
		"it relatively easy to work with. They can also be left, center, right aligned or justified as illustrated below:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	// Example paragraphs:
	loremTxt := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt" +
		"ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore" +
		"eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt " +
		"mollit anim id est laborum."
	alignments := []creator.TextAlignment{creator.TextAlignmentLeft, creator.TextAlignmentCenter,
		creator.TextAlignmentRight, creator.TextAlignmentJustify}
	for j := 0; j < 4; j++ {
		p = c.NewParagraph(loremTxt)
		p.SetFont(normalFont)
		p.SetFontSize(normalFontSize)
		p.SetColor(normalFontColor)
		p.SetMargins(20, 0, 10, 10)
		p.SetTextAlignment(alignments[j%4])

		sc.Add(p)
	}

	sc = ch.NewSubchapter("Tables")
	// Mock table: Priority table.
	priTable := c.NewTable(2)
	priTable.SetMargins(40, 40, 10, 0)
	// Column headers:
	tableCols := []string{"Priority", "Items fulfilled / available"}
	for _, tableCol := range tableCols {
		p = c.NewParagraph(tableCol)
		p.SetFont(fontBold)
		p.SetFontSize(10)
		p.SetColor(creator.ColorWhite)
		cell := priTable.NewCell()
		cell.SetBackgroundColor(bgColor)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetContent(p)
	}
	items := [][]string{
		[]string{"High", "52/80"},
		[]string{"Medium", "32/100"},
		[]string{"Low", "10/90"},
	}
	for _, lineItems := range items {
		for _, item := range lineItems {
			p = c.NewParagraph(item)
			p.SetFont(fontBold)
			p.SetFontSize(10)
			p.SetColor(creator.ColorWhite)
			cell := priTable.NewCell()
			cell.SetBackgroundColor(bgColor)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			cell.SetContent(p)
		}
	}
	sc.Add(priTable)

	sc = ch.NewSubchapter("Images")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Images can be loaded from multiple file formats, example from a PNG image:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 5)
	sc.Add(p)

	// Show logo.
	//img, err := c.NewImageFromFile("./unidoc-logo.png")
	//if err != nil {
	//	panic(err)
	//}
	//img.ScaleToHeight(50)
	//sc.Add(img)

	sc = ch.NewSubchapter("QR Codes / Barcodes")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Example of a QR code generated with package github.com/boombuler/barcode:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 5)
	sc.Add(p)

	qrCode, _ := makeQrCodeImage("HELLO", 40, 5)
	img, err := c.NewImageFromGoImage(qrCode)
	if err != nil {
		panic(err)
	}
	img.SetWidth(40)
	img.SetHeight(40)
	sc.Add(img)

	sc = ch.NewSubchapter("Graphing / Charts")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Graphs can be generated via packages such as github.com/wcharczuk/go-chart as illustrated " +
		"in the following plot:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	//graph := chart.PieChart{
	//	Width:  200,
	//	Height: 200,
	//	Values: []chart.Value{
	//		{Value: 70, Label: "Compliant"},
	//		{Value: 30, Label: "Non-Compliant"},
	//	},
	//}
	//
	//buffer := bytes.NewBuffer([]byte{})
	//err = graph.Render(chart.PNG, buffer)
	//if err != nil {
	//	panic(err)
	//}
	//img, err = c.NewImageFromData(buffer.Bytes())
	//if err != nil {
	//	panic(err)
	//}
	//img.SetMargins(0, 0, 10, 0)
	//sc.Add(img)

	sc = ch.NewSubchapter("Headers and footers")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Convenience functions are provided to generate headers and footers, see: " +
		"https://godoc.org/github.com/unidoc/unipdf/creator#Creator.DrawHeader and " +
		"https://godoc.org/github.com/unidoc/unipdf/creator#Creator.DrawFooter " +
		"They both set a function that accepts a block which the header/footer is drawn on for each page. " +
		"More information is provided in the arguments, allowing to skip header/footer on specific pages and " +
		"showing page number and count.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	sc = ch.NewSubchapter("Table of contents generation")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("A convenience function is provided to generate table of contents " +
		"as can be seen on https://godoc.org/github.com/unidoc/unipdf/creator#Creator.CreateTableOfContents and " +
		"in our example code on unidoc.io.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	c.Draw(ch)
}

// Helper function to make the QR code image with a specified oversampling factor.
// The oversampling specifies how many pixels/point. Standard PDF resolution is 72 points/inch.
func makeQrCodeImage(text string, width float64, oversampling int) (image.Image, error) {
	qrCode, err := qr.Encode(text, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	pixelWidth := oversampling * int(math.Ceil(width))
	qrCode, err = barcode.Scale(qrCode, pixelWidth, pixelWidth)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}
