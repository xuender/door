// package: chat
// file: chat.proto

import * as jspb from "google-protobuf";

export class Chat extends jspb.Message {
  getNick(): string;
  setNick(value: string): void;

  getContext(): string;
  setContext(value: string): void;

  getTimestamp(): number;
  setTimestamp(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Chat.AsObject;
  static toObject(includeInstance: boolean, msg: Chat): Chat.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Chat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Chat;
  static deserializeBinaryFromReader(message: Chat, reader: jspb.BinaryReader): Chat;
}

export namespace Chat {
  export type AsObject = {
    nick: string,
    context: string,
    timestamp: number,
  }
}

export class Chats extends jspb.Message {
  clearChatsList(): void;
  getChatsList(): Array<Chat>;
  setChatsList(value: Array<Chat>): void;
  addChats(value?: Chat, index?: number): Chat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Chats.AsObject;
  static toObject(includeInstance: boolean, msg: Chats): Chats.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Chats, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Chats;
  static deserializeBinaryFromReader(message: Chats, reader: jspb.BinaryReader): Chats;
}

export namespace Chats {
  export type AsObject = {
    chatsList: Array<Chat.AsObject>,
  }
}

