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

Analytics.PeriodicStoreSalesAmount = {
  methodName: "PeriodicStoreSalesAmount",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.PeriodicStoreSalesAmountRequest,
  responseType: analytics_pb.PeriodicStoreSalesAmountResponse
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

Analytics.TopRevenueItems = {
  methodName: "TopRevenueItems",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.TopRevenueItemsRequest,
  responseType: analytics_pb.TopRevenueItemsResponse
};

Analytics.LeastRevenueItems = {
  methodName: "LeastRevenueItems",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.LeastRevenueItemsRequest,
  responseType: analytics_pb.LeastRevenueItemsResponse
};

Analytics.StoreRevenueData = {
  methodName: "StoreRevenueData",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.StoreRevenueDataRequest,
  responseType: analytics_pb.StoreRevenueDataResponse
};

Analytics.StoreSalesData = {
  methodName: "StoreSalesData",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.StoreSalesDataRequest,
  responseType: analytics_pb.StoreSalesDataResponse
};

Analytics.AverageSalesAmount = {
  methodName: "AverageSalesAmount",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.AverageSalesAmountRequest,
  responseType: analytics_pb.AverageSalesAmountResponse
};

Analytics.AverageSoldQuantity = {
  methodName: "AverageSoldQuantity",
  service: Analytics,
  requestStream: false,
  responseStream: false,
  requestType: analytics_pb.AverageSoldQuantityRequest,
  responseType: analytics_pb.AverageSoldQuantityResponse
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

AnalyticsClient.prototype.periodicStoreSalesAmount = function periodicStoreSalesAmount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.PeriodicStoreSalesAmount, {
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

AnalyticsClient.prototype.topRevenueItems = function topRevenueItems(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.TopRevenueItems, {
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

AnalyticsClient.prototype.leastRevenueItems = function leastRevenueItems(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.LeastRevenueItems, {
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

AnalyticsClient.prototype.storeRevenueData = function storeRevenueData(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.StoreRevenueData, {
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

AnalyticsClient.prototype.storeSalesData = function storeSalesData(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.StoreSalesData, {
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

AnalyticsClient.prototype.averageSalesAmount = function averageSalesAmount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.AverageSalesAmount, {
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

AnalyticsClient.prototype.averageSoldQuantity = function averageSoldQuantity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Analytics.AverageSoldQuantity, {
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

