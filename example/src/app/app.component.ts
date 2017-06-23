import { Component, Input } from '@angular/core';
import { $WebSocket, WebSocketSendMode } from 'angular2-websocket/angular2-websocket'
import { Chat, Chats } from './chat_pb';
import { Door, Event, MethodEnum, Context } from 'ws-door';

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
		this.nick = '路人';
		this.context = '';
		this.chats = [];
		this.ws = new $WebSocket('ws://localhost:8888/ws');
		this.door = new Door();
		this.door.PUT('send', (c: Context) => {
			this.chats = c.toObject().chatsList;
		});
		this.door.POST('send', (c: Context) => {
			this.chats.push(c.toObject());
		});
		this.ws.onMessage((m: MessageEvent) => this.door.onMessage(m), {autoApply: false});
	}

	send() {
		console.log('text', this.context);
		const c = new Chat();
		c.setNick(this.nick);
		c.setContext(this.context);
		const e = new Event();
		e.setMethod(MethodEnum.POST);
		e.setPath('send');
		e.setData(c.serializeBinary());
		this.ws.send(e.serializeBinary(), WebSocketSendMode.Direct, true);
		this.context = '';
	}

	key(event: any) {
		if (event.keyCode === 13) {
			this.send()
		}
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
