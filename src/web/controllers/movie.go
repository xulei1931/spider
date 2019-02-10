package controllers

import (
	"web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"

	//"fmt"
	"time"
)

type MovieController struct {
	// 继承
	beego.Controller
}

func (c *MovieController) Create() {
	var movieInfo models.MovieInfo
	// 爬虫入口
	surl := `https://movie.douban.com`

	// 连接redis
	models.ConnectRedis()
	models.PutQueue(surl)

	//rsp := httplib.Get(surl)
	//html, _ := rsp.String()
	//c.Ctx.WriteString(html)
	for {
		length := models.GetQueueLength()
		if length == 0 {
			break
		}

		surl = models.PopQueue()
		if models.IsHave(surl) {
			continue
		}
		rsp := httplib.Get(surl)
		html, err := rsp.String()
		if err != nil {
			continue
		}
		movieInfo.Movie_name = models.GetMovieName(html)
		if movieInfo.Movie_name != "" {
			// 爬取电影
			movieInfo.Id = 0
			movieInfo.Movie_id = models.MovieId(surl)
			movieInfo.Movie_name = models.GetMovieName(html)
			movieInfo.Movie_director = models.GetMovieDirectory(html)
			movieInfo.Movie_writer = models.GetMovieWrite(html)
			movieInfo.Movie_conutry = models.GetMovieCountry(html)
			movieInfo.Movie_language = models.GetMovieLanguage(html)
			movieInfo.Movie_type = models.GetMovieType(html)
			movieInfo.Movie_image = models.GetMoviePic(html)
			movieInfo.Movie_main_charactor = models.GetMovieMainActor(html)
			movieInfo.Movie_on_line = models.GetMovieOnLineTime(html)
			movieInfo.Movie_span = models.GetMovieSpan(html)
			movieInfo.Movie_grade = models.GetMovieScores(html)
			movieInfo.Movie_name_as = models.GetMovieNameAs(html)
			models.AddMovie(&movieInfo)
		}
		// 提取网页的URL加到队列
		urls := models.GetMovieUrl(html)
		for _, url := range urls {
			models.PutQueue(url)
			c.Ctx.WriteString(url)

		}
		// 把爬取的url加到集合
		models.AddSet(surl)
		time.Sleep(time.Second)

	}
	c.Ctx.WriteString("end.......")

}
