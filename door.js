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
    Door.prototype.openBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.OPEN, paths);
    };
    Door.prototype.closeBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.CLOSE, paths);
    };
    Door.prototype.getBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.GET, paths);
    };
    Door.prototype.putBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.PUT, paths);
    };
    Door.prototype.postBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.POST, paths);
    };
    Door.prototype.deleteBind = function (handler) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        this.add(handler, event_pb_1.MethodEnum.DELETE, paths);
    };
    Door.prototype.add = function (handler, method, paths) {
        var path = paths.join('/');
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
    Door.prototype.serializeBinary = function (pb, method, paths) {
        var e = new event_pb_1.Event();
        e.setMethod(method);
        e.setPath(paths.join('/'));
        e.setData(pb.serializeBinary());
        return e.serializeBinary();
    };
    Door.prototype.openBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.OPEN, paths);
    };
    Door.prototype.closeBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.CLOSE, paths);
    };
    Door.prototype.getBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.GET, paths);
    };
    Door.prototype.postBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.POST, paths);
    };
    Door.prototype.putBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.PUT, paths);
    };
    Door.prototype.deleteBinary = function (pb) {
        var paths = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            paths[_i - 1] = arguments[_i];
        }
        return this.serializeBinary(pb, event_pb_1.MethodEnum.DELETE, paths);
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