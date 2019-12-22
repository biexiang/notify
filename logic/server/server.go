package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	Duration = 1
	NotifyTypeInterval = 1
	NotifyTypeTimePoint = 2
)

var (
	NotifyList []*Notify
	ConfigPath string
)

type Message struct {
	Title string
	Content string
}

type Notify struct {
	Message Message
	Type int64
	NotifyInterval int64
	NotifyTimePoint int64
	LastNotifyTimePoint int64
	IsClosed bool
}

func (n *Notify) canNotify() (bNotify bool) {
	nowTimesStamp := time.Now().Unix()
	switch n.Type {
	case NotifyTypeInterval:
		if nowTimesStamp > n.LastNotifyTimePoint + n.NotifyInterval {
			bNotify = true
		}
	case NotifyTypeTimePoint:
		if nowTimesStamp > n.LastNotifyTimePoint {
			bNotify = true
		}
	default:
		bNotify = false
	}
	return
}

/**
	Just Support Mac For Now
 */
func (n *Notify) doNotify() error {
	if runtime.GOOS == "darwin" {
		sNotify := fmt.Sprintf(`display notification "%s" with title "%s"`,n.Message.Title,n.Message.Content)
		cmd := exec.Command("/usr/bin/osascript","-e",sNotify)
		err := cmd.Run()
		if err != nil {
			return err
		}
		if n.Type == NotifyTypeTimePoint {
			n.IsClosed = true
		}
	}
	return nil
}


func Init(configPath string) error {
	ConfigPath = configPath
	bConfig,err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println("configPath is not valid：",err)
		return err
	}
	if len(bConfig) == 0 {
		return nil
	}
	err = yaml.Unmarshal(bConfig,NotifyList)
	if err != nil {
		log.Println("configFile is not valid：",err)
		return err
	}
	return nil
}

func Start() {
	t := time.NewTicker(time.Duration(time.Second * Duration))
	defer func() {
		bConfig,err := yaml.Marshal(NotifyList)
		if err != nil {
			log.Println("marshalConfig Fail：",err)
		}
		err = ioutil.WriteFile(ConfigPath,bConfig,os.ModeAppend)
		if err != nil {
			log.Println("marshalConfig Fail：",err)
		}
		t.Stop()
	}()
	for {
		<- t.C
		for _,notifyItem := range NotifyList {
			if !notifyItem.IsClosed && notifyItem.canNotify() {
				if err := notifyItem.doNotify(); err != nil {
					log.Println("Notify Fail：",err)
				}
			}
		}
	}
}














