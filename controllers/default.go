package controllers

import (
	"beeblog/models"
	"beeblog/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "github.com/beego/beego/v2/client/orm"
	// bee "github.com/beego/beego/v2/server/web"
)

type Err struct {
	HasError bool
	ErrMsg   string
}

type HomeController struct {
	BaseController
	Articles     *[]models.Article
	HasInitiated bool
}

// todo: show the latest articles in the homepage, and all articles in other pages
// and also page division
func (c *HomeController) Get() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	if !c.HasInitiated {
		models.GetAllArticle(c.Articles)
		c.HasInitiated = true
	}
	utils.RenderFlash(&c.Controller, "base.layout.tpl")

	pageStr := c.Ctx.Input.Param(":page")
	if pageStr == "" {
		pageStr = "1"
	}
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Printf("Invalid page number input: %v", err)
		return
	}

	c.TplName = "home.page.tpl"
	c.Data["Articles"] = utils.GetPaginatedArticles(c.Articles, pageNum, 8)
	c.Data["CurrentPage"] = pageNum
	c.Data["TotalPages"] = (len(*c.Articles) + 7) / 8
}

type ShowController struct {
	BaseController
	Articles *[]models.Article
}

func (c *ShowController) Get() {
	utils.RenderFlash(&c.Controller, "base.layout.tpl")
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")

	c.TplName = "show.page.tpl"
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid article id input: %v", err)
		return
	}
	log.Println("get id: ", id)

	article, err := models.GetArticle(id)
	if err != nil {
		log.Printf("GetArticle failed: %v", err)
		return
	}
	log.Println(article)

	c.Data["Title"] = article.Title
	c.Data["Id"] = article.Id
	c.Data["CreateTime"] = article.CreateTime
	c.Data["ExpireTime"] = article.ExpireTime
	c.Data["Content"] = article.Content
	c.Render()
}

type CreateController struct {
	BaseController
	Articles *[]models.Article
}

func (c *CreateController) Get() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "create.page.tpl"
}

func (c *CreateController) Post() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "create.page.tpl"
	title := c.GetString("title")
	content := c.GetString("content")
	expires, err := c.GetInt("expires")
	if err != nil {
		log.Printf("Error getting expires: %v", err)
		return
	}
	createTime := time.Now()
	expiresTime := time.Now().Add(time.Duration(expires) * time.Hour * 24)

	// fmt.Println(createTime)
	// fmt.Println(expiresTime)
	// user := new(models.User)
	// user.Username = "fyerfyer"
	// user.CreateTime = time.Now()
	// user.Status = 1
	// user.Password = "dfkjahfk"

	article := models.Article{
		Title:      title,
		Content:    content,
		CreateTime: createTime,
		ExpireTime: expiresTime,
		// User:       user,
	}
	var id int
	if id, err = models.CreateArticle(&article); err != nil {
		log.Printf("CreateArticle failed: %v", err)
		c.Data["Error"] = "Failed to create article."
		c.TplName = "create.page.tpl"
		return
	}
	*c.Articles = append(*c.Articles, article)
	utils.SetFlash(&c.Controller, "success", "Successfully create an article")
	c.Redirect(fmt.Sprintf("/beegoblog/%d", id), http.StatusFound)
}

type SignUpController struct {
	BaseController
}

func (c *SignUpController) Get() {
	// log.Println("Rendering signup page")
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "signup.page.tpl"
	c.Data["NameError"] = Err{HasError: false}
	c.Data["EmailError"] = Err{HasError: false}
	c.Data["PasswordError"] = Err{HasError: false}
}

func (c *SignUpController) Post() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "signup.page.tpl"

	isValid := true
	NameError := Err{HasError: false}
	EmailError := Err{HasError: false}
	PasswordError := Err{HasError: false}

	name := c.GetString("name")
	email := c.GetString("email")
	password := c.GetString("password")
	c.Data["Name"] = name
	c.Data["Email"] = email
	c.Data["Password"] = password

	user := &models.User{
		Username:       name,
		Email:          email,
		HashedPassword: password,
		Status:         1,
		CreateTime:     time.Now(),
	}

	errMsg := user.Valid()
	if len(errMsg) > 0 {
		isValid = false
		for k, v := range errMsg {
			if k == "Name" {
				c.Data[k] = ""
				NameError.HasError = true
				NameError.ErrMsg = v
			}
			if k == "Email" {
				c.Data[k] = ""
				EmailError.HasError = true
				EmailError.ErrMsg = v
			}
			if k == "Password" {
				c.Data[k] = ""
				PasswordError.HasError = true
				PasswordError.ErrMsg = v
			}
		}
	}

	c.Data["NameError"] = NameError
	c.Data["PasswordError"] = PasswordError
	c.Data["EmailError"] = EmailError

	// if the input is invalid, there's no need to hash
	if !isValid {
		c.Render()
		log.Println("Failed to sign up")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return
	}
	user.HashedPassword = string(hashedPassword)
	models.CreateUser(user)
	utils.SetFlash(&c.Controller, "success", "Successfully create a user")
	c.Redirect("/user/login", http.StatusFound)
}

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	utils.RenderFlash(&c.Controller, "base.layout.tpl")
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "login.page.tpl"
	c.Data["Error"] = Err{HasError: false}
}

func (c *LoginController) Post() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	c.TplName = "login.page.tpl"
	email := c.GetString("email")
	password := c.GetString("password")
	c.Data["Email"] = email
	e := Err{HasError: false}
	userID, hasFound, isAuthenticated := models.UserAuthenticate(email, password)

	if !hasFound || !isAuthenticated {
		e.HasError = true
		if !hasFound {
			if email == "" {
				e.ErrMsg = "Email cannot be empty"
			} else {
				e.ErrMsg = "Cannot find active user with this email"
			}
			c.Data["Email"] = ""
		} else {
			e.ErrMsg = "Password incorrect"
		}
		c.Data["Error"] = e
		c.Render()
		return
	}

	c.Data["Error"] = e
	// authenticatation
	c.SetSession("authenticatedUserID", userID)
	utils.SetFlash(&c.Controller, "success", "Successfully log in")
	c.Redirect("/beegoblog/create", http.StatusFound)
}

type LogoutController struct {
	BaseController
}

func (c *LogoutController) Post() {
	// utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")
	sessionValue := c.GetSession("authenticatedUserID")
	id, ok := sessionValue.(int)
	if !ok {
		log.Println("Failed to convert session value to string")
		return
	}

	if err := c.DelSession("authenticatedUserID"); err != nil {
		log.Printf("Failed to log out: %v", err)
		return
	}

	models.DeactiveUser(id)
	utils.SetFlash(&c.Controller, "success", "Successfully log out")
	c.Redirect("/", http.StatusFound)
}
