package gopher

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/programmingpercy/gopheragency/job"
)

type Resolver interface {
	// ResolveGophers should return a list of all gophers in the repository
	ResolveGophers(p graphql.ResolveParams) (interface{}, error)
	// ResolveGopher is used to respond to single queries for gophers
	ResolveGopher(p graphql.ResolveParams) (interface{}, error)
	// ResolveJobs is used to find Jobs
	ResolveJobs(p graphql.ResolveParams) (interface{}, error)
}

// GopherService is the service that holds all repositories
type GopherService struct {
	gophers Repository
	// Jobs are reachable by the Repository
	jobs job.Repository
}

// NewService is a factory that creates a new GopherService
func NewService(repo Repository, jobrepo job.Repository) GopherService {
	return GopherService{
		gophers: repo,
		jobs:    jobrepo,
	}
}

// ResolveGophers will be used to retrieve all available Gophers
func (gs GopherService) ResolveGophers(p graphql.ResolveParams) (interface{}, error) {
	// Fetch gophers from the Repository
	gophers, err := gs.gophers.GetGophers()
	if err != nil {
		return nil, err
	}
	return gophers, nil
}

// ResolveGopher is used to find a single gopher using the id argument
func (gs GopherService) ResolveGopher(p graphql.ResolveParams) (interface{}, error) {
	// show them if cast to int instead how error ir returned in message
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("id has to be a string")
	}
	// Call the Repository for the ID
	gopher, err := gs.gophers.GetGopher(id)
	if err != nil {
		return nil, err
	}
	return gopher, nil
}

// ResolveJobs is used to find all jobs related to a gopher
func (gs *GopherService) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {
	// Fetch Source Value
	g, ok := p.Source.(Gopher)

	if !ok {
		return nil, errors.New("source was not a Gopher")
	}
	// Here we extract the Argument Company
	company := ""
	if value, ok := p.Args["company"]; ok {
		company, ok = value.(string)
		if !ok {
			return nil, errors.New("id has to be a string")
		}
	}

	// Find Jobs Based on the Gophers ID
	jobs, err := gs.jobs.GetJobs(g.ID, company)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// MutateJobs is used to modify jobs based on a mutation request
// Available params are
// employeeid! -- the id of the employee, required
// jobid! -- job to modify, required
// start -- the date to set as start date
// end -- the date to set as end
func (gs *GopherService) MutateJobs(p graphql.ResolveParams) (interface{}, error) {
	employee, err := grabStringArgument("employeeid", p.Args, true)
	if err != nil {
		return nil, err
	}
	jobid, err := grabStringArgument("jobid", p.Args, true)
	if err != nil {
		return nil, err
	}
	start, err := grabStringArgument("start", p.Args, false)
	if err != nil {
		return nil, err
	}
	end, err := grabStringArgument("end", p.Args, false)
	if err != nil {
		return nil, err
	}

	// Get the job
	job, err := gs.jobs.GetJob(employee, jobid)
	if err != nil {
		return nil, err
	}
	// Modify start and end date if they are set
	if start != "" {
		job.Start = start
	}

	if end != "" {
		job.End = end
	}
	// Update with new values
	return gs.jobs.Update(job)
}

// grabStringArgument is used to grab a string argument
func grabStringArgument(k string, args map[string]interface{}, required bool) (string, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}
