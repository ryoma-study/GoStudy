package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	// quickStart()

	//len10()
	// 观察 used_memory（PS：由于通过 GO 程序和直接使用 redis-cli 命令读取的 info memory 不一致，这里手动比对)
	// 运行之前：1295792 -> 2226800：增长 931008，预计每个 key 93 字节

	//len20()
	// 3495808：增长 1269008；预计每个 key 126

	//len40()
	// 4935808：增长 1440000：预计每个 key 144

	//len45()
	// 6957952：增长 2022144：预计每个 key 202

	//key16value30()
	// 8237952：增长 1280000：预计每个 key 128

	//key16value40()
	// 9517952：增长 1280000：预计每个 key 128

	//key16value43()
	// 11322240：增长 1804648：预计每个 key 180

	key16value42()
	// 12602240：增长 1280000：预计每个 key 128

	// 可以看出，对于 string 类型，低于 43 字节的 value 占用的是 64 字节，一旦超过，有近 50 字节的增长
}

func key16value42() {
	for i := 0; i < 10000; i++ {
		key := "000000" + "0703-3" + (fmt.Sprintf("%04d", i))
		value := "00000000000000000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func key16value43() {
	for i := 0; i < 10000; i++ {
		key := "000000" + "0703-2" + (fmt.Sprintf("%04d", i))
		value := "000000000000000000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func key16value40() {
	for i := 0; i < 10000; i++ {
		key := "000000" + "0703-1" + (fmt.Sprintf("%04d", i))
		value := "000000000000000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func key16value30() {
	for i := 0; i < 10000; i++ {
		key := "000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		value := "00000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// 10000 个：key 长度为 45，value 长度为 45
func len45() {
	for i := 0; i < 10000; i++ {
		key := "00000000000000000000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, key, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// 10000 个：key 长度为 40，value 长度为 40
func len40() {
	for i := 0; i < 10000; i++ {
		key := "000000000000000000000000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, key, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// 10000 个：key 长度为 20，value 长度为 20
func len20() {
	for i := 0; i < 10000; i++ {
		key := "0000000000" + "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, key, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// 10000 个：key 长度为 10，value 长度为 10
func len10() {
	for i := 0; i < 10000; i++ {
		key := "0703-0" + (fmt.Sprintf("%04d", i))
		err := rdb.Set(ctx, key, key, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func quickStart() {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
