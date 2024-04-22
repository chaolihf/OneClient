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