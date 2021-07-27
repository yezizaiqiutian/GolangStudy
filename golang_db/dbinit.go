package golang_db

import (
	"GolangStudy/golang_gin/bean"
	"strconv"
)

func InitNewsDb() {
	//新建表
	//DB.AutoMigrate(&bean.News{})

	for i, n := 0, 100; i < n; i++ {

		news := bean.News{
			Title:   "新闻标题" + strconv.Itoa(i),
			Content: "新闻内容新闻内容新闻内容新闻内容新闻内容" + strconv.Itoa(i),
			Url:     "http://www.baidu.com",
		}

		DB.Create(&news)

	}

}
