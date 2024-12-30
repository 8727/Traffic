package trafficarchive

import (
	"net/http"
    "io/ioutil"
    "time"
    "strings"
    "encoding/json"
)

type ViolationsType struct {
    Violations string
    Quantity string
}

type Jsons struct {
    UpTime string
    DiskTotalSize string
    DiskTotalFreeSpace string
    DiskPercentTotalSize string
    DiskPercentTotalFreeSpace string
	NetworkNetspeed string
	NetworkSent string
	NetworkReceived string
    ArchiveDepthSeconds string
    ArchiveDepthCount string
	LastReplicationSeconds string
	UnprocessedViolationsCount string
	UnprocessedViolationsSeconds string
	UnexportedCount string
	UnexportedSeconds string
	AllViolationsPrevioushour string
	ViolationsYesterday[] ViolationsType
}

func GetJson() (json string){
    http.DefaultClient.Timeout = 2 * time.Second
    resp, err := http.Get("http://192.168.4.30:8030/")
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
func LastReplicationSeconds() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.LastReplicationSeconds
        return
    }else{
        response = "ERROR"
        return
    }
}
func UnprocessedViolationsCount() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.UnprocessedViolationsCount
        return
    }else{
        response = "ERROR"
        return
    }
}
func UnprocessedViolationsSeconds() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.UnprocessedViolationsSeconds
        return
    }else{
        response = "ERROR"
        return
    }
}
func UnexportedCount() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.UnexportedCount
        return
    }else{
        response = "ERROR"
        return
    }
}
func UnexportedSeconds() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.UnexportedSeconds
        return
    }else{
        response = "ERROR"
        return
    }
}
func AllViolationsPrevioushour() (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
        response = status.AllViolationsPrevioushour
        return
    }else{
        response = "ERROR"
        return
    }
}

func ViolationsYesterday(message string) (response string){
    data := GetJson()
    var status Jsons
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &status)
		for i := 0; i < len(status.ViolationsYesterday); i++ {
			if status.ViolationsYesterday[i].Violations == message {
				response = status.ViolationsYesterday[i].Quantity
				break
			}else{
				response = "ERROR no violation type"
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
	case "networknetspeed":
        return NetworkNetspeed()
	case "networksent":
        return NetworkSent()
	case "networkreceived":
        return NetworkReceived()
	case "archivedepthseconds":
        return ArchiveDepthSeconds()
	case "archivedepthcount":
        return ArchiveDepthCount()
	case "lastreplicationseconds":
        return LastReplicationSeconds()
	case "unprocessedviolationscount":
        return UnprocessedViolationsCount()
	case "unprocessedviolationsseconds":
        return UnprocessedViolationsSeconds()
	case "unexportedcount":
        return UnexportedCount()
	case "unexportedseconds":
        return UnexportedSeconds()
	case "allviolationsprevioushour":
        return AllViolationsPrevioushour()
    default:
      return "ERROR Unsupported Metric"
    }
}


func TwoRequest(key string, message string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "violationsyesterday":
        return ViolationsYesterday(message)
    default:
      return "ERROR Unsupported Metric"
    }
}
