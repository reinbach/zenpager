var contacts = {
    init: function() {
        $(".contacts-link").addClass("active");
    },
    close: function() {
        $(".contacts-link").removeClass("active");
    },
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

var SettingsContactsMixin = {
    componentDidMount: function() {
        contacts.init();
    },
    componentWillUnmount: function() {
        contacts.close();
    }
};

var SettingsContactsHolder = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    render: function() {
        return (
            <div>
                <h1 className="page-header">Contacts</h1>
                <RouteHandler />
            </div>
        );
    }
});

var SettingsContacts = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
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
                <Link to="s_contacts_add"><Button bsStyle="primary">Add Contact</Button></Link>
                <p>Settings Contacts... we need more</p>
            </div>
        );
    }
});

var SettingsContactsAdd = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    render: function() {
        return (
            <div>
                add the contacts here!!!
            </div>
        );
    }
});
