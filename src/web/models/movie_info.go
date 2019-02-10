package models

import (
	"regexp"

	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int
	Movie_name           string
	Movie_director       string
	Movie_writer         string
	Movie_conutry        string
	Movie_language       string
	Movie_type           string
	Movie_image          string
	Movie_main_charactor string
	Movie_on_line        string
	Movie_span           string
	Movie_grade          string
	remark               string
	_create_time         string
	Movie_name_as        string
}

func init() {
	// 需要在init中注册定义的model
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/beego?charset=utf8", 30, 30) //注册默认数据库
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}
func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err

}
func GetMovieDirectory(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*? rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func GetMovieWrite(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a href="/celebrity/.*?">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[0]
	}
	return main
	//return string(result[0][1])
}
func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 主演
func GetMovieMainActor(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
}

//类型==

func GetMovieType(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
}

// 制片国家movie_conutry
func GetMovieCountry(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">制片国家/地区:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 语言
func GetMovieLanguage(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">语言:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//<span property="v:initialReleaseDate" content="2018-11-09(中国大陆)">2018-11-09(中国大陆)</span>
// 上映时间 movie_on_line

func GetMovieOnLineTime(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:initialReleaseDate" content="(.*?)">`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	main := ""
	for _, v := range result {
		main += v[1] + "/"
	}
	return strings.Trim(main, "/")
	//return string(result[0][1])
}

// 片长
//<span property="v:runtime" content="114">114分钟</span>
func GetMovieSpan(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 又名
func GetMovieNameAs(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">又名:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

// 图片

func GetMoviePic(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	//reg := regexp.MustCompile(`<a class="nbgnbg".*?><img src="(.*?)" title="点击看更多海报".*?></a>`)
	reg := regexp.MustCompile(`<img src="(.*?)" title="点击看更多海报"`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//<strong class="ll rating_num" property="v:average">6.8</strong>
// 评分
func GetMovieScores(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//desc

func GetMovieDesc(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:summary" class="">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func GetMovieUrl(movieHtml string) []string {
	var movieurls []string
	reg := regexp.MustCompile(`<a.*href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(string(movieHtml), -1)
	if len(result) == 0 {
		return movieurls
	}
	for _, v := range result {
		movieurls = append(movieurls, v[1])
	}
	return movieurls
}

// 获取电影id
func MovieId(url string) int {
	reg := regexp.MustCompile(`subject/(\d+)/.*`)
	result := reg.FindAllStringSubmatch(string(url), -1)

	if len(result) == 0 {
		return 0
	}
	int, _ := strconv.Atoi(result[0][1])
	return int

}
