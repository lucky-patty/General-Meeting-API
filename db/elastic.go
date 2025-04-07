package db

import (
  "fmt"
  "github.com/elastic/go-elasticsearch/v8"
)

func ElasticNewClient(addr string) (*elasticsearch.Client, error) {
  fmt.Println("Elastic Address: ", addr)
  cfg := elasticsearch.Config{
    Addresses: []string{
      addr,
    },
  }

  con, err := elasticsearch.NewClient(cfg)
  if err != nil {
    fmt.Println("Error creating the client: %s", err)
    return nil, err 
  }

  res, err := con.Info()
  if err != nil {
    fmt.Println("error getting response: %s", err)
    return nil, err
  }
  defer res.Body.Close()

  fmt.Println("Elasticsearch connected")
  return con, nil 
}
