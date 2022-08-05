/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "blob.v1";

export interface CreateBlobRequest {
  accountId: string;
  uploadUrlTimeoutSec: number;
}

export interface CreateBlobResponse {
  id: string;
  uploadUrl: string;
}

export interface GetBlobRequest {
  id: string;
}

export interface GetBlobResponse {
  data: Uint8Array;
}

export interface GetBlobURLRequest {
  id: string;
  timeoutSec: number;
}

export interface GetBlobURLResponse {
  url: string;
}

function createBaseCreateBlobRequest(): CreateBlobRequest {
  return { accountId: "", uploadUrlTimeoutSec: 0 };
}

export const CreateBlobRequest = {
  encode(
    message: CreateBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.uploadUrlTimeoutSec !== 0) {
      writer.uint32(16).int32(message.uploadUrlTimeoutSec);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        case 2:
          message.uploadUrlTimeoutSec = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateBlobRequest {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      uploadUrlTimeoutSec: isSet(object.uploadUrlTimeoutSec)
        ? Number(object.uploadUrlTimeoutSec)
        : 0,
    };
  },

  toJSON(message: CreateBlobRequest): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.uploadUrlTimeoutSec !== undefined &&
      (obj.uploadUrlTimeoutSec = Math.round(message.uploadUrlTimeoutSec));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBlobRequest>, I>>(
    object: I
  ): CreateBlobRequest {
    const message = createBaseCreateBlobRequest();
    message.accountId = object.accountId ?? "";
    message.uploadUrlTimeoutSec = object.uploadUrlTimeoutSec ?? 0;
    return message;
  },
};

function createBaseCreateBlobResponse(): CreateBlobResponse {
  return { id: "", uploadUrl: "" };
}

export const CreateBlobResponse = {
  encode(
    message: CreateBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.uploadUrl !== "") {
      writer.uint32(18).string(message.uploadUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.uploadUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateBlobResponse {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      uploadUrl: isSet(object.uploadUrl) ? String(object.uploadUrl) : "",
    };
  },

  toJSON(message: CreateBlobResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.uploadUrl !== undefined && (obj.uploadUrl = message.uploadUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBlobResponse>, I>>(
    object: I
  ): CreateBlobResponse {
    const message = createBaseCreateBlobResponse();
    message.id = object.id ?? "";
    message.uploadUrl = object.uploadUrl ?? "";
    return message;
  },
};

function createBaseGetBlobRequest(): GetBlobRequest {
  return { id: "" };
}

export const GetBlobRequest = {
  encode(
    message: GetBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetBlobRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: GetBlobRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetBlobRequest>, I>>(
    object: I
  ): GetBlobRequest {
    const message = createBaseGetBlobRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetBlobResponse(): GetBlobResponse {
  return { data: new Uint8Array() };
}

export const GetBlobResponse = {
  encode(
    message: GetBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data.length !== 0) {
      writer.uint32(10).bytes(message.data);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetBlobResponse {
    return {
      data: isSet(object.data)
        ? bytesFromBase64(object.data)
        : new Uint8Array(),
    };
  },

  toJSON(message: GetBlobResponse): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetBlobResponse>, I>>(
    object: I
  ): GetBlobResponse {
    const message = createBaseGetBlobResponse();
    message.data = object.data ?? new Uint8Array();
    return message;
  },
};

function createBaseGetBlobURLRequest(): GetBlobURLRequest {
  return { id: "", timeoutSec: 0 };
}

export const GetBlobURLRequest = {
  encode(
    message: GetBlobURLRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.timeoutSec !== 0) {
      writer.uint32(16).int32(message.timeoutSec);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBlobURLRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBlobURLRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.timeoutSec = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetBlobURLRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      timeoutSec: isSet(object.timeoutSec) ? Number(object.timeoutSec) : 0,
    };
  },

  toJSON(message: GetBlobURLRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.timeoutSec !== undefined &&
      (obj.timeoutSec = Math.round(message.timeoutSec));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetBlobURLRequest>, I>>(
    object: I
  ): GetBlobURLRequest {
    const message = createBaseGetBlobURLRequest();
    message.id = object.id ?? "";
    message.timeoutSec = object.timeoutSec ?? 0;
    return message;
  },
};

function createBaseGetBlobURLResponse(): GetBlobURLResponse {
  return { url: "" };
}

export const GetBlobURLResponse = {
  encode(
    message: GetBlobURLResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBlobURLResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBlobURLResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetBlobURLResponse {
    return {
      url: isSet(object.url) ? String(object.url) : "",
    };
  },

  toJSON(message: GetBlobURLResponse): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetBlobURLResponse>, I>>(
    object: I
  ): GetBlobURLResponse {
    const message = createBaseGetBlobURLResponse();
    message.url = object.url ?? "";
    return message;
  },
};

export interface BlobService {
  CreateBlob(request: CreateBlobRequest): Promise<CreateBlobResponse>;
  GetBlob(request: GetBlobRequest): Promise<GetBlobResponse>;
  GetBlobURL(request: GetBlobURLRequest): Promise<GetBlobURLResponse>;
}

export class BlobServiceClientImpl implements BlobService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateBlob = this.CreateBlob.bind(this);
    this.GetBlob = this.GetBlob.bind(this);
    this.GetBlobURL = this.GetBlobURL.bind(this);
  }
  CreateBlob(request: CreateBlobRequest): Promise<CreateBlobResponse> {
    const data = CreateBlobRequest.encode(request).finish();
    const promise = this.rpc.request("blob.v1.BlobService", "CreateBlob", data);
    return promise.then((data) =>
      CreateBlobResponse.decode(new _m0.Reader(data))
    );
  }

  GetBlob(request: GetBlobRequest): Promise<GetBlobResponse> {
    const data = GetBlobRequest.encode(request).finish();
    const promise = this.rpc.request("blob.v1.BlobService", "GetBlob", data);
    return promise.then((data) => GetBlobResponse.decode(new _m0.Reader(data)));
  }

  GetBlobURL(request: GetBlobURLRequest): Promise<GetBlobURLResponse> {
    const data = GetBlobURLRequest.encode(request).finish();
    const promise = this.rpc.request("blob.v1.BlobService", "GetBlobURL", data);
    return promise.then((data) =>
      GetBlobURLResponse.decode(new _m0.Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  arr.forEach((byte) => {
    bin.push(String.fromCharCode(byte));
  });
  return btoa(bin.join(""));
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
