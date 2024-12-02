# Operation
## ConnDB
```go
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlCon.Username, mysqlCon.Password, mysqlCon.Addr, mysqlCon.port, mysqlCon.Db, "10s")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}
    if mysqlCon.BatchSize != 0 {
        db = db.Session(&gorm.Session{CreateBatchSize: mysqlCon.BatchSize})
    }
```
## Create Table
`Gorm` 基于数据结构体创建新的数据表，表结构就是结构体结构

定义 `owner` 数据表，内部嵌套 `gorm` 模型以及自定义的三个数据类型
```go
type Owner struct {
	gorm.Model
	Address string `json:"address"`
	Owner   string `json:"owner"`
	Tokenid int64  `json:"tokenid"`
}
```
创建数据表或者基于结构体更新表结构(在旧表的基础上进行新增新的结构体参数)
```go
	// 基于定义的结构体创建数据表
	err = db.AutoMigrate(&table.Owner{}, &table.Transfer{})
	if err != nil {
		log.Printf("failed to create database, got error %v", err)
		return
	}
```
## Insert
### Insert Owner By Struct
结构体中的主键会采用默认的递增值
```go
	var email = new(string)
	*email = "string"
	var ts = new(time.Time)
	*ts = time.Now()
	owner := table.Owner{
		Address: "",
		Token:   0,
		Owner: table.User{
			Name:     "Sam",
			Email:    email,
			Age:      18,
			Birthday: ts,
		},
	}
    result := db.Create(&owner)
    log.Printf("Insert new log with ID = %d, affected rows count = %d, with err = %v", owner.ID, result.RowsAffected, result.Error)
```
### Insert Owner By Table and Map
采用 `map` 插值的方式，不会补全默认主键的值
```go
// key is the column name of the table, must equal to the struct definition 
	var owner = map[string]interface{}{
		"Name": "Sam16",
	}
    result := db.Model(&table.Owner{}).Create(owner)
	log.Printf("Insert new log within affected rows count = %d, with err = %v", result.RowsAffected, result.Error)
```
### Insert owners By Structs
`Gorm` 支持批量插入数据
```go
	var owners = make([]table.Owner, 0)
	var owner table.Owner
	var email = new(string)
	*email = "string"
	var ts = new(time.Time)
	*ts = time.Now()
	for index := 1; index < 5; index++ {
		owner = table.Owner{
			Address: "",
			Token:   int64(index),
			Owner: table.User{
				Name:     fmt.Sprintf("%s%d", "Sam", index),
				Email:    email,
				Age:      18,
				Birthday: ts,
			},
		}
		owners = append(owners, owner)
	}
    result := db.Create(&owner)
    log.Printf("Insert new log with affected rows count = %d, with err = %v", result.RowsAffected, result.Error)
```
`CreateInBatches` 支持传入结构体数组和单次限额执行数据的批量插入
- `CreateInBatches(owner, 10)`
  - 单次插入仅操作 `10` 条数据
  - 循环多次，每次操作 `10` 条，直到将全部 `owner` 数据插入到库
```go
func BatchInsertOwners(db *gorm.DB, owner []table.Owner, limit int) {
	result := db.CreateInBatches(owner, limit)
	log.Printf("Insert new log with affected rows count = %d, with err = %v", result.RowsAffected, result.Error)
}
```
单次操作限额可以加在连接上，直接在此次连接中进行限制
```go
db := db.Session(&gorm.Session{CreateBatchSize: 1000})
```

