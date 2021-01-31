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

