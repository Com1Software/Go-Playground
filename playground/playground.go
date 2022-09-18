package playground

import (
	"playground/css"
	"playground/js"
)

//
//------------------ (c) 1992-2022 Com1 Software Development
//


// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

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
	xdata = xdata + "<p>XStart Page</p>"
	xdata = xdata + "<BR> Go Playground for Programers <BR> (c) 1992-2022 Com1 Software Development"
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
	xdata = xdata + "<p>Test Page 1</p>"
	xdata = xdata + "<BR> Test page 1 <BR> (c) 1992-2022 Com1 Software Development"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}
