package worker

type Job struct {
	name   string
	Object interface{}
}

func NewJob(name string, object interface{}) *Job {
	return &Job{
		name:   name,
		Object: object,
	}
}
