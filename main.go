package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/marcellribeiro/odoscope_attrrange3/calculator/delivery"
	"github.com/marcellribeiro/odoscope_attrrange3/calculator/usecase"
	"github.com/marcellribeiro/odoscope_attrrange3/entities"
	"github.com/marcellribeiro/odoscope_attrrange3/repository"
)

func main() {
	g := gin.Default()

	if len(os.Args) < 2 {
		r, err := repository.MakeSampleData()
		if err != nil {
			fmt.Println(err)
		}
		process(r, g)
	} else {
		r, err := repository.GetDataFromFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		process(r, g)
	}

	g.Run()
}

func process(r entities.Datasets, g *gin.Engine) {
	calculator := usecase.NewCalculatorUsecase(r)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	service, err := calculator.CalculatorService(ctx)
	if err != nil {
		fmt.Println(err)
	}

	delivery.CalculatorDelivery(service, g)
}
