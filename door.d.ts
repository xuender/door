import { MethodEnum } from './event_pb';
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
    routes: Router[][];
    constructor();
    OPEN(path: string, handler: (c: Context) => void): void;
    CLOSE(path: string, handler: (c: Context) => void): void;
    GET(path: string, handler: (c: Context) => void): void;
    POST(path: string, handler: (c: Context) => void): void;
    PUT(path: string, handler: (c: Context) => void): void;
    DELETE(path: string, handler: (c: Context) => void): void;
    add(method: MethodEnum, path: string, handler: (c: Context) => void): void;
    onMessage(msg: MessageEvent): void;
}
