# Hooks
## Create
`Gorm` 支持用户自定义回调 `hooks` 例如 `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate`
```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlCon.Username, mysqlCon.Password, mysqlCon.Addr, mysqlCon.port, mysqlCon.Db, "10s")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}
```