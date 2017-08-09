package controllers

import (
  "os" // local
  "net/http"
  "encoding/json"
)

type Map struct {
}

func NewMap() Map {
  return Map{}
}

// クエリのオートコンプリートメソッド
func QueryAutoComplete(place string) interface{} {
  url := "https://maps.googleapis.com/maps/api/place/queryautocomplete/json?&key=" + os.Getenv("GOOGLE_API_KEY") + "&input=" + place
  // url := "https://maps.googleapis.com/maps/api/place/queryautocomplete/json?&key=" + ENV["GOOGLE_API_KEY"] + "&input=" + place
  res, err := http.Get(url)
  defer res.Body.Close()
  if err == nil {
    var data map[string]interface{}
    error := json.NewDecoder(res.Body).Decode(&data)
    if error == nil {
      return data["predictions"]
    }
  }
  return nil
}

// 最短経路を求めるメソッド
func SearchOptiRoute(route string) interface{} {
  // url := "https://maps.googleapis.com/maps/api/directions/json?language=ja&" + route + "&key=" + ENV["GOOGLE_API_KEY"]
  url := "https://maps.googleapis.com/maps/api/directions/json?language=ja&" + route + "&key=" + os.Getenv("GOOGLE_API_KEY") //local
  // TODO:  DBに検索情報を保存する
  res, err := http.Get(url)
  defer res.Body.Close()
  if err == nil {
    var data map[string]interface{}
    error := json.NewDecoder(res.Body).Decode(&data)
    if error == nil {
      return data["routes"]
    }
  }
  return nil
}
