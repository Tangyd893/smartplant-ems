package controllers

import (
	"smartplant-ems/common/utils"
	"smartplant-ems/device-service/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.ReturnBadRequest(&c.Controller, "用户名和密码不能为空")
		return
	}

	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(new(models.User)).Filter("username", req.Username).One(&user)

	if err != nil {
		utils.ReturnError(&c.Controller, 401, "用户名或密码错误")
		return
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		utils.ReturnError(&c.Controller, 401, "用户名或密码错误")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		utils.ReturnServerError(&c.Controller, "生成令牌失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":         user.ID,
			"username":   user.Username,
			"real_name":  user.RealName,
			"role":       user.Role,
			"phone":      user.Phone,
			"email":      user.Email,
			"status":     user.Status,
		},
	})
}

func (c *AuthController) Logout() {
	c.DestroySession()
	utils.ReturnSuccess(&c.Controller, nil)
}

func (c *AuthController) GetCurrentUser() {
	userID := c.GetSession("user_id")
	if userID == nil {
		utils.ReturnUnauthorized(&c.Controller, "未登录")
		return
	}

	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(new(models.User)).Filter("id", userID).One(&user)
	if err != nil {
		utils.ReturnNotFound(&c.Controller, "用户不存在")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"real_name":  user.RealName,
		"role":       user.Role,
		"phone":      user.Phone,
		"email":      user.Email,
		"status":     user.Status,
	})
}