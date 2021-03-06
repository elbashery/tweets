package tweet

import (
	"bufio"
	"fmt"
	"time"

	"github.com/bashery/tweets/dbs"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
)

type Tweet struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetOne(c *fiber.Ctx) {

	//c.Accepts("application/json")
	id := c.Params("id")
	db := dbs.Conn
	var tweet Tweet
	db.Find(&tweet, id)

	if tweet.Body == "" {
		c.Send("tweet is no round")
		return
	}

	c.JSON(tweet)
}

func New(c *fiber.Ctx) {

	db := dbs.Conn

	tweet := &Tweet{}
	// TODO : hand zero paramse,
	//tweet.Title = c.Params("title")
	//tweet.Body = c.Params("body")
	if err := c.BodyParser(tweet); err != nil {
		c.Status(503).JSON(err)
		fmt.Println("error")
		return
	}
	db.Create(tweet)
	c.JSON(tweet) // or c.Send("success")
}

func Update(c *fiber.Ctx) {
	db := dbs.Conn
	var tweet Tweet
	db.First(&tweet, c.Params("id"))
	// I will use params instead BodyParser just for semple and test performece
	// if err := c.BodyParser(&tweet); err != nil {c.Status(503).Send(err);	fmt.Println(err);return;}
	// TODO check params is no zero val and handel it
	db.Model(&tweet).Update("body", c.Params("body")) // &tweet.Body)

	c.Send("success") // or c.JSON("success")
}

func Remove(c *fiber.Ctx) {
	var tweet Tweet
	db := dbs.Conn
	id := c.Params("id")
	db.First(&tweet, id)

	if tweet.Title == "" {
		c.Status(500).JSON("no tweet find with given ID ")
		return
	}
	db.Delete(&tweet)
	c.JSON("tweet is deleted")
}

func GetAll(c *fiber.Ctx) {
	db := dbs.Conn
	var tweets []Tweet
	db.Find(&tweets)
	c.JSON(tweets)
}
