var commands = {
    get: function(id, cb) {
        request.get("/api/v1/commands/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/commands/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, command, cb) {
        request.post(
            "/api/v1/commands/",
            {name: name, command: command},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added command."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/commands/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed command."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, command, cb) {
        request.put(
            "/api/v1/commands/" + id,
            {name: name, command: command},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated command."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getGroups: function(id, cb) {
        request.get(
            "/api/v1/commands/" + id + "/groups/",
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
