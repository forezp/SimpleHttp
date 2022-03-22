package routers

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/forezp/SimpleHttp/app/api"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	testMetrics()

	r.GET("/metrics", func(c *gin.Context) {
		metrics.WritePrometheus(c.Writer, true)
	})
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", api.GetTags)
		//新建标签

	}

	return r
}

var (
	// Register counter without labels.
	requestsTotal = metrics.NewCounter("requests_total")

	// Register summary with a single label.
	requestDuration = metrics.NewSummary(`requests_duration_seconds{path="/foobar/baz"}`)

	// Register gauge with two labels.
	queueSize = metrics.NewGauge(`queue_size{queue="foobar",topic="baz"}`, func() float64 {
		return float64(123)
	})

	// Register histogram with a single label.
	responseSize = metrics.NewHistogram(`response_size{path="/foo/bar"}`)
)

func testMetrics() {
	// Increment requestTotal counter.
	requestsTotal.Inc()

	startTime := time.Now()
	//processRequest()
	time.Sleep(time.Millisecond * 200)
	// Update requestDuration summary.
	requestDuration.UpdateDuration(startTime)

	// Update responseSize histogram.
	responseSize.Update(float64(12332))
}
