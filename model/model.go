package model

import (
	"cli/util"
	"encoding/json"
	"io/ioutil"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

func Conn(arg string) *gjson.Json {
	url := util.BASE_URL + util.ObjectList[arg] + "/detail"

	rsp, err := g.Client().Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	gs, err := gjson.DecodeToJson(body)
	if err != nil {
		panic(err)
	}
	return gs
}

func GetStat(args []string, action string) (data []util.Rsp) {
	for _, v := range args {
		e := util.Rsp{}
		gs := Conn(v)
		err := json.Unmarshal([]byte(gs.Export()), &e)
		if err != nil {
			panic(err)
		}
		data = append(data, e)
	}
	return
}

func Do(args []string, action string, file string) interface{} {
	return nil
}
