package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"time"
)

var rowHeight = float64(5)

func main() {
	page := pdf.NewMaroto(consts.Portrait, consts.A4)
	page.SetPageMargins(20, 10, 20)
	page.SetDefaultFontFamily(consts.Arial)

	buildHeader(page)

	buildTitle(page, "CONTRATO DE HONORÁRIOS ADVOCATÍCIOS")
	buildArea(page, "DAS PARTES")
	buildText(page, "CONTRATANTE(S): [NOME CONTRATANTE], [Nacionalidade], [Estado Civil], [Profissão], [com Cédula de Identidade (RG) sob o nº [RG], inscrito no CPF sob o nº [CPF] [ou] [inscrito no CNPJ sob o nº [CNPJ]], [residente e domiciliado] [ou] [situada/sediada] na [Rua], [nº], [Complemento], [Bairro], [Cidade], [Sigla Estado], [CEP], doravante designado CONTRATANTE;")
	buildText(page, "CONTRATADO: [NOME EMPRESARIAL], pessoa jurídica de direito privado, [inscrita na Ordem dos Advogados do Brasil (OAB), seção de [Estado], sob o nº [nº OAB]], inscrita no Cadastro Nacional de Pessoas Jurídicas (CNPJ) sob o nº [CNPJ], com escritório situado na [Rua], [nº], [Complemento], [Bairro], [Cidade], [Sigla Estado], [CEP], [tendo como seu sócio], [NOME], [Nacionalidade], [Estado Civil], Advogado, inscrito na Ordem dos Advogados do Brasil (OAB) sob o nº [OAB], portador de Cédula de Identidade (RG) nº [RG], inscrito no CPF sob o nº [CPF], [com endereço na  Rua], [nº], [Complemento], [Bairro], [Cidade], [Sigla Estado], [CEP]] [OU] [com o mesmo endereço anteriormente mencionado], doravante designado CONTRATADO.")
	buildText(page, "As partes acima identificadas têm, entre si, justo e acertado o presente Contrato de Honorários, que será regido pelas cláusulas seguintes e pelas condições descritas no presente.`)")

	err := page.OutputFileAndClose(fmt.Sprintf("./pdfs/%s.pdf", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}
}

func buildTitle(page pdf.Maroto, title string) {
	page.Row(rowHeight, func() {
		page.Col(12, func() {
			page.Text(title, props.Text{
				Align: consts.Center,
				Style: consts.Bold,
				Size:  14,
			})
		})
	})
}

func buildArea(page pdf.Maroto, area string) {
	page.Row(rowHeight, func() {
		page.Col(12, func() {
			page.Text(area, props.Text{
				Align: consts.Left,
				Style: consts.Bold,
				Size:  12,
			})
		})
	})
}

func buildText(page pdf.Maroto, text string) {
	page.Row(rowHeight, func() {
		page.Col(12, func() {
			page.Text(text, props.Text{
				Size:            13,
				Align:           consts.Left,
				Top:             50,
				VerticalPadding: 2.0,
			})
		})
	})
}

func buildHeader(page pdf.Maroto) {
	page.RegisterHeader(func() {
		page.Row(15, func() {
			page.Col(3, func() {
				page.Text("lorem ipsum dolor", props.Text{Align: consts.Left})
			})
			page.Col(3, func() {
				page.Text(time.Now().Format("02-January-2006"),
					props.Text{Align: consts.Right})
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
