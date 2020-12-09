## 作业

**Q:** 

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


**A:**

不应该直接wrap `sql.ErrNoRows`，而是包装一层，新增一个`ErrRecordNotFound`错误类型并wrap，然后返回给上层

#### 理由：

首先要明确：`sql`包提供了保证**SQL**或**类SQL**数据库的泛用接口，也就是说是针对关系型数据库。

如果后续服务需要扩展mongodb等NoSQL，那么需要提供一个通用的错误来保证上层逻辑的兼容性和可扩展性，因此需要封装一个`Sentinel Error`。

#### 核心思想：
- 数据访问层要wrap `ErrRecordNotFound`和其他错误
- 业务逻辑层不需要知道数据访问层查询的是哪个类型的数据库，error透传即可
- 最顶层需要记录`error`日志

## 代码


```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/pkg/errors"
)

const _selectUserByUserIDSQL = `select user_id,user_name from user where user_id = ?`

var (
    db                *sqlx.DB
    ErrRecordNotFound = errors.New("dao: record not found")
)

type User struct {
    UserID   int64  `db:"user_id"`
    UserName string `db:"user_name"`
}

func Dao(userID int64) (user *User, err error) {
    user = new(User)
    
    if err = db.Get(user, _selectUserByUserIDSQL, userID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            err = errors.Wrapf(ErrRecordNotFound, "by userID:(%d)", userID)
        } else {
            err = errors.Wrapf(err, "by userID:(%d)", userID)
        }
        return
    }
    return
}

func Biz(userID int64) (user *User, err error) {
    return Dao(userID)
}

func Http() {
    res, err := Biz(1)
    if err != nil {
        if errors.Is(err, ErrRecordNotFound) {
            log.Println("no match data:", res)
            return
        }
        log.Printf("get user err:%+v \r\n", err)
        return
    }
    fmt.Println("res:", res)
}

func main() {
    if err := initDB(); err != nil {
        log.Printf("init db failed:%+v \r\n", err)
        return
    }
    Http()
}

func initDB() (err error) {
    dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
    if db, err = sqlx.Connect("mysql", dsn); err != nil {
        return
    }
    return
}
```