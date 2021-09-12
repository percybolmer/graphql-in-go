// memory is a in memory data storage solution for Job
package job

import (
	"errors"
	"sync"
)

// InMemoryRepository is a storage for jobs that uses a map to store them
type InMemoryRepository struct {
	// jobs is used to store jobs
	jobs map[string][]Job
	sync.Mutex
}

// NewMemoryRepository initializes a memory with mock data
func NewMemoryRepository() *InMemoryRepository {
	jobs := make(map[string][]Job)

	jobs["1"] = []Job{
		{
			ID:         "123-123",
			EmployeeID: "1",
			Company:    "Google",
			Title:      "Logo",
			Start:      "2021-01-01",
			End:        "",
		},
	}
	jobs["2"] = []Job{
		{
			ID:         "124-124",
			EmployeeID: "2",
			Company:    "Google",
			Title:      "Janitor",
			Start:      "2021-05-03",
			End:        "",
		}, {
			ID:         "125-125",
			EmployeeID: "2",
			Company:    "Microsoft",
			Title:      "Janitor",
			Start:      "1980-03-04",
			End:        "2021-05-02",
		},
	}
	return &InMemoryRepository{
		jobs: jobs,
	}
}

// GetJobs returns all jobs for a certain Employee
func (imr *InMemoryRepository) GetJobs(employeeID, companyName string) ([]Job, error) {
	if jobs, ok := imr.jobs[employeeID]; ok {
		filtered := make([]Job, 0)
		// Filter out companyName
		for _, job := range jobs {
			// If Company Is Empty accept it, If Company matches filter accept it
			if (job.Company == companyName) || companyName == "" {
				filtered = append(filtered, job)
			}
		}
		return filtered, nil
	}
	return nil, errors.New("no such employee exist")

}

// GetJob will return a job based on the ID
func (imr *InMemoryRepository) GetJob(employeeID, jobID string) (Job, error) {
	if jobs, ok := imr.jobs[employeeID]; ok {
		for _, job := range jobs {
			// If Company Is Empty accept it, If Company matches filter accept it
			if job.ID == jobID {
				return job, nil
			}
		}
		return Job{}, errors.New("no such job exists for that employee")
	}
	return Job{}, errors.New("no such employee exist")
}

// Update will update a job and return the new state of it
func (imr *InMemoryRepository) Update(j Job) (Job, error) {
	imr.Lock()
	defer imr.Unlock()
	// Grab the employees jobs and locate the job and change the value
	if jobs, ok := imr.jobs[j.EmployeeID]; ok {
		// Find correct job
		for i, job := range jobs {
			if job.ID == j.ID {
				// Replace the whole instance by index
				imr.jobs[j.EmployeeID][i] = j
				// Return Job, we can Image changes from input Job, like CreateJob which will generate an ID etc etc.
				return j, nil
			}
		}
	}
	return Job{}, errors.New("no such employee exist")
}
