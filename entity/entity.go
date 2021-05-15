package entity

// 这三个对象是 bilibili返回的对象
type ItemList struct {
	Aid       int64  `json:"aid"`
	Bvid      string `json:"bvid"`
	Dynamic   string `json:"dynamic"`
	ShortLink string `json:"short_link"`
	Title     string `json:"title"`
	Tname     string `json:"tname"`
}
type DataList struct {
	List []ItemList `json:"list"`
	Note string     `json:"note"`
}
type ResObj struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Ttl     int64    `json:"ttl"`
	Data    DataList `json:"data"`
}

type Resp struct {
	Token           string `json:"token"`
	UserDisplayName string `json:"user_display_name"`
	Name            string `json:"name"`
	Id              string `json:"id"`
	Avatar          string `json:"avatar"`
	UserLink        string `json:"user_link"`
	IsAdmin         bool   `json:"is_admin"`
	Verify          string `json:"verify"`
	UserCode        string `json:"user_code"`
}

// 访问 的 url
var Url string = ""

// READEME.md 文件名
var FileName = ""
