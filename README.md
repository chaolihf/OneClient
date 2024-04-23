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

# 运行
需要将发布包下面的resource，release目录拷贝到此项目的release目录下。如果出现异常情况下可以将debug目录下文件拷贝过来，这样可以报一些异常