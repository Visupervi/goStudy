记录学习go,如何从入门到放弃

go如何拉取第三方的包?

```shell
go get github.com/garyburd/redigo/redis

go get github.com/google/uuid   

go get -u github.com/go-sql-driver/mysql
```
ps: 要去项目的gopath下面下载


go web开发

go 服务端开发

切片

协程

管道

错误补货

switch

for/for range


bookStore goWeb项目

麻雀虽小五脏俱全，增、删、改、查都有。

// 删除切片的其中一个
cart.Items = append(cart.Items[:k], cart.Items[k+1:]...)
