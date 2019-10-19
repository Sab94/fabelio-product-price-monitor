// package: priceMonitor
// file: services/priceMonitorpb/priceMonitor.proto

var services_priceMonitorpb_priceMonitor_pb = require("../../proto/priceMonitorpb/priceMonitor_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PriceMonitorService = (function () {
  function PriceMonitorService() {}
  PriceMonitorService.serviceName = "priceMonitor.PriceMonitorService";
  return PriceMonitorService;
}());

PriceMonitorService.AddProduct = {
  methodName: "AddProduct",
  service: PriceMonitorService,
  requestStream: false,
  responseStream: false,
  requestType: services_priceMonitorpb_priceMonitor_pb.AddProductRequest,
  responseType: services_priceMonitorpb_priceMonitor_pb.AddProductResponse
};

PriceMonitorService.GetProduct = {
  methodName: "GetProduct",
  service: PriceMonitorService,
  requestStream: false,
  responseStream: false,
  requestType: services_priceMonitorpb_priceMonitor_pb.GetProductRequest,
  responseType: services_priceMonitorpb_priceMonitor_pb.ProductResponse
};

PriceMonitorService.GetProducts = {
  methodName: "GetProducts",
  service: PriceMonitorService,
  requestStream: false,
  responseStream: false,
  requestType: services_priceMonitorpb_priceMonitor_pb.GetProductsRequest,
  responseType: services_priceMonitorpb_priceMonitor_pb.ProductsResponse
};

exports.PriceMonitorService = PriceMonitorService;

function PriceMonitorServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PriceMonitorServiceClient.prototype.addProduct = function addProduct(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PriceMonitorService.AddProduct, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

PriceMonitorServiceClient.prototype.getProduct = function getProduct(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PriceMonitorService.GetProduct, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

PriceMonitorServiceClient.prototype.getProducts = function getProducts(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PriceMonitorService.GetProducts, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.PriceMonitorServiceClient = PriceMonitorServiceClient;

