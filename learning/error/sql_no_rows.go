package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var NotFound = errors.New("data not found")
var notFoundCode = 40001
var systemErr = 50001

type rowData struct {
	id   int
	name string
}

func mockError() error {
	return sql.ErrNoRows
	//return sql.ErrTxDone
}

/**
解答：
1. dao 层：如果是 sql.ErrNoRows 抛出业务可识的错误；否则则抛出系统错误
2. 业务层：针对系统错误，抛出对应的问题；如果是 origin ErrNoRows 的错误，则根据业务场景处理
*/
func main() {
	err := biz1()
	if err != nil {
		log.Println(err)
	}

	err = biz2()
	if err != nil {
		log.Println(err)
	}

	err = biz3()
	if err != nil {
		log.Println(err)
	}
}

/**
dao 采用 opaque error，利用错误码来断定是什么错误
biz 可以使用方法来判定是否属于某个错误，灵活性更好
*/
func biz3() error {
	_, err := dao3("select * from user")

	if isNoRow(err) {
		// 站在业务角度考虑：
		// 1. 如果找不到是正常的，则返回 nil
		// 2. 如果此时应该找到，可以转为另一个错误，或返回错误响应
		return nil
	}

	if err != nil {
		// 可以转为业务领域错误，也可以继续上抛
		return err
	}

	return nil
}

func isNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}

func dao3(query string) (*rowData, error) {
	err := mockError()

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("%d, not found", notFoundCode)
	}

	if err != nil {
		return nil, fmt.Errorf("%d, not found", systemErr)
	}
	// do something
	return nil, nil
}

/**
dao 只转换了第三方的 NoRows错误，转为自身的 NoFound。其余错误则
是附加上了必要的 DEBUG 信息
biz 使用 errors.Is 来判断
*/
func biz2() error {
	_, err := dao2("select * from user")

	if errors.Is(err, NotFound) {
		// 站在业务角度考虑：
		// 1. 如果找不到是正常的，则返回 nil
		// 2. 如果此时应该找到，可以转为另一个错误，或返回错误响应
		return nil
	}

	if err != nil {
		// 可以转为业务领域错误，也可以继续上抛
		return err
	}

	return nil
}

func dao2(query string) (*rowData, error) {
	err := mockError()

	if err == sql.ErrNoRows {
		return nil, errors.Wrap(NotFound, fmt.Sprintf("data not found: sql: %s", query))
	}

	if err != nil {
		// 不细究具体问题，直接抛错
		return nil, errors.Wrap(err, fmt.Sprintf("db query system error: sql: %s", query))
	}

	return nil, nil
}

/**
dao 任意错误包一下，转为 NotFound，就给 biz，
biz 也懒得判断具体是什么错误
*/
func biz1() error {
	_, err := dao1("select * from user")

	if err != nil {
		// 可以转为业务领域错误，也可以继续上抛
		return err
	}

	return nil
}

func dao1(query string) (*rowData, error) {
	err := mockError()

	if err != nil {
		// 不细究具体问题，直接抛错
		return nil, errors.Wrap(NotFound, fmt.Sprintf("data not found: sql: %s", query))
	}

	return nil, nil
}
