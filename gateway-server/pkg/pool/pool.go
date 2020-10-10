package pkg

import "fmt"

var (
	MaxWorker = 2
	MaxQueue  = 2
)

type Payload struct {
	// [redacted]

}

// Job represents the job to be run

type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.

var JobQueue chan Job

// Worker represents the worker that executes the job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}
func (p *Payload) UploadToS3() error {
	fmt.Printf("xxxxxxxxxxxxx")
	return nil
	// the storageFolder method ensures that there are no name collision in

	// case we get same timestamp in the key name

	// storage_path := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())

	// bucket := S3Bucket

	// b := new(bytes.Buffer)
	// encodeErr := json.NewEncoder(b).Encode(payload)
	// if encodeErr != nil {
	// 	return encodeErr
	// }

	// // Everything we post to the S3 bucket should be marked 'private'

	// var acl = s3.Private
	// var contentType = "application/octet-stream"

	// return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.

			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.

				if err := job.Payload.UploadToS3(); err != nil {
					// log.Errorf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop

				return

			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.

func (w Worker) Stop() {
	go func() {
		w.quit <- true

	}()
}
