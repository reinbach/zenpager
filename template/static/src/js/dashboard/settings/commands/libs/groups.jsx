var commandgroups = {
    get: function(id, cb) {
        request.get("/api/v1/commands/groups/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/commands/groups/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, cb) {
        request.post(
            "/api/v1/commands/groups/",
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added command group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/commands/groups/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed command group."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, cb) {
        request.put(
            "/api/v1/commands/groups/" + id,
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated command group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getCommands: function(id, cb) {
        request.get(
            "/api/v1/commands/groups/" + id + "/commands/",
            function(data) {
                if (data.Result === "success") {
                    cb(data.Data, []);
                } else {
                    cb([], data.Messages);
                }
            }
        );
    },
    addCommand: function(id, command_id, cb) {
        request.post(
            "/api/v1/commands/groups/" + id + "/commands/",
            {id: parseInt(command_id, 10)},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully added command to group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    removeCommand: function(id, command_id, cb) {
        request.remove(
            "/api/v1/commands/groups/" + id + "/commands/" + command_id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully removed command from group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    }
}
