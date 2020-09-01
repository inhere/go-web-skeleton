package controller

import (
	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/app/errcode"
	"github.com/inhere/go-web-skeleton/model/form"
	"github.com/inhere/go-web-skeleton/model/logic"
)

// UserApi
type UserApi struct {
	BaseApi
}

// AddRoutes for the API controller
func (u *UserApi) AddRoutes(g *rux.Router) {
	g.GET("/users/{id}", u.GetOne)
}

// @Tags UserApi
// @Summary Get one user detail by ID
// @Description get data by ID
// @Param   id     path    int     true        "user ID"
// @Failure 200 {object} model.JsonMapData "We need ID!!"
// @Failure 404 {object} model.JsonMapData "Can not find ID"
// @Router /users/{id} [get]
func (u *UserApi) GetOne(c *rux.Context) {
	u.JSON(c, 200, "hello")
}

// @Tags UserApi
// @Summary create an new user
// @Description get data by ID
// @Param   bodyData     body    form.CreateUserForm     true  "new user data"
// @Failure 200 {object} model.JsonMapData "We need ID!!"
// @Failure 404 {object} model.JsonMapData "Can not find ID"
// @Router /users [post]
func (u *UserApi) Save(c *rux.Context) {
	var f form.CreateUserForm

	// c.PostParams()

	if err := c.Bind(&f); err != nil {
		c.AbortThen().JSON(406, u.MakeRes(errcode.ErrParam, err, map[int]int{}))
		return
	}

	code, err, id := logic.CreateUser(&f)

	c.JSON(200, u.MakeRes(code, err, rux.M{"id": id}))
}

