// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: payments/v1/payments.proto

package payments

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_Deposit_FullMethodName               = "/payments.v1.PaymentService/Deposit"
	PaymentService_Withdraw_FullMethodName              = "/payments.v1.PaymentService/Withdraw"
	PaymentService_GetTransactionHistory_FullMethodName = "/payments.v1.PaymentService/GetTransactionHistory"
	PaymentService_UploadTransactions_FullMethodName    = "/payments.v1.PaymentService/UploadTransactions"
	PaymentService_RealTimeTransaction_FullMethodName   = "/payments.v1.PaymentService/RealTimeTransaction"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PaymentService provides a set of APIs for processing payments and transactions.
type PaymentServiceClient interface {
	// Deposit processes a deposit request.
	//
	// This method expects a POST request with the user's ID and the amount to be deposited.
	//
	// Request:
	//   - `user_id`: The ID of the user making the deposit.
	//   - `amount`: The amount to be deposited.
	//
	// Response:
	//   - `transaction_id`: The ID of the transaction.
	//   - `status`: The status of the deposit (e.g., "success", "failure").
	//   - `message`: Additional information about the transaction.
	//
	// Possible HTTP responses:
	//   - 200: The deposit was successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	// Withdraw processes a withdrawal request.
	//
	// This method expects a POST request with the user's ID and the amount to be withdrawn.
	//
	// Request:
	//   - `user_id`: The ID of the user making the withdrawal.
	//   - `amount`: The amount to be withdrawn.
	//
	// Response:
	//   - `transaction_id`: The ID of the transaction.
	//   - `status`: The status of the withdrawal (e.g., "success", "failure").
	//   - `message`: Additional information about the transaction.
	//
	// Possible HTTP responses:
	//   - 200: The withdrawal was successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	// GetTransactionHistory retrieves the transaction history for a user.
	//
	// This method expects a GET request with the user's ID as a path parameter.
	//
	// Request:
	//   - `user_id`: The ID of the user whose transaction history is being retrieved.
	//
	// Response (streaming):
	//   - `transaction_id`: The ID of the transaction.
	//   - `user_id`: The ID of the user associated with the transaction.
	//   - `amount`: The amount involved in the transaction.
	//   - `type`: The type of transaction (e.g., "deposit", "withdrawal").
	//   - `status`: The status of the transaction.
	//   - `timestamp`: The timestamp of when the transaction occurred.
	//
	// Possible HTTP responses:
	//   - 200: The transaction history was successfully retrieved.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	GetTransactionHistory(ctx context.Context, in *TransactionHistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Transaction], error)
	// UploadTransactions uploads multiple transactions for processing.
	//
	// This method expects a POST request with a list of transactions to be processed.
	//
	// Request:
	//   - `transactions`: A list of transactions to be uploaded.
	//
	// Response:
	//   - `success_count`: The number of transactions successfully processed.
	//   - `failure_count`: The number of transactions that failed to process.
	//   - `errors`: A list of error messages for the failed transactions.
	//
	// Possible HTTP responses:
	//   - 200: The transactions were successfully uploaded.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	UploadTransactions(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadTransactionsRequest, UploadTransactionsResponse], error)
	// RealTimeTransaction processes transactions in real-time.
	//
	// This method expects a bidirectional stream of transactions to be processed in real-time.
	//
	// Request/Response (streaming):
	//   - `transaction_id`: The ID of the transaction.
	//   - `user_id`: The ID of the user associated with the transaction.
	//   - `amount`: The amount involved in the transaction.
	//   - `type`: The type of transaction (e.g., "deposit", "withdrawal").
	//   - `status`: The status of the transaction.
	//   - `timestamp`: The timestamp of when the transaction occurred.
	//
	// Possible HTTP responses:
	//   - 200: The transactions were successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	RealTimeTransaction(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Transaction, Transaction], error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, PaymentService_Deposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, PaymentService_Withdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetTransactionHistory(ctx context.Context, in *TransactionHistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Transaction], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PaymentService_ServiceDesc.Streams[0], PaymentService_GetTransactionHistory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TransactionHistoryRequest, Transaction]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_GetTransactionHistoryClient = grpc.ServerStreamingClient[Transaction]

func (c *paymentServiceClient) UploadTransactions(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadTransactionsRequest, UploadTransactionsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PaymentService_ServiceDesc.Streams[1], PaymentService_UploadTransactions_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadTransactionsRequest, UploadTransactionsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_UploadTransactionsClient = grpc.ClientStreamingClient[UploadTransactionsRequest, UploadTransactionsResponse]

func (c *paymentServiceClient) RealTimeTransaction(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Transaction, Transaction], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PaymentService_ServiceDesc.Streams[2], PaymentService_RealTimeTransaction_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Transaction, Transaction]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_RealTimeTransactionClient = grpc.BidiStreamingClient[Transaction, Transaction]

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
//
// PaymentService provides a set of APIs for processing payments and transactions.
type PaymentServiceServer interface {
	// Deposit processes a deposit request.
	//
	// This method expects a POST request with the user's ID and the amount to be deposited.
	//
	// Request:
	//   - `user_id`: The ID of the user making the deposit.
	//   - `amount`: The amount to be deposited.
	//
	// Response:
	//   - `transaction_id`: The ID of the transaction.
	//   - `status`: The status of the deposit (e.g., "success", "failure").
	//   - `message`: Additional information about the transaction.
	//
	// Possible HTTP responses:
	//   - 200: The deposit was successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	// Withdraw processes a withdrawal request.
	//
	// This method expects a POST request with the user's ID and the amount to be withdrawn.
	//
	// Request:
	//   - `user_id`: The ID of the user making the withdrawal.
	//   - `amount`: The amount to be withdrawn.
	//
	// Response:
	//   - `transaction_id`: The ID of the transaction.
	//   - `status`: The status of the withdrawal (e.g., "success", "failure").
	//   - `message`: Additional information about the transaction.
	//
	// Possible HTTP responses:
	//   - 200: The withdrawal was successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	// GetTransactionHistory retrieves the transaction history for a user.
	//
	// This method expects a GET request with the user's ID as a path parameter.
	//
	// Request:
	//   - `user_id`: The ID of the user whose transaction history is being retrieved.
	//
	// Response (streaming):
	//   - `transaction_id`: The ID of the transaction.
	//   - `user_id`: The ID of the user associated with the transaction.
	//   - `amount`: The amount involved in the transaction.
	//   - `type`: The type of transaction (e.g., "deposit", "withdrawal").
	//   - `status`: The status of the transaction.
	//   - `timestamp`: The timestamp of when the transaction occurred.
	//
	// Possible HTTP responses:
	//   - 200: The transaction history was successfully retrieved.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	GetTransactionHistory(*TransactionHistoryRequest, grpc.ServerStreamingServer[Transaction]) error
	// UploadTransactions uploads multiple transactions for processing.
	//
	// This method expects a POST request with a list of transactions to be processed.
	//
	// Request:
	//   - `transactions`: A list of transactions to be uploaded.
	//
	// Response:
	//   - `success_count`: The number of transactions successfully processed.
	//   - `failure_count`: The number of transactions that failed to process.
	//   - `errors`: A list of error messages for the failed transactions.
	//
	// Possible HTTP responses:
	//   - 200: The transactions were successfully uploaded.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	UploadTransactions(grpc.ClientStreamingServer[UploadTransactionsRequest, UploadTransactionsResponse]) error
	// RealTimeTransaction processes transactions in real-time.
	//
	// This method expects a bidirectional stream of transactions to be processed in real-time.
	//
	// Request/Response (streaming):
	//   - `transaction_id`: The ID of the transaction.
	//   - `user_id`: The ID of the user associated with the transaction.
	//   - `amount`: The amount involved in the transaction.
	//   - `type`: The type of transaction (e.g., "deposit", "withdrawal").
	//   - `status`: The status of the transaction.
	//   - `timestamp`: The timestamp of when the transaction occurred.
	//
	// Possible HTTP responses:
	//   - 200: The transactions were successfully processed.
	//   - 400: The request was malformed, usually due to missing or invalid parameters.
	//   - 401: Unauthorized request, typically due to missing or invalid authentication credentials.
	//   - 500: Internal server error, indicating a problem on the server side.
	RealTimeTransaction(grpc.BidiStreamingServer[Transaction, Transaction]) error
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) Deposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedPaymentServiceServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedPaymentServiceServer) GetTransactionHistory(*TransactionHistoryRequest, grpc.ServerStreamingServer[Transaction]) error {
	return status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
func (UnimplementedPaymentServiceServer) UploadTransactions(grpc.ClientStreamingServer[UploadTransactionsRequest, UploadTransactionsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadTransactions not implemented")
}
func (UnimplementedPaymentServiceServer) RealTimeTransaction(grpc.BidiStreamingServer[Transaction, Transaction]) error {
	return status.Errorf(codes.Unimplemented, "method RealTimeTransaction not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetTransactionHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PaymentServiceServer).GetTransactionHistory(m, &grpc.GenericServerStream[TransactionHistoryRequest, Transaction]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_GetTransactionHistoryServer = grpc.ServerStreamingServer[Transaction]

func _PaymentService_UploadTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaymentServiceServer).UploadTransactions(&grpc.GenericServerStream[UploadTransactionsRequest, UploadTransactionsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_UploadTransactionsServer = grpc.ClientStreamingServer[UploadTransactionsRequest, UploadTransactionsResponse]

func _PaymentService_RealTimeTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaymentServiceServer).RealTimeTransaction(&grpc.GenericServerStream[Transaction, Transaction]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PaymentService_RealTimeTransactionServer = grpc.BidiStreamingServer[Transaction, Transaction]

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payments.v1.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _PaymentService_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _PaymentService_Withdraw_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTransactionHistory",
			Handler:       _PaymentService_GetTransactionHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadTransactions",
			Handler:       _PaymentService_UploadTransactions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RealTimeTransaction",
			Handler:       _PaymentService_RealTimeTransaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "payments/v1/payments.proto",
}
