package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationData struct {
	ActivateDate string `json:"activateDate"`
	Battery      string `json:"battery"`
	CommandCode  string `json:"commandCode"`
	ExpireDate   string `json:"expireDate"`
	GpsTime      string `json:"gpsTime"`
	Imei         string `json:"imei"`
	Imsi         string `json:"imsi"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Model        string `json:"model"`
	PosType      string `json:"posType"`
	Reqid        string `json:"reqid"`
	SignalTime   string `json:"signalTime"`
	Speed        string `json:"speed"`
	Status       string `json:"status"`
}

type HealthData struct {
	BloodOxygen      string `json:"bloodOxygen"`
	BloodOxygenTime  string `json:"bloodOxygenTime"`
	BloodPressureMax string `json:"bloodPressureMax"`
	BloodPressureMin string `json:"bloodPressureMin"`
	BpTime           string `json:"bpTime"`
	Calorie          string `json:"calorie"`
	CommandCode      string `json:"commandCode"`
	DeepSleep        string `json:"deepSleep"`
	Distance         string `json:"distance"`
	HeartRate        int    `json:"heartRate"`
	HrTime           string `json:"hrTime"`
	Imei             string `json:"imei"`
	LightSleep       string `json:"lightSleep"`
	Reqid            string `json:"reqid"`
	SleepTime        string `json:"sleepTime"`
	Steps            string `json:"steps"`
	TotalSleep       string `json:"totalSleep"`
}

type StepRollData struct {
	Step     string `json:"step"`
	Roll     string `json:"roll"`
	DataTime string `json:"dataTime"`
	Imei     string `json:"imei"`
}

type AlarmData struct {
	ExceptionID      int      `json:"ExceptionID"`
	SerialNumber     string   `json:"SerialNumber"`
	GeoFenceID       *int     `json:"GeoFenceID"`
	GeoFenceName     *string  `json:"GeoFenceName"`
	NotificationType int      `json:"NotificationType"`
	Message          *string  `json:"Message"`
	Created          string   `json:"Created"`
	Deleted          *bool    `json:"Deleted"`
	ClearDate        *string  `json:"ClearDate"`
	ClearBy          *int     `json:"ClearBy"`
	Note             *string  `json:"Note"`
	Lat              *float64 `json:"Lat"`
	Lng              *float64 `json:"Lng"`
	Power            *int     `json:"Power"`
	Address          *string  `json:"Address"`
	AccON            *string  `json:"AccON"`
	EngineON         *string  `json:"EngineON"`
	OtherStatus      *string  `json:"OtherStatus"`
	GSM              *string  `json:"GSM"`
	DeviceUTCTime    string   `json:"DeviceUTCTime"`
	OLat             *float64 `json:"OLat"`
	OLng             *float64 `json:"OLng"`
	BaiduLat         *float64 `json:"BaiduLat"`
	BaiduLng         *float64 `json:"BaiduLng"`
	SerialNo         *string  `json:"SerialNo"`
	DataType         *string  `json:"DataType"`
}

type TemperatureData struct {
	Imei            string   `json:"imei"`
	Temperature     *float64 `json:"temperature"`
	TemperatureTime *string  `json:"temperatureTime"`
}

type Request struct {
	DataType   string      `json:"DataType"`
	ResultData interface{} `json:"ResultData"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.Any("/*path", func(c *gin.Context) {
			// 打印请求头
			fmt.Println("Headers:")
			for k, v := range c.Request.Header {
				fmt.Printf("%s: %s\n", k, v)
			}

			// 读取请求体
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid request"})
				return
			}
			// 打印请求体
			fmt.Println("Request Body:", string(bodyBytes))

			// 将请求体重新设置到请求中，以便后续处理器可以读取
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			var req Request
			if err := json.Unmarshal(bodyBytes, &req); err != nil {
				c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid request"})
				return
			}

			switch req.DataType {
			case "Location":
				handleLocation(c, req.ResultData)
			case "Health":
				handleHealth(c, req.ResultData)
			case "StepRoll":
				handleStepRoll(c, req.ResultData)
			case "alarm":
				handleAlarm(c, req.ResultData)
			case "Temperature":
				handleTemperature(c, req.ResultData)
			default:
				c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Unknown DataType"})
			}
		})
	}

	r.Run("0.0.0.0:8011") // 启动服务器并监听在端口8011
}

func handleLocation(c *gin.Context, data interface{}) {
	var locationData LocationData
	if err := mapToStruct(data, &locationData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid Location data"})
		return
	}
	// 处理locationData逻辑
	c.JSON(http.StatusOK, Response{Code: 1, Msg: "操作成功"})
}

func handleHealth(c *gin.Context, data interface{}) {
	var healthData HealthData
	if err := mapToStruct(data, &healthData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid Health data"})
		return
	}
	// 处理healthData逻辑
	c.JSON(http.StatusOK, Response{Code: 1, Msg: "操作成功"})
}

func handleStepRoll(c *gin.Context, data interface{}) {
	var stepRollData StepRollData
	if err := mapToStruct(data, &stepRollData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid StepRoll data"})
		return
	}
	// 处理stepRollData逻辑
	c.JSON(http.StatusOK, Response{Code: 1, Msg: "操作成功"})
}

func handleAlarm(c *gin.Context, data interface{}) {
	var alarmData AlarmData
	if err := mapToStruct(data, &alarmData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid Alarm data"})
		return
	}
	// 处理alarmData逻辑
	c.JSON(http.StatusOK, Response{Code: 1, Msg: "操作成功"})
}

func handleTemperature(c *gin.Context, data interface{}) {
	var temperatureData TemperatureData
	if err := mapToStruct(data, &temperatureData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 0, Msg: "Invalid Temperature data"})
		return
	}
	// 处理temperatureData逻辑
	c.JSON(http.StatusOK, Response{Code: 1, Msg: "操作成功"})
}

func mapToStruct(data interface{}, out interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}
