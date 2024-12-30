package crossroads

import (
	"net/http"
    "io/ioutil"
    "time"
    "strings"
    "encoding/json"
)

type CameraAndOverviewImages struct {
    Ip string
    Quantity string
}

type ViewCameras struct {
    Ip string
    Status string
}

type Jsons struct {
    UpTime string
    DiskTotalSize string
    DiskTotalFreeSpace string
    DiskPercentTotalSize string
    DiskPercentTotalFreeSpace string
    ArchiveDepthSeconds string
    ArchiveDepthCount string
    ArchiveNumberOfCarsOfTheFuture string
	ArchiveNumberOfCarsOfThePast string
	NetworkNetspeed string
	NetworkSent string
	NetworkReceived string
	TrafficLight string
	RecognizingCamera[] CameraAndOverviewImages
	TrafficLightStatus[] ViewCameras
	TimeAccuracy[] CameraAndOverviewImages
	PercentageRedZone[] CameraAndOverviewImages
	ViewCamera[] ViewCameras
	NumberOfOverviewImages[] CameraAndOverviewImages
}

func GetJson() (json string){
    http.DefaultClient.Timeout = 2 * time.Second
    resp, err := http.Get("http://192.168.4.100:8010/")
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
func DiskTotalSize() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.DiskTotalSize
        return
    }else{
        response = "ERROR"
        return
    }
}
func DiskTotalFreeSpace() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.DiskTotalFreeSpace
        return
    }else{
        response = "ERROR"
        return
    }
}
func DiskPercentTotalSize() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.DiskPercentTotalSize
        return
    }else{
        response = "ERROR"
        return
    }
}
func DiskPercentTotalFreeSpace() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.DiskPercentTotalFreeSpace
        return
    }else{
        response = "ERROR"
        return
    }
}
func ArchiveDepthSeconds() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.ArchiveDepthSeconds
        return
    }else{
        response = "ERROR"
        return
    }
}
func ArchiveDepthCount() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.ArchiveDepthCount
        return
    }else{
        response = "ERROR"
        return
    }
}
func ArchiveNumberOfCarsOfTheFuture() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.ArchiveNumberOfCarsOfTheFuture
        return
    }else{
        response = "ERROR"
        return
    }
}
func ArchiveNumberOfCarsOfThePast() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.ArchiveNumberOfCarsOfThePast
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
func TrafficLight() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.TrafficLight
        return
    }else{
        response = "ERROR"
        return
    }
}

func RecognizingCamera(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.RecognizingCamera); i++ {
			if status.RecognizingCamera[i].Ip == message {
				response = status.RecognizingCamera[i].Quantity
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

func TrafficLightStatus(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.TrafficLightStatus); i++ {
			if status.TrafficLightStatus[i].Ip == message {
				response = status.TrafficLightStatus[i].Status
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

func TimeAccuracy(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.TimeAccuracy); i++ {
			if status.TimeAccuracy[i].Ip == message {
				response = status.TimeAccuracy[i].Quantity
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

func PercentageRedZone(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.PercentageRedZone); i++ {
			if status.PercentageRedZone[i].Ip == message {
				response = status.PercentageRedZone[i].Quantity
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

func NumberOfOverviewImages(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.NumberOfOverviewImages); i++ {
			if status.NumberOfOverviewImages[i].Ip == message {
				response = status.NumberOfOverviewImages[i].Quantity
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
    case "disktotalsize":
        return DiskTotalSize()
    case "disktotalfreespace":
        return DiskTotalFreeSpace()
	case "diskpercenttotalsize":
        return DiskPercentTotalSize()
	case "diskpercenttotalfreespace":
        return DiskPercentTotalFreeSpace()
	case "archivedepthseconds":
        return ArchiveDepthSeconds()
	case "archivedepthcount":
        return ArchiveDepthCount()
	case "archivenumberofcarsofthefuture":
        return ArchiveNumberOfCarsOfTheFuture()
	case "archivenumberofcarsofthepast":
        return ArchiveNumberOfCarsOfThePast()
	case "networknetspeed":
        return NetworkNetspeed()
	case "networksent":
        return NetworkSent()
	case "networkreceived":
        return NetworkReceived()
	case "trafficlight":
        return TrafficLight()
    default:
      return "ERROR Unsupported Metric"
    }
}

func TwoRequest(key string, message string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "recognizingcamera":
        return RecognizingCamera(message)
	case "trafficlightstatus":
        return TrafficLightStatus(message)	
	case "timeaccuracy":
        return TimeAccuracy(message)	
	case "percentageredzone":
        return PercentageRedZone(message)
    case "viewcamera":
        return ViewCamera(message)
	case "numberofoverviewimages":
        return NumberOfOverviewImages(message)
    default:
      return "ERROR Unsupported Metric"
    }
}