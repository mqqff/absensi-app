package main

import "github.com/mqqff/absensi-app/internal/bootstrap"

func main() {
	//hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	//fmt.Println(string(hash))
	bootstrap.Init()
}
