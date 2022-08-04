/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "rental.v1";

export enum TripStatus {
  NOT_SPECIFIED = 0,
  IN_PROGRESS = 1,
  FINISHED = 2,
  UNRECOGNIZED = -1,
}

export function tripStatusFromJSON(object: any): TripStatus {
  switch (object) {
    case 0:
    case "NOT_SPECIFIED":
      return TripStatus.NOT_SPECIFIED;
    case 1:
    case "IN_PROGRESS":
      return TripStatus.IN_PROGRESS;
    case 2:
    case "FINISHED":
      return TripStatus.FINISHED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TripStatus.UNRECOGNIZED;
  }
}

export function tripStatusToJSON(object: TripStatus): string {
  switch (object) {
    case TripStatus.NOT_SPECIFIED:
      return "NOT_SPECIFIED";
    case TripStatus.IN_PROGRESS:
      return "IN_PROGRESS";
    case TripStatus.FINISHED:
      return "FINISHED";
    case TripStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface TripEntity {
  id: string;
  trip: Trip | undefined;
}

export interface Trip {
  accountId: string;
  carId: string;
  start: LocationStatus | undefined;
  current: LocationStatus | undefined;
  end: LocationStatus | undefined;
  status: TripStatus;
  identityId: string;
}

export interface CreateTripRequest {
  start: Location | undefined;
  carId: string;
  avatarUrl: string;
}

export interface CreateTripResponse {
  start: Location | undefined;
  carId: string;
  avatarUrl: string;
}

export interface GetTripRequest {
  id: string;
}

export interface GetTripsRequest {
  status: TripStatus;
}

export interface GetTripsResponse {
  trips: TripEntity[];
}

export interface UpdateTripRequest {
  id: string;
  current: Location | undefined;
  endTrip: boolean;
}

export interface Location {
  latitude: number;
  longitude: number;
}

export interface LocationStatus {
  location: Location | undefined;
  feeCent: number;
  kmDriven: number;
  poiName: string;
  timestampSec: number;
}

function createBaseTripEntity(): TripEntity {
  return { id: "", trip: undefined };
}

export const TripEntity = {
  encode(
    message: TripEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.trip !== undefined) {
      Trip.encode(message.trip, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TripEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTripEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.trip = Trip.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TripEntity {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      trip: isSet(object.trip) ? Trip.fromJSON(object.trip) : undefined,
    };
  },

  toJSON(message: TripEntity): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.trip !== undefined &&
      (obj.trip = message.trip ? Trip.toJSON(message.trip) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TripEntity>, I>>(
    object: I
  ): TripEntity {
    const message = createBaseTripEntity();
    message.id = object.id ?? "";
    message.trip =
      object.trip !== undefined && object.trip !== null
        ? Trip.fromPartial(object.trip)
        : undefined;
    return message;
  },
};

function createBaseTrip(): Trip {
  return {
    accountId: "",
    carId: "",
    start: undefined,
    current: undefined,
    end: undefined,
    status: 0,
    identityId: "",
  };
}

export const Trip = {
  encode(message: Trip, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.carId !== "") {
      writer.uint32(18).string(message.carId);
    }
    if (message.start !== undefined) {
      LocationStatus.encode(message.start, writer.uint32(26).fork()).ldelim();
    }
    if (message.current !== undefined) {
      LocationStatus.encode(message.current, writer.uint32(34).fork()).ldelim();
    }
    if (message.end !== undefined) {
      LocationStatus.encode(message.end, writer.uint32(42).fork()).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    if (message.identityId !== "") {
      writer.uint32(58).string(message.identityId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Trip {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTrip();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        case 2:
          message.carId = reader.string();
          break;
        case 3:
          message.start = LocationStatus.decode(reader, reader.uint32());
          break;
        case 4:
          message.current = LocationStatus.decode(reader, reader.uint32());
          break;
        case 5:
          message.end = LocationStatus.decode(reader, reader.uint32());
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.identityId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Trip {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      carId: isSet(object.carId) ? String(object.carId) : "",
      start: isSet(object.start)
        ? LocationStatus.fromJSON(object.start)
        : undefined,
      current: isSet(object.current)
        ? LocationStatus.fromJSON(object.current)
        : undefined,
      end: isSet(object.end) ? LocationStatus.fromJSON(object.end) : undefined,
      status: isSet(object.status) ? tripStatusFromJSON(object.status) : 0,
      identityId: isSet(object.identityId) ? String(object.identityId) : "",
    };
  },

  toJSON(message: Trip): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.carId !== undefined && (obj.carId = message.carId);
    message.start !== undefined &&
      (obj.start = message.start
        ? LocationStatus.toJSON(message.start)
        : undefined);
    message.current !== undefined &&
      (obj.current = message.current
        ? LocationStatus.toJSON(message.current)
        : undefined);
    message.end !== undefined &&
      (obj.end = message.end ? LocationStatus.toJSON(message.end) : undefined);
    message.status !== undefined &&
      (obj.status = tripStatusToJSON(message.status));
    message.identityId !== undefined && (obj.identityId = message.identityId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Trip>, I>>(object: I): Trip {
    const message = createBaseTrip();
    message.accountId = object.accountId ?? "";
    message.carId = object.carId ?? "";
    message.start =
      object.start !== undefined && object.start !== null
        ? LocationStatus.fromPartial(object.start)
        : undefined;
    message.current =
      object.current !== undefined && object.current !== null
        ? LocationStatus.fromPartial(object.current)
        : undefined;
    message.end =
      object.end !== undefined && object.end !== null
        ? LocationStatus.fromPartial(object.end)
        : undefined;
    message.status = object.status ?? 0;
    message.identityId = object.identityId ?? "";
    return message;
  },
};

function createBaseCreateTripRequest(): CreateTripRequest {
  return { start: undefined, carId: "", avatarUrl: "" };
}

export const CreateTripRequest = {
  encode(
    message: CreateTripRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.start !== undefined) {
      Location.encode(message.start, writer.uint32(10).fork()).ldelim();
    }
    if (message.carId !== "") {
      writer.uint32(18).string(message.carId);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(26).string(message.avatarUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTripRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTripRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.start = Location.decode(reader, reader.uint32());
          break;
        case 2:
          message.carId = reader.string();
          break;
        case 3:
          message.avatarUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateTripRequest {
    return {
      start: isSet(object.start) ? Location.fromJSON(object.start) : undefined,
      carId: isSet(object.carId) ? String(object.carId) : "",
      avatarUrl: isSet(object.avatarUrl) ? String(object.avatarUrl) : "",
    };
  },

  toJSON(message: CreateTripRequest): unknown {
    const obj: any = {};
    message.start !== undefined &&
      (obj.start = message.start ? Location.toJSON(message.start) : undefined);
    message.carId !== undefined && (obj.carId = message.carId);
    message.avatarUrl !== undefined && (obj.avatarUrl = message.avatarUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateTripRequest>, I>>(
    object: I
  ): CreateTripRequest {
    const message = createBaseCreateTripRequest();
    message.start =
      object.start !== undefined && object.start !== null
        ? Location.fromPartial(object.start)
        : undefined;
    message.carId = object.carId ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    return message;
  },
};

function createBaseCreateTripResponse(): CreateTripResponse {
  return { start: undefined, carId: "", avatarUrl: "" };
}

export const CreateTripResponse = {
  encode(
    message: CreateTripResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.start !== undefined) {
      Location.encode(message.start, writer.uint32(10).fork()).ldelim();
    }
    if (message.carId !== "") {
      writer.uint32(18).string(message.carId);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(26).string(message.avatarUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTripResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTripResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.start = Location.decode(reader, reader.uint32());
          break;
        case 2:
          message.carId = reader.string();
          break;
        case 3:
          message.avatarUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateTripResponse {
    return {
      start: isSet(object.start) ? Location.fromJSON(object.start) : undefined,
      carId: isSet(object.carId) ? String(object.carId) : "",
      avatarUrl: isSet(object.avatarUrl) ? String(object.avatarUrl) : "",
    };
  },

  toJSON(message: CreateTripResponse): unknown {
    const obj: any = {};
    message.start !== undefined &&
      (obj.start = message.start ? Location.toJSON(message.start) : undefined);
    message.carId !== undefined && (obj.carId = message.carId);
    message.avatarUrl !== undefined && (obj.avatarUrl = message.avatarUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateTripResponse>, I>>(
    object: I
  ): CreateTripResponse {
    const message = createBaseCreateTripResponse();
    message.start =
      object.start !== undefined && object.start !== null
        ? Location.fromPartial(object.start)
        : undefined;
    message.carId = object.carId ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    return message;
  },
};

function createBaseGetTripRequest(): GetTripRequest {
  return { id: "" };
}

export const GetTripRequest = {
  encode(
    message: GetTripRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTripRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTripRequest();
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

  fromJSON(object: any): GetTripRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: GetTripRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetTripRequest>, I>>(
    object: I
  ): GetTripRequest {
    const message = createBaseGetTripRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetTripsRequest(): GetTripsRequest {
  return { status: 0 };
}

export const GetTripsRequest = {
  encode(
    message: GetTripsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTripsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTripsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetTripsRequest {
    return {
      status: isSet(object.status) ? tripStatusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: GetTripsRequest): unknown {
    const obj: any = {};
    message.status !== undefined &&
      (obj.status = tripStatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetTripsRequest>, I>>(
    object: I
  ): GetTripsRequest {
    const message = createBaseGetTripsRequest();
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseGetTripsResponse(): GetTripsResponse {
  return { trips: [] };
}

export const GetTripsResponse = {
  encode(
    message: GetTripsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.trips) {
      TripEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTripsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTripsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.trips.push(TripEntity.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetTripsResponse {
    return {
      trips: Array.isArray(object?.trips)
        ? object.trips.map((e: any) => TripEntity.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetTripsResponse): unknown {
    const obj: any = {};
    if (message.trips) {
      obj.trips = message.trips.map((e) =>
        e ? TripEntity.toJSON(e) : undefined
      );
    } else {
      obj.trips = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetTripsResponse>, I>>(
    object: I
  ): GetTripsResponse {
    const message = createBaseGetTripsResponse();
    message.trips = object.trips?.map((e) => TripEntity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateTripRequest(): UpdateTripRequest {
  return { id: "", current: undefined, endTrip: false };
}

export const UpdateTripRequest = {
  encode(
    message: UpdateTripRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.current !== undefined) {
      Location.encode(message.current, writer.uint32(18).fork()).ldelim();
    }
    if (message.endTrip === true) {
      writer.uint32(24).bool(message.endTrip);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTripRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTripRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.current = Location.decode(reader, reader.uint32());
          break;
        case 3:
          message.endTrip = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateTripRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      current: isSet(object.current)
        ? Location.fromJSON(object.current)
        : undefined,
      endTrip: isSet(object.endTrip) ? Boolean(object.endTrip) : false,
    };
  },

  toJSON(message: UpdateTripRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.current !== undefined &&
      (obj.current = message.current
        ? Location.toJSON(message.current)
        : undefined);
    message.endTrip !== undefined && (obj.endTrip = message.endTrip);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTripRequest>, I>>(
    object: I
  ): UpdateTripRequest {
    const message = createBaseUpdateTripRequest();
    message.id = object.id ?? "";
    message.current =
      object.current !== undefined && object.current !== null
        ? Location.fromPartial(object.current)
        : undefined;
    message.endTrip = object.endTrip ?? false;
    return message;
  },
};

function createBaseLocation(): Location {
  return { latitude: 0, longitude: 0 };
}

export const Location = {
  encode(
    message: Location,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.latitude !== 0) {
      writer.uint32(9).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(17).double(message.longitude);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Location {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.latitude = reader.double();
          break;
        case 2:
          message.longitude = reader.double();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Location {
    return {
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
    };
  },

  toJSON(message: Location): unknown {
    const obj: any = {};
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Location>, I>>(object: I): Location {
    const message = createBaseLocation();
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
    return message;
  },
};

function createBaseLocationStatus(): LocationStatus {
  return {
    location: undefined,
    feeCent: 0,
    kmDriven: 0,
    poiName: "",
    timestampSec: 0,
  };
}

export const LocationStatus = {
  encode(
    message: LocationStatus,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.location !== undefined) {
      Location.encode(message.location, writer.uint32(10).fork()).ldelim();
    }
    if (message.feeCent !== 0) {
      writer.uint32(16).int32(message.feeCent);
    }
    if (message.kmDriven !== 0) {
      writer.uint32(25).double(message.kmDriven);
    }
    if (message.poiName !== "") {
      writer.uint32(34).string(message.poiName);
    }
    if (message.timestampSec !== 0) {
      writer.uint32(40).int64(message.timestampSec);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LocationStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocationStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.location = Location.decode(reader, reader.uint32());
          break;
        case 2:
          message.feeCent = reader.int32();
          break;
        case 3:
          message.kmDriven = reader.double();
          break;
        case 4:
          message.poiName = reader.string();
          break;
        case 5:
          message.timestampSec = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LocationStatus {
    return {
      location: isSet(object.location)
        ? Location.fromJSON(object.location)
        : undefined,
      feeCent: isSet(object.feeCent) ? Number(object.feeCent) : 0,
      kmDriven: isSet(object.kmDriven) ? Number(object.kmDriven) : 0,
      poiName: isSet(object.poiName) ? String(object.poiName) : "",
      timestampSec: isSet(object.timestampSec)
        ? Number(object.timestampSec)
        : 0,
    };
  },

  toJSON(message: LocationStatus): unknown {
    const obj: any = {};
    message.location !== undefined &&
      (obj.location = message.location
        ? Location.toJSON(message.location)
        : undefined);
    message.feeCent !== undefined &&
      (obj.feeCent = Math.round(message.feeCent));
    message.kmDriven !== undefined && (obj.kmDriven = message.kmDriven);
    message.poiName !== undefined && (obj.poiName = message.poiName);
    message.timestampSec !== undefined &&
      (obj.timestampSec = Math.round(message.timestampSec));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LocationStatus>, I>>(
    object: I
  ): LocationStatus {
    const message = createBaseLocationStatus();
    message.location =
      object.location !== undefined && object.location !== null
        ? Location.fromPartial(object.location)
        : undefined;
    message.feeCent = object.feeCent ?? 0;
    message.kmDriven = object.kmDriven ?? 0;
    message.poiName = object.poiName ?? "";
    message.timestampSec = object.timestampSec ?? 0;
    return message;
  },
};

export interface TripService {
  CreateTrip(request: CreateTripRequest): Promise<TripEntity>;
  GetTrip(request: GetTripRequest): Promise<Trip>;
  GetTrips(request: GetTripsRequest): Promise<GetTripsResponse>;
  UpdateTrip(request: UpdateTripRequest): Promise<Trip>;
}

export class TripServiceClientImpl implements TripService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateTrip = this.CreateTrip.bind(this);
    this.GetTrip = this.GetTrip.bind(this);
    this.GetTrips = this.GetTrips.bind(this);
    this.UpdateTrip = this.UpdateTrip.bind(this);
  }
  CreateTrip(request: CreateTripRequest): Promise<TripEntity> {
    const data = CreateTripRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rental.v1.TripService",
      "CreateTrip",
      data
    );
    return promise.then((data) => TripEntity.decode(new _m0.Reader(data)));
  }

  GetTrip(request: GetTripRequest): Promise<Trip> {
    const data = GetTripRequest.encode(request).finish();
    const promise = this.rpc.request("rental.v1.TripService", "GetTrip", data);
    return promise.then((data) => Trip.decode(new _m0.Reader(data)));
  }

  GetTrips(request: GetTripsRequest): Promise<GetTripsResponse> {
    const data = GetTripsRequest.encode(request).finish();
    const promise = this.rpc.request("rental.v1.TripService", "GetTrips", data);
    return promise.then((data) =>
      GetTripsResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateTrip(request: UpdateTripRequest): Promise<Trip> {
    const data = UpdateTripRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rental.v1.TripService",
      "UpdateTrip",
      data
    );
    return promise.then((data) => Trip.decode(new _m0.Reader(data)));
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
