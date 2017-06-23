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

	openBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.OPEN, path, handler);
	}

	closeBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.CLOSE, path, handler);
	}

	getBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.GET, path, handler);
	}

	putBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.PUT, path, handler);
	}

	postBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.POST, path, handler);
	}

	deleteBind(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.DELETE, path, handler);
	}

	private add(method: MethodEnum, path: string, handler: (c: Context) => void) {
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
		console.log('onMessage');
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

	private serializeBinary(method: MethodEnum, path: string, pb: jspb.Message): Uint8Array{
		const e = new Event();
		e.setMethod(method);
		e.setPath(path);
		e.setData(pb.serializeBinary());
		return e.serializeBinary()
	}

	openBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.OPEN, path, pb)
	}

	closeBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.CLOSE, path, pb)
	}

	getBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.GET, path, pb)
	}

	putBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.PUT, path, pb)
	}

	postBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.POST, path, pb)
	}

	deleteBinary(path: string, pb: jspb.Message): Uint8Array {
		return this.serializeBinary(MethodEnum.DELETE, path, pb)
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
