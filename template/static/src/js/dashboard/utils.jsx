var request = {
    call : function(method, url, cb) {
        var r = new XMLHttpRequest();
        r.open(method, url, true);
        r.setRequestHeader("Content-Type", "application/json");
        r.setRequestHeader("X-Access-Token", auth.getToken());
        r.onreadystatechange = function() {
            if (r.readyState === 4) {
                if (r.status == 401) {
                    auth.logout();
                } else {
                    data = JSON.parse(r.responseText);
                    if (cb) {
                        cb(data);
                    }
                }
            }
        };
        return r
    },
    remove: function(url, cb) {
        var r = this.call("DELETE", url, cb);
        r.send();
    },
    get: function(url, cb) {
        var r = this.call("GET", url, cb);
        r.send();
    },
    patch: function(url, data, cb) {
        var r = this.call("PATH", url, cb);
        r.send(JSON.stringify(data));
    },
    post: function(url, data, cb) {
        var r = this.call("POST", url, cb);
        r.send(JSON.stringify(data));
    },
    put: function(url, data, cb) {
        var r = this.call("PUT", url, cb);
        r.send(JSON.stringify(data));
    }
}

function removeFromList(l, o) {
    var n = [];
    for (var i = 0; i < l.length; i++) {
        if (l[i] != o) {
            n.push(l[i]);
        }
    }
    return n;
}

function removeFromListByKey(l, o) {
    var n = [];
    for (var i = 0; i < l.length; i++) {
        if (l[i].key != o.id) {
            n.push(l[i]);
        }
    }
    return n;
}

function excludeByKey(l, o) {
    var n = [];
    var k = [];
    for (var i = 0; i < o.length; i++) {
        k.push(o[i].key);
    }

    for (var j = 0; j < l.length; j++) {
        if (k.indexOf(l[j].key) === -1) {
            n.push(l[j]);
        }
    }

    return n;
}
