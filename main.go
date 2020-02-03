package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"encoding/json"
	"io/ioutil"
	"text/template"
	"strings"

	"github.com/zserge/lorca"
)

// Set HTTP port
const HTTP_PORT = 8080

// structures to hold the json configuration
type CtrlGroup struct {
	Name string
	Color string
	Elements []CtrlGroupElement
}
type CtrlGroupElement struct {
	Name string
	Color string
	Endpoint string
	Parameter string
}

// Find my IP address
func firstIpv4Address() string {
	// Get all addresses from all interfaces
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	// Find the first IPv4 address that is not a loopback address
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// read the whole file at "path" into a byte array
func readFileToString(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file to read")
	}
	defer file.Close()
	fileText, _ := ioutil.ReadAll(file)
	return fileText
}

// Handle Ajax calls from the ctrl view by loading differrent UIs with different parameters
func handleAjax(ui lorca.UI) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Extract the endpoint from the url and the parameter from a GET/POST parameter
		endpoint := strings.TrimPrefix(request.URL.Path, "/ajax/")
		parameter := request.FormValue("parameter")

		log.Println("Endpoint: " + endpoint)
		log.Println("Parameter: " + parameter)

		switch endpoint {
			case "black":
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/black.html", HTTP_PORT))
			case "welcome":
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/welcome.html", HTTP_PORT))
			case "marquee":
				ui.Bind("initMarquee", func() {
					ui.Eval(fmt.Sprintf("setupMarquee('%s', %d)", parameter, 500))
				})
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/marquee.html", HTTP_PORT))
			case "video":
				ui.Bind("initVideo", func() {
					ui.Eval(fmt.Sprintf("setupVideo('%s')", parameter))
				})
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/video.html", HTTP_PORT))
			case "video-loop":
				ui.Bind("initVideo", func() {
					ui.Eval(fmt.Sprintf("setupVideoLoop('%s')", parameter))
				})
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/video-loop.html", HTTP_PORT))
			case "image":
				ui.Bind("initImage", func() {
					ui.Eval(fmt.Sprintf("setupImage('%s')", parameter))
				})
				ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/image.html", HTTP_PORT))
			default:
				fmt.Fprintf(writer, "%s", "Unknown Endpoint")
		}
		
		fmt.Fprintf(writer, "%s", "OK")
	}
}

// Answer /ctrl/ requests with a controller web view (usually used from a mobile device)
func handleCtrl(jsonContents []CtrlGroup) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Prepare HTML building blocks
		htmlHeader := readFileToString("assets/ctrl.head.tpl")
		htmlFooter := readFileToString("assets/ctrl.foot.tpl")
		tplGroupBegin, err := template.New("groupBegin").Parse(`{{define "groupBegin"}}<div class="group"><div class="card card--outlined card--{{ .Color | html }} color-{{ .Color | html }}"><h3>{{ .Name | html }}</h3><div class="card__content">{{end}}`)
		if err != nil {
			log.Fatal("Error setting up template")
		}
		tplGroupElement, err := template.New("groupElement").Parse(`{{define "groupElement"}}<p><a href="javascript:void(0);" class="button button--filled button--{{ .Color | html }}" onclick="sendPostRequest(this, '{{ .Endpoint | urlquery }}', '{{ .Parameter | html }}')">{{ .Name | html }}</a></p>{{end}}`)
		if err != nil {
			log.Fatal("Error setting up template")
		}
		htmlGroupEnd := "</div></div></div>"
		
		// Iterate over the config data structure and generate the actual HTML
		fmt.Fprintf(writer, "%s", htmlHeader)
		for _, group := range jsonContents {
			err = tplGroupBegin.ExecuteTemplate(writer, "groupBegin", group)
			for _, element := range group.Elements {
				err = tplGroupElement.ExecuteTemplate(writer, "groupElement", element)
			}
			fmt.Fprintf(writer, "%s", htmlGroupEnd)
		}
		fmt.Fprintf(writer, "%s", htmlFooter)
	}
}

func main() {
	// Parse configuration file
	// Extract file path and read the file into jsonText
	if len(os.Args) < 2 {
		log.Fatal("Please supply a configuration file as command line argument!")
	}
	jsonText := readFileToString(os.Args[1])
	
	// Parse JSON (config file)
	var jsonContents []CtrlGroup
	err := json.Unmarshal(jsonText, &jsonContents)
	if err != nil {
		log.Fatal("Could not parse config file")
	}

	// Create new lorca window
	ui, err := lorca.New("", "", 1280, 720, "--autoplay-policy=no-user-gesture-required")
	ui.SetBounds(lorca.Bounds{
		Left: 0,
		Top: 0,
		Width: 0,
		Height: 0,
		WindowState: lorca.WindowStateFullscreen,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Serve HTML files, videos, and the control endpoint
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/conf/", http.StripPrefix("/conf/", http.FileServer(http.Dir("conf"))))
	http.HandleFunc("/ajax/", handleAjax(ui))
	http.HandleFunc("/ctrl/", handleCtrl(jsonContents))
	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", HTTP_PORT), nil)

	// Show connection information as soon as the welcome UI is ready
	ui.Bind("initWelcome", func() {
		ui.Eval(fmt.Sprintf("setIpAndPort('%s', %d)", firstIpv4Address(), HTTP_PORT))
	})
	ui.Load(fmt.Sprintf("http://127.0.0.1:%d/assets/welcome.html", HTTP_PORT))

	// Wait until an interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
		case <-sigc:
		case <-ui.Done():
	}

	log.Println("exiting...")
}
