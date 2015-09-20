package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jondot/armor"
	"time"
	"net/url"
	"net/http/httputil"
	"net/http"
)

type HeaderPokingTransport struct {
	AddHeaders map[interface{}]interface{}
	RemoveHeaders map[interface{}]interface{}

}

func (t *HeaderPokingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = "" // force using url host
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil{
		return resp, err
	}

	for k, _ := range(t.RemoveHeaders){
		resp.Header.Del(k.(string))
	}

	for k, v := range(t.AddHeaders){
		resp.Header.Add(k.(string), v.(string))
	}

	return resp, nil
}


func main() {
	m := armor.New("sticker", "1.0")
	r := m.GinRouter()

	backend := m.Config.GetString("backend")
	addHeaders := m.Config.Get("add_headers").(map[interface{}]interface{})
	removeHeaders := m.Config.Get("remove_headers").(map[interface{}]interface{})

	u, err := url.Parse(backend)
	if err != nil {
		m.Log.Fatalf("Cannot parse backend url %s, error: %s", u, err)
	}


	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = &HeaderPokingTransport{
		RemoveHeaders:removeHeaders,
		AddHeaders:addHeaders,
	}

	r.GET("/api/*any", func(c *gin.Context) {
		defer m.Metrics.Timed("timed.request", time.Now())
		m.Metrics.Inc("request")
		c.Request.URL.Path = c.Request.URL.Path[5:]
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	m.Run(r)
}
