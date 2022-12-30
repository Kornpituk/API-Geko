package gecko

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) Gecko(c echo.Context) error {



	res, err := s.gecko.Gecko()
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)


}
