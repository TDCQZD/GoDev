# BoltDB
bolt是一个纯go语言实现的键值数据库，支持完全的ACID事务操作，尽管不像SQLite那样有完善的查询语言，但是接口简单易用。bolt本身通过使用一个内存映射的磁盘文件来管理数据，逻辑清晰，接口简单易用。
## Installing
```
$ go get github.com/boltdb/bolt/
```
## 打开数据库 
### bolt.Open()
```
db, err := bolt.Open("my.db", 0600, nil)
```
请注意，Bolt在数据文件上获取文件锁定，因此多个进程无法同时打开同一个数据库。
 ### 设置超时时间

 打开一个已经Open的Bolt数据库会导致当前进程挂起直到其他进程关闭该Bolt数据库。(栈特性)

 为了避免无限等待你可以在打开数据库文件的时候设置一个超时时间。
```
db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})//设置超时时间
```
### 只读Bolt数据库
使用Options.ReadOnly创建一个共享的只读Bolt数据库。
```
db, err := bolt.Open("my.db", 0666, &bolt.Options{ReadOnly: true})
if err != nil {
	log.Fatal(err)
}
```
只读模式使用共享锁来允许多个进程从数据库读取，但它将阻止任何进程以读写模式打开数据库。
## 事务-Transactions
Bolt数据库同时只支持一个read-write transaction或者多个read-only transactions。

- 每个独立的transaction以及在这个transaction中创建的所有对象（buckerts，keys等）都不是线程安全(thread safe)的。如果要在多个routine中处理数据，那么必须在每个routine中单独使用一个transaction或者显式的使用lock以确保在每个时刻只有一个routine访问这个transaction.

- read-only的transaction和read-write的transaction不应该相互之间有依赖，一般来说在同一个goroutine中不要同时打开这两种transaction，因为read-write transaction需要周期性的re-map数据文件，但是由于read-only transaction打开导致read-write transaction的re-map操作无法进行造成死锁。
### DB.Update() 启动读写事务
通过DB.UPdate()打开一个read-write transaction.
```
err := db.Update(func(tx *bolt.Tx) error {
	...
	return nil
})
```
- 在closure闭包内部，获取一个数据库的视图view。
- 在closure最后返回nil完成commit的事务操作，
- 可以在任何地方通过返回error完成rolleback的操作。
- 在read-write transaction中允许所有的数据库操作
### DB.View()启动只读事务
通过 DB.View()函数打开一个read-only transaction。
```
err := db.View(func(tx *bolt.Tx) error {
	...
	return nil
})
```
在只读事务中不允许进行变更操作。只能获取buckets，查询value，复制数据库。 
### Batch read-write transactions
每个DB.Update()都等待磁盘提交写入。通过将多个更新与DB.Batch()函数组合，可以最大限度地减少此开销：
```
err := db.Batch(func(tx *bolt.Tx) error {
	...
	return nil
})
```
- 并发批量处理会随机地组合成更大的事务。
- 批处理仅在有多个goroutine调用它时才有用。
### 手动管理事务
DB.View()和DB.UPdate()函数都包括DB.Begin()函数,这些函数都会启动事务,执行函数，然后在返回error的时候安全的关闭事务. 

DB.View()和DB.UPdate()是Bolt推荐的使用方法。

**手动启动和结束transactionn**

可以调用DB.Begin()函数来启动一个transaction然后调用Commmit或Rollback()来结束transaction. 
```
// Start a writable transaction.
tx, err := db.Begin(true)
if err != nil {
    return err
}
defer tx.Rollback()

// Use the transaction...
_, err := tx.CreateBucket([]byte("MyBucket"))
if err != nil {
    return err
}

// Commit the transaction and check for error.
if err := tx.Commit(); err != nil {
    return err
}
```

## Buckets

Buckets是bolt数据库中存放key/value对的地方。一个bucket 中的所有key必须是唯一的，可以通过DB.CreateBucket()或CreateBucketIfNotExists来创建。
### DB.CreateBucket()
```
db.Update(func(tx *bolt.Tx) error {
	b, err := tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		return fmt.Errorf("create bucket: %s", err)
	}
	return nil
})
```
### CreateBucketIfNotExists()
CreateBucketIfNotExists用来创建一个不存在的bucket，如果已经存在就不会创建。 
```
Tx.CreateBucketIfNotExists()
```
## Tx.DeleteBucket()
```
Tx.DeleteBucket()
```
###  bucket 整数自增
NextSequence()可以让Bolt确定一个序列，该序列可以用作键/值对的唯一标识符。
```
// CreateUser saves u to the store. The new user ID is set on u once the data is persisted.
func (s *Store) CreateUser(u *User) error {
    return s.db.Update(func(tx *bolt.Tx) error {
        // Retrieve the users bucket.
        // This should be created when the DB is first opened.
        b := tx.Bucket([]byte("users"))

        // Generate ID for the user.
        // This returns an error only if the Tx is closed or not writeable.
        // That can't happen in an Update() call so I ignore the error check.
        id, _ := b.NextSequence()
        u.ID = int(id)

        // Marshal user data into bytes.
        buf, err := json.Marshal(u)
        if err != nil {
            return err
        }

        // Persist bytes to users bucket.
        return b.Put(itob(u.ID), buf)
    })
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}

type User struct {
    ID int
    ...
}
```
### 嵌套 buckets
You can also store a bucket in a key to create nested buckets. The API is the same as the bucket management API on the DB object:
```
func (*Bucket) CreateBucket(key []byte) (*Bucket, error)
func (*Bucket) CreateBucketIfNotExists(key []byte) (*Bucket, error)
func (*Bucket) DeleteBucket(key []byte) error
```
```
// createUser creates a new user in the given account.
func createUser(accountID int, u *User) error {
    // Start the transaction.
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Retrieve the root bucket for the account.
    // Assume this has already been created when the account was set up.
    root := tx.Bucket([]byte(strconv.FormatUint(accountID, 10)))

    // Setup the users bucket.
    bkt, err := root.CreateBucketIfNotExists([]byte("USERS"))
    if err != nil {
        return err
    }

    // Generate an ID for the new user.
    userID, err := bkt.NextSequence()
    if err != nil {
        return err
    }
    u.ID = userID

    // Marshal and save the encoded user.
    if buf, err := json.Marshal(u); err != nil {
        return err
    } else if err := bkt.Put([]byte(strconv.FormatUint(u.ID, 10)), buf); err != nil {
        return err
    }

    // Commit the transaction.
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}
```
## key/value 键值对

### Bucket.Put() 保存
```
db.Update(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte("MyBucket"))
	err := b.Put([]byte("answer"), []byte("42"))
	return err
})
```
### Bucket.Get() 读取
```
db.View(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte("MyBucket"))
	v := b.Get([]byte("answer"))
	fmt.Printf("The answer is: %s\n", v)
	return nil
})
```
Get()函数不会返回错误，如果key存在，则返回byte slice值，如果不存在就会返回nil。
> 注意 ：从Get()返回的值仅在事务处于打开状态时有效。如果需要在事务之外使用值，则必须使用copy()将其复制到另一个字节片。
### 删除Key
```
Bucket.Delete()
```

## 迭代遍历 keys
Bolt在Bucket中以字节排序的顺序存储Key，这使得对这些键的连续迭代非常快。
### Cursor

```
db.View(func(tx *bolt.Tx) error {
	// Assume bucket exists and has keys
	b := tx.Bucket([]byte("MyBucket"))

	c := b.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}

	return nil
})
The cursor allows 
```
bolt允许Cursor移动到特定的一点，当迭代到最后一个key的时候，再次调用Next会返回nil.

**Cursor 中函数**
```
First()  Move to the first key.
Last()   Move to the last key.
Seek()   Move to a specific key.
Next()   Move to the next key.
Prev()   Move to the previous key.
```
- Each of those functions has a return signature of (key []byte, value []byte). When you have iterated to the end of the cursor then Next() will return a nil key. You must seek to a position using First(), Last(), or Seek() before calling Next() or Prev(). If you do not seek to a position then these functions will return a nil key.

如果Key是非零但value为nil，则表示Key是指Bucket而不是value。使用Bucket.Bucket()访问子Bucket。
### 前缀扫描 Prefix scans
搜索含有特定前缀的key,结合Seek和bytes.HasPrefix来实现
```
db.View(func(tx *bolt.Tx) error {
	// Assume bucket exists and has keys
	c := tx.Bucket([]byte("MyBucket")).Cursor()

	prefix := []byte("1234")
	for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}

	return nil
})
```
### 范围搜索 Range scans

```
db.View(func(tx *bolt.Tx) error {
	// Assume our events bucket exists and has RFC3339 encoded time keys.
	c := tx.Bucket([]byte("Events")).Cursor()

	// Our time range spans the 90's decade.
	min := []byte("1990-01-01T00:00:00Z")
	max := []byte("2000-01-01T00:00:00Z")

	// Iterate over the 90's.
	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		fmt.Printf("%s: %s\n", k, v)
	}

	return nil
})
```
### ForEach()
通过ForEach()来遍历bucket的所有key
```
db.View(func(tx *bolt.Tx) error {
	// Assume bucket exists and has keys
	b := tx.Bucket([]byte("MyBucket"))

	b.ForEach(func(k, v []byte) error {
		fmt.Printf("key=%s, value=%s\n", k, v)
		return nil
	})
	return nil
})
```
## 数据库备份
Bolt是一个单独的文件，因此很容易备份。 可以使用Tx.WriteTo()函数将一致的数据库视图写入writer。 如果从只读事务中调用它，它将执行热备份而不阻止其他数据库读写。

一个常见的用例是通过HTTP进行备份，以便您可以使用cURL等工具进行数据库备份：
```
func BackupHandleFunc(w http.ResponseWriter, req *http.Request) {
	err := db.View(func(tx *bolt.Tx) error {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="my.db"`)
		w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(w)
		return err
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```
```
$ curl http://localhost/backup > my.db
```
如果要备份到另一个文件，可以使用Tx.CopyFile()函数。
## 数据库 统计
数据库保持其执行的许多内部操作的运行计数，以便您可以更好地了解正在发生的事情。通过在两个时间点抓取这些统计数据的快照，我们可以看到在该时间范围内执行了哪些操作。

例如，我们可以每隔10秒启动一个goroutine来记录统计信息：
```
go func() {
	// Grab the initial stats.
	prev := db.Stats()

	for {
		// Wait for 10s.
		time.Sleep(10 * time.Second)

		// Grab the current stats and diff them.
		stats := db.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
	}
}()
```