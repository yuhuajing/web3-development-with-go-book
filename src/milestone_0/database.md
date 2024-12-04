# Database
## mysql
### mysql Docker
下载 `mysql` `Docker` 镜像
```shell
docker pull mysql:latest
```
`Docker` 启动本地 `mysql` 容器
```shell
docker run -p 3306:3306 --name root -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```
进入 `Docker` 容器，通过用户名和密码操作数据库
```shell
docker exec -it root bash
```
输入用户名/密码，执行 `mysql` 的命令行代码创建数据库 `create database XXX;`
```shell
mysql -u root -p123456

create database eventLog;
```
### Gorm
`Gorm` 通过 `golang` 结构体简化数据库建表和表查询的过程
- `Gorm` `Golang` 结构体能够包含的数据类型
    - `Golang` 基础类型，`uint,int,bool,string`
    - 指针,`*string`
    - 类型别名
    - 自定义类型
- 结构体参数中首字母大写表示支持导出该表字段，否则该值不支持导出
```go
type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
  Birthday     *time.Time     // A pointer to time.Time, can be null
  MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
  ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
  CreatedAt    time.Time      // Automatically managed by GORM for creation time
  UpdatedAt    time.Time      // Automatically managed by GORM for update time
  ignored      string         // fields that aren't exported are ignored
}
```
### gorm.Model
`GORM` 有预定义的数据类型，包含 `PrimaryKey` 主键、数据被创建、更新和软删除时间
```go
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time // Set to current time if it is zero on creating
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

}
```
自定义时间：
```go
  Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
  Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
```
`gorm` 关于结构体数据的权限设置：
```go
type User struct {
  Name string `gorm:"<-:create"` // allow read and create
  Name string `gorm:"<-:update"` // allow read and update
  Name string `gorm:"<-"`        // allow read and write (create and update)
  Name string `gorm:"<-:false"`  // allow read, disable write permission
  Name string `gorm:"->"`        // readonly (disable write permission unless it configured)
  Name string `gorm:"->;<-:create"` // allow read and create
  Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
  Name string `gorm:"-"`            // ignore this field when write and read with struct
  Name string `gorm:"-:all"`        // ignore this field when write, read and migrate with struct
  Name string `gorm:"-:migration"`  // ignore this field when migrate with struct
}
```
### GormEmbeddedStruct
`GORM` 数据结构体支持多结构体嵌套
```go
type Author struct {
  Name  string
  Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// equals
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```
`embeddedPrefix` 为嵌入的结构体设置统一的前缀：
```go
type Blog struct {
  ID      int
  Author  Author `gorm:"embedded;embeddedPrefix:author_"`
  Upvotes int32
}
// equals
type Blog struct {
  ID          int64
  AuthorName  string
  AuthorEmail string
  Upvotes     int32
}
```
### Mysql Preference:

[MySQLGormCase](https://github.com/yuhuajing/MySQLGormCase)

[StoreChainTxsInMySQL](https://github.com/yuhuajing/StoreChainTransactionInMySQL)

[GormDoc](https://gorm.io/docs/)

## MongoDB
### MongoDB Docker
下载 `mongo` `Docker` 镜像
```shell
docker pull mongo:latest
```
`Docker` 启动本地 `mysql` 容器
```shell
docker run -d -p 27017:27017 --name some-mongo \
	-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
	-e MONGO_INITDB_ROOT_PASSWORD=secret \
	mongo:latest
```

### Download
[https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)

勾选 `MongoDB Compass`工具，可视化数据库数据

### MongoDB
MongoDB 将数据按照 golang 结构体定义存储成 document,并汇集在集合表中
```go
db.collection.insertOne()
db.collection.insertMany()
db.collection.updateOne() 与 upsert: true 选项一起使用时
db.collection.updateMany() 与 upsert: true 选项一起使用时
db.collection.findAndModify() 与 upsert: true 选项一起使用时
db.collection.findOneAndUpdate() 与 upsert: true 选项一起使用时
db.collection.findOneAndReplace() 与 upsert: true 选项一起使用时
db.collection.bulkWrite()
```
### MongoDB Preference:
[MongoDBCase](https://github.com/yuhuajing/MongoDBCase)

[MongoDoc](https://www.mongodb.com/zh-cn/docs/manual/reference/insert-methods/)


