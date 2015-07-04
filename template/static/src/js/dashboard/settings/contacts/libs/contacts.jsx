var contacts = {
    get: function(id, cb) {
        request.get("/api/v1/contacts/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/contacts/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, email, cb) {
        request.post(
            "/api/v1/contacts/",
            {name: name, email: email},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added contact."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/contacts/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed contact."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, email, cb) {
        request.put(
            "/api/v1/contacts/" + id,
            {name: name, email: email},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated contact."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getGroups: function(id, cb) {
        request.get(
            "/api/v1/contacts/" + id + "/groups/",
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
