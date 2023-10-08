package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Servive interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Servive
}

func New(s Servive) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()

	s := fmt.Sprintf("Quantity day: %d" ,d)
 	err :=	ctx.String(http.StatusOK, s)

	if err != nil {
		return err
	}

	return nil
}