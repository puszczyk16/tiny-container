package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	version = "0.1"
)

var (
	conMeta   meta
	counter   int
	imReady   bool = false
)

type meta struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip",omitempty`
	ExtIP    string `json:"extip",omitempty`
}

func (m *meta) newMeta() {
	m.Hostname, _ = os.Hostname()
}

func (m *meta) getIP() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting interfaces: %s", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("Error getting ip from %v: %s", i, err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// fmt.Printf("%s\n", ip)
			if !ip.IsLoopback() {
				m.IP = fmt.Sprintf("%s", ip)
				break
			}
		}
	}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if imReady {
		w.WriteHeader(200)
		fmt.Fprintf(w, "OK")
	} else {
		w.WriteHeader(503)
		fmt.Fprintf(w, "Not ready")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "%-15s %s\n", "Hostname:", conMeta.Hostname)
	fmt.Fprintf(w, "%-15s %s\n", "IP:", conMeta.IP)
	fmt.Fprintf(w, "%-15s %s\n", "External IP:", conMeta.ExtIP)
	fmt.Fprintf(w, "%-15s %d\n", "Request count:", counter)
	fmt.Fprintf(w, "%-15s %s\n", "Version:", version)
}

/*
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	counter += 1
	cj, _ := json.Marshal(conMeta)
	fmt.Fprintf(w, "%s", string(cj))
}
*/

func main() {
	conMeta.newMeta()
	conMeta.getIP()
	counter = 0

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/", handler)
	// http.HandleFunc("/json", jsonHandler)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
