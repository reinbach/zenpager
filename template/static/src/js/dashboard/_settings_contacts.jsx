var contacts = {
    get: function(cb) {
        var r = new XMLHttpRequest();
        r.open("GET", "/api/v1/contacts/", true);
        r.setRequestHeader("Content-Type", "application/json");
        r.setRequestHeader("X-Access-Token", auth.getToken());
        r.onreadystatechange = function() {
            if (r.readyState === 4) {
                data = JSON.parse(r.responseText);
                if (data.Result === "success") {
                    cb({data: data});
                } else {
                    cb({data: [], errors: data.Messages});
                }
            }
        };
        r.send();
    }
}

var SettingsContacts = React.createClass({
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
    render: function() {
        return (
            <div>
                add the contacts here!!!
            </div>
        );
    }
});
