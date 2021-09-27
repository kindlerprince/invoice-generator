package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("roboto-bold", "Roboto/Roboto-Bold.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont("roboto-regular", "Roboto/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("roboto-light", "Roboto/Roboto-Light.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont("roboto-italic", "Roboto/Roboto-Italic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("roboto-bold", "", 18)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX(30)
	pdf.SetY(40)
	pdf.Text("TAX INVOICE")
	//keeping this for sometime needs to change this font
	err = pdf.SetFontSize(10)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(430)
	pdf.SetY(55)
	pdf.Text("ORIGINAL for RECIPIENT")

	pdf.SetX(30)
	pdf.SetY(70)
	err = pdf.SetFont("roboto-bold", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Text("No:- 0011030023 | ")
	err = pdf.SetFont("roboto-regular", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Text("Issue Date ")
	err = pdf.SetFont("roboto-bold", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Text("25.05.2020")

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("solid")
	pdf.Line(30, 80, 555, 80)

	pdf.SetX(30)
	pdf.SetY(100)
	pdf.Text("Alliance Broadband Services Pvt. Ltd.")

	pdf.SetX(30)
	pdf.SetY(125)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("City: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Patna")
	pdf.SetX(30)
	pdf.SetY(135)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("Address: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("United Colors of Benetton, East GandhiMaidan Road, Patna, Bihar-800004")
	pdf.SetX(30)
	pdf.SetY(145)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("PAN No: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("AAECA3151B")
	pdf.SetX(30)
	pdf.SetY(155)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("GST No: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("10AAECA3151B1ZP")
	pdf.SetX(30)
	pdf.SetY(165)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("State: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Bihar")
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text(" Code: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("10")
	pdf.SetX(30)
	pdf.SetY(175)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("SAC No: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("998422")
	pdf.SetX(30)
	pdf.SetY(185)
	pdf.SetFont("roboto-regular", "", 8)
	pdf.Text("Phone: 033-71002000, Toll Free No: 1800 1200 300 ")
	pdf.Text("www.alliancebroadband.co.in")
	pdf.AddExternalLink("https://alliancebroadband.co.in/", 27.5, 28, 125, 15)

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("solid")
	pdf.Line(30, 190, 555, 190)
	pdf.SetX(30)
	pdf.SetY(208)
	pdf.SetFont("roboto-italic", "", 14)
	pdf.Text("TO: ")
	pdf.SetFont("roboto-bold", "", 14)
	pdf.Text("PRINCE KUMAR")
	pdf.SetX(30)
	pdf.SetY(228)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("Address: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("CHHOTA TELPA CHAPRA SARAN BIHAR PIN-841301 7542908927")
	pdf.SetX(30)
	pdf.SetY(238)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("State: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Bihar")
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text(" Code: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("10")

	rectFillColor(&pdf, "N", 8, 30, 250, 18, 14, color{0, 0, 0}, color{255, 255, 255}, alignCenter, valignMiddle, 0.1)
	rectFillColor(&pdf, " Description of goods or services", 8, 49, 250, 450, 14, color{0, 0, 0}, color{255, 255, 255}, alignLeft, valignMiddle, 0.1)
	rectFillColor(&pdf, "amount  ", 8, 500, 250, 55, 14, color{0, 0, 0}, color{255, 255, 255}, alignRight, valignMiddle, 0.1)

	grey := color{197, 197, 197}
	black := color{0, 0, 0}
	borderWidth := 0.5
	rectFillColor(&pdf, "1", 8, 30, 265, 18, 16, grey, black, alignCenter, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, " fee \"BROWSE+\" (25.05.2020 to 23.06.2020)", 8, 48, 265, 452, 16, grey, black, alignLeft, valignMiddle, borderWidth)
	rectFillColor(&pdf, "400.000 ", 8, 500, 265, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	//rectFillColor(&pdf, "1", 8, 30, 265, 18, 16, grey, black, alignCenter, valignMiddle, borderWidth)
	pdf.SetFont("roboto-italic", "", 8)
	rectFillColor(&pdf, "TOTAL AMOUNT ", 8, 30, 281, 470, 16, grey, black, alignRight, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, "400.00 ", 8, 500, 281, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	pdf.SetFont("roboto-italic", "", 8)
	rectFillColor(&pdf, "CGST (9%) ", 8, 30, 297, 470, 16, grey, black, alignRight, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, "36.00 ", 8, 500, 297, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	pdf.SetFont("roboto-italic", "", 8)
	rectFillColor(&pdf, "SGST (9%) ", 8, 30, 313, 470, 16, grey, black, alignRight, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, "36.00 ", 8, 500, 313, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	pdf.SetFont("roboto-italic", "", 8)
	rectFillColor(&pdf, "TOTAL ", 8, 30, 329, 470, 16, grey, black, alignRight, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, "472.00 ", 8, 500, 329, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	pdf.SetFont("roboto-italic", "", 8)
	rectFillColor(&pdf, "Rounded off ", 8, 30, 345, 470, 16, grey, black, alignRight, valignMiddle, borderWidth)
	pdf.SetFont("roboto-regular", "", 8)
	rectFillColor(&pdf, "472.00 ", 8, 500, 345, 55, 16, grey, black, alignRight, valignMiddle, borderWidth)

	rectFillColor(&pdf, " IN WORDS: INR Four hundred and seventy two rupee", 8, 30, 361, 525, 16, grey, black, alignLeft, valignMiddle, borderWidth)

	rectFillColor(&pdf, " Payment method: [_] Cheque [_] D.D/P.O. [_] Cash", 8, 30, 380, 525, 16, grey, black, alignLeft, valignMiddle, borderWidth)
	rectFillColor(&pdf, " Date of occurrence of chargeable event / payment: 25.05.2020 / 25.05.2020", 8, 30, 396, 525, 16, grey, black, alignLeft, valignMiddle, borderWidth)

	pdf.SetX(30)
	pdf.SetY(423)
	pdf.SetFont("roboto-bold", "", 6.8)
	pdf.Text("TERMS AND CONDITIONS")
	pdf.SetX(30)
	pdf.SetY(430)
	pdf.SetFont("roboto-italic", "", 6)
	pdf.Text("1) It will be deemed that you have accepted this Invoice in full in the event you have not lodged any written objection with us within 20 days of receipt of this Invoice.")
	pdf.SetX(30)
	pdf.SetY(437)
	pdf.Text("2) To avoid disconnection of service you are requested to pay the full amount by the due date mentioned in the invoice. An interest of 18% per annum will be charged on the amount")
	pdf.SetX(30)
	pdf.SetY(444)
	pdf.Text("remaining unpaid after the due date.")
	pdf.SetX(30)
	pdf.SetY(451)
	pdf.Text("3) All Cheques/Demand Drafts in payment of Invoice should be drawn in favour of \"Alliance Broadband Services Pvt. Ltd.\".")
	pdf.SetX(30)
	pdf.SetY(458)
	pdf.Text("4) Kindly mention invoice number along with your payment to ensure correct and timely processing.")
	pdf.SetX(30)
	pdf.SetY(465)
	pdf.Text("5) Cheque Return Charges of Rs. 250 would be charged extra")
	pdf.SetX(30)
	pdf.SetY(472)
	pdf.Text("6) E. & O. E.")

	pdf.SetX(30)
	pdf.SetY(490)
	pdf.SetFontSize(8)
	pdf.Text("Issuer: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Rajesh Cable TV (Bihar) - 04")
	pdf.SetX(320)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("Reciever:")
	pdf.SetX(320)
	pdf.SetY(500)
	pdf.Text("Client ID: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("6534169655")
	pdf.SetX(320)
	pdf.SetY(510)
	pdf.SetFont("roboto-italic", "", 8)
	pdf.Text("Authorized Signature:")

	pdf.SetX(30)
	pdf.SetY(530)
	pdf.Text("Authorized Signature")

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("solid")
	pdf.Line(30, 580, 555, 580)

	pdf.SetX(30)
	pdf.SetY(590)
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Additional user details: ")
	pdf.SetFont("roboto-regular", "", 8)
	pdf.Text("Username: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("princek_rcbt")
	pdf.SetX(30)
	pdf.SetY(600)
	pdf.SetFont("roboto-regular", "", 8)
	pdf.Text("IP Address: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("172.21.181.154")
	pdf.SetX(30)
	pdf.SetY(610)
	pdf.SetFont("roboto-regular", "", 8)
	pdf.Text("Zone: ")
	pdf.SetFont("roboto-bold", "", 8)
	pdf.Text("Rajesh Cable TV (Bihar) - 04")
	pdf.SetLineWidth(0.5)
	pdf.SetLineType("solid")
	pdf.Line(30, 630, 555, 630)

	pdf.Image("QRCode.jpeg", 490, 500, nil) //print image
	pdf.Image("sign.jpeg", 30, 500, nil)
	pdf.WritePdf("invoice.pdf")
}

const (
	valignTop    = 1
	valignMiddle = 2
	valignBottom = 3
)

const (
	alignLeft   = 4
	alignCenter = 5
	alignRight  = 6
)

type color struct {
	r, g, b uint8
}

// rectFillColor takes pdf pointer, text , text's font size,
// rectange with x coordinate from left, y coordinate from top corner, width of rectangle, height of rectangle, background color
// foreground color, horizontal alignment and vertical alignment of text to be filled in rectangle

func rectFillColor(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	bg color,
	fg color,
	align, valign int, lineWidth float64,
) {
	pdf.SetLineWidth(lineWidth)
	pdf.SetFillColor(bg.r, bg.g, bg.b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)

	if align == alignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == alignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	}

	pdf.SetX(x)

	if valign == valignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == valignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetY(y)
	pdf.SetTextColor(fg.r, fg.g, fg.b)
	pdf.Cell(nil, text)
}
