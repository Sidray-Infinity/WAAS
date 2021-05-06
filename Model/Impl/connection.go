package Impl

import "sync"

var address string = "root:root@tcp(127.0.0.1:3306)/waas?charset=utf8&parseTime=True&loc=Local"
var err error
var walletMutx = make(map[int]*sync.Mutex)
