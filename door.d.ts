import * as jspb from 'google-protobuf';
export declare class Context {
    private bytes;
    constructor(bytes: Uint8Array);
    toObject(pb: any): any;
}
export interface Router {
    path: string;
    handler(c: Context): void;
}
export declare class Door {
    private routes;
    constructor();
    openBind(handler: (c: Context) => void, ...paths: string[]): void;
    closeBind(handler: (c: Context) => void, ...paths: string[]): void;
    getBind(handler: (c: Context) => void, ...paths: string[]): void;
    putBind(handler: (c: Context) => void, ...paths: string[]): void;
    postBind(handler: (c: Context) => void, ...paths: string[]): void;
    deleteBind(handler: (c: Context) => void, ...paths: string[]): void;
    private add(handler, method, paths);
    onMessage(msg: MessageEvent): void;
    private serializeBinary(pb, method, paths);
    openBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
    closeBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
    getBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
    postBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
    putBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
    deleteBinary(pb: jspb.Message, ...paths: string[]): Uint8Array;
}
