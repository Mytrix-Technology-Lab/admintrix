package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// Middleware
	//s.Use(middleware.CORS)
	//s.Use(middleware.CheckLogin)
	//s.Use(middleware.OperLog)
	//s.Use(middleware.LoginLog)

	/* Upload */
	s.Group("/upload", func(group *ghttp.RouterGroup) {
		//group.POST("/uploadImage", controller.Upload.UploadImage)
	})

	/* Index and Root */
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.GET("/", controller.Login.Login)
		//group.ALL("/login", controller.Login.Login)
		//group.GET("/captcha", controller.Login.Captcha)
		//group.ALL("/updateUserInfo", controller.Index.UpdateUserInfo)
		//group.ALL("/updatePwd", controller.Index.UpdatePwd)
		//group.GET("/logout", controller.Index.Logout)
	})

	s.Group("index", func(group *ghttp.RouterGroup) {
		//group.GET("/menu", controller.Index.Menu)
		//group.GET("/user", controller.Index.User)
	})

	/* User */
	s.Group("user", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.User.List)
		//group.GET("/detail", controller.User.Detail)
		//group.POST("/add", controller.User.Add)
		//group.PUT("/update", controller.User.Update)
		//group.DELETE("/delete/:ids", controller.User.Delete)
		//group.PUT("/status", controller.User.Status)
		//group.PUT("/resetPwd", controller.User.ResetPwd)
		//group.GET("/checkUser", controller.User.CheckUser)
	})

	/* Level */
	s.Group("level", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Level.List)
		//group.POST("/add", controller.Level.Add)
		//group.PUT("/update", controller.Level.Update)
		//group.DELETE("/delete/:ids", controller.Level.Delete)
		//group.PUT("/status", controller.Level.Status)
		//group.GET("/getLevelList", controller.Level.GetLevelList)
		//group.GET("/exportExcel", controller.Level.ExportExcel)
		//group.POST("/importExcel", controller.Level.ImportExcel)
		//group.GET("/downloadExcel", controller.Level.DownloadExcel)
	})

	/* Position */
	s.Group("position", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Position.List)
		//group.POST("/add", controller.Position.Add)
		//group.PUT("/update", controller.Position.Update)
		//group.DELETE("/delete/:ids", controller.Position.Delete)
		//group.PUT("/status", controller.Position.Status)
		//group.GET("/getPositionList", controller.Position.GetPositionList)
	})

	/* Role */
	s.Group("role", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Role.List)
		//group.POST("/add", controller.Role.Add)
		//group.PUT("/update", controller.Role.Update)
		//group.DELETE("/delete/:ids", controller.Role.Delete)
		//group.PUT("/status", controller.Role.Status)
		//group.GET("/getRoleList", controller.Role.GetRoleList)
	})

	/* Role Menu */
	s.Group("rolemenu", func(group *ghttp.RouterGroup) {
		//group.GET("/index", controller.RoleMenu.Index)
		//group.POST("/save", controller.RoleMenu.Save)
	})

	/* Department */
	s.Group("dept", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Dept.List)
		//group.POST("/add", controller.Dept.Add)
		//group.PUT("/update", controller.Dept.Update)
		//group.DELETE("/delete/:ids", controller.Dept.Delete)
		//group.GET("/getDeptList", controller.Dept.GetDeptList)
	})

	/* Menu */
	s.Group("menu", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Menu.List)
		//group.GET("/detail", controller.Menu.Detail)
		//group.POST("/add", controller.Menu.Add)
		//group.PUT("/update", controller.Menu.Update)
		//group.DELETE("/delete/:ids", controller.Menu.Delete)
	})

	/* Operation Log */
	s.Group("operlog", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.OperLog.List)
	})

	/* Login Log */
	s.Group("loginlog", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.LoginLog.List)
		//group.DELETE("/delete/:ids", controller.LoginLog.Delete)
	})

	/* City */
	s.Group("city", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.City.List)
		//group.POST("/add", controller.City.Add)
		//group.PUT("/update", controller.City.Update)
		//group.DELETE("/delete/:ids", controller.City.Delete)
		//group.POST("/getChilds", controller.City.GetChilds)
	})

	/* Dictionary */
	s.Group("dict", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Dict.List)
		//group.POST("/add", controller.Dict.Add)
		//group.PUT("/update", controller.Dict.Update)
		//group.DELETE("/delete/:ids", controller.Dict.Delete)
	})

	/* Dictionary Data */
	s.Group("dictdata", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.DictData.List)
		//group.POST("/add", controller.DictData.Add)
		//group.PUT("/update", controller.DictData.Update)
		//group.DELETE("/delete/:ids", controller.DictData.Delete)
	})

	/* Configuration */
	s.Group("config", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Config.List)
		//group.POST("/add", controller.Config.Add)
		//group.PUT("/update", controller.Config.Update)
		//group.DELETE("/delete/:ids", controller.Config.Delete)
	})

	/* Configuration Data */
	s.Group("configdata", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.ConfigData.List)
		//group.POST("/add", controller.ConfigData.Add)
		//group.PUT("/update", controller.ConfigData.Update)
		//group.DELETE("/delete/:ids", controller.ConfigData.Delete)
		//group.PUT("/status", controller.ConfigData.Status)
	})

	/* Link */
	s.Group("link", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Link.List)
		//group.POST("/add", controller.Link.Add)
		//group.PUT("/update", controller.Link.Update)
		//group.DELETE("/delete/:ids", controller.Link.Delete)
		//group.PUT("/status", controller.Link.Status)
	})

	/* Item */
	s.Group("item", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Item.List)
		//group.POST("/add", controller.Item.Add)
		//group.PUT("/update", controller.Item.Update)
		//group.DELETE("/delete/:ids", controller.Item.Delete)
		//group.PUT("/status", controller.Item.Status)
		//group.GET("/getItemList", controller.Item.GetItemList)
	})

	/* Item Category */
	s.Group("itemcate", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.ItemCate.List)
		//group.POST("/add", controller.ItemCate.Add)
		//group.PUT("/update", controller.ItemCate.Update)
		//group.DELETE("/delete/:ids", controller.ItemCate.Delete)
		////group.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
		//group.GET("/getCateList", controller.ItemCate.GetCateList)
	})

	/* Advertising Sorting */
	s.Group("adsort", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.AdSort.List)
		//group.POST("/add", controller.AdSort.Add)
		//group.PUT("/update", controller.AdSort.Update)
		//group.DELETE("/delete/:ids", controller.AdSort.Delete)
		//group.GET("/getAdSortList", controller.AdSort.GetAdSortList)
	})

	/* Advertising */
	s.Group("ad", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Ad.List)
		//group.POST("/add", controller.Ad.Add)
		//group.PUT("/update", controller.Ad.Update)
		//group.DELETE("/delete/:ids", controller.Ad.Delete)
		//group.PUT("/status", controller.Ad.Status)
	})

	/* Notice */
	s.Group("notice", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Notice.List)
		//group.POST("/add", controller.Notice.Add)
		//group.PUT("/update", controller.Notice.Update)
		//group.DELETE("/delete/:ids", controller.Notice.Delete)
		//group.PUT("/status", controller.Notice.Status)
	})

	/* Configuration Web */
	s.Group("configweb", func(group *ghttp.RouterGroup) {
		//group.GET("/index", controller.ConfigWeb.Index)
		//group.PUT("/save", controller.ConfigWeb.Save)
	})

	/* Member Level */
	s.Group("memberlevel", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.MemberLevel.List)
		//group.POST("/add", controller.MemberLevel.Add)
		//group.PUT("/update", controller.MemberLevel.Update)
		//group.DELETE("/delete/:ids", controller.MemberLevel.Delete)
		//group.GET("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	})

	/* Member */
	s.Group("member", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Member.List)
		//group.POST("/add", controller.Member.Add)
		//group.PUT("/update", controller.Member.Update)
		//group.DELETE("/delete/:ids", controller.Member.Delete)
		//group.PUT("/status", controller.Member.Status)
	})

	/* Analysis */
	s.Group("analysis", func(group *ghttp.RouterGroup) {
		//group.GET("/index", controller.Analysis.Index)
	})

	/* Generate */
	s.Group("generate", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Generate.List)
		//group.POST("/generate", controller.Generate.Generate)
		//group.POST("/batchGenerate", controller.Generate.BatchGenerate)
	})

}
