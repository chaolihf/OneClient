openssl req -new -newkey rsa:2048 -sha256 -nodes -out server.csr -keyout server.key -subj "/C=CN/ST=Anhui/L=Hefei/O=Telecom/OU=IT/CN=127.0.0.1"
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt -extfile http.ext
rem 在管理员权限下或者手工打开证书设置，在根目录下导入server.crt文件
rem certutil -addstore -f "Root" server.crt
