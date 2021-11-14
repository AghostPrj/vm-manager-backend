/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:10
 * @Desc:
 */

package requestUtils

import (
	"encoding/json"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func DecodeJsonRequestBody(c *gin.Context, payload interface{}) error {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.BaseResponse{
			Code: 500,
			Desc: "body error",
		})
		return err
	}

	err = json.Unmarshal(bodyBytes, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.BaseResponse{
			Code: 500,
			Desc: "body error",
		})
		return err
	}
	return nil
}
