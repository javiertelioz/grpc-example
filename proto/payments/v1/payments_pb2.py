# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: payments/v1/payments.proto
# Protobuf Python Version: 5.27.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    3,
    '',
    'payments/v1/payments.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1apayments/v1/payments.proto\x12\x0bpayments.v1\x1a\x1cgoogle/api/annotations.proto\"A\n\x0e\x44\x65positRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12\x16\n\x06\x61mount\x18\x02 \x01(\x01R\x06\x61mount\"j\n\x0f\x44\x65positResponse\x12%\n\x0etransaction_id\x18\x01 \x01(\tR\rtransactionId\x12\x16\n\x06status\x18\x02 \x01(\tR\x06status\x12\x18\n\x07message\x18\x03 \x01(\tR\x07message\"B\n\x0fWithdrawRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12\x16\n\x06\x61mount\x18\x02 \x01(\x01R\x06\x61mount\"k\n\x10WithdrawResponse\x12%\n\x0etransaction_id\x18\x01 \x01(\tR\rtransactionId\x12\x16\n\x06status\x18\x02 \x01(\tR\x06status\x12\x18\n\x07message\x18\x03 \x01(\tR\x07message\"4\n\x19TransactionHistoryRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\"\xaf\x01\n\x0bTransaction\x12%\n\x0etransaction_id\x18\x01 \x01(\tR\rtransactionId\x12\x17\n\x07user_id\x18\x02 \x01(\tR\x06userId\x12\x16\n\x06\x61mount\x18\x03 \x01(\x01R\x06\x61mount\x12\x12\n\x04type\x18\x04 \x01(\tR\x04type\x12\x16\n\x06status\x18\x05 \x01(\tR\x06status\x12\x1c\n\ttimestamp\x18\x06 \x01(\tR\ttimestamp\"Y\n\x19UploadTransactionsRequest\x12<\n\x0ctransactions\x18\x01 \x03(\x0b\x32\x18.payments.v1.TransactionR\x0ctransactions\"~\n\x1aUploadTransactionsResponse\x12#\n\rsuccess_count\x18\x01 \x01(\x05R\x0csuccessCount\x12#\n\rfailure_count\x18\x02 \x01(\x05R\x0c\x66\x61ilureCount\x12\x16\n\x06\x65rrors\x18\x03 \x03(\tR\x06\x65rrors2\x94\x05\n\x0ePaymentService\x12\x65\n\x07\x44\x65posit\x12\x1b.payments.v1.DepositRequest\x1a\x1c.payments.v1.DepositResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1/payments/deposit:\x01*\x12i\n\x08Withdraw\x12\x1c.payments.v1.WithdrawRequest\x1a\x1d.payments.v1.WithdrawResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/v1/payments/withdraw:\x01*\x12\x8f\x01\n\x15GetTransactionHistory\x12&.payments.v1.TransactionHistoryRequest\x1a\x18.payments.v1.Transaction\"2\x82\xd3\xe4\x93\x02,\x12*/v1/payments/transaction-history/{user_id}0\x01\x12\x9f\x01\n\x12UploadTransactions\x12&.payments.v1.UploadTransactionsRequest\x1a\'.payments.v1.UploadTransactionsResponse\"6\x82\xd3\xe4\x93\x02\x30\" /v1/payments/upload-transactions:\x0ctransactions(\x01\x12|\n\x13RealTimeTransaction\x12\x18.payments.v1.Transaction\x1a\x18.payments.v1.Transaction\"-\x82\xd3\xe4\x93\x02\'\"\"/v1/payments/real-time-transaction:\x01*(\x01\x30\x01\x42\x38Z6github.com/javiertelioz/grpc-service-template/paymentsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'payments.v1.payments_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z6github.com/javiertelioz/grpc-service-template/payments'
  _globals['_PAYMENTSERVICE'].methods_by_name['Deposit']._loaded_options = None
  _globals['_PAYMENTSERVICE'].methods_by_name['Deposit']._serialized_options = b'\202\323\344\223\002\031\"\024/v1/payments/deposit:\001*'
  _globals['_PAYMENTSERVICE'].methods_by_name['Withdraw']._loaded_options = None
  _globals['_PAYMENTSERVICE'].methods_by_name['Withdraw']._serialized_options = b'\202\323\344\223\002\032\"\025/v1/payments/withdraw:\001*'
  _globals['_PAYMENTSERVICE'].methods_by_name['GetTransactionHistory']._loaded_options = None
  _globals['_PAYMENTSERVICE'].methods_by_name['GetTransactionHistory']._serialized_options = b'\202\323\344\223\002,\022*/v1/payments/transaction-history/{user_id}'
  _globals['_PAYMENTSERVICE'].methods_by_name['UploadTransactions']._loaded_options = None
  _globals['_PAYMENTSERVICE'].methods_by_name['UploadTransactions']._serialized_options = b'\202\323\344\223\0020\" /v1/payments/upload-transactions:\014transactions'
  _globals['_PAYMENTSERVICE'].methods_by_name['RealTimeTransaction']._loaded_options = None
  _globals['_PAYMENTSERVICE'].methods_by_name['RealTimeTransaction']._serialized_options = b'\202\323\344\223\002\'\"\"/v1/payments/real-time-transaction:\001*'
  _globals['_DEPOSITREQUEST']._serialized_start=73
  _globals['_DEPOSITREQUEST']._serialized_end=138
  _globals['_DEPOSITRESPONSE']._serialized_start=140
  _globals['_DEPOSITRESPONSE']._serialized_end=246
  _globals['_WITHDRAWREQUEST']._serialized_start=248
  _globals['_WITHDRAWREQUEST']._serialized_end=314
  _globals['_WITHDRAWRESPONSE']._serialized_start=316
  _globals['_WITHDRAWRESPONSE']._serialized_end=423
  _globals['_TRANSACTIONHISTORYREQUEST']._serialized_start=425
  _globals['_TRANSACTIONHISTORYREQUEST']._serialized_end=477
  _globals['_TRANSACTION']._serialized_start=480
  _globals['_TRANSACTION']._serialized_end=655
  _globals['_UPLOADTRANSACTIONSREQUEST']._serialized_start=657
  _globals['_UPLOADTRANSACTIONSREQUEST']._serialized_end=746
  _globals['_UPLOADTRANSACTIONSRESPONSE']._serialized_start=748
  _globals['_UPLOADTRANSACTIONSRESPONSE']._serialized_end=874
  _globals['_PAYMENTSERVICE']._serialized_start=877
  _globals['_PAYMENTSERVICE']._serialized_end=1537
# @@protoc_insertion_point(module_scope)
