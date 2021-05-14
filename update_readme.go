package main

import (
	"github.com/iuuuuuaena/entity"
	"github.com/iuuuuuaena/util"
	"os"
	"strconv"
)

func init() {

}

func main() {
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
	// 写入 文件
	for i := 0; i < len(rankingList.Data.List); i++ {
		if i != 1 && i%9 == 1 {
			channel.WriteString(util.WriteSelectItem1(strconv.Itoa(i)+" ~ "+strconv.Itoa((i / 10+1)*10)) + "\n")
		}
		channel.WriteString(strconv.Itoa(i + 1)   + ". " + rankingList.Data.List[i].Title + " [:link:](//www.bilibili.com/video/" + rankingList.Data.List[i].Bvid + ") <br>\n")
		if i != 0 && i != 9 && i%9 == 0 {
			channel.WriteString(util.WriteSelectItem2() + "\n")
		}
	}
	// 确保写入到磁盘
	err := channel.Sync()
	util.DropErr(err)
}

// 请求获取 列表
func getRankingList(url string) *entity.ResObj {
	get := util.Get(entity.Url)
	// 2. 解析返回的json
	rankList := util.Unmarshal(get)
	return rankList
}
