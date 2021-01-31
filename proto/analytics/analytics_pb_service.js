// package: analytics
// file: analytics.proto

var analytics_pb = require("./analytics_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Analytics = (function () {
  function Analytics() {}
  Analytics.serviceName = "analytics.Analytics";
  return Analytics;
}());

Analytics.Ping = {
  methodName: "Ping",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.PingRequest,
  responseType: analytics_pb.PingResponse
};

Analytics.PeriodicSalesAmount = {
  methodName: "PeriodicSalesAmount",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.PeriodicSalesAmountRequest,
  responseType: analytics_pb.PeriodicSalesAmountResponse
};

Analytics.TopSoldItems = {
  methodName: "TopSoldItems",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.TopSoldItemsRequest,
  responseType: analytics_pb.TopSoldItemsResponse
};

Analytics.LeastSoldItems = {
  methodName: "LeastSoldItems",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.LeastSoldItemsRequest,
  responseType: analytics_pb.LeastSoldItemsResponse
};

Analytics.TotalSalesAmount = {
  methodName: "TotalSalesAmount",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.TotalSalesAmountRequest,
  responseType: analytics_pb.TotalSalesAmountResponse
};

exports.Analytics = Analytics;

function AnalyticsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AnalyticsClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.Ping, {
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

AnalyticsClient.prototype.periodicSalesAmount = function periodicSalesAmount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.PeriodicSalesAmount, {
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

AnalyticsClient.prototype.topSoldItems = function topSoldItems(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.TopSoldItems, {
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

AnalyticsClient.prototype.leastSoldItems = function leastSoldItems(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.LeastSoldItems, {
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

AnalyticsClient.prototype.totalSalesAmount = function totalSalesAmount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.TotalSalesAmount, {
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

exports.AnalyticsClient = AnalyticsClient;

