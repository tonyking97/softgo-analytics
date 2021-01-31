// package: analytics
// file: analytics.proto

import * as analytics_pb from "./analytics_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AnalyticsPing = {
  readonly methodName: string;
  readonly service: typeof Analytics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof analytics_pb.PingRequest;
  readonly responseType: typeof analytics_pb.PingResponse;
};

type AnalyticsPeriodicSalesAmount = {
  readonly methodName: string;
  readonly service: typeof Analytics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof analytics_pb.PeriodicSalesAmountRequest;
  readonly responseType: typeof analytics_pb.PeriodicSalesAmountResponse;
};

type AnalyticsTopSoldItems = {
  readonly methodName: string;
  readonly service: typeof Analytics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof analytics_pb.TopSoldItemsRequest;
  readonly responseType: typeof analytics_pb.TopSoldItemsResponse;
};

type AnalyticsLeastSoldItems = {
  readonly methodName: string;
  readonly service: typeof Analytics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof analytics_pb.LeastSoldItemsRequest;
  readonly responseType: typeof analytics_pb.LeastSoldItemsResponse;
};

type AnalyticsTotalSalesAmount = {
  readonly methodName: string;
  readonly service: typeof Analytics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof analytics_pb.TotalSalesAmountRequest;
  readonly responseType: typeof analytics_pb.TotalSalesAmountResponse;
};

export class Analytics {
  static readonly serviceName: string;
  static readonly Ping: AnalyticsPing;
  static readonly PeriodicSalesAmount: AnalyticsPeriodicSalesAmount;
  static readonly TopSoldItems: AnalyticsTopSoldItems;
  static readonly LeastSoldItems: AnalyticsLeastSoldItems;
  static readonly TotalSalesAmount: AnalyticsTotalSalesAmount;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AnalyticsClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  ping(
    requestMessage: analytics_pb.PingRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.PingResponse|null) => void
  ): UnaryResponse;
  ping(
    requestMessage: analytics_pb.PingRequest,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.PingResponse|null) => void
  ): UnaryResponse;
  periodicSalesAmount(
    requestMessage: analytics_pb.PeriodicSalesAmountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.PeriodicSalesAmountResponse|null) => void
  ): UnaryResponse;
  periodicSalesAmount(
    requestMessage: analytics_pb.PeriodicSalesAmountRequest,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.PeriodicSalesAmountResponse|null) => void
  ): UnaryResponse;
  topSoldItems(
    requestMessage: analytics_pb.TopSoldItemsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.TopSoldItemsResponse|null) => void
  ): UnaryResponse;
  topSoldItems(
    requestMessage: analytics_pb.TopSoldItemsRequest,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.TopSoldItemsResponse|null) => void
  ): UnaryResponse;
  leastSoldItems(
    requestMessage: analytics_pb.LeastSoldItemsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.LeastSoldItemsResponse|null) => void
  ): UnaryResponse;
  leastSoldItems(
    requestMessage: analytics_pb.LeastSoldItemsRequest,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.LeastSoldItemsResponse|null) => void
  ): UnaryResponse;
  totalSalesAmount(
    requestMessage: analytics_pb.TotalSalesAmountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.TotalSalesAmountResponse|null) => void
  ): UnaryResponse;
  totalSalesAmount(
    requestMessage: analytics_pb.TotalSalesAmountRequest,
    callback: (error: ServiceError|null, responseMessage: analytics_pb.TotalSalesAmountResponse|null) => void
  ): UnaryResponse;
}

