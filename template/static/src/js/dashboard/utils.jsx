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
    post: function(url, data, cb) {
        var r = this.call("POST", url, cb);
        r.send(JSON.stringify(data));
    },
    patch: function(url, data, cb) {
        var r = this.call("PATH", data, cb);
        r.send(JSON.stringify(data));
    }
}

function removeItem(l, o) {
    var n = [];
    for (var i = 0; i < l.length; i++) {
        if (l[i] != o) {
            n.push(l[i]);
        }
    }
    return n;
}
