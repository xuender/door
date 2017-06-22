import { Component, Input } from '@angular/core';
import { $WebSocket } from 'angular2-websocket/angular2-websocket'
import { Chat } from './chat_pb';
import { Event } from './event_pb';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.css']
})
export class AppComponent {
	@Input() nick: string;
	@Input() context: string;
	@Input() chats: Chat[];
	private ws: $WebSocket;
	constructor() {
		this.nick = '路人';
		this.context = '';
		this.chats = [];
		console.log('...........')
		this.ws = new $WebSocket('ws://localhost:8888/ws');
		this.ws.onMessage((msg: MessageEvent) => {
			const reader = new FileReader();
			reader.readAsArrayBuffer(msg.data);
			reader.onload = function () {
				const buf = new Uint8Array(reader.result);
				const pb = Event.deserializeBinary(buf);
				const e = pb.toObject();
				console.log('event', e);
			}
		}, {autoApply: false});
	}
	send() {
		console.log('text', this.context);
		const c = new Chat();
		c.setNick(this.nick);
		c.setContext(this.context);
		this.chats.push(c);
		this.context = '';
	}
	key(event: any) {
		if (event.keyCode === 13) {
			this.send()
		}
	}
}
