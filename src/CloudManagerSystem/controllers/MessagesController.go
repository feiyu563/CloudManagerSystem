package controllers

import (
	"encoding/json"
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/astaxie/beego"
)

type MessagesController struct {
	BaseController
}

func (c *MessagesController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

//日志管理
func (c *MessagesController) Logs() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="日志管理"
	//页面模板设置
	c.setTpl("messages/logs.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "messages/logs_footerjs.html"
	c.LayoutSections["headcssjs"] = "messages/logs_headcssjs.html"
}
//定义前端传来的query结构
type TableQuery struct {
	StartTime string
	EndTime string
	PodName string
	QuerySting string
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Order string `json:"order"`
}
//定义接收es结构
type EsSting struct {
	Scroll_id string `json:"_scroll_id"`
	Hits Hits `json:"hits"`
}
type Hits struct {
	Total int64 `json:"total"`
	Hits HitsH `json:"hits"`
}
type HitsH []Message
type Message struct {
	Source Source `json:"_source"`
}
type Source struct {
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Level int64 `json:"level"`
	Tag string `json:"tag"`
}

//日志检索
func (c *MessagesController) GetLogs() {
	//http://es.zxbike.top/_search?scroll=60m
	tablequery:=TableQuery{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &tablequery)
	//定义返回的数据结构
	result := make(map[string]interface{})
	if tablequery.PodName=="" {
		result["err"]="No POD Name"
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		if tablequery.Offset == 0 {
			queryString:="(tag:*"+tablequery.PodName+"*)"
			if tablequery.QuerySting!="" {
				queryString=queryString+" AND ("+tablequery.QuerySting+")"
			}
			beego.Error("\n\n\n\n")
			beego.Error(queryString)
			beego.Error("\n\n\n\n")
			jsonstring:=`{
			  "from": 0,
			  "size": 150,
			  "query": {
				"bool": {
				  "must": {
					"query_string": {
					  "query": "`+queryString+`",
					  "allow_leading_wildcard": true
					}
				  },
				  "filter": {
					"bool": {
					  "must": {
						"range": {
						  "timestamp": {
							"from": "`+tablequery.StartTime+`.000",
							"to": "`+tablequery.EndTime+`.000",
							"include_lower": true,
							"include_upper": true
						  }
						}
					  }
					}
				  }
				}
			  },
			  "sort": [
				{
				  "timestamp": {
					"order": "`+tablequery.Order+`"
				  }
				}
			  ],
			  "highlight": {
				"fragment_size": 0,
				"number_of_fragments": 0,
				"require_field_match": false,
				"fields": {
				  "*": {}
				}
			  }
			}`
			//开始向es读取数据
			url := "http://es.zxbike.top/_search?scroll=60m"
			reader := bytes.NewReader([]byte(jsonstring))
			err,respBytes:=Post(url,reader)
			if err != nil {
				result["err"] = err.Error()
				c.Data["json"] = result
				c.ServeJSON()
				return
			}
			essting:=EsSting{}
			json.Unmarshal(respBytes,&essting)
			c.SetSession("Scroll_id",essting.Scroll_id)
			result["total"]=essting.Hits.Total
			result["rows"] = essting.Hits.Hits
			c.Data["json"] = result
			c.ServeJSON()
		} else {
			url:="http://es.zxbike.top/_search/scroll"
			scrollid:=c.GetSession("Scroll_id")
			fmt.Println(scrollid)
			jsonstring:=`{"scroll" : "60m", "scroll_id" : "`+scrollid.(string)+`"}`
			reader:=bytes.NewReader([]byte(jsonstring))
			err,respBytes:=Post(url,reader)
			if err != nil {
				result["err"] = err.Error()
				c.Data["json"] = result
				c.ServeJSON()
				return
			}
			essting:=EsSting{}
			json.Unmarshal(respBytes,&essting)
			scrollid=essting.Scroll_id
			result["total"]=essting.Hits.Total
			result["rows"] = essting.Hits.Hits
			c.Data["json"] = result
			c.ServeJSON()
		}
	}
}

func Post(url string,reader *bytes.Reader)(error,[]byte)  {
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return err,nil
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err,nil
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err,nil
	}
	return nil,respBytes
}
