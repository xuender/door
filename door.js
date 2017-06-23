"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var event_pb_1 = require("./event_pb");
var Context = (function () {
    function Context(bytes) {
        this.bytes = bytes;
    }
    Context.prototype.toObject = function (pb) {
        var m = pb.deserializeBinary(this.bytes);
        return m.toObject();
    };
    return Context;
}());
exports.Context = Context;
var Door = (function () {
    function Door() {
        this.routes = [];
    }
    Door.prototype.OPEN = function (path, handler) {
        this.add(event_pb_1.MethodEnum.OPEN, path, handler);
    };
    Door.prototype.CLOSE = function (path, handler) {
        this.add(event_pb_1.MethodEnum.CLOSE, path, handler);
    };
    Door.prototype.GET = function (path, handler) {
        this.add(event_pb_1.MethodEnum.GET, path, handler);
    };
    Door.prototype.POST = function (path, handler) {
        this.add(event_pb_1.MethodEnum.POST, path, handler);
    };
    Door.prototype.PUT = function (path, handler) {
        this.add(event_pb_1.MethodEnum.PUT, path, handler);
    };
    Door.prototype.DELETE = function (path, handler) {
        this.add(event_pb_1.MethodEnum.DELETE, path, handler);
    };
    Door.prototype.add = function (method, path, handler) {
        var routes = this.routes[method];
        if (!routes) {
            routes = [];
            this.routes[method] = routes;
        }
        for (var _i = 0, routes_1 = routes; _i < routes_1.length; _i++) {
            var r = routes_1[_i];
            if (r.path === path) {
                r.handler = handler;
                return;
            }
        }
        routes.push({
            path: path,
            handler: handler,
        });
    };
    Door.prototype.onMessage = function (msg) {
        var _this = this;
        console.log('onMessage');
        readFile(msg.data).then(function (buffer) {
            var array = new Uint8Array(buffer);
            var event = event_pb_1.Event.deserializeBinary(array).toObject();
            var routes = _this.routes[event.method];
            if (!routes)
                return;
            for (var _i = 0, routes_2 = routes; _i < routes_2.length; _i++) {
                var r = routes_2[_i];
                if (r.path === event.path) {
                    r.handler(new Context(event.data));
                    return;
                }
            }
        });
    };
    return Door;
}());
exports.Door = Door;
function readFile(blob) {
    return new Promise(function (resolve, reject) {
        var reader = new FileReader();
        reader.onloadend = function () {
            resolve(reader.result);
        };
        reader.onerror = function (event) {
            reject(event.error);
        };
        reader.readAsArrayBuffer(blob);
    });
}
//# sourceMappingURL=door.js.map