// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: submission/v1/submission.proto

package submissionv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/muhlemmer/kudofine/gen/submission/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SubmissionServiceName is the fully-qualified name of the SubmissionService service.
	SubmissionServiceName = "submission.v1.SubmissionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SubmissionServiceGetSubmissionProcedure is the fully-qualified name of the SubmissionService's
	// GetSubmission RPC.
	SubmissionServiceGetSubmissionProcedure = "/submission.v1.SubmissionService/GetSubmission"
	// SubmissionServiceCreateSubmissionProcedure is the fully-qualified name of the SubmissionService's
	// CreateSubmission RPC.
	SubmissionServiceCreateSubmissionProcedure = "/submission.v1.SubmissionService/CreateSubmission"
	// SubmissionServiceUpdateSubmissionProcedure is the fully-qualified name of the SubmissionService's
	// UpdateSubmission RPC.
	SubmissionServiceUpdateSubmissionProcedure = "/submission.v1.SubmissionService/UpdateSubmission"
	// SubmissionServiceDeleteSubmissionProcedure is the fully-qualified name of the SubmissionService's
	// DeleteSubmission RPC.
	SubmissionServiceDeleteSubmissionProcedure = "/submission.v1.SubmissionService/DeleteSubmission"
)

// SubmissionServiceClient is a client for the submission.v1.SubmissionService service.
type SubmissionServiceClient interface {
	GetSubmission(context.Context, *connect.Request[v1.GetSubmissionRequest]) (*connect.Response[v1.GetSubmissionResponse], error)
	CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreateSubmissionResponse], error)
	UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.UpdateSubmissionResponse], error)
	DeleteSubmission(context.Context, *connect.Request[v1.DeleteSubmissionRequest]) (*connect.Response[v1.DeleteSubmissionResponse], error)
}

// NewSubmissionServiceClient constructs a client for the submission.v1.SubmissionService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSubmissionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SubmissionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	submissionServiceMethods := v1.File_submission_v1_submission_proto.Services().ByName("SubmissionService").Methods()
	return &submissionServiceClient{
		getSubmission: connect.NewClient[v1.GetSubmissionRequest, v1.GetSubmissionResponse](
			httpClient,
			baseURL+SubmissionServiceGetSubmissionProcedure,
			connect.WithSchema(submissionServiceMethods.ByName("GetSubmission")),
			connect.WithClientOptions(opts...),
		),
		createSubmission: connect.NewClient[v1.CreateSubmissionRequest, v1.CreateSubmissionResponse](
			httpClient,
			baseURL+SubmissionServiceCreateSubmissionProcedure,
			connect.WithSchema(submissionServiceMethods.ByName("CreateSubmission")),
			connect.WithClientOptions(opts...),
		),
		updateSubmission: connect.NewClient[v1.UpdateSubmissionRequest, v1.UpdateSubmissionResponse](
			httpClient,
			baseURL+SubmissionServiceUpdateSubmissionProcedure,
			connect.WithSchema(submissionServiceMethods.ByName("UpdateSubmission")),
			connect.WithClientOptions(opts...),
		),
		deleteSubmission: connect.NewClient[v1.DeleteSubmissionRequest, v1.DeleteSubmissionResponse](
			httpClient,
			baseURL+SubmissionServiceDeleteSubmissionProcedure,
			connect.WithSchema(submissionServiceMethods.ByName("DeleteSubmission")),
			connect.WithClientOptions(opts...),
		),
	}
}

// submissionServiceClient implements SubmissionServiceClient.
type submissionServiceClient struct {
	getSubmission    *connect.Client[v1.GetSubmissionRequest, v1.GetSubmissionResponse]
	createSubmission *connect.Client[v1.CreateSubmissionRequest, v1.CreateSubmissionResponse]
	updateSubmission *connect.Client[v1.UpdateSubmissionRequest, v1.UpdateSubmissionResponse]
	deleteSubmission *connect.Client[v1.DeleteSubmissionRequest, v1.DeleteSubmissionResponse]
}

// GetSubmission calls submission.v1.SubmissionService.GetSubmission.
func (c *submissionServiceClient) GetSubmission(ctx context.Context, req *connect.Request[v1.GetSubmissionRequest]) (*connect.Response[v1.GetSubmissionResponse], error) {
	return c.getSubmission.CallUnary(ctx, req)
}

