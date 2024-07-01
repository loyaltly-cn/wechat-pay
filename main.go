package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-pay/io"
	"wechat-pay/pojo"
	"wechat-pay/sdk"
)

func payment(c *gin.Context) {

	b, _ := c.GetRawData()

	body := pojo.Pay{}
	_ = json.Unmarshal(b, &body)
	res := sdk.Payment(body)
	c.JSONP(http.StatusOK, res)

}

func status(c *gin.Context) {
	b, _ := c.GetRawData()

	status := pojo.Status{}

	_ = json.Unmarshal(b, &status)
	res := sdk.Status(status.OrderId)

	c.JSONP(http.StatusOK, res)
}

func _close(c *gin.Context) {
	b, _ := c.GetRawData()

	status := pojo.Status{}

	_ = json.Unmarshal(b, &status)
	res := sdk.Close(status.OrderId)

	c.JSONP(http.StatusOK, res)
}
func refund(c *gin.Context) {
	b, _ := c.GetRawData()
	body := pojo.Refund{}
	_ = json.Unmarshal(b, &body)
	res := sdk.Refund(body)
	c.JSONP(http.StatusOK, res)
}

func main() {
	r := gin.Default()
	r.POST("/payment", payment) //下单
	r.POST("/status", status)   //查询订单状态
	r.POST("/close", _close)    //关闭订单
	r.POST("/refund", refund)   //退款

	conf, _ := io.ReadFile()
	port := ":" + conf["port"].(string)
	err := r.Run(port)
	if err != nil {
		return
	}
}
