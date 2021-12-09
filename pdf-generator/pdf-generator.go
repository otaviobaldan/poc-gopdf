package main

import (
	"fmt"
	go_pdf_generator "github.com/otaviobaldan/go-pdf-generator"
	"github.com/otaviobaldan/go-pdf-generator/config"
	"github.com/otaviobaldan/go-pdf-generator/constants"
	"io/ioutil"
	"time"
)

func main() {
	pdfConfig := config.NewPdfConfig(
		constants.OrientationPortrait,
		constants.UnitsPoints,
		constants.PaperSizeA4,
		85,
		85,
		56.7,
		true,
	)
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
	bytes, _ := ioutil.ReadFile("./text/parte1.txt")
	pdf.GenerateText(string(bytes))

	pdf.GenerateSubtitle("CHAPTER 2")
	bytes, _ = ioutil.ReadFile("./text/parte2.txt")
	pdf.GenerateText(string(bytes))

	//htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
	//	`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
	//	`<center>You can also center text.</center>` +
	//	`<right>Or align it to the right.</right>` +
	//	`You can also insert links on text, such as ` +
	//	`<a href="http://www.fpdf.org">www.fpdf.org</a>, or on an image: click on the logo.`
	pdf.GenerateSubtitle("CHAPTER 3")
	bytes, _ = ioutil.ReadFile("./text/parte3.txt")
	html := pdf.Pdf.HTMLBasicNew()
	html.Write(12, string(bytes))

	err = pdf.Pdf.OutputFileAndClose(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}
}
