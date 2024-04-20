# 说明
CorePlugins.go： 核心插件，核心插件不允许删除
CustomPlugins.go： 自定义插件，可以根据需求自行添加插件
# windows编译安装
go install github.com/akavel/rsrc@latest
rsrc -manifest main.manifest -ico client.ico -o main.syso
