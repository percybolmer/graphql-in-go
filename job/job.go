package job

// Repository is used to specify whats needed to fulfill the job storage requirements
type Repository interface {
	// GetJobs will search for all jobs related to and EmployeeID
	GetJobs(employeeID, company string) ([]Job, error)
	// GetJob will search for a certain job based on ID
	GetJob(employeeID, jobid string) (Job, error)
	// Update will take in a job and update the repository, it will return the new state of the job
	Update(Job) (Job, error)
}

// Job is how a job is presented
type Job struct {
	ID string `json:"id"`
	// EmployeeID is the employee related to the job
	EmployeeID string `json:"employeeID"`
	Company    string `json:"company"`
	Title      string `json:"title"`
	// Start is when the job started
	Start string `json:"start"`
	// End is when the employment ended
	End string `json:"end"`
}
