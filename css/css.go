package css

//	"fmt"

// BaseStyle : This is a base style sheet
func BaseStyle(xdata string) string {
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: lightblue;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	return xdata
}

// Test1Style : Setting style sheet
func Test1Style(xdata string) string {
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	return xdata
}
