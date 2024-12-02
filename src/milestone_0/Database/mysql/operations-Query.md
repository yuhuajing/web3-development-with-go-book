# Operation
## QueryDB
根据条件查询的三种方式，当条件无效时会返回 `ErrRecordNotFound` 错误
- First
  - `SELECT * FROM users ORDER BY id LIMIT 1;`
- Take
  - `SELECT * FROM users LIMIT 1;`
- Last
  - `SELECT * FROM users ORDER BY id DESC LIMIT 1;`



```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlCon.Username, mysqlCon.Password, mysqlCon.Addr, mysqlCon.port, mysqlCon.Db, "10s")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}
```
