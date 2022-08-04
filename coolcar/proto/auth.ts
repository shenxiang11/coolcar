/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "auth.v1";

export interface LoginRequest {
  code: string;
}

export interface LoginResponse {
  accessToken: string;
  expiresIn: number;
}

function createBaseLoginRequest(): LoginRequest {
  return { code: "" };
}

export const LoginRequest = {
  encode(
    message: LoginRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginRequest {
    return {
      code: isSet(object.code) ? String(object.code) : "",
    };
  },

  toJSON(message: LoginRequest): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginRequest>, I>>(
    object: I
  ): LoginRequest {
    const message = createBaseLoginRequest();
    message.code = object.code ?? "";
    return message;
  },
};

function createBaseLoginResponse(): LoginResponse {
  return { accessToken: "", expiresIn: 0 };
}

export const LoginResponse = {
  encode(
    message: LoginResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.accessToken !== "") {
      writer.uint32(10).string(message.accessToken);
    }
    if (message.expiresIn !== 0) {
      writer.uint32(16).int32(message.expiresIn);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accessToken = reader.string();
          break;
        case 2:
          message.expiresIn = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginResponse {
    return {
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : "",
      expiresIn: isSet(object.expiresIn) ? Number(object.expiresIn) : 0,
    };
  },

  toJSON(message: LoginResponse): unknown {
    const obj: any = {};
    message.accessToken !== undefined &&
      (obj.accessToken = message.accessToken);
    message.expiresIn !== undefined &&
      (obj.expiresIn = Math.round(message.expiresIn));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginResponse>, I>>(
    object: I
  ): LoginResponse {
    const message = createBaseLoginResponse();
    message.accessToken = object.accessToken ?? "";
    message.expiresIn = object.expiresIn ?? 0;
    return message;
  },
};

export interface AuthService {
  Login(request: LoginRequest): Promise<LoginResponse>;
}

export class AuthServiceClientImpl implements AuthService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Login = this.Login.bind(this);
  }
  Login(request: LoginRequest): Promise<LoginResponse> {
    const data = LoginRequest.encode(request).finish();
    const promise = this.rpc.request("auth.v1.AuthService", "Login", data);
    return promise.then((data) => LoginResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
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
