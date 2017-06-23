import { Component, Input } from '@angular/core';
import { $WebSocket, WebSocketSendMode } from 'angular2-websocket/angular2-websocket'
import { Chat, Chats } from './chat_pb';
import { Door, Context } from 'ws-door';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.css']
})
export class AppComponent {
	@Input() nick: string;
	@Input() context: string;
	@Input() chats: Chat.AsObject[];
	private ws: $WebSocket;
	private door: Door;
	constructor() {
		this.nick = '';
		this.context = '';
		this.chats = [];
		this.ws = new $WebSocket('ws://localhost:8888/ws');
		this.door = new Door(this.ws);
		this.door.putBind('send', (c: Context) => {
			this.chats = c.toObject(Chats).chatsList.reverse();
		});
		this.door.postBind('send', (c: Context) => {
			this.chats.unshift(c.toObject(Chat));
		});
		this.ws.onMessage((m: MessageEvent) => this.door.onMessage(m), { autoApply: false });
	}

	send() {
		const c = new Chat();
		c.setNick(this.nick);
		c.setContext(this.context);
		this.ws.send(this.door.postBinary('send', c), WebSocketSendMode.Direct, true);
		this.context = '';
	}

	setNick(nick: string) {
		this.nick = nick;
		const c = new Chat();
		c.setNick(nick);
		this.ws.send(this.door.putBinary('nick', c), WebSocketSendMode.Direct, true);
	}

	key(event: any) {
		if (event.keyCode === 13) {
			this.send()
		}
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
