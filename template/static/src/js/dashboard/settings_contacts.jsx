var contacts = {
    get: function(cb) {
        callback = cb;
        request.get("/api/v1/contacts/", this.processGet);
    },
    processGet: function(data) {
        if (data.Result === "success") {
            callback({data: data});
        } else {
            callback({data: [], errors: data.Messages});
        }
    }
}

var SettingsContacts = React.createClass({
    mixins: [Authentication],
    propTypes: {
        contacts: React.PropTypes.array,
        messages: React.PropTypes.array
    },
    getDefaultProps: function() {
        contacts.get(function(data, messages) {
            return {
                contacts: data,
                messages: messages
            };
        });
    },
    render: function() {
        return (
            <div>
                Settings Contacts... we need more
            </div>
        );
    }
});

var SettingsContactsAdd = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div>
                add the contacts here!!!
            </div>
        );
    }
});
