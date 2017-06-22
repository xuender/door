// package: door
// file: event.proto

import * as jspb from "google-protobuf";

export class Event extends jspb.Message {
  getPath(): string;
  setPath(value: string): void;

  getMethod(): MethodEnum;
  setMethod(value: MethodEnum): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    path: string,
    method: MethodEnum,
    data: Uint8Array | string,
  }
}

export enum MethodEnum {
  GET = 0,
  POST = 1,
  PUT = 2,
  DELETE = 3,
}

