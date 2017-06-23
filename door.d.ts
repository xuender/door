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
    openBind(path: string, handler: (c: Context) => void): void;
    closeBind(path: string, handler: (c: Context) => void): void;
    getBind(path: string, handler: (c: Context) => void): void;
    putBind(path: string, handler: (c: Context) => void): void;
    postBind(path: string, handler: (c: Context) => void): void;
    deleteBind(path: string, handler: (c: Context) => void): void;
    private add(method, path, handler);
    onMessage(msg: MessageEvent): void;
    private serializeBinary(method, path, pb);
    openBinary(path: string, pb: jspb.Message): Uint8Array;
    closeBinary(path: string, pb: jspb.Message): Uint8Array;
    getBinary(path: string, pb: jspb.Message): Uint8Array;
    putBinary(path: string, pb: jspb.Message): Uint8Array;
    postBinary(path: string, pb: jspb.Message): Uint8Array;
    deleteBinary(path: string, pb: jspb.Message): Uint8Array;
}
