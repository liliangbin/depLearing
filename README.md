# depLearing
go后端的开发。
- 来开始  搞事情  搞事情  搞事情
##### 使用的库 
- "github.com/codegangsta/negroni" ====>中间件的一个库。
- "github.com/go-sql-driver/mysql" ====> MySQL的库
- 	"github.com/gorilla/mux"  =====> 这个是一个路由管理的库
##### 目录结构

#### 使用说明。使用go的官方依赖管理工具dep。(用着的感受是真的垃圾)
- 加载依赖，dep ensure

#### 如何使用gorm
```go
//创建一个数据
type Animal struct {
    ID   int64
    Name string `gorm:"default:'galeone'"`
    Age  int64
}

var animal = Animal{Age: 99, Name: ""}
db.Create(&animal)
// INSERT INTO animals("age") values('99');
// SELECT name from animals WHERE ID=111; // 返回主键为 111
// animal.Name => 'galeone'



//查询

// 获取第一条记录，按主键排序
db.First(&user)
//// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取最后一条记录，按主键排序
db.Last(&user)
//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 获取所有记录
db.Find(&users)
//// SELECT * FROM users;

// 使用主键获取记录
db.First(&user, 10)
//// SELECT * FROM users WHERE id = 10;


// 获取第一个匹配记录
db.Where("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

// 获取所有匹配记录
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';

db.Where("name <> ?", "jinzhu").Find(&users)


//更新
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)


// 删除存在的记录
db.Delete(&email)
//// DELETE from emails where id=10;


//
db.Where("age = ?", 20).Delete(&User{})




```

- 更多的gorm操作细节请查看http://gorm.book.jasperxu.com/crud.html#u

##### 现在主要是加上了jwt验证。用户的权限认证这个我们可以后期来讨论。"自己是实现一套权限认证小框架"