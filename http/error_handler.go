package http

import (
	"net/http"

	errorCommon "github.com/aryahmph/skripsi-common/error"
	"github.com/labstack/echo/v4"
)

func (h HTTPServer) errorHandler(err error, c echo.Context) {
	if he, ok := err.(*errorCommon.ClientError); ok {
		err = &echo.HTTPError{
			Code: he.StatusCode,
			Message: map[string]interface{}{
				"message": he.Message,
			},
		}
	} else if _, ok := err.(*echo.HTTPError); !ok {
		err = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusInternalServerError,
		}
	}

	h.E.DefaultHTTPErrorHandler(err, c)
}
