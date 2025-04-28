package submissionv1_impl

import (
	"context"
	"net/http"

	"github.com/muhlemmer/kudofine/gen/submission/v1/submissionv1connect"
	"github.com/muhlemmer/kudofine/internal/data"
)

type SubmissionService struct {
	submissionv1connect.UnimplementedSubmissionServiceHandler
	data *data.Data
}

func NewSubmissionService(background context.Context, data *data.Data) func() (string, http.Handler) {
	submissionService := &SubmissionService{
		data: data,
	}
	return func() (string, http.Handler) {
		return submissionv1connect.NewSubmissionServiceHandler(submissionService)
	}
}
