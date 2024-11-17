package controllers

import (
	"beeblog/models"
	"beeblog/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "github.com/beego/beego/v2/client/orm"
	bee "github.com/beego/beego/v2/server/web"
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

func (c *HomeController) Get() {
	utils.RenderFlash(&c.Controller, "base.layout.tpl")

	hasAuthenticatedUser, userId := utils.GetIntSession(&c.Controller, "authenticatedUserID")
	if !hasAuthenticatedUser {
		c.Articles = &[]models.Article{}
	} else if !c.HasInitiated {
		err := models.GetArticleByUserId(userId, c.Articles)
		if err != nil {
			log.Printf("Failed to get articles of current user: %v", err)
			return
		}

		c.HasInitiated = true
	}

	for _, article := range *c.Articles {
		utils.SetTagString(&article)
	}

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

	c.TplName = "show.page.tpl"
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid article id input: %v", err)
		return
	}

	article, err := models.GetArticleById(id)
	if err != nil {
		log.Printf("GetArticle failed: %v", err)
		return
	}

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
	utils.RenderFlash(&c.Controller, "base.layout.tpl")
	c.TplName = "create.page.tpl"
}

func (c *CreateController) Post() {
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
	tagContext := c.GetString("tags")

	_, userId := utils.GetIntSession(&c.Controller, "authenticatedUserID")
	if userId == 0 {
		log.Println("Failed to get userId")
		return
	}

	article := models.Article{
		Title:      title,
		Content:    content,
		CreateTime: createTime,
		ExpireTime: expiresTime,
		UserId:     userId,
		TagString:  tagContext,
	}

	var articleId int
	if articleId, err = models.CreateArticle(&article, tagContext); err != nil {
		log.Printf("CreateArticle failed: %v", err)
		c.Data["Error"] = "Failed to create article."
		c.TplName = "create.page.tpl"
		return
	}
	*c.Articles = append(*c.Articles, article)
	utils.SetFlash(&c.Controller, "success", "Successfully create an article")
	c.Redirect(fmt.Sprintf("/beegoblog/%d", articleId), http.StatusFound)
}

type SignUpController struct {
	BaseController
}

func (c *SignUpController) Get() {
	c.TplName = "signup.page.tpl"
	c.Data["NameError"] = Err{HasError: false}
	c.Data["EmailError"] = Err{HasError: false}
	c.Data["PasswordError"] = Err{HasError: false}
}

func setError(err *Err, c *bee.Controller, k string, v string) {
	c.Data[k] = ""
	err.HasError = true
	err.ErrMsg = v
}

func (c *SignUpController) Post() {
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
		Status:         0,
		CreateTime:     time.Now(),
	}

	errMsg := user.Valid()
	if len(errMsg) > 0 {
		isValid = false
		for k, v := range errMsg {
			switch k {
			case "Name":
				setError(&NameError, &c.Controller, k, v)
			case "Email":
				setError(&EmailError, &c.Controller, k, v)
			case "Password":
				setError(&PasswordError, &c.Controller, k, v)
			}
		}
	}

	c.Data["NameError"] = NameError
	c.Data["PasswordError"] = PasswordError
	c.Data["EmailError"] = EmailError

	// if the input is invalid, there's no need to hash
	if !isValid {
		log.Println("Failed to sign up")
		c.Render()
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
	c.TplName = "login.page.tpl"
	c.Data["Error"] = Err{HasError: false}
}

func handleNotFound(e *Err, email string, c *bee.Controller) {
	e.ErrMsg = "Email cannot be empty"
	e.HasError = true
	if email != "" {
		e.ErrMsg = "Cannot find active user with this email"
	}

	c.Data["Email"] = ""
	c.Data["Error"] = e
}

func handleNotAuthenticate(e *Err, c *bee.Controller) {
	e.HasError = true
	e.ErrMsg = "Password incorrect"
	c.Data["Error"] = e
}

func (c *LoginController) Post() {
	c.TplName = "login.page.tpl"
	email := c.GetString("email")
	password := c.GetString("password")
	c.Data["Email"] = email
	e := Err{HasError: false}
	userId, hasFound, isAuthenticated := models.UserAuthenticate(email, password)

	flag := true
	if !isAuthenticated {
		flag = false
		handleNotAuthenticate(&e, &c.Controller)
	}

	if !hasFound {
		flag = false
		handleNotFound(&e, email, &c.Controller)
	}

	if !flag {
		c.Render()
		return
	}

	c.Data["Error"] = e
	// authenticatation
	models.SetUserStatusById(userId, 1)
	c.SetSession("authenticatedUserID", userId)
	utils.SetFlash(&c.Controller, "success", "Successfully log in")
	c.Redirect("/beegoblog/create", http.StatusFound)
}

type LogoutController struct {
	BaseController
}

func (c *LogoutController) Post() {
	_, id := utils.GetIntSession(&c.Controller, "authenticatedUserID")
	if id == 0 {
		log.Println("Failed to get userId")
		return
	}

	if err := c.DelSession("authenticatedUserID"); err != nil {
		log.Printf("Failed to log out: %v", err)
		return
	}

	log.Println("logout !!!!!!!")
	models.SetUserStatusById(id, 0)
	utils.SetFlash(&c.Controller, "success", "Successfully log out")
	c.Redirect("/", http.StatusFound)
}

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Get() {
	c.TplName = "profile.page.tpl"
	_, id := utils.GetIntSession(&c.Controller, "authenticatedUserID")
	if id == 0 {
		log.Println("Failed to get userId")
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		log.Printf("Failed to get the user: %v", err)
		return
	}

	c.Data["User"] = user
}
