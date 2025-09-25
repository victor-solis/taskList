package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := getOrmDb()
	//db.AutoMigrate(
	//	&User{},
	//	&Post{},
	//	&Comment{},
	//)
	var user User
	db.Preload("Post").Preload("Post.Comment").First(&user, 1)
	fmt.Println(user)

	var post Post
	//查找评论最多的文章
	db.Model(&Post{}).Select("posts.*,count(comments.id) as comment_count ").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count desc").
		First(&post)
	fmt.Println(post)

}

func getOrmDb() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印 SQL 日志
	})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return nil
	}
	return db
}

type User struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255" json:"name"`
	Post []Post `json:"post"`
}

type Post struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"size:255" json:"content"`
	UserId    int       `json:"user_id"`
	Comment   []Comment `json:"comment"`
	WordCount int       `json:"word_count"`
}

type Comment struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Remark string `gorm:"size:255" json:"remark"`
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	User   User   `json:"user"`
	Post   Post   `json:"post"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 在创建之前自动调用
	p.WordCount = len(p.Comment)
	return
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if count == 0 {
		defaultComment := Comment{
			PostId: c.PostId,
			Remark: "暂无评论",
		}
		return db.Create(&defaultComment).Error
	}
	return nil
}
