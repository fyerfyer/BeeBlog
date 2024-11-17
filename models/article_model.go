package models

import (
	"log"
	"strings"
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
	Tags       []*Tag `orm:"rel(m2m)"`
	TagString  string
}

func init() {
	orm.RegisterModel(new(Article))
}

func CreateArticle(article *Article, tagString string) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(article)
	if err != nil {
		log.Printf("CreateArticle failed: %v", err)
		return 0, err
	}

	// set the m2m relation
	m2m := o.QueryM2M(article, "Tags")

	tagStrings := strings.Split(tagString, ",")
	// create and add tag first
	for _, tagString := range tagStrings {
		tag := &Tag{Name: tagString}
		_, _, err := o.ReadOrCreate(tag, "Name")
		if err != nil {
			log.Printf("CreateTag failed: %v", err)
			return 0, err
		}

		// build the relationship
		_, err = m2m.Add(tag)
		if err != nil {
			log.Printf("Build relationship failed: %v", err)
			return 0, err
		}
	}

	return int(id), err
}

func GetArticleById(id int) (*Article, error) {
	o := orm.NewOrm()
	article := Article{Id: id}
	err := o.Read(&article, "Id")
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

func GetArticlesWithTags(tags *[]Tag) (*[]Article, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	var articles *[]Article
	for _, tag := range *tags {
		qs = qs.Filter("Tags__Tag__Id", tag.Id)
	}
	_, err := qs.Distinct().All(articles)
	if err != nil {
		log.Printf("GetArticleWithTags failed: %v", err)
		return nil, err
	}

	return articles, nil
}
