import * as jspb from 'google-protobuf';
import { Event, MethodEnum } from './event_pb';

export class Context {
	private bytes: Uint8Array;

	constructor(bytes: Uint8Array) {
		this.bytes = bytes;
	}

	toObject(pb: any): any {
		const m = pb.deserializeBinary(this.bytes);
		return m.toObject();
	}
}

export interface Router {
	path: string;
	handler(c: Context): void;
}

export class Door {
	private routes: Router[][];

	constructor() {
		this.routes = [];
	}

	openBind(handler: (c: Context) => void, ...paths: string[])  {
		this.add(handler, MethodEnum.OPEN, paths);
	}

	closeBind(handler: (c: Context) => void, ...paths: string[]) {
		this.add(handler, MethodEnum.CLOSE, paths);
	}

	getBind(handler: (c: Context) => void, ...paths: string[]) {
		this.add(handler, MethodEnum.GET, paths);
	}

	putBind(handler: (c: Context) => void,...paths: string[]) {
		this.add(handler, MethodEnum.PUT, paths);
	}

	postBind(handler: (c: Context) => void,...paths: string[]) {
		this.add(handler, MethodEnum.POST, paths);
	}

	deleteBind(handler: (c: Context) => void,...paths: string[]) {
		this.add(handler, MethodEnum.DELETE, paths);
	}

	private add(handler: (c: Context) => void, method: MethodEnum, paths: string[]) {
		const path = paths.join('/');
		let routes = this.routes[method];
		if (!routes) {
			routes = [];
			this.routes[method] = routes;
		}
		for (let r of routes) {
			if (r.path === path) {
				r.handler = handler;
				return;
			}
		}
		routes.push({
			path: path,
			handler: handler,
		});
	}

	onMessage(msg: MessageEvent) {
		readFile(msg.data).then((buffer) => {
			const array = new Uint8Array(buffer);
			const event = Event.deserializeBinary(array).toObject();
			const routes = this.routes[event.method];
			if (!routes) return;
			for (let r of routes) {
				if (r.path === event.path) {
					r.handler(new Context(event.data as Uint8Array));
					return;
				}
			}
		});
	}

	private serializeBinary(pb: jspb.Message, method: MethodEnum, paths: string[]): Uint8Array{
		const e = new Event();
		e.setMethod(method);
		e.setPath(paths.join('/'));
		if (pb) {
			e.setData(pb.serializeBinary());
		}
		return e.serializeBinary()
	}

	openBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.OPEN, paths)
	}

	closeBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.CLOSE, paths)
	}

	getBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.GET, paths)
	}

	postBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.POST, paths)
	}

	putBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.PUT, paths)
	}

	deleteBinary(pb: jspb.Message, ...paths: string[]): Uint8Array {
		return this.serializeBinary(pb, MethodEnum.DELETE, paths)
	}
}

function readFile(blob: Blob): Promise<ArrayBuffer> {
	return new Promise<ArrayBuffer>((resolve, reject) => {
		const reader = new FileReader();
		reader.onloadend = () => {
			resolve(reader.result);
		}
		reader.onerror = (event) => {
			reject(event.error);
		}
		reader.readAsArrayBuffer(blob);
	});
}
