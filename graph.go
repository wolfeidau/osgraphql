package osgraphql

import (
	"log"

	"github.com/graphql-go/graphql"
)

var (
	// PsSchema graphql schema for osgraphql
	PsSchema graphql.Schema

	systemInfo = NewLocalSystemInfo()

	cpuType       *graphql.Object
	partitionType *graphql.Object
	processType   *graphql.Object
)

func init() {
	cpuType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Cpu",
		Description: "CPU in a machine",
		Fields: graphql.Fields{
			"cpu": &graphql.Field{
				Type:        graphql.String,
				Description: "The index of the CPU.",
			},
			"vendorId": &graphql.Field{
				Type:        graphql.String,
				Description: "The vendor ID of the CPU.",
			},
			"family": &graphql.Field{
				Type:        graphql.String,
				Description: "The family of the CPU.",
			},
			"model": &graphql.Field{
				Type:        graphql.String,
				Description: "The model of the CPU.",
			},
			"modelName": &graphql.Field{
				Type:        graphql.String,
				Description: "The model name of the CPU.",
			},
		},
	})

	processType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Process",
		Description: "A process running on a machine",
		Fields: graphql.Fields{
			"pid": &graphql.Field{
				Type:        graphql.Int,
				Description: "The pid of process.",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of process.",
			},
			"rss": &graphql.Field{
				Type:        graphql.Int,
				Description: "The rss of process.",
			},
			"vms": &graphql.Field{
				Type:        graphql.Int,
				Description: "The vms of process.",
			},
			"swap": &graphql.Field{
				Type:        graphql.Int,
				Description: "The swap of process.",
			},
		},
	})

	partitionType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Partition",
		Description: "A partition mounted on the host.",
		Fields: graphql.Fields{
			"path": &graphql.Field{
				Type:        graphql.String,
				Description: "The path of the partition.",
			},
			"fstype": &graphql.Field{
				Type:        graphql.String,
				Description: "The filesystem type of the partition.",
			},
			"total": &graphql.Field{
				Type:        graphql.String,
				Description: "The total space of the partition.",
			},
			"free": &graphql.Field{
				Type:        graphql.String,
				Description: "The free space of the partition.",
			},
			"usedPercent": &graphql.Field{
				Type:        graphql.Int,
				Description: "The used percentage of the partition.",
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "viewer",
		Fields: graphql.Fields{
			"cpus": &graphql.Field{
				Type: graphql.NewList(cpuType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return systemInfo.GetCPUInfo()
				},
			},
			"partitions": &graphql.Field{
				Type: graphql.NewList(partitionType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return systemInfo.GetPartitions(true)
				},
			},
			"processes": &graphql.Field{
				Type: graphql.NewList(processType),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					nameQuery, ok := params.Args["name"].(string)

					if ok {
						return systemInfo.GetProcessesByName(nameQuery)
					}

					return systemInfo.GetProcesses()
				},
			},
		},
	})

	var err error

	PsSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("schema created.")
}

// ExecuteQuery execute a graphql query
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
