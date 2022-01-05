package delivery

import (
	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CalculatorHandler struct {
	CalcUsecase entities.CalculatorUsecase
}

// func NewCalculatorHandler(engine *gin.Engine, us entities.CalculatorUsecase) {
// 	handler := &CalculatorHandler{
// 		CalcUsecase: us,
// 	}
// 	engine.GET("/attrrange3", handler.ThreeSigmasMethod)
// }

// func (a *CalculatorHandler) ThreeSigmasMethod(c gin.Context) {
// 	data := ""
// 	result, err := a.CalcUsecase.ThreeSigmasMethod(data)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
// 	}

// 	c.JSON(http.StatusOK, result)
// }
