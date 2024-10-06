# 说明
CorePlugins.go： 核心插件，核心插件不允许删除
CustomPlugins.go： 自定义插件，可以根据需求自行添加插件
# windows编译安装
go install github.com/akavel/rsrc@latest
rsrc -manifest main.manifest -ico client.ico -o main.syso
# 引用
探测部分主要代码来源于blackbox项目(https://github.com/prometheus/blackbox_exporter)，适当去除普米相关的代码

# 调试
启动VSCode的时候需要使用管理员权限模式运行

# windows下安装
1. 下载msys2安装包，解压到任意目录
2. pacman -Syu
3. pacman -Su
4. pacman -S --needed base-devel mingw-w64-x86_64-toolchain
如果出现无法找到atomic_load等错误时，可以尝试使用visual studio 2019+来进行编译
# 运行
网站下载https://cef-builds.spotifycdn.com/index.html版本包，将发布包下面的resource，release目录拷贝到此项目的release目录下。如果出现异常情况下可以将debug目录下文件拷贝过来，这样可以报一些异常

# 其他说明
## 引入并修改https://github.com/deluan/bring，实现对guacd的代理访问和渲染
## cefclient编译和调试
参考cef版本包中的CMakeLists.txt的说明说明进行安装
安装Visual Studio 2022，选择desktop开发，选择 C++
进入到版本目录
mkdir build 
cd build
cmake -G "Visual Studio 17" -A x64 ..
使用vs2022打开build\cef.sln，选择生成>生成解决方案
如果需要调试，可以在cefclient_win.cc中RunMain方法后增加Sleep(20000);
使用调试菜单>附加到进程，选择cefclient.exe，设置断点进行调试

## 增加一个处理器
在cef_base中声明一个结构

typedef struct _render_handler {
	cef_render_handler_t handler;
	atomic_int ref_count;
} render_handler;

IMPLEMENT_REFCOUNTING(render_handler)
GENERATE_CEF_BASE_INITIALIZER(render_handler)

在initialize_cef_base中增加
struct _render_handler;
struct _render_handler*: initialize_render_handler_base, \
void initialize_render_handler_base(struct _render_handler *object);


增加一个cef_render_handler.h文件
增加对原始头文件的引用
#include "include/capi/cef_render_handler_capi.h"
拷贝cef_render_handler_capi中cef_render_handler_t中的方法到cef_render_handler.h中
如声明为如下的函数
struct _cef_accessibility_handler_t*(CEF_CALLBACK* get_accessibility_handler)(
      struct _cef_render_handler_t* self);
修改为
struct _cef_accessibility_handler_t* CEF_CALLBACK get_accessibility_handler(
      struct _cef_render_handler_t* self){

}
在main_win.c中增加
#include "cefapi/cef_render_handler.h"
声明全局变量
render_handler *g_render_handler;
在startCef中进行初始化
g_render_handler=initialize_cef_render_handler();
在cef_client.h中增加
extern render_handler *g_render_handler;
在获取实例方法处实现相关方法
struct _cef_render_handler_t* CEF_CALLBACK get_render_handler(
        struct _cef_client_t* self) {
    DEBUG_CALLBACK("get render handler\n");
    cef_render_handler_t * handler=(cef_render_handler_t *)g_render_handler;
    handler->base.add_ref((cef_base_ref_counted_t *)g_render_handler);
    return handler;
}

## 服务端证书
需要使用cmd/uiExperienceService/generateKey.bat生成证书，注意其中需要在http.ext中配置证书绑定的域名或IP地址，否则导入windows下的根证书也是没有用的，仍然显示不安全的网站。
在管理员权限下执行certutil -addstore -f "Root" server.crt
或者手工打开证书设置，在根目录下导入server.crt文件

