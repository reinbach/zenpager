var servers = {
    get: function(id, cb) {
        request.get("/api/v1/servers/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/servers/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, url, cb) {
        request.post(
            "/api/v1/servers/",
            {name: name, url: url},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added server."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/servers/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed server."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, url, cb) {
        request.put(
            "/api/v1/servers/" + id,
            {name: name, url: url},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated server."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getGroups: function(id, cb) {
        request.get(
            "/api/v1/servers/" + id + "/groups/",
            function(data) {
                if (data.Result === "success") {
                    cb(data.Data, []);
                } else {
                    cb([], data.Messages);
                }
            }
        );
    },
}
