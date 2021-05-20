package main

import (
	"fmt"
	"github.com/iuuuuuaena/entity"
	"github.com/iuuuuuaena/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func init() {

}

func main() {
	//util.Get2("http://www.baidu.com")
	// 1. 访问排行榜 url
	entity.Url = "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"
	// 2. http请求获取 ranking List
	rankingList := getRankingList(entity.Url)
	// 3. 文件名
	entity.FileName = "README.md"
	// 4.清空文件
	util.EmptyFile(entity.FileName)
	// 6. 获取文件 channel，可以持续写入
	channel := util.OpenFileChannel(entity.FileName, os.O_APPEND|os.O_WRONLY, 0664)
	defer channel.Close()
	// 6.5 引入这个开源 项目
	//[![Anurag's GitHub stats](https://github-readme-stats.vercel.app/api?username=anuraghazra)](https://github.com/anuraghazra/github-readme-stats)
	theme := "buefy"
	username := "iuuuuuaena"
	channel.WriteString(
		`<div >
	<a style="float:left;width:55%;" href = "https://github.com/anuraghazra/github-readme-stats">
	 <img src = "https://github-readme-stats.vercel.app/api?username=` + username + `&theme=` + theme + `&show_icons=true"/>
	</a>
	<a  style="float:right;width:45%" href = "https://github.com/anuraghazra/github-readme-stats">
	 <img  src="https://github-readme-stats.vercel.app/api/top-langs/?username=anuraghazra&layout=compact"/>
	</a>
	</div>`)
	channel.WriteString("\n\n[![](https://img.shields.io/badge/jxd-@jxdgogogo.xyz-yellowgreen.svg)](https://www.jxdgogogo.xyz)<br>\n")
	// 写入 文件
	for i := 0; i < len(rankingList.Data.List); i++ {
		if i != 1 && i%9 == 1 {
			channel.WriteString(util.WriteSelectItem1(strconv.Itoa(i)+" ~ "+strconv.Itoa((i/10+1)*10)) + "\n\n")
		}
		channel.WriteString(strconv.Itoa(i+1) + ". " + rankingList.Data.List[i].Title + " [:link:](//www.bilibili.com/video/" + rankingList.Data.List[i].Bvid + ") <br>\n")
		if i != 0 && i != 9 && i%9 == 0 {
			channel.WriteString(util.WriteSelectItem2() + "\n")
		}
	}
	// 确保写入到磁盘
	err := channel.Sync()
	util.DropErr(err)
	// 登录 并签到

	loginAndSignIn()
}

// 请求获取 列表
func getRankingList(url string) *entity.ResObj {
	get := util.Get(entity.Url)
	// 2. 解析返回的json
	rankList := util.Unmarshal(get)
	return rankList
}

// 登录并签到
func loginAndSignIn() {
	// 1. 登录
	urls := "https://sijishe.wtf/wp-json/jwt-auth/v1/token"
	//这里添加post的body内容
	data := make(url.Values)
	data["nickname"] = []string{""}
	data["code"] = nil
	data["img_code"] = nil
	data["invitation_code"] = []string{""}
	data["token"] = []string{""}
	data["smsToken"] = []string{""}
	data["luoToken"] = []string{""}
	data["confirmPassword"] = []string{""}
	data["loginType"] = []string{""}
	data["username"] = []string{"1459555323@qq.com"}
	data["password"] = []string{"qwertyuiop"}
	form, err := http.PostForm(urls, data)
	if err != nil {
		// handle error
	}
	body, err := ioutil.ReadAll(form.Body)
	if err != nil {
		// handle error
	}
	unmarshal2 := util.Unmarshal2(string(body))
	// 2. 签到
	url2 := "https://sijishe.wtf/wp-json/b2/v1/userMission"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url2, strings.NewReader(""))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("authorization", "Bearer "+unmarshal2.Token)
	req.Header.Set("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	defer form.Body.Close()
	res2, err := client.Do(req)
	defer res2.Body.Close()
	all, err := ioutil.ReadAll(res2.Body)
	//fmt.Println("签到获得积分："+string(all))


	// 使用 QMSG 酱 发送qq消息 提示成功
	qUrl := "https://qmsg.zendee.cn/send/"
	token := "27d00079e5bc89df4228f0a4f90eba03"
	// 用 qUrl + token + get参数 msg 就可以发送消息了

	scores  := strings.Replace(string(all), "\"", "", -1)
	xx := qUrl + token + fmt.Sprintf("?msg=gotit")
	fmt.Println(xx)
	// 发请求通知qq成功，带上scores
	util.Get2(xx + scores)

}
