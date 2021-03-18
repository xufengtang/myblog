package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"myblog/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var pageSize int = 10
var session sessions.Session

//Init Init
func Init(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", map[string]interface{}{})
}

//dologin dologin
func dologin(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()

	userName := c.PostForm("username")
	userPwd := c.PostForm("pwd")

	user, err := models.Loginusers(qm.Where("userName = ?", userName), qm.And("userPwd = ?", userPwd)).One(ctx, db)

	if err != nil {
		msg := "ユーザが存在しておりません"
		c.JSON(http.StatusOK, gin.H{
			"result": "false",
			"msg":    msg,
		})
	} else {

		if user != nil {

			session = sessions.Default(c)
			session.Set("loginUser", user.UserName)

			c.JSON(http.StatusOK, gin.H{
				"result": "ok",
			})
		} else {
			msg := "ユーザが存在しておりません"
			c.JSON(http.StatusOK, gin.H{
				"result": "false",
				"msg":    msg,
			})
		}
	}

}

func doLogout(c *gin.Context) {

	session.Clear()

	c.HTML(http.StatusOK, "login.html", map[string]interface{}{})
}

//doSearchList doSearchList
func doSearchList(c *gin.Context) {

	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()

	searchKey := c.PostForm("searchKey")
	currentPage, _ := strconv.Atoi(c.PostForm("currentPage"))
	searchType := c.PostForm("searchType")

	switch searchType {
	case "1":
		currentPage = 1
	case "2":
		currentPage = currentPage - 1
	case "3":
		currentPage = currentPage + 1
	}

	currentPageSearch := currentPage
	if currentPageSearch != 0 {
		currentPageSearch = currentPageSearch - 1
	}

	var iswhere qm.QueryMod
	if searchKey == "" {

		iswhere = qm.Where("1=?", 1)
	} else {
		iswhere = qm.Where("title like ?", "%"+searchKey+"%")
	}

	sumCount, _ := models.Mainblogs(iswhere).Count(ctx, db)

	mainblogs, _ := models.Mainblogs(iswhere, qm.Limit(pageSize), qm.Offset(currentPageSearch*pageSize)).All(ctx, db)

	pageCount := sumCount/int64(pageSize) + 1

	if mainblogs != nil {

		mainblogsJSON, _ := json.Marshal(mainblogs)

		c.JSON(http.StatusOK, gin.H{
			"result":      "ok",
			"pageCount":   pageCount,
			"currentPage": currentPage,
			"sumCount":    sumCount,
			"mainblogs":   string(mainblogsJSON),
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
			"msg":    "データが存在しません",
		})
	}

}

//indexInit indexInit
func indexInit(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
}

//samllInit samllInit
func samllInit(c *gin.Context) {

	selectID := c.Query("selectID")

	c.HTML(http.StatusOK, "smallSelect.html", map[string]interface{}{
		"selectID": selectID,
	})
}

func smallSelect(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()

	selectID := c.PostForm("selectID")
	searchType := c.PostForm("searchType")
	currentPage, _ := strconv.Atoi(c.PostForm("currentPage"))

	switch searchType {
	case "1":
		currentPage = 1
	case "2":
		currentPage = currentPage - 1
	case "3":
		currentPage = currentPage + 1
	}

	currentPageSearch := currentPage
	if currentPageSearch != 0 {
		currentPageSearch = currentPageSearch - 1
	}

	sumCount, _ := models.Smallblogs(qm.Where("parentId=?", selectID)).Count(ctx, db)
	smallblogs, _ := models.Smallblogs(qm.Where("parentId=?", selectID), qm.Limit(pageSize), qm.Offset(currentPageSearch*pageSize)).All(ctx, db)
	pageCount := sumCount/int64(pageSize) + 1

	if smallblogs != nil {
		smallblogsJSON, _ := json.Marshal(smallblogs)

		c.JSON(http.StatusOK, gin.H{
			"result":      "ok",
			"pageCount":   pageCount,
			"currentPage": currentPage,
			"sumCount":    sumCount,
			"smallblogs":  string(smallblogsJSON),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
			"msg":    "データが存在しません",
		})
	}
}

func doSave(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()

	selectID := c.PostForm("selectID")
	parentID, _ := strconv.Atoi(selectID)
	titleEdit := c.PostForm("titleEdit")
	contentEdit := c.PostForm("contentEdit")

	timeValue := time.Now()

	var smallblog models.Smallblog
	smallblog.ParentId.SetValid(parentID)
	smallblog.Title.SetValid(titleEdit)
	smallblog.Content.SetValid(contentEdit)
	smallblog.CreateTime.SetValid(timeValue)
	smallblog.UpdateTime.SetValid(timeValue)

	err := smallblog.Insert(ctx, db, boil.Infer())

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	}

}

func doUpdate(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()
	idEdit := c.PostForm("idEdit")
	titleEdit := c.PostForm("titleEdit")
	contentEdit := c.PostForm("contentEdit")

	timeValue := time.Now()

	updateID, _ := strconv.Atoi(idEdit)

	smallblog, _ := models.FindSmallblog(ctx, db, updateID)

	smallblog.Title.SetValid(titleEdit)
	smallblog.Content.SetValid(contentEdit)
	smallblog.UpdateTime.SetValid(timeValue)

	rowsAff, err := smallblog.Update(ctx, db, boil.Infer())

	if rowsAff == 1 && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
		})
	}
}

func doDelete(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "my_blog_id:my_blog_id@tcp(localhost:3306)/myblog?parseTime=true")
	defer db.Close()
	deleteID := c.PostForm("deleteId")

	updateID, _ := strconv.Atoi(deleteID)

	smallblog, _ := models.FindSmallblog(ctx, db, updateID)

	rowsAff, err := smallblog.Delete(ctx, db)

	if rowsAff == 1 && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
		})
	}
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUser := session.Get("loginUser")
		// セッションがない場合、ログインフォームをだす
		if loginUser == nil {
			//c.HTML(http.StatusOK, "login.html", map[string]interface{}{})
			c.Redirect(http.StatusMovedPermanently, "/")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLFiles("./templent/login.html",
		"./templent/index.html",
		"./templent/smallSelect.html")
	r.Static("/css", "./templent/css")
	r.Static("/images", "./templent/images")
	r.Static("/fonts", "./templent/fonts")
	/**ログイン**/
	r.GET("/", Init)
	r.POST("/dologin", dologin)
	/**ログアウト**/
	r.GET("/doLogout", doLogout)

	menu := r.Group("menu")

	menu.Static("/css", "./templent/css")
	menu.Static("/images", "./templent/images")
	menu.Static("/fonts", "./templent/fonts")

	/**大項目**/
	r.GET("/indexInit", indexInit)
	menu.Use(sessionCheck())
	{

		//menu.GET("/indexInit", indexInit)
		menu.POST("/indexSearchList", doSearchList)
		/**小項目**/
		menu.GET("/samllInit", samllInit)
		menu.POST("/smallSelect", smallSelect)
		menu.POST("/doSave", doSave)
		menu.POST("/doUpdate", doUpdate)
		menu.POST("/doDelete", doDelete)
	}

	// /**大項目**/
	// r.GET("/indexInit", indexInit)
	// r.POST("/indexSearchList", doSearchList)
	/**小項目**/
	// r.GET("/samllInit", samllInit)
	// r.POST("smallSelect", smallSelect)
	// r.POST("doSave", doSave)
	// r.POST("doUpdate", doUpdate)
	// r.POST("doDelete", doDelete)

	r.Run(":8088") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
