package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type Tag struct {
	Id       int        `orm:"auto"`
	Name     string     `orm:"size(30)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Tag))
}

func InsertTags(id int, tags []*Tag) error {
	o := orm.NewOrm()
	article := Article{Id: id}
	article.Tags = tags
	_, err := o.Update(&article, "Tags")
	if err != nil {
		log.Printf("InsertTagsByArticleId failed: %v", err)
		return err
	}

	return nil
}

func GetTagsByArticleId(id int, tags *[]Tag) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article_tags").Filter("article_id", id).All(tags)
	if err != nil {
		log.Printf("GetTagsByArticleId failed: %v", err)
		return err
	}

	return nil
}
