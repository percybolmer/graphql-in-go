package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/programmingpercy/gopheragency/gopher"
)

// We can initialize Objects like this unless they need a special resolver
var jobType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Job",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"employeeID": &graphql.Field{
			Type: graphql.String,
		},
		"company": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"start": &graphql.Field{
			Type: graphql.String,
		},
		"end": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)

// GenerateSchema will create a GraphQL Schema and set the Resolvers found in the GopherService
// For all the needed fields
func GenerateSchema(gs *gopher.GopherService) (*graphql.Schema, error) {
	gopherType := generateGopherType(gs)
	// RootQuery
	fields := graphql.Fields{
		// We define the Gophers query
		"gophers": &graphql.Field{
			// It will return a list of GopherTypes, a List is an Slice
			Type: graphql.NewList(gopherType),
			// We change the Resolver to use the gopherRepo instead, allowing us to access all Gophers
			Resolve: gs.ResolveGophers,
			// Description explains the field
			Description: "Query all Gophers",
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// RootMutation
	rootMutation := generateRootMutation(gs)

	// Now combine all Objects into a Schema Configuration
	schemaConfig := graphql.SchemaConfig{
		// Query is the root object query schema
		Query: graphql.NewObject(rootQuery),
		// Appliy the Mutation to the schema
		Mutation: rootMutation,
	}
	// Create a new GraphQL Schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

// generateGraphQLField is a generic builder factory to create graphql fields
func generateGraphQLField(output graphql.Output, resolver graphql.FieldResolveFn, description string, args graphql.FieldConfigArgument) *graphql.Field {
	return &graphql.Field{
		Type:        output,
		Resolve:     resolver,
		Description: description,
		Args:        args,
	}
}

// genereateGopherType will assemble the Gophertype and all related fields
func generateGopherType(gs *gopher.GopherService) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Gopher",
		// Fields is the field values to declare the structure of the object
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID that is used to identify unique gophers",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the gopher",
			},
			"hired": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "True if the Gopher is employeed",
			},
			"profession": &graphql.Field{
				Type:        graphql.String,
				Description: "The gophers last/current profession",
			},
			// Here we create a graphql.Field which is depending on the jobs repository, notice how the Gopher struct does not contain any information about jobs
			// But this still works
			"jobs": generateJobsField(gs),
		}},
	)
}

// generateJobsField will build the GraphQL Field for jobs
func generateJobsField(gs *gopher.GopherService) *graphql.Field {
	return &graphql.Field{
		// Return a list of Jobs
		Type:        graphql.NewList(jobType),
		Description: "A list of all jobs the gopher had",
		Resolve:     gs.ResolveJobs,
		// Args are the possible arguments.
		Args: graphql.FieldConfigArgument{
			"company": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
	}
}
