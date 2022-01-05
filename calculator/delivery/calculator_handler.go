package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

type CalcDelivery struct {
	OutliersResult     entities.OutliersResult
	OutliersResultJson string
}

func CalculatorDelivery(a entities.OutliersResult, g *gin.Engine) {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	delivery := &CalcDelivery{OutliersResult: a, OutliersResultJson: string(b)}
	g.GET("/", delivery.attrrange3)
	g.GET("/attrrange3", delivery.attrrange3)
}

func (c *CalcDelivery) attrrange3(ctx *gin.Context) {
	jsonStr := c.OutliersResultJson
	ctx.DataFromReader(http.StatusOK, int64(len(jsonStr)), gin.MIMEJSON, strings.NewReader(jsonStr), nil)
}
