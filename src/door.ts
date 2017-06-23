import * as jspb from "google-protobuf";
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
	routes: Router[][];
	constructor() {
		this.routes = [];
	}

	OPEN(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.OPEN, path, handler);
	}

	CLOSE(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.CLOSE, path, handler);
	}

	GET(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.GET, path, handler);
	}

	POST(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.POST, path, handler);
	}

	PUT(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.PUT, path, handler);
	}

	DELETE(path: string, handler: (c: Context) => void) {
		this.add(MethodEnum.DELETE, path, handler);
	}

	add (method: MethodEnum, path: string, handler:(c: Context) => void) {
		let routes = this.routes[method];
		if (!routes) {
			routes = [];
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
					if (r.path === event.path){
						r.handler(new Context(event.data as Uint8Array));
						return;
					}
				}
			});
		}
}

function readFile (blob: Blob): Promise<ArrayBuffer> {
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
