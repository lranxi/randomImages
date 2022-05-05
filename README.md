### 随机图片API

go实现的随机图片API，所有图片均存放在服务器本地。<br>
通过计算图片的dhash实现图片指纹功能，防止出现重复图片。

### feat
 - [x] 随机获取一张图片
 - [x] 指定图片尺寸
 - [ ]  图片分类获取（接口支持）
 - [ ] 图片上传与审核（接口支持）


### 技术栈
 - gin
 - redis
 - go-redis

### API文档

##### 请求方式：GET
##### URL: https://images.lyranxi.com/images?width=1024&height=768