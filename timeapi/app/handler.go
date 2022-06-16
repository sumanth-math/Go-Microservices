package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(c *gin.Context) {
	timeZones := c.Query("tz")
	timeZoneResponse := make(map[string]string)
	if timeZones == "" {
		timeZoneResponse["current_time"] = time.Now().UTC().String()
	} else {
		tzs := strings.Split(timeZones, ",")
		for _, tz := range tzs {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				c.JSON(http.StatusNotFound,
					fmt.Sprintf("invalid timezone %s", tz))
				return
			}
			timeZoneResponse[tz] = time.Now().In(loc).String()
		}
	}

	c.JSON(http.StatusOK, timeZoneResponse)
}
