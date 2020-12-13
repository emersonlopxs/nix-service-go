package handler

import (
	"net/http"
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
)

func LaunchesGet() gin.HandlerFunc {
	return func (c *gin.Context)  {

		resp, err := http.Get("https://api.spacexdata.com/v4/launches/past")

		if err != nil {
      log.Fatalln(err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)


		c.JSON(http.StatusOK, map[string]string{
			"data": sb,
		})
	}
}