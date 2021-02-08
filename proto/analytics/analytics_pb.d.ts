// package: analytics
// file: analytics.proto

import * as jspb from "google-protobuf";

export class PingRequest extends jspb.Message {
  getPing(): string;
  setPing(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingRequest): PingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingRequest;
  static deserializeBinaryFromReader(message: PingRequest, reader: jspb.BinaryReader): PingRequest;
}

export namespace PingRequest {
  export type AsObject = {
    ping: string,
  }
}

export class PingResponse extends jspb.Message {
  getPong(): string;
  setPong(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PingResponse): PingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingResponse;
  static deserializeBinaryFromReader(message: PingResponse, reader: jspb.BinaryReader): PingResponse;
}

export namespace PingResponse {
  export type AsObject = {
    pong: string,
  }
}

export class SalesData extends jspb.Message {
  getTimestamp(): number;
  setTimestamp(value: number): void;

  getBillcount(): number;
  setBillcount(value: number): void;

  getTotalamount(): number;
  setTotalamount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SalesData.AsObject;
  static toObject(includeInstance: boolean, msg: SalesData): SalesData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SalesData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SalesData;
  static deserializeBinaryFromReader(message: SalesData, reader: jspb.BinaryReader): SalesData;
}

export namespace SalesData {
  export type AsObject = {
    timestamp: number,
    billcount: number,
    totalamount: number,
  }
}

export class PeriodicSalesAmountRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getInterval(): string;
  setInterval(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeriodicSalesAmountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PeriodicSalesAmountRequest): PeriodicSalesAmountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeriodicSalesAmountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeriodicSalesAmountRequest;
  static deserializeBinaryFromReader(message: PeriodicSalesAmountRequest, reader: jspb.BinaryReader): PeriodicSalesAmountRequest;
}

export namespace PeriodicSalesAmountRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    interval: string,
    from: number,
    to: number,
  }
}

export class PeriodicSalesAmountResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getTotal(): number;
  setTotal(value: number): void;

  clearSalesdataList(): void;
  getSalesdataList(): Array<SalesData>;
  setSalesdataList(value: Array<SalesData>): void;
  addSalesdata(value?: SalesData, index?: number): SalesData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeriodicSalesAmountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PeriodicSalesAmountResponse): PeriodicSalesAmountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeriodicSalesAmountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeriodicSalesAmountResponse;
  static deserializeBinaryFromReader(message: PeriodicSalesAmountResponse, reader: jspb.BinaryReader): PeriodicSalesAmountResponse;
}

export namespace PeriodicSalesAmountResponse {
  export type AsObject = {
    success: boolean,
    total: number,
    salesdataList: Array<SalesData.AsObject>,
  }
}

export class StoreSalesData extends jspb.Message {
  getStoreid(): string;
  setStoreid(value: string): void;

  getBillcount(): number;
  setBillcount(value: number): void;

  getTotal(): number;
  setTotal(value: number): void;

  clearSalesdataList(): void;
  getSalesdataList(): Array<SalesData>;
  setSalesdataList(value: Array<SalesData>): void;
  addSalesdata(value?: SalesData, index?: number): SalesData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreSalesData.AsObject;
  static toObject(includeInstance: boolean, msg: StoreSalesData): StoreSalesData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreSalesData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreSalesData;
  static deserializeBinaryFromReader(message: StoreSalesData, reader: jspb.BinaryReader): StoreSalesData;
}

export namespace StoreSalesData {
  export type AsObject = {
    storeid: string,
    billcount: number,
    total: number,
    salesdataList: Array<SalesData.AsObject>,
  }
}

export class PeriodicStoreSalesAmountRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStorecount(): number;
  setStorecount(value: number): void;

  getInterval(): string;
  setInterval(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeriodicStoreSalesAmountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PeriodicStoreSalesAmountRequest): PeriodicStoreSalesAmountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeriodicStoreSalesAmountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeriodicStoreSalesAmountRequest;
  static deserializeBinaryFromReader(message: PeriodicStoreSalesAmountRequest, reader: jspb.BinaryReader): PeriodicStoreSalesAmountRequest;
}

export namespace PeriodicStoreSalesAmountRequest {
  export type AsObject = {
    channelname: string,
    storecount: number,
    interval: string,
    from: number,
    to: number,
  }
}

export class PeriodicStoreSalesAmountResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearStoresalesdataList(): void;
  getStoresalesdataList(): Array<StoreSalesData>;
  setStoresalesdataList(value: Array<StoreSalesData>): void;
  addStoresalesdata(value?: StoreSalesData, index?: number): StoreSalesData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeriodicStoreSalesAmountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PeriodicStoreSalesAmountResponse): PeriodicStoreSalesAmountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeriodicStoreSalesAmountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeriodicStoreSalesAmountResponse;
  static deserializeBinaryFromReader(message: PeriodicStoreSalesAmountResponse, reader: jspb.BinaryReader): PeriodicStoreSalesAmountResponse;
}

export namespace PeriodicStoreSalesAmountResponse {
  export type AsObject = {
    success: boolean,
    storesalesdataList: Array<StoreSalesData.AsObject>,
  }
}

export class SoldItemsData extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SoldItemsData.AsObject;
  static toObject(includeInstance: boolean, msg: SoldItemsData): SoldItemsData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SoldItemsData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SoldItemsData;
  static deserializeBinaryFromReader(message: SoldItemsData, reader: jspb.BinaryReader): SoldItemsData;
}

export namespace SoldItemsData {
  export type AsObject = {
    key: string,
    count: number,
  }
}

export class TopSoldItemsRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopSoldItemsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TopSoldItemsRequest): TopSoldItemsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TopSoldItemsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopSoldItemsRequest;
  static deserializeBinaryFromReader(message: TopSoldItemsRequest, reader: jspb.BinaryReader): TopSoldItemsRequest;
}

export namespace TopSoldItemsRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
  }
}

export class TopSoldItemsResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearSolditemsdataList(): void;
  getSolditemsdataList(): Array<SoldItemsData>;
  setSolditemsdataList(value: Array<SoldItemsData>): void;
  addSolditemsdata(value?: SoldItemsData, index?: number): SoldItemsData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopSoldItemsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TopSoldItemsResponse): TopSoldItemsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TopSoldItemsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopSoldItemsResponse;
  static deserializeBinaryFromReader(message: TopSoldItemsResponse, reader: jspb.BinaryReader): TopSoldItemsResponse;
}

export namespace TopSoldItemsResponse {
  export type AsObject = {
    success: boolean,
    solditemsdataList: Array<SoldItemsData.AsObject>,
  }
}

export class LeastSoldItemsRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeastSoldItemsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LeastSoldItemsRequest): LeastSoldItemsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LeastSoldItemsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeastSoldItemsRequest;
  static deserializeBinaryFromReader(message: LeastSoldItemsRequest, reader: jspb.BinaryReader): LeastSoldItemsRequest;
}

export namespace LeastSoldItemsRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
  }
}

export class LeastSoldItemsResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearSolditemsdataList(): void;
  getSolditemsdataList(): Array<SoldItemsData>;
  setSolditemsdataList(value: Array<SoldItemsData>): void;
  addSolditemsdata(value?: SoldItemsData, index?: number): SoldItemsData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeastSoldItemsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LeastSoldItemsResponse): LeastSoldItemsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LeastSoldItemsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeastSoldItemsResponse;
  static deserializeBinaryFromReader(message: LeastSoldItemsResponse, reader: jspb.BinaryReader): LeastSoldItemsResponse;
}

export namespace LeastSoldItemsResponse {
  export type AsObject = {
    success: boolean,
    solditemsdataList: Array<SoldItemsData.AsObject>,
  }
}

export class RevenueItemsData extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getBillcount(): number;
  setBillcount(value: number): void;

  getTotalsalesprice(): number;
  setTotalsalesprice(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RevenueItemsData.AsObject;
  static toObject(includeInstance: boolean, msg: RevenueItemsData): RevenueItemsData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RevenueItemsData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RevenueItemsData;
  static deserializeBinaryFromReader(message: RevenueItemsData, reader: jspb.BinaryReader): RevenueItemsData;
}

export namespace RevenueItemsData {
  export type AsObject = {
    key: string,
    billcount: number,
    totalsalesprice: number,
  }
}

export class TopRevenueItemsRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopRevenueItemsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TopRevenueItemsRequest): TopRevenueItemsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TopRevenueItemsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopRevenueItemsRequest;
  static deserializeBinaryFromReader(message: TopRevenueItemsRequest, reader: jspb.BinaryReader): TopRevenueItemsRequest;
}

export namespace TopRevenueItemsRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
  }
}

export class TopRevenueItemsResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearRevenueitemsdataList(): void;
  getRevenueitemsdataList(): Array<RevenueItemsData>;
  setRevenueitemsdataList(value: Array<RevenueItemsData>): void;
  addRevenueitemsdata(value?: RevenueItemsData, index?: number): RevenueItemsData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopRevenueItemsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TopRevenueItemsResponse): TopRevenueItemsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TopRevenueItemsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopRevenueItemsResponse;
  static deserializeBinaryFromReader(message: TopRevenueItemsResponse, reader: jspb.BinaryReader): TopRevenueItemsResponse;
}

export namespace TopRevenueItemsResponse {
  export type AsObject = {
    success: boolean,
    revenueitemsdataList: Array<RevenueItemsData.AsObject>,
  }
}

export class LeastRevenueItemsRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeastRevenueItemsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LeastRevenueItemsRequest): LeastRevenueItemsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LeastRevenueItemsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeastRevenueItemsRequest;
  static deserializeBinaryFromReader(message: LeastRevenueItemsRequest, reader: jspb.BinaryReader): LeastRevenueItemsRequest;
}

export namespace LeastRevenueItemsRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
  }
}

export class LeastRevenueItemsResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearRevenueitemsdataList(): void;
  getRevenueitemsdataList(): Array<RevenueItemsData>;
  setRevenueitemsdataList(value: Array<RevenueItemsData>): void;
  addRevenueitemsdata(value?: RevenueItemsData, index?: number): RevenueItemsData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeastRevenueItemsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LeastRevenueItemsResponse): LeastRevenueItemsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LeastRevenueItemsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeastRevenueItemsResponse;
  static deserializeBinaryFromReader(message: LeastRevenueItemsResponse, reader: jspb.BinaryReader): LeastRevenueItemsResponse;
}

export namespace LeastRevenueItemsResponse {
  export type AsObject = {
    success: boolean,
    revenueitemsdataList: Array<RevenueItemsData.AsObject>,
  }
}

export class StoreRevenueData extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getBillcount(): number;
  setBillcount(value: number): void;

  getTotalrevenue(): number;
  setTotalrevenue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRevenueData.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRevenueData): StoreRevenueData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreRevenueData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRevenueData;
  static deserializeBinaryFromReader(message: StoreRevenueData, reader: jspb.BinaryReader): StoreRevenueData;
}

export namespace StoreRevenueData {
  export type AsObject = {
    key: string,
    billcount: number,
    totalrevenue: number,
  }
}

export class StoreRevenueDataRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  getAscending(): boolean;
  setAscending(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRevenueDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRevenueDataRequest): StoreRevenueDataRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreRevenueDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRevenueDataRequest;
  static deserializeBinaryFromReader(message: StoreRevenueDataRequest, reader: jspb.BinaryReader): StoreRevenueDataRequest;
}

export namespace StoreRevenueDataRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
    ascending: boolean,
  }
}

export class StoreRevenueDataResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearStorerevenuedataList(): void;
  getStorerevenuedataList(): Array<StoreRevenueData>;
  setStorerevenuedataList(value: Array<StoreRevenueData>): void;
  addStorerevenuedata(value?: StoreRevenueData, index?: number): StoreRevenueData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRevenueDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRevenueDataResponse): StoreRevenueDataResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreRevenueDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRevenueDataResponse;
  static deserializeBinaryFromReader(message: StoreRevenueDataResponse, reader: jspb.BinaryReader): StoreRevenueDataResponse;
}

export namespace StoreRevenueDataResponse {
  export type AsObject = {
    success: boolean,
    storerevenuedataList: Array<StoreRevenueData.AsObject>,
  }
}

export class StoreBillData extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getBillcount(): number;
  setBillcount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreBillData.AsObject;
  static toObject(includeInstance: boolean, msg: StoreBillData): StoreBillData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreBillData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreBillData;
  static deserializeBinaryFromReader(message: StoreBillData, reader: jspb.BinaryReader): StoreBillData;
}

