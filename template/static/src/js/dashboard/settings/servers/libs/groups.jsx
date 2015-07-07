var servergroups = {
    get: function(id, cb) {
        request.get("/api/v1/servers/groups/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/servers/groups/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, cb) {
        request.post(
            "/api/v1/servers/groups/",
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added server group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/servers/groups/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed server group."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, cb) {
        request.put(
            "/api/v1/servers/groups/" + id,
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated server group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getServers: function(id, cb) {
        request.get(
            "/api/v1/servers/groups/" + id + "/servers/",
            function(data) {
                if (data.Result === "success") {
                    cb(data.Data, []);
                } else {
                    cb([], data.Messages);
                }
            }
        );
    },
    addServer: function(id, server_id, cb) {
        request.post(
            "/api/v1/servers/groups/" + id + "/servers/",
            {id: parseInt(server_id, 10)},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully added server to group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    removeServer: function(id, server_id, cb) {
        request.remove(
            "/api/v1/servers/groups/" + id + "/servers/" + server_id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully removed server from group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    }
}
