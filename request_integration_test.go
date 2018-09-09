// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestRequestService_CreateAndFindById(t *testing.T) {
	t.Parallel()

	createdRequest, res, err := integrationClient.Requests.CreateByRepoId(context.TODO(), integrationRepoId, &CreateRequestsOption{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	request, res, err := integrationClient.Request.FindByRepoId(context.TODO(), integrationRepoId, createdRequest.Id)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if request.Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", request.Id, createdRequest.Id)
	}
}

func TestRequestService_CreateAndFindBySlug(t *testing.T) {
	t.Parallel()

	createdRequest, res, err := integrationClient.Requests.CreateByRepoSlug(context.TODO(), integrationRepo, &CreateRequestsOption{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	request, res, err := integrationClient.Request.FindByRepoSlug(context.TODO(), integrationRepo, createdRequest.Id)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if request.Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", request.Id, createdRequest.Id)
	}
}