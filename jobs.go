package travis

import (
	"fmt"
	"net/http"

	"context"

	"github.com/fatih/structs"
)

// JobsService handles communication with the jobs
// related methods of the Travis CI API.
type JobsService struct {
	client *Client
}

type findJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// getJobResponse represents the response of a call
// to the Travis CI get build endpoint.
type getJobResponse struct {
	Job Job `json:"job"`
}

// JobListOptions specifies the optional parameters to the
// JobsService.List method. You need to provide exactly one
// of the below attribute. If you provide State or Queue, a
// maximum of 250 jobs will be returned.
type JobFindOptions struct {
	ListOptions

	// List of job ids
	Ids []uint `url:"ids,omitempty"`

	// Job state to filter by
	State string `url:"state,omitempty"`

	// Job queue to filter by
	Queue string `url:"queue,omitempty"`
}

// IsValid asserts the JobFindOptions instance has one
// and only one value set to a non-zero value.
//
// This method is particularly useful to check a JobFindOptions
// instance before passing it to JobsService.FindByID method.
func (jfo *JobFindOptions) IsValid() bool {
	s := structs.New(jfo)
	f := s.Fields()

	nonZeroValues := 0

	for _, field := range f {
		if !field.IsZero() {
			nonZeroValues += 1
		}
	}

	return nonZeroValues == 0 || nonZeroValues == 1
}

// Get fetches job with the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#jobs
func (js *JobsService) Get(ctx context.Context, id uint) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jobResp getJobResponse
	resp, err := js.client.Do(ctx, req, &jobResp)
	if err != nil {
		return nil, resp, err
	}

	return &jobResp.Job, resp, err
}

// ListByBuild retrieve a build jobs from its provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#jobs
func (js *JobsService) ListFromBuild(ctx context.Context, buildId uint) ([]Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/builds/%d", buildId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var buildResp getBuildResponse
	resp, err := js.client.Do(ctx, req, &buildResp)
	if err != nil {
		return nil, resp, err
	}

	return buildResp.Jobs, resp, err
}

// FindByID jobs using the provided options.
// You need to provide exactly one of the opt fields value.
// If you provide State or Queue, a maximum of 250 jobs will be returned.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#jobs
func (js *JobsService) Find(ctx context.Context, opt *JobFindOptions) ([]Job, *http.Response, error) {
	if opt != nil && !opt.IsValid() {
		return nil, nil, fmt.Errorf(
			"more than one value set in provided JobFindOptions instance",
		)
	}

	u, err := urlWithOptions("/jobs", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jobsResp findJobsResponse
	resp, err := js.client.Do(ctx, req, &jobsResp)
	if err != nil {
		return nil, resp, err
	}

	return jobsResp.Jobs, resp, err
}

// Cancel job with the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#jobs
func (js *JobsService) Cancel(ctx context.Context, id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs/%d/cancel", id), nil)
	if err != nil {
		return nil, err
	}

	req, err := js.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := js.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Restart job with the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#jobs
func (js *JobsService) Restart(ctx context.Context, id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs/%d/restart", id), nil)
	if err != nil {
		return nil, err
	}

	req, err := js.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := js.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
