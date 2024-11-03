package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id         int       `orm:"auto"` // self-added primary key
	Title      string    `orm:"size(128)"`
	Content    string    `orm:"type(text)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ExpireTime time.Time `orm:"auto_now_add;type(datetime)"`
	UserId     int
}

func init() {
	orm.RegisterModel(new(Article))
}

func CreateArticle(article *Article) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(article)
	if err != nil {
		log.Println("CreateArticle failed")
	}
	return int(id), err
}

func GetArticleById(id int) (*Article, error) {
	o := orm.NewOrm()
	article := Article{Id: id}
	err := o.Read(&article)
	if err != nil {
		if err == orm.ErrNoRows {
			log.Println("Cannot find article")
		} else if err == orm.ErrMissPK {
			log.Println("Invalid id input")
		} else {
			log.Println("GetArticle fail")
		}
		return nil, err
	}

	return &article, nil
}

func GetArticleByUserId(id int, articles *[]Article) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("user_id", id).All(articles)
	if err != nil {
		log.Printf("GetArticleByUserId failed: %v", err)
		return err
	}

	return nil
}
