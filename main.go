package main

import (
    "fmt"
    
    "traffic/crossroads"
	"traffic/tahion"
	"traffic/trafficarchive"
    
    "git.zabbix.com/ap/plugin-support/plugin"
    "git.zabbix.com/ap/plugin-support/plugin/container"
)

const pluginName = "Traffic"

type Plugin struct {
    plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
    p.Infof("received request to handle %s key with %d parameters", key, len(params))
    if len(params) == 0 || params[0] == "" {
        return nil, fmt.Errorf("Invalid first parameter.")
    }
    
    if len(params) > 2 {
        return nil, fmt.Errorf("Too many parameters.")
    }
    
    switch key {
      
    case "crossroads":
        if len(params) > 1 {
            return crossroads.TwoRequest(params[0], params[1]), nil
        } else {
			if len(params) == 1 {
				return crossroads.Request(params[0]), nil
			}else{
				return nil, fmt.Errorf("Invalid parameter.")
			}
        }
		
	case "tahion":
        if len(params) > 1 {
            return tahion.TwoRequest(params[0], params[1]), nil
        } else {
			if len(params) == 1 {
				return tahion.Request(params[0]), nil
			}else{
				return nil, fmt.Errorf("Invalid parameter.")
			}
        }
		
	case "trafficarchive":
        if len(params) > 1 {
            return trafficarchive.TwoRequest(params[0], params[1]), nil
        } else {
			if len(params) == 1 {
				return trafficarchive.Request(params[0]), nil
			}else{
				return nil, fmt.Errorf("Invalid parameter.")
			}
        }
      
    default:
        return nil, plugin.UnsupportedMetricError
    }
}

func init() {
    plugin.RegisterMetrics(&impl, pluginName,
		"crossroads", "Status Crossroads.",
		"tahion", "Status Tahion.",
		"trafficarchive", "Status Traffic Archive.")
}

func main() {
    h, err := container.NewHandler(impl.Name())
    if err != nil {
        panic(fmt.Sprintf("failed to create plugin handler %s", err.Error()))
    }
    impl.Logger = &h
    
err = h.Execute()
    if err != nil {
        panic(fmt.Sprintf("failed to execute plugin handler %s", err.Error()))
    }
}