package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/programmingpercy/gopheragency/gopher"
)

// modifyJobArgs are arguments available for the modifyJob Mutation request
var modifyJobArgs = graphql.FieldConfigArgument{
	"employeeid": &graphql.ArgumentConfig{
		// Create a string argument that cannot be empty
		Type: graphql.NewNonNull(graphql.String),
	},
	"jobid": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	// The new start date to apply if set
	"start": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	// The new end date to apply if set
	"end": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// generateRootMutation will create the root mutation object
func generateRootMutation(gs *gopher.GopherService) *graphql.Object {

	mutationFields := graphql.Fields{
		// Create a mutation named modifyJob which accepts a JobType
		"modifyJob": generateGraphQLField(jobType, gs.MutateJobs, "Modify a job for a gopher", modifyJobArgs),
	}
	mutationConfig := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}

	return graphql.NewObject(mutationConfig)
}
