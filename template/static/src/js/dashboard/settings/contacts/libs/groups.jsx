var contactgroups = {
    get: function(id, cb) {
        request.get("/api/v1/contacts/groups/" + id, function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb([], data.Messages);
            }
        });
    },
    getAll: function(cb) {
        request.get("/api/v1/contacts/groups/", function(data) {
            if (data.Result === "success") {
                cb(data.Data, []);
            } else {
                cb({data: [], errors: data.Messages});
            }
        });
    },
    add: function(name, cb) {
        request.post(
            "/api/v1/contacts/groups/",
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully added contact group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    remove: function(id, cb) {
        request.remove(
            "/api/v1/contacts/groups/" + id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb([{
                        Type: "success",
                        Content: "Successfully removed contact group."
                    }]);
                } else {
                    if (cb) cb(res.Messages);
                }
            }
        );
    },
    update: function(id, name, cb) {
        request.put(
            "/api/v1/contacts/groups/" + id,
            {name: name},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(true, [{
                        Type: "success",
                        Content: "Successfully updated contact group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    getContacts: function(id, cb) {
        request.get(
            "/api/v1/contacts/groups/" + id + "/contacts/",
            function(data) {
                if (data.Result === "success") {
                    cb(data.Data, []);
                } else {
                    cb([], data.Messages);
                }
            }
        );
    },
    addContact: function(id, contact_id, cb) {
        request.post(
            "/api/v1/contacts/groups/" + id + "/contacts/",
            {id: parseInt(contact_id, 10)},
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully added contact to group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    },
    removeContact: function(id, contact_id, cb) {
        request.remove(
            "/api/v1/contacts/groups/" + id + "/contacts/" + contact_id,
            function(res) {
                if (res.Result == "success") {
                    if (cb) cb(res.Data, [{
                        Type: "success",
                        Content: "Successfully removed contact from group."
                    }]);
                } else {
                    if (cb) cb(false, res.Messages);
                }
            }
        );
    }
}
