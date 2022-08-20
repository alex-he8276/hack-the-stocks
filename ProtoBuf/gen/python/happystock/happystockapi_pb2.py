# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: happystock/happystockapi.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ehappystock/happystockapi.proto\x12\x18happystock.happystockapi\x1a\x1fgoogle/protobuf/timestamp.proto\"r\n\x0estockSentiment\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12.\n\x04\x64\x61te\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x64\x61te\x12\x1c\n\tsentiment\x18\x03 \x01(\x05R\tsentiment\"f\n\nstockPrice\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12.\n\x04\x64\x61te\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x64\x61te\x12\x14\n\x05price\x18\x03 \x01(\x01R\x05price\"d\n\x12listStockSentiment\x12N\n\rsentimentList\x18\x01 \x03(\x0b\x32(.happystock.happystockapi.stockSentimentR\rsentimentList\"T\n\x0elistStockPrice\x12\x42\n\tpriceList\x18\x01 \x03(\x0b\x32$.happystock.happystockapi.stockPriceR\tpriceListB\xde\x01\n\x1c\x63om.happystock.happystockapiB\x12HappystockapiProtoP\x01Z)buf.build/mewcifer/happystocks/happystock\xa2\x02\x03HHX\xaa\x02\x18Happystock.Happystockapi\xca\x02\x18Happystock\\Happystockapi\xe2\x02$Happystock\\Happystockapi\\GPBMetadata\xea\x02\x19Happystock::Happystockapib\x06proto3')



_STOCKSENTIMENT = DESCRIPTOR.message_types_by_name['stockSentiment']
_STOCKPRICE = DESCRIPTOR.message_types_by_name['stockPrice']
_LISTSTOCKSENTIMENT = DESCRIPTOR.message_types_by_name['listStockSentiment']
_LISTSTOCKPRICE = DESCRIPTOR.message_types_by_name['listStockPrice']
stockSentiment = _reflection.GeneratedProtocolMessageType('stockSentiment', (_message.Message,), {
  'DESCRIPTOR' : _STOCKSENTIMENT,
  '__module__' : 'happystock.happystockapi_pb2'
  # @@protoc_insertion_point(class_scope:happystock.happystockapi.stockSentiment)
  })
_sym_db.RegisterMessage(stockSentiment)

stockPrice = _reflection.GeneratedProtocolMessageType('stockPrice', (_message.Message,), {
  'DESCRIPTOR' : _STOCKPRICE,
  '__module__' : 'happystock.happystockapi_pb2'
  # @@protoc_insertion_point(class_scope:happystock.happystockapi.stockPrice)
  })
_sym_db.RegisterMessage(stockPrice)

listStockSentiment = _reflection.GeneratedProtocolMessageType('listStockSentiment', (_message.Message,), {
  'DESCRIPTOR' : _LISTSTOCKSENTIMENT,
  '__module__' : 'happystock.happystockapi_pb2'
  # @@protoc_insertion_point(class_scope:happystock.happystockapi.listStockSentiment)
  })
_sym_db.RegisterMessage(listStockSentiment)

listStockPrice = _reflection.GeneratedProtocolMessageType('listStockPrice', (_message.Message,), {
  'DESCRIPTOR' : _LISTSTOCKPRICE,
  '__module__' : 'happystock.happystockapi_pb2'
  # @@protoc_insertion_point(class_scope:happystock.happystockapi.listStockPrice)
  })
_sym_db.RegisterMessage(listStockPrice)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\034com.happystock.happystockapiB\022HappystockapiProtoP\001Z)buf.build/mewcifer/happystocks/happystock\242\002\003HHX\252\002\030Happystock.Happystockapi\312\002\030Happystock\\Happystockapi\342\002$Happystock\\Happystockapi\\GPBMetadata\352\002\031Happystock::Happystockapi'
  _STOCKSENTIMENT._serialized_start=93
  _STOCKSENTIMENT._serialized_end=207
  _STOCKPRICE._serialized_start=209
  _STOCKPRICE._serialized_end=311
  _LISTSTOCKSENTIMENT._serialized_start=313
  _LISTSTOCKSENTIMENT._serialized_end=413
  _LISTSTOCKPRICE._serialized_start=415
  _LISTSTOCKPRICE._serialized_end=499
# @@protoc_insertion_point(module_scope)
