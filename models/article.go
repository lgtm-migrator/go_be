package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       string `orm:"column(id);pk"`
	Content  string
	Title    string
	OkNumber int64
	Momment  string
}

func init() {
	orm.RegisterModel(new(Article))
}

//添加新的文章
func Add(Content string, Title string) (int64, error) {
	o := orm.NewOrm()
	var article Article
	article.Content = Content
	article.Title = Title
	Id, err := o.Insert(&article)
	return Id, err
}

//获取所有文章
func GetAll() ([]Article, error) {
	o := orm.NewOrm()
	var article []Article
	qs := o.QueryTable("article")
	message, err := qs.All(&article)
	fmt.Println(message)
	return article, err
}

//根据ID获取文章
func GetOne(Id string) (*Article, error) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable("article")
	err := qs.Filter("Id", Id).One(article)
	return article, err
}

//根据ID修改文章
func UpdateOne(Id string, Momment string, Title string) Article {
	o := orm.NewOrm()
	article := Article{Id: Id}
	article.Momment = Momment
	article.Title = Title
	o.Update(&article)
	fmt.Println(article)
	return article
}

//根据ID删除文章
func DeleteOne(Id string) Article {
	o := orm.NewOrm()
	article := Article{Id: Id}
	o.Delete(&article)
	fmt.Println(article)
	return article
}