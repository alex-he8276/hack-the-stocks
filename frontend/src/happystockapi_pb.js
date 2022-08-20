// source: happystock/happystockapi.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.happystock.happystockapi.listStockPrice', null, global);
goog.exportSymbol('proto.happystock.happystockapi.listStockSentiment', null, global);
goog.exportSymbol('proto.happystock.happystockapi.stockPrice', null, global);
goog.exportSymbol('proto.happystock.happystockapi.stockSentiment', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.happystock.happystockapi.stockSentiment = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.happystock.happystockapi.stockSentiment, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.happystock.happystockapi.stockSentiment.displayName = 'proto.happystock.happystockapi.stockSentiment';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.happystock.happystockapi.stockPrice = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.happystock.happystockapi.stockPrice, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.happystock.happystockapi.stockPrice.displayName = 'proto.happystock.happystockapi.stockPrice';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.happystock.happystockapi.listStockSentiment = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.happystock.happystockapi.listStockSentiment.repeatedFields_, null);
};
goog.inherits(proto.happystock.happystockapi.listStockSentiment, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.happystock.happystockapi.listStockSentiment.displayName = 'proto.happystock.happystockapi.listStockSentiment';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.happystock.happystockapi.listStockPrice = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.happystock.happystockapi.listStockPrice.repeatedFields_, null);
};
goog.inherits(proto.happystock.happystockapi.listStockPrice, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.happystock.happystockapi.listStockPrice.displayName = 'proto.happystock.happystockapi.listStockPrice';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.happystock.happystockapi.stockSentiment.prototype.toObject = function(opt_includeInstance) {
  return proto.happystock.happystockapi.stockSentiment.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.happystock.happystockapi.stockSentiment} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.stockSentiment.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    date: (f = msg.getDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    sentiment: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.happystock.happystockapi.stockSentiment}
 */
proto.happystock.happystockapi.stockSentiment.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.happystock.happystockapi.stockSentiment;
  return proto.happystock.happystockapi.stockSentiment.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.happystock.happystockapi.stockSentiment} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.happystock.happystockapi.stockSentiment}
 */
proto.happystock.happystockapi.stockSentiment.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDate(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSentiment(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.happystock.happystockapi.stockSentiment.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.happystock.happystockapi.stockSentiment.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.happystock.happystockapi.stockSentiment} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.stockSentiment.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDate();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getSentiment();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.happystock.happystockapi.stockSentiment.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.happystock.happystockapi.stockSentiment} returns this
 */
proto.happystock.happystockapi.stockSentiment.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp date = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.happystock.happystockapi.stockSentiment.prototype.getDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.happystock.happystockapi.stockSentiment} returns this
*/
proto.happystock.happystockapi.stockSentiment.prototype.setDate = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.happystock.happystockapi.stockSentiment} returns this
 */
proto.happystock.happystockapi.stockSentiment.prototype.clearDate = function() {
  return this.setDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.happystock.happystockapi.stockSentiment.prototype.hasDate = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int32 sentiment = 3;
 * @return {number}
 */
proto.happystock.happystockapi.stockSentiment.prototype.getSentiment = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.happystock.happystockapi.stockSentiment} returns this
 */
proto.happystock.happystockapi.stockSentiment.prototype.setSentiment = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.happystock.happystockapi.stockPrice.prototype.toObject = function(opt_includeInstance) {
  return proto.happystock.happystockapi.stockPrice.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.happystock.happystockapi.stockPrice} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.stockPrice.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    date: (f = msg.getDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    price: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.happystock.happystockapi.stockPrice}
 */
proto.happystock.happystockapi.stockPrice.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.happystock.happystockapi.stockPrice;
  return proto.happystock.happystockapi.stockPrice.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.happystock.happystockapi.stockPrice} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.happystock.happystockapi.stockPrice}
 */
proto.happystock.happystockapi.stockPrice.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDate(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setPrice(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.happystock.happystockapi.stockPrice.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.happystock.happystockapi.stockPrice.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.happystock.happystockapi.stockPrice} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.stockPrice.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDate();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPrice();
  if (f !== 0.0) {
    writer.writeDouble(
      3,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.happystock.happystockapi.stockPrice.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.happystock.happystockapi.stockPrice} returns this
 */
proto.happystock.happystockapi.stockPrice.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp date = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.happystock.happystockapi.stockPrice.prototype.getDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.happystock.happystockapi.stockPrice} returns this
*/
proto.happystock.happystockapi.stockPrice.prototype.setDate = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.happystock.happystockapi.stockPrice} returns this
 */
proto.happystock.happystockapi.stockPrice.prototype.clearDate = function() {
  return this.setDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.happystock.happystockapi.stockPrice.prototype.hasDate = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional double price = 3;
 * @return {number}
 */
proto.happystock.happystockapi.stockPrice.prototype.getPrice = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.happystock.happystockapi.stockPrice} returns this
 */
proto.happystock.happystockapi.stockPrice.prototype.setPrice = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.happystock.happystockapi.listStockSentiment.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.happystock.happystockapi.listStockSentiment.prototype.toObject = function(opt_includeInstance) {
  return proto.happystock.happystockapi.listStockSentiment.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.happystock.happystockapi.listStockSentiment} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.listStockSentiment.toObject = function(includeInstance, msg) {
  var f, obj = {
    sentimentlistList: jspb.Message.toObjectList(msg.getSentimentlistList(),
    proto.happystock.happystockapi.stockSentiment.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.happystock.happystockapi.listStockSentiment}
 */
proto.happystock.happystockapi.listStockSentiment.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.happystock.happystockapi.listStockSentiment;
  return proto.happystock.happystockapi.listStockSentiment.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.happystock.happystockapi.listStockSentiment} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.happystock.happystockapi.listStockSentiment}
 */
proto.happystock.happystockapi.listStockSentiment.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.happystock.happystockapi.stockSentiment;
      reader.readMessage(value,proto.happystock.happystockapi.stockSentiment.deserializeBinaryFromReader);
      msg.addSentimentlist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.happystock.happystockapi.listStockSentiment.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.happystock.happystockapi.listStockSentiment.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.happystock.happystockapi.listStockSentiment} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.listStockSentiment.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSentimentlistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.happystock.happystockapi.stockSentiment.serializeBinaryToWriter
    );
  }
};


/**
 * repeated stockSentiment sentimentList = 1;
 * @return {!Array<!proto.happystock.happystockapi.stockSentiment>}
 */
proto.happystock.happystockapi.listStockSentiment.prototype.getSentimentlistList = function() {
  return /** @type{!Array<!proto.happystock.happystockapi.stockSentiment>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.happystock.happystockapi.stockSentiment, 1));
};


/**
 * @param {!Array<!proto.happystock.happystockapi.stockSentiment>} value
 * @return {!proto.happystock.happystockapi.listStockSentiment} returns this
*/
proto.happystock.happystockapi.listStockSentiment.prototype.setSentimentlistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.happystock.happystockapi.stockSentiment=} opt_value
 * @param {number=} opt_index
 * @return {!proto.happystock.happystockapi.stockSentiment}
 */
proto.happystock.happystockapi.listStockSentiment.prototype.addSentimentlist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.happystock.happystockapi.stockSentiment, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.happystock.happystockapi.listStockSentiment} returns this
 */
proto.happystock.happystockapi.listStockSentiment.prototype.clearSentimentlistList = function() {
  return this.setSentimentlistList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.happystock.happystockapi.listStockPrice.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.happystock.happystockapi.listStockPrice.prototype.toObject = function(opt_includeInstance) {
  return proto.happystock.happystockapi.listStockPrice.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.happystock.happystockapi.listStockPrice} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.listStockPrice.toObject = function(includeInstance, msg) {
  var f, obj = {
    pricelistList: jspb.Message.toObjectList(msg.getPricelistList(),
    proto.happystock.happystockapi.stockPrice.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.happystock.happystockapi.listStockPrice}
 */
proto.happystock.happystockapi.listStockPrice.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.happystock.happystockapi.listStockPrice;
  return proto.happystock.happystockapi.listStockPrice.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.happystock.happystockapi.listStockPrice} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.happystock.happystockapi.listStockPrice}
 */
proto.happystock.happystockapi.listStockPrice.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.happystock.happystockapi.stockPrice;
      reader.readMessage(value,proto.happystock.happystockapi.stockPrice.deserializeBinaryFromReader);
      msg.addPricelist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.happystock.happystockapi.listStockPrice.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.happystock.happystockapi.listStockPrice.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.happystock.happystockapi.listStockPrice} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.happystock.happystockapi.listStockPrice.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPricelistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.happystock.happystockapi.stockPrice.serializeBinaryToWriter
    );
  }
};


/**
 * repeated stockPrice priceList = 1;
 * @return {!Array<!proto.happystock.happystockapi.stockPrice>}
 */
proto.happystock.happystockapi.listStockPrice.prototype.getPricelistList = function() {
  return /** @type{!Array<!proto.happystock.happystockapi.stockPrice>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.happystock.happystockapi.stockPrice, 1));
};


/**
 * @param {!Array<!proto.happystock.happystockapi.stockPrice>} value
 * @return {!proto.happystock.happystockapi.listStockPrice} returns this
*/
proto.happystock.happystockapi.listStockPrice.prototype.setPricelistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.happystock.happystockapi.stockPrice=} opt_value
 * @param {number=} opt_index
 * @return {!proto.happystock.happystockapi.stockPrice}
 */
proto.happystock.happystockapi.listStockPrice.prototype.addPricelist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.happystock.happystockapi.stockPrice, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.happystock.happystockapi.listStockPrice} returns this
 */
proto.happystock.happystockapi.listStockPrice.prototype.clearPricelistList = function() {
  return this.setPricelistList([]);
};


goog.object.extend(exports, proto.happystock.happystockapi);
