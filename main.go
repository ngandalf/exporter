package main
 
import (
    "flag"
    "log"
    "net/http"
	
    "math/rand"
    "time"
	"github.com/ngandalf/exporter/collector/foo"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)
 
var (
    listenAddress = flag.String("web.listen-address", ":9888", "Address to listen on for web interface.")
    metricPath    = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
)
 
func main() {
    log.Fatal(serverMetrics(*listenAddress, *metricPath))
}
 
func serverMetrics(listenAddress, metricsPath string) error {
    foo := newFooCollector()
    prometheus.MustRegister(foo)

    http.Handle(metricsPath, promhttp.Handler())
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <html>
            <head><title>Exporter Metrics</title></head>
            <body>
            <h1>ConfigMap Reload</h1>
            <p><a href='` + metricsPath + `'>Metrics</a></p>
            </body>
            </html>
        `))
    })
    return http.ListenAndServe(listenAddress, nil)
}
