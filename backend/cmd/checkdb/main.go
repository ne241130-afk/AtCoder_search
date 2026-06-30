package main

import (
  "fmt"
  "github.com/ne241130/atcoder-learning-hub/backend/internal/database"
  "github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

func main() {
  if err := database.Connect(); err != nil {
    panic(err)
  }
  fmt.Println("has table model:", database.DB.Migrator().HasTable(&model.Problem{}))
  fmt.Println("has table name:", database.DB.Migrator().HasTable("problems"))
  var count int64
  if err := database.DB.Model(&model.Problem{}).Count(&count).Error; err != nil {
    panic(err)
  }
  fmt.Println("count:", count)
  rows, err := database.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='public'").Rows()
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil { panic(err) }
    fmt.Println("table:", name)
  }
}
