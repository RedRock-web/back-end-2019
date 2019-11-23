package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
)

func main() {
	var s student

	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		xh := c.Query("xh")
		url := "http://jwzx.node3.cqupt.co/data/json_StudentSearch.php?searchKey=" + xh
		//c.String(http.StatusOK, url)
		resp, _ := http.Get(url)
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &s)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v\n", s)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data": gin.H{
				"studentId": s.ReturnData[0].Xh,
				"name":      s.ReturnData[0].Xm,
				"gender":    s.ReturnData[0].Xb,
				"classId":   s.ReturnData[0].Bj,
				"major":     s.ReturnData[0].Zym,
				"college":   s.ReturnData[0].Yxm,
			},
		})

	})
	r.Run()

}

type student struct {
	Code       int    `json:"code"`
	Info       string `json:"info"`
	ReturnData []struct {
		Xh    string      `json:"xh"`
		Xm    string      `json:"xm"`
		XmEn  interface{} `json:"xmEn"`
		Xb    string      `json:"xb"`
		Bj    string      `json:"bj"`
		Zyh   string      `json:"zyh"`
		Zym   string      `json:"zym"`
		Yxh   string      `json:"yxh"`
		Yxm   string      `json:"yxm"`
		Nj    string      `json:"nj"`
		Csrq  string      `json:"csrq"`
		Xjzt  string      `json:"xjzt"`
		Rxrq  string      `json:"rxrq"`
		Yxmen string      `json:"yxmen"`
		ZymEn string      `json:"zymEn"`
		Xz    int         `json:"xz"`
		Mz    string      `json:"mz"`
	} `json:"returnData"`
}
