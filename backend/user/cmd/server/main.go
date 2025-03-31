package main

import (
	"fmt"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/db"
	"github.com/pkg/errors"
)

func main() {
	// データベース接続の初期化
	if _, err := db.InitDB(); err != nil {
		fmt.Println(errors.Wrap(err, "failed to initialize database"))
		return
	}
	_, err := fmt.Println("hello world")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to print"))
	}
}