export namespace StoreBillData {
  export type AsObject = {
    key: string,
    billcount: number,
  }
}

export class StoreSalesDataRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getCount(): number;
  setCount(value: number): void;

  getAscending(): boolean;
  setAscending(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreSalesDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreSalesDataRequest): StoreSalesDataRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreSalesDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreSalesDataRequest;
  static deserializeBinaryFromReader(message: StoreSalesDataRequest, reader: jspb.BinaryReader): StoreSalesDataRequest;
}

export namespace StoreSalesDataRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
    count: number,
    ascending: boolean,
  }
}

export class StoreSalesDataResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  clearStorebilldataList(): void;
  getStorebilldataList(): Array<StoreBillData>;
  setStorebilldataList(value: Array<StoreBillData>): void;
  addStorebilldata(value?: StoreBillData, index?: number): StoreBillData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreSalesDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StoreSalesDataResponse): StoreSalesDataResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreSalesDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreSalesDataResponse;
  static deserializeBinaryFromReader(message: StoreSalesDataResponse, reader: jspb.BinaryReader): StoreSalesDataResponse;
}

export namespace StoreSalesDataResponse {
  export type AsObject = {
    success: boolean,
    storebilldataList: Array<StoreBillData.AsObject>,
  }
}

export class AverageSalesAmountRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AverageSalesAmountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AverageSalesAmountRequest): AverageSalesAmountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AverageSalesAmountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AverageSalesAmountRequest;
  static deserializeBinaryFromReader(message: AverageSalesAmountRequest, reader: jspb.BinaryReader): AverageSalesAmountRequest;
}

export namespace AverageSalesAmountRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
  }
}

export class AverageSalesAmountResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getAmount(): number;
  setAmount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AverageSalesAmountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AverageSalesAmountResponse): AverageSalesAmountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AverageSalesAmountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AverageSalesAmountResponse;
  static deserializeBinaryFromReader(message: AverageSalesAmountResponse, reader: jspb.BinaryReader): AverageSalesAmountResponse;
}

export namespace AverageSalesAmountResponse {
  export type AsObject = {
    success: boolean,
    amount: number,
  }
}

export class AverageSoldQuantityRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AverageSoldQuantityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AverageSoldQuantityRequest): AverageSoldQuantityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AverageSoldQuantityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AverageSoldQuantityRequest;
  static deserializeBinaryFromReader(message: AverageSoldQuantityRequest, reader: jspb.BinaryReader): AverageSoldQuantityRequest;
}

export namespace AverageSoldQuantityRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
  }
}

export class AverageSoldQuantityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getQuantity(): number;
  setQuantity(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AverageSoldQuantityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AverageSoldQuantityResponse): AverageSoldQuantityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AverageSoldQuantityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AverageSoldQuantityResponse;
  static deserializeBinaryFromReader(message: AverageSoldQuantityResponse, reader: jspb.BinaryReader): AverageSoldQuantityResponse;
}

export namespace AverageSoldQuantityResponse {
  export type AsObject = {
    success: boolean,
    quantity: number,
  }
}

export class TotalSalesAmountRequest extends jspb.Message {
  getChannelname(): string;
  setChannelname(value: string): void;

  getStoreid(): string;
  setStoreid(value: string): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TotalSalesAmountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TotalSalesAmountRequest): TotalSalesAmountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TotalSalesAmountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TotalSalesAmountRequest;
  static deserializeBinaryFromReader(message: TotalSalesAmountRequest, reader: jspb.BinaryReader): TotalSalesAmountRequest;
}

export namespace TotalSalesAmountRequest {
  export type AsObject = {
    channelname: string,
    storeid: string,
    from: number,
    to: number,
  }
}

export class TotalSalesAmountResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getAmount(): number;
  setAmount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TotalSalesAmountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TotalSalesAmountResponse): TotalSalesAmountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TotalSalesAmountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TotalSalesAmountResponse;
  static deserializeBinaryFromReader(message: TotalSalesAmountResponse, reader: jspb.BinaryReader): TotalSalesAmountResponse;
}

export namespace TotalSalesAmountResponse {
  export type AsObject = {
    success: boolean,
    amount: number,
  }
}

