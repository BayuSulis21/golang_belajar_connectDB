package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service/config"
	"service/database"
	"service/models"
	"service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

/*
noted:
jika get value dalam request pakai
1. c.ShouldBindJSON (jika untuk request = json )
2. c.Bind (jika untuk request = json & form-data )
3. c.param (jika untuk request = dari param  --> contoh pulsa/product/XL)
3. c.query (jika untuk request = dari query --> contoh pulsa?product=XL )
*/

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Println("[Info:] status:", status, "|time:", latency, "|method:", c.Request.Method, "|path:", c.Request.URL.Path)
		log.Println("[Header:] ", c.Request.Header)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("[ErrorGetBody:] ", err)
		}
		log.Println("[RequestBody:] ", string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		//c.Next()
		log.Println("[ResponseBody:] ", c.Request.Response)

	}
}

// router
func main() {
	//koneksi db
	conf := config.Config
	dbConn = database.InitDB(conf.Db)

	router := gin.Default()

	router.Use(Logger())
	router.GET("/test", RootHandler)
	router.GET("/product/:billerType/:product", GetProductHandler)
	router.GET("/queryString", QueryStringHandler)
	router.POST("/login", LoginHandler)
	router.POST("/denomList", PostdenomListHandler)

	router.Run(":8888")
}

// rooter default
func RootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "Bayu Sulistyo",
		"bio":  "Software Engineer",
	})
}

// rooter with param
func GetProductHandler(c *gin.Context) {
	billerType := c.Param("billerType")
	product := c.Param("product")
	c.JSON(http.StatusOK, gin.H{"billerType": billerType, "product": product})
}

// rooter with param pakai tanda tanya (?)
func QueryStringHandler(c *gin.Context) {
	billType := c.Query("billerType")
	c.JSON(http.StatusOK, gin.H{"billerType": billType})
}

func LoginHandler(c *gin.Context) {
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "bayu" || json.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "user": json.User})
}

func PostdenomListHandler(c *gin.Context) {
	var denom models.DenomList
	//var errorResp data.ErrorResponse

	/*
		ini digunakan jika error default dr go
		if err := c.ShouldBindJSON(&denom); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
	*/

	//Validasi request mandatory
	c.ShouldBindJSON(&denom)
	if denom.BilleryType == "" {
		c.JSON(utils.ResponseError(400, "Field Type is mandatory"))
		return
	}

	appsID := c.Request.Header.Get("appsID")
	var stringQuery string = fmt.Sprintf(`SELECT aggregatorProduct.aggregatorProductId ,application.appName ,productName,denom,billerTypeName,aggregatorName,aggregatorProductCode,applicationProductPrice,
	additionalProductInfo,aggregatorPrice,aggregatorDenomAdditionalPrice,aggregatorProduct.statusBiller, aggregatorProduct.statusProduct
	FROM applicationProduct
		INNER JOIN aggregatorProduct ON aggregatorProduct.aggregatorProductId = applicationProduct.aggregatorProductId
		LEFT JOIN application ON application.appId =applicationProduct.appId  AND application.appName = %s
		JOIN aggregatorBillerType ON aggregatorProduct.aggregatorBillerTypeId = aggregatorBillerType.aggregatorBillerTypeId 
		JOIN aggregator ON aggregator.aggregatorId = aggregatorBillerType.aggregatorId 
		JOIN billerType ON billerType.billerTypeId = aggregatorBillerType.billerTypeId 
		JOIN product ON aggregatorProduct.productId = product.productId 
		WHERE  aggregatorProduct.statusBiller = 'ACTIVE' AND aggregatorProduct.statusProduct = 'ACTIVE' AND billerTypeName = %s
	ORDER BY product.denom ASC, product.productName ASC, applicationProduct.applicationProductPrice ASC LIMIT 3`, "'"+appsID+"'", "'"+denom.BilleryType+"'")
	var biller []models.Biller
	errDB := dbConn.Raw(stringQuery).Scan(&biller)
	fmt.Println(errDB.Error)
	/*
		String := map[string]interface{}{
			"appsID":   appsID,
			"status":   "success",
			"listData": biller,
		}
	*/
	c.JSON(http.StatusOK, models.DenomListresponse{
		AppsID:          appsID,
		BillerType:      denom.BilleryType,
		ResponseCode:    "200",
		ResponseMessage: "success",
		ListData:        biller,
	})
}
