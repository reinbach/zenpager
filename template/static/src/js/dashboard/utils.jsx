var request = {
    get: function(url, cb) {
        var r = new XMLHttpRequest();
        r.open("GET", url, true);
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
        r.send();
    },
    post: function(url, data, cb) {
        var r = new XMLHttpRequest();
        r.open("POST", url, true);
        r.setRequestHeader("Content-Type", "application/json");
        r.setRequestHeader("X-Access-Token", auth.getToken());
        r.onreadystatechange = function() {
            if (r.readyState === 4) {
                if (r.status == 401) {
                    auth.logout();
                } else {
                    rdata = JSON.parse(r.responseText);
                    if (cb) {
                        cb(rdata);
                    }
                }
            }
        };
        r.send(JSON.stringify(data));
    },
    patch: function(url, data, cb) {
        var r = new XMLHttpRequest();
        r.open("PATCH", url, true);
        r.setRequestHeader("Content-Type", "application/json");
        r.setRequestHeader("X-Access-Token", auth.getToken());
        r.onreadystatechange = function() {
            if (r.readyState === 4) {
                if (r.status == 401) {
                    auth.logout();
                } else {
                    rdata = JSON.parse(r.responseText);
                    if (cb) {
                        cb(rdata);
                    }
                }
            }
        };
        r.send(JSON.stringify(data));
    }
}