// CreateSubmission calls submission.v1.SubmissionService.CreateSubmission.
func (c *submissionServiceClient) CreateSubmission(ctx context.Context, req *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreateSubmissionResponse], error) {
	return c.createSubmission.CallUnary(ctx, req)
}

// UpdateSubmission calls submission.v1.SubmissionService.UpdateSubmission.
func (c *submissionServiceClient) UpdateSubmission(ctx context.Context, req *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.UpdateSubmissionResponse], error) {
	return c.updateSubmission.CallUnary(ctx, req)
}

// DeleteSubmission calls submission.v1.SubmissionService.DeleteSubmission.
func (c *submissionServiceClient) DeleteSubmission(ctx context.Context, req *connect.Request[v1.DeleteSubmissionRequest]) (*connect.Response[v1.DeleteSubmissionResponse], error) {
	return c.deleteSubmission.CallUnary(ctx, req)
}

// SubmissionServiceHandler is an implementation of the submission.v1.SubmissionService service.
type SubmissionServiceHandler interface {
	GetSubmission(context.Context, *connect.Request[v1.GetSubmissionRequest]) (*connect.Response[v1.GetSubmissionResponse], error)
	CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreateSubmissionResponse], error)
	UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.UpdateSubmissionResponse], error)
	DeleteSubmission(context.Context, *connect.Request[v1.DeleteSubmissionRequest]) (*connect.Response[v1.DeleteSubmissionResponse], error)
}

// NewSubmissionServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSubmissionServiceHandler(svc SubmissionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	submissionServiceMethods := v1.File_submission_v1_submission_proto.Services().ByName("SubmissionService").Methods()
	submissionServiceGetSubmissionHandler := connect.NewUnaryHandler(
		SubmissionServiceGetSubmissionProcedure,
		svc.GetSubmission,
		connect.WithSchema(submissionServiceMethods.ByName("GetSubmission")),
		connect.WithHandlerOptions(opts...),
	)
	submissionServiceCreateSubmissionHandler := connect.NewUnaryHandler(
		SubmissionServiceCreateSubmissionProcedure,
		svc.CreateSubmission,
		connect.WithSchema(submissionServiceMethods.ByName("CreateSubmission")),
		connect.WithHandlerOptions(opts...),
	)
	submissionServiceUpdateSubmissionHandler := connect.NewUnaryHandler(
		SubmissionServiceUpdateSubmissionProcedure,
		svc.UpdateSubmission,
		connect.WithSchema(submissionServiceMethods.ByName("UpdateSubmission")),
		connect.WithHandlerOptions(opts...),
	)
	submissionServiceDeleteSubmissionHandler := connect.NewUnaryHandler(
		SubmissionServiceDeleteSubmissionProcedure,
		svc.DeleteSubmission,
		connect.WithSchema(submissionServiceMethods.ByName("DeleteSubmission")),
		connect.WithHandlerOptions(opts...),
	)
	return "/submission.v1.SubmissionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SubmissionServiceGetSubmissionProcedure:
			submissionServiceGetSubmissionHandler.ServeHTTP(w, r)
		case SubmissionServiceCreateSubmissionProcedure:
			submissionServiceCreateSubmissionHandler.ServeHTTP(w, r)
		case SubmissionServiceUpdateSubmissionProcedure:
			submissionServiceUpdateSubmissionHandler.ServeHTTP(w, r)
		case SubmissionServiceDeleteSubmissionProcedure:
			submissionServiceDeleteSubmissionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSubmissionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSubmissionServiceHandler struct{}

func (UnimplementedSubmissionServiceHandler) GetSubmission(context.Context, *connect.Request[v1.GetSubmissionRequest]) (*connect.Response[v1.GetSubmissionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("submission.v1.SubmissionService.GetSubmission is not implemented"))
}

func (UnimplementedSubmissionServiceHandler) CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreateSubmissionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("submission.v1.SubmissionService.CreateSubmission is not implemented"))
}

func (UnimplementedSubmissionServiceHandler) UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.UpdateSubmissionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("submission.v1.SubmissionService.UpdateSubmission is not implemented"))
}

func (UnimplementedSubmissionServiceHandler) DeleteSubmission(context.Context, *connect.Request[v1.DeleteSubmissionRequest]) (*connect.Response[v1.DeleteSubmissionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("submission.v1.SubmissionService.DeleteSubmission is not implemented"))
}
