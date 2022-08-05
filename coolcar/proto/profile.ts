/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "profile.v1";

export enum Gender {
  NOT_SPECIFIED = 0,
  MALE = 1,
  FEMALE = 2,
  UNRECOGNIZED = -1,
}

export function genderFromJSON(object: any): Gender {
  switch (object) {
    case 0:
    case "NOT_SPECIFIED":
      return Gender.NOT_SPECIFIED;
    case 1:
    case "MALE":
      return Gender.MALE;
    case 2:
    case "FEMALE":
      return Gender.FEMALE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Gender.UNRECOGNIZED;
  }
}

export function genderToJSON(object: Gender): string {
  switch (object) {
    case Gender.NOT_SPECIFIED:
      return "NOT_SPECIFIED";
    case Gender.MALE:
      return "MALE";
    case Gender.FEMALE:
      return "FEMALE";
    case Gender.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum IdentityStatus {
  UNSUBMITTED = 0,
  PENDING = 1,
  VERIFIED = 2,
  UNRECOGNIZED = -1,
}

export function identityStatusFromJSON(object: any): IdentityStatus {
  switch (object) {
    case 0:
    case "UNSUBMITTED":
      return IdentityStatus.UNSUBMITTED;
    case 1:
    case "PENDING":
      return IdentityStatus.PENDING;
    case 2:
    case "VERIFIED":
      return IdentityStatus.VERIFIED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IdentityStatus.UNRECOGNIZED;
  }
}

export function identityStatusToJSON(object: IdentityStatus): string {
  switch (object) {
    case IdentityStatus.UNSUBMITTED:
      return "UNSUBMITTED";
    case IdentityStatus.PENDING:
      return "PENDING";
    case IdentityStatus.VERIFIED:
      return "VERIFIED";
    case IdentityStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Profile {
  identity: Identity | undefined;
  identityStatus: IdentityStatus;
}

export interface Identity {
  licNumber: string;
  name: string;
  gender: Gender;
  birthDateMillis: number;
}

export interface GeProfileRequest {}

export interface ClearProfileRequest {}

export interface GetProfilePhotoRequest {}

export interface GetProfilePhotoResponse {
  url: string;
}

export interface CreateProfilePhotoRequest {}

export interface CreateProfilePhotoResponse {
  uploadUrl: string;
}

export interface CompleteProfilePhotoRequest {}

export interface ClearProfilePhotoRequest {}

export interface ClearProfilePhotoResponse {}

function createBaseProfile(): Profile {
  return { identity: undefined, identityStatus: 0 };
}

export const Profile = {
  encode(
    message: Profile,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.identity !== undefined) {
      Identity.encode(message.identity, writer.uint32(10).fork()).ldelim();
    }
    if (message.identityStatus !== 0) {
      writer.uint32(16).int32(message.identityStatus);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Profile {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProfile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.identity = Identity.decode(reader, reader.uint32());
          break;
        case 2:
          message.identityStatus = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Profile {
    return {
      identity: isSet(object.identity)
        ? Identity.fromJSON(object.identity)
        : undefined,
      identityStatus: isSet(object.identityStatus)
        ? identityStatusFromJSON(object.identityStatus)
        : 0,
    };
  },

  toJSON(message: Profile): unknown {
    const obj: any = {};
    message.identity !== undefined &&
      (obj.identity = message.identity
        ? Identity.toJSON(message.identity)
        : undefined);
    message.identityStatus !== undefined &&
      (obj.identityStatus = identityStatusToJSON(message.identityStatus));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Profile>, I>>(object: I): Profile {
    const message = createBaseProfile();
    message.identity =
      object.identity !== undefined && object.identity !== null
        ? Identity.fromPartial(object.identity)
        : undefined;
    message.identityStatus = object.identityStatus ?? 0;
    return message;
  },
};

function createBaseIdentity(): Identity {
  return { licNumber: "", name: "", gender: 0, birthDateMillis: 0 };
}

export const Identity = {
  encode(
    message: Identity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.licNumber !== "") {
      writer.uint32(10).string(message.licNumber);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.gender !== 0) {
      writer.uint32(24).int32(message.gender);
    }
    if (message.birthDateMillis !== 0) {
      writer.uint32(32).int64(message.birthDateMillis);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Identity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.licNumber = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.gender = reader.int32() as any;
          break;
        case 4:
          message.birthDateMillis = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Identity {
    return {
      licNumber: isSet(object.licNumber) ? String(object.licNumber) : "",
      name: isSet(object.name) ? String(object.name) : "",
      gender: isSet(object.gender) ? genderFromJSON(object.gender) : 0,
      birthDateMillis: isSet(object.birthDateMillis)
        ? Number(object.birthDateMillis)
        : 0,
    };
  },

  toJSON(message: Identity): unknown {
    const obj: any = {};
    message.licNumber !== undefined && (obj.licNumber = message.licNumber);
    message.name !== undefined && (obj.name = message.name);
    message.gender !== undefined && (obj.gender = genderToJSON(message.gender));
    message.birthDateMillis !== undefined &&
      (obj.birthDateMillis = Math.round(message.birthDateMillis));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Identity>, I>>(object: I): Identity {
    const message = createBaseIdentity();
    message.licNumber = object.licNumber ?? "";
    message.name = object.name ?? "";
    message.gender = object.gender ?? 0;
    message.birthDateMillis = object.birthDateMillis ?? 0;
    return message;
  },
};

function createBaseGeProfileRequest(): GeProfileRequest {
  return {};
}

export const GeProfileRequest = {
  encode(
    _: GeProfileRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GeProfileRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGeProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): GeProfileRequest {
    return {};
  },

  toJSON(_: GeProfileRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GeProfileRequest>, I>>(
    _: I
  ): GeProfileRequest {
    const message = createBaseGeProfileRequest();
    return message;
  },
};

function createBaseClearProfileRequest(): ClearProfileRequest {
  return {};
}

export const ClearProfileRequest = {
  encode(
    _: ClearProfileRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClearProfileRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClearProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): ClearProfileRequest {
    return {};
  },

  toJSON(_: ClearProfileRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClearProfileRequest>, I>>(
    _: I
  ): ClearProfileRequest {
    const message = createBaseClearProfileRequest();
    return message;
  },
};

function createBaseGetProfilePhotoRequest(): GetProfilePhotoRequest {
  return {};
}

export const GetProfilePhotoRequest = {
  encode(
    _: GetProfilePhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): GetProfilePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetProfilePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): GetProfilePhotoRequest {
    return {};
  },

  toJSON(_: GetProfilePhotoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetProfilePhotoRequest>, I>>(
    _: I
  ): GetProfilePhotoRequest {
    const message = createBaseGetProfilePhotoRequest();
    return message;
  },
};

function createBaseGetProfilePhotoResponse(): GetProfilePhotoResponse {
  return { url: "" };
}

export const GetProfilePhotoResponse = {
  encode(
    message: GetProfilePhotoResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): GetProfilePhotoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetProfilePhotoResponse();
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

  fromJSON(object: any): GetProfilePhotoResponse {
    return {
      url: isSet(object.url) ? String(object.url) : "",
    };
  },

  toJSON(message: GetProfilePhotoResponse): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetProfilePhotoResponse>, I>>(
    object: I
  ): GetProfilePhotoResponse {
    const message = createBaseGetProfilePhotoResponse();
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseCreateProfilePhotoRequest(): CreateProfilePhotoRequest {
  return {};
}

export const CreateProfilePhotoRequest = {
  encode(
    _: CreateProfilePhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateProfilePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateProfilePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): CreateProfilePhotoRequest {
    return {};
  },

  toJSON(_: CreateProfilePhotoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateProfilePhotoRequest>, I>>(
    _: I
  ): CreateProfilePhotoRequest {
    const message = createBaseCreateProfilePhotoRequest();
    return message;
  },
};

function createBaseCreateProfilePhotoResponse(): CreateProfilePhotoResponse {
  return { uploadUrl: "" };
}

export const CreateProfilePhotoResponse = {
  encode(
    message: CreateProfilePhotoResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uploadUrl !== "") {
      writer.uint32(10).string(message.uploadUrl);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateProfilePhotoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateProfilePhotoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uploadUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateProfilePhotoResponse {
    return {
      uploadUrl: isSet(object.uploadUrl) ? String(object.uploadUrl) : "",
    };
  },

  toJSON(message: CreateProfilePhotoResponse): unknown {
    const obj: any = {};
    message.uploadUrl !== undefined && (obj.uploadUrl = message.uploadUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateProfilePhotoResponse>, I>>(
    object: I
  ): CreateProfilePhotoResponse {
    const message = createBaseCreateProfilePhotoResponse();
    message.uploadUrl = object.uploadUrl ?? "";
    return message;
  },
};

function createBaseCompleteProfilePhotoRequest(): CompleteProfilePhotoRequest {
  return {};
}

export const CompleteProfilePhotoRequest = {
  encode(
    _: CompleteProfilePhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CompleteProfilePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCompleteProfilePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): CompleteProfilePhotoRequest {
    return {};
  },

  toJSON(_: CompleteProfilePhotoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CompleteProfilePhotoRequest>, I>>(
    _: I
  ): CompleteProfilePhotoRequest {
    const message = createBaseCompleteProfilePhotoRequest();
    return message;
  },
};

function createBaseClearProfilePhotoRequest(): ClearProfilePhotoRequest {
  return {};
}

export const ClearProfilePhotoRequest = {
  encode(
    _: ClearProfilePhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ClearProfilePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClearProfilePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): ClearProfilePhotoRequest {
    return {};
  },

  toJSON(_: ClearProfilePhotoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClearProfilePhotoRequest>, I>>(
    _: I
  ): ClearProfilePhotoRequest {
    const message = createBaseClearProfilePhotoRequest();
    return message;
  },
};

function createBaseClearProfilePhotoResponse(): ClearProfilePhotoResponse {
  return {};
}

export const ClearProfilePhotoResponse = {
  encode(
    _: ClearProfilePhotoResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ClearProfilePhotoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClearProfilePhotoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): ClearProfilePhotoResponse {
    return {};
  },

  toJSON(_: ClearProfilePhotoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClearProfilePhotoResponse>, I>>(
    _: I
  ): ClearProfilePhotoResponse {
    const message = createBaseClearProfilePhotoResponse();
    return message;
  },
};

export interface ProfileService {
  GetProfile(request: GeProfileRequest): Promise<Profile>;
  SubmitProfile(request: Identity): Promise<Profile>;
  ClearProfile(request: ClearProfileRequest): Promise<Profile>;
  GetProfilePhoto(
    request: GetProfilePhotoRequest
  ): Promise<GetProfilePhotoResponse>;
  CreateProfilePhoto(
    request: CreateProfilePhotoRequest
  ): Promise<CreateProfilePhotoResponse>;
  CompleteProfilePhoto(request: CompleteProfilePhotoRequest): Promise<Identity>;
  ClearProfilePhoto(
    request: ClearProfilePhotoRequest
  ): Promise<ClearProfilePhotoResponse>;
}

export class ProfileServiceClientImpl implements ProfileService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetProfile = this.GetProfile.bind(this);
    this.SubmitProfile = this.SubmitProfile.bind(this);
    this.ClearProfile = this.ClearProfile.bind(this);
    this.GetProfilePhoto = this.GetProfilePhoto.bind(this);
    this.CreateProfilePhoto = this.CreateProfilePhoto.bind(this);
    this.CompleteProfilePhoto = this.CompleteProfilePhoto.bind(this);
    this.ClearProfilePhoto = this.ClearProfilePhoto.bind(this);
  }
  GetProfile(request: GeProfileRequest): Promise<Profile> {
    const data = GeProfileRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "GetProfile",
      data
    );
    return promise.then((data) => Profile.decode(new _m0.Reader(data)));
  }

  SubmitProfile(request: Identity): Promise<Profile> {
    const data = Identity.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "SubmitProfile",
      data
    );
    return promise.then((data) => Profile.decode(new _m0.Reader(data)));
  }

  ClearProfile(request: ClearProfileRequest): Promise<Profile> {
    const data = ClearProfileRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "ClearProfile",
      data
    );
    return promise.then((data) => Profile.decode(new _m0.Reader(data)));
  }

  GetProfilePhoto(
    request: GetProfilePhotoRequest
  ): Promise<GetProfilePhotoResponse> {
    const data = GetProfilePhotoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "GetProfilePhoto",
      data
    );
    return promise.then((data) =>
      GetProfilePhotoResponse.decode(new _m0.Reader(data))
    );
  }

  CreateProfilePhoto(
    request: CreateProfilePhotoRequest
  ): Promise<CreateProfilePhotoResponse> {
    const data = CreateProfilePhotoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "CreateProfilePhoto",
      data
    );
    return promise.then((data) =>
      CreateProfilePhotoResponse.decode(new _m0.Reader(data))
    );
  }

  CompleteProfilePhoto(
    request: CompleteProfilePhotoRequest
  ): Promise<Identity> {
    const data = CompleteProfilePhotoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "CompleteProfilePhoto",
      data
    );
    return promise.then((data) => Identity.decode(new _m0.Reader(data)));
  }

  ClearProfilePhoto(
    request: ClearProfilePhotoRequest
  ): Promise<ClearProfilePhotoResponse> {
    const data = ClearProfilePhotoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "profile.v1.ProfileService",
      "ClearProfilePhoto",
      data
    );
    return promise.then((data) =>
      ClearProfilePhotoResponse.decode(new _m0.Reader(data))
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
