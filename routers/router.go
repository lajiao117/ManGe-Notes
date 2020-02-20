package routers

import (
	"man/ManNotes/controllers"
	"github.com/astaxie/beego"
)

func init() {

    //pg 页面请求
    beego.Router("/", &controllers.PGController{},"get:LoginPG")//登录页面
    beego.Router("/index", &controllers.PGController{},"get:IndexPG")//主页
    beego.Router("/mdeditor", &controllers.PGController{},"get:MdEditorPG")//MD编辑器
    beego.Router("/home", &controllers.PGController{},"get:HomePG")//首页
    beego.Router("/tool", &controllers.PGController{},"get:ToolPG")//工具页面

    //登录注册
    beego.Router("/userreg", &controllers.LoginController{},"post:UserRegistered")//用户注册
    beego.Router("/userlogin", &controllers.LoginController{},"post:UserLogin")//用户登录
    beego.Router("/outlogin", &controllers.LoginController{},"get:OutLogin")//退出登录

    //笔记本
    beego.Router("/cnotes", &controllers.NotesController{}, "post:CreateNotes")//创建笔记本
    beego.Router("/noteslist", &controllers.NotesController{}, "get:GetNotesList")//获取当前笔记本列表

    //MD笔记
    beego.Router("/cmd", &controllers.MDController{}, "post:CreateMD")//创建MD笔记
    beego.Router("/allnote", &controllers.MDController{}, "get:GetAllNote")//获取所有笔记
    beego.Router("/notes/:id:int", &controllers.MDController{}, "get:NotesMDList")//获取笔记本对应的笔记  
    beego.Router("/noteshow/:mdid:*", &controllers.MDController{}, "get:MDShow")//显示MD笔记内容
    beego.Router("/editmdnote/:mdid:*", &controllers.MDController{}, "get:MDEditPG")//修改MD笔记内容编辑页面
    beego.Router("/modifynote/:mdid:*", &controllers.MDController{}, "post:MDNoteModify")//提交修改笔记内容
    beego.Router("/searchnote", &controllers.MDController{}, "get:SearchNote")//搜索笔记
    beego.Router("/delnote/:mdid:*", &controllers.MDController{}, "get:DelNote")//删除笔记到回收站
    beego.Router("/recycler", &controllers.MDController{}, "get:NoteRecycler")//回收站
    beego.Router("/rnoteshow/:mdid:*", &controllers.MDController{}, "get:RMDShow")//显示回收站MD笔记内容
    beego.Router("/schen/:mdid:*", &controllers.MDController{}, "get:SchenNote")//笔记永久删除
    beego.Router("/restore/:mdid:*", &controllers.MDController{}, "get:RestoreNote")//笔记永久删除

    //收藏（链与工具）
    beego.Router("/collectlink", &controllers.TandLController{}, "post:AddCollectLink")//添加网络资源
    beego.Router("/tandllist", &controllers.TandLController{}, "get:GetTandL")//获取网络工具列表
}
