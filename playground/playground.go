package playground

import (
	"playground/css"
	"playground/js"
	//	"comone/css"
	//	"comone/js"
)

// Start Page :
func StartPage() string {
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "Start Page</title>"
	xdata = css.SettingsStyle(xdata)
	xdata = js.DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>X12 Viewer</p>"
	xdata = xdata + "<BR> X12 Viewer <BR> (c) 1992-2018 Com1 Software Development"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// TestPage1 :
func TestPage1() string {
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>TestPage1</title>"
	xdata = css.SettingsStyle(xdata)
	xdata = js.DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>X12 Parser</p>"
	xdata = xdata + "<BR> X12 Parser <BR> (c) 1992-2018 Com1 Software Development"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}
