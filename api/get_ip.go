// ipify-api/api
//
// This package holds our API handlers which we use to service REST API
// requests.

package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tech10/ipify-api/models"
	"net"
	"net/http"
	"strings"
)

// GetIP returns a user's public facing IP address (IPv4 OR IPv6).
//
// By default, it will return the IP address in plain text, but can also return
// data in JSON, JSONP, and XML if requested to.
func GetIP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// We'll always grab the first IP address in the X-Forwarded-For header
	// list.  We do this because this is always the *origin* IP address, which
	// is the *true* IP of the user.  For more information on this, see the
	// Wikipedia page: https://en.wikipedia.org/wiki/X-Forwarded-For
	ip := strings.Split(r.Header.Get("X-Forwarded-For"), ", ")[0]
fmt.Println("IP from forwarded header:", ip)
	//  Fall back to the request ip.
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
fmt.Println("IP from host:", ip)
	}
	// If the user specifies a 'format' querystring, we'll try to return the
	// user's IP address in the specified format.
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch format[0] {
		case "json":
			jsonStr, _ := json.Marshal(models.IPAddress{ip})
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprintf(w, string(jsonStr))
			return
		case "jsonp":
			// If the user specifies a 'callback' parameter, we'll use that as
			// the name of our JSONP callback.
			callback := "callback"
			if val, ok := r.Form["callback"]; ok && len(val) > 0 {
				callback = val[0]
			}
			jsonStr, _ := json.Marshal(models.IPAddress{ip})
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			fmt.Fprintf(w, callback+"("+string(jsonStr)+");")
			return
		case "xml":
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			xmlStr, _ := xml.MarshalIndent(models.IPAddress{ip}, " ", "  ")
			fmt.Fprintf(w, xml.Header+string(xmlStr))
			return
		}
	}

	// If no 'format' querystring was specified, we'll default to returning the
	// IP in plain text.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, ip)
}
