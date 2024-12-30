package tahion

import (
	"net/http"
    "io/ioutil"
    "time"
    "strings"
    "encoding/json"
)

type ViewCameras struct {
    Ip string
    Status string
}

type Jsons struct {
    UpTime string
	NetworkNetspeed string
	NetworkSent string
	NetworkReceived string
    ViewCamera[] ViewCameras
}

func GetJson() (json string){
    http.DefaultClient.Timeout = 2 * time.Second
    resp, err := http.Get("http://192.168.4.20:8020/")
    if err != nil {
        json = "ERROR The requested host is not responding"
        return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        json = "ERROR Requested host no data"
        return
    }
    json = string(body)
    return
}

func UpTime() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.UpTime
        return
    }else{
        response = "ERROR"
        return
    }
}
func NetworkNetspeed() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.NetworkNetspeed
        return
    }else{
        response = "ERROR"
        return
    }
}
func NetworkSent() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.NetworkSent
        return
    }else{
        response = "ERROR"
        return
    }
}
func NetworkReceived() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.NetworkReceived
        return
    }else{
        response = "ERROR"
        return
    }
}

func ViewCamera(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.ViewCamera); i++ {
			if status.ViewCamera[i].Ip == message {
				response = status.ViewCamera[i].Status
				break
			}else{
				response = "ERROR IP Address"
			}
		}
        return
    }else{
        response = "ERROR"
        return
    }
}

func Request(key string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "uptime":
        return UpTime()
	case "networknetspeed":
        return NetworkNetspeed()
	case "networksent":
        return NetworkSent()
	case "networkreceived":
        return NetworkReceived()
    default:
      return "ERROR Unsupported Metric"
    }
}


func TwoRequest(key string, message string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "viewcamera":
        return ViewCamera(message)
    default:
      return "ERROR Unsupported Metric"
    }
}
