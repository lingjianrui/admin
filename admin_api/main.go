package main

import (
	"backend/api"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
func main() {
	//util.Query()
	server := &api.Server{}
	server.Initialize("mysql", "root", "xiaohei", "3306", "localhost", "backend")
	server.Run("0.0.0.0:8811")
}
