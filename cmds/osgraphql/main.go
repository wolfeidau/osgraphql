package main

import (
	"encoding/json"

	"github.com/Unknwon/macaron"
	"github.com/alecthomas/kingpin"
	"github.com/macaron-contrib/binding"
	"github.com/wolfeidau/osgraphql"
)

type queryReq struct {
	Query string `json:"query" binding:"Required"`
}

var (
	port = kingpin.Flag("port", "Listen port.").Short('p').Default("9000").Int()
)

func main() {

	kingpin.Parse()

	m := macaron.Classic()

	m.Use(osgraphql.CORS())

	m.Options("/graphql", func() string {
		return ""
	})

	m.Post("/graphql", binding.Json(queryReq{}), func(qr queryReq) string {
		result := osgraphql.ExecuteQuery(qr.Query, osgraphql.PsSchema)
		data, _ := json.Marshal(result)
		return string(data)
	})

	m.Run(*port)
}

var query = `
{
	cpus {
		vendorId
		family
	}
	processes {
		pid
	}
}`

// var q = `
// viewer {
// 	machines {
// 		edges {
// 			node {
// 				cpus {
// 					edges {
// 						node {
// 							vendorId
// 							family
// 						}
// 					}
// 				}
// 				processes {
// 					edges {
// 						node {
// 							pid
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
// `
