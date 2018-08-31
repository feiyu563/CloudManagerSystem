package controllers

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"net/url"
)

type KubeImagesController struct {
	BaseController
}

func (c *KubeImagesController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

type Project struct {
	Project_id int64 `json:"project_id"`
	Name string `json:"name"`
	Repo_count int64 `json:"repo_count"`
	Update_time string `json:"update_time"`
}

//Projects
func (c *KubeImagesController) Projects()  {
	Harbor_url := "http://"+beego.AppConfig.String("harbor::harbor_ip")+"/"

	lens,_:=strconv.Atoi(c.Input().Get("limit"))
	start,_:=strconv.Atoi(c.Input().Get("offset"))
	err,resp:=GetSwagger(Harbor_url+"api/projects")
	if err != nil {
		beego.Error(err)
	}
	project:=[]Project{}
	json.Unmarshal(resp,&project)
	//定义返回的数据结构
	result := make(map[string]interface{})
	lenp:=len(project)
	result["total"] = lenp
	if (start+lens)>lenp || lens>lenp  {
		result["rows"] = project[start:lenp]
	} else {
		result["rows"] = project[start:start+lens]
	}
	c.Data["json"] = result
	c.ServeJSON()
}
type Repositorie struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tags_count int64 `json:"tags_count"`
	Update_time string `json:"update_time"`
}
//Repositories
func (c *KubeImagesController) Repositories() {
	Harbor_url := "http://"+beego.AppConfig.String("harbor::harbor_ip")+"/"
	Id:=c.Input().Get("Id")
	lens,_:=strconv.Atoi(c.Input().Get("limit"))
	start,_:=strconv.Atoi(c.Input().Get("offset"))
	err,resp:=GetSwagger(Harbor_url+"api/repositories?project_id="+Id)
	if err != nil {
		beego.Error(err)
		return
	}
	repositorie:=[]Repositorie{}
	json.Unmarshal(resp,&repositorie)
	//定义返回的数据结构
	result := make(map[string]interface{})
	lenp:=len(repositorie)
	result["total"] = lenp
	if (start+lens)>lenp || lens>lenp  {
		result["rows"] = repositorie[start:lenp]
	} else {
		result["rows"] = repositorie[start:start+lens]
	}
	c.Data["json"] = result
	c.ServeJSON()
}

type Tag struct {
	Name string `json:"name"`
	Docker_version string `json:"docker_version"`
	Architecture string `json:"architecture"`
	Os string `json:"os"`
	Author string `json:"author"`
	Created string `json:"created"`
}
//Tags
func (c *KubeImagesController) Tags() {
	Harbor_url := "http://"+beego.AppConfig.String("harbor::harbor_ip")+"/"
	Name:=c.Input().Get("name")
	lens,_:=strconv.Atoi(c.Input().Get("limit"))
	start,_:=strconv.Atoi(c.Input().Get("offset"))
	err,resp:=GetSwagger(Harbor_url+"api/repositories/"+Name+"/tags")
	if err != nil {
		beego.Error(err)
		return
	}
	tag:=[]Tag{}
	json.Unmarshal(resp,&tag)
	//定义返回的数据结构
	result := make(map[string]interface{})
	lenp:=len(tag)
	result["total"] = lenp
	if (start+lens)>lenp || lens>lenp  {
		result["rows"] = tag[start:lenp]
	} else {
		result["rows"] = tag[start:start+lens]
	}
	c.Data["json"] = result
	c.ServeJSON()
}


func GetSwagger(url string) (error,[]byte) {
	SessionId,err:=GetSessionId()
	if err != nil {
		beego.Error(err)
		return err,nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(SessionId))
	//req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		beego.Error(err)
		return err,nil
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", "rem-username=admin; beegosessionID=3a7f697427caed8eb17f14368bb8c832")
	resp, err := client.Do(req)
	if err != nil {
		return err,nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err,nil
	}
	//fmt.Println(url,string(body))
	return nil,body
}

func GetSessionId() (string,error) {
	Harbor_url := "http://"+beego.AppConfig.String("harbor::harbor_ip")+"/"
	Harbor_login_url:=Harbor_url+"login"
	Harbor_username := beego.AppConfig.String("harbor::harbor_username")
	Harbor_password := beego.AppConfig.String("harbor::harbor_password")
	resp, err :=http.PostForm(Harbor_login_url,url.Values{"principal": {Harbor_username}, "password": {Harbor_password}})
	if err != nil {
		beego.Error(err)
		return "",err
	}
	defer resp.Body.Close()
	return "beegosessionID="+resp.Cookies()[0].Value,nil
}