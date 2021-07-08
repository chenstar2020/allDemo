package main

import "fmt"

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string)([]byte, error)

func (f GetterFunc)Get(key string)([]byte, error){
	return f(key)
}

func main() {
	var f1 = GetterFunc(func(key string) ([]byte, error) {
		fmt.Println("Getter", key)
		return nil, nil
	})

	f1("key1")
	f1.Get("key2")

	res := GetFromSource(GetterFunc(func(key string)([]byte, error){
		return []byte(key + "star"), nil
	}), "abcdd")

	fmt.Println("res:", res)

	GetFromSource(new(DB), "hello")
}

func GetFromSource(getter Getter, key string)[]byte{
	buf, err := getter.Get(key)
	if err == nil{
		return buf
	}
	return nil
}

type DB struct{  //DB实现Getter接口
	url string
}

func (db *DB)Query(sql string, args ...string) string{
	return "hello"
}

func (db *DB)Get(key string)([]byte, error){
	v := db.Query("SELECT name FROM TABLE WHERE name = ?", key)
	return []byte(v), nil
}



