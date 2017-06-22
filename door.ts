import { MethodEnum } from './event_pb';

export class Context {

}

interface Router {
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
}
