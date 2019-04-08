package app

import (
	"net/http"

	"../core"
	"../data"
	"github.com/gin-gonic/gin"
)

func PostComputeDistribution(c *gin.Context) {
	var request data.ComputeDistributionRequest
	c.BindJSON(&request)
	northWestSolution := core.NorthWest(request.Providers, request.Consumers)
	result := core.Potencial(request.Costs, northWestSolution, 0)
	c.JSON(
		http.StatusOK, gin.H{
			"optimal": result,
		})
}

// GetInstructions curl -i http://localhost:8080/api/v1/test
func GetTest(c *gin.Context) {

	c.JSON(200, gin.H{"ok": "Hello world"})

}
