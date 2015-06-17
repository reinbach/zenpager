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
    },
    add: function(name, email, cb) {
        callback = cb;
        request.post(
            "/api/v1/contacts/",
            {name: name, email: email},
            this.processAdd
        );
    },
    processAdd: function(res) {
        if (res.Result == "success") {
            if (callback) callback(true, [{
                Type: "success",
                Content: "Successfully added contact."
            }]);
        } else {
            if (callback) callback(false, res.Messages);
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
    getInitialState: function() {
        return {
            contacts: [],
            messages: []
        };
    },
    componentWillMount: function() {
        contacts.get(function(data, messages) {
            this.setState({
                contacts: data,
                messages: messages
            });
        }.bind(this));
    },
    render: function() {
        return (
            <div>
                <Link to="s_contacts_add" className="btn btn-primary">Add Contact</Link>
                <p>Settings Contacts... we need more</p>
            </div>
        );
    }
});

var SettingsContactsAdd = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    getInitialState: function() {
        return {
            name: "",
            email: "",
            messages: []
        }
    },
    validateNameState: function() {
        if (this.state.name.length > 0) {
            if (this.state.name.length > 2) {
                return "success";
            }
            return "error"
        }
    },
    validateEmailState: function() {
        if (this.state.email.length > 0) {
            if (validateEmail(this.state.email) === true) {
                return "success";
            }
            return "error"
        }
    },
    handleChange: function() {
        this.setState({
            name: this.refs.name.getValue(),
            email: this.refs.email.getValue()
        });
    },
    handleSubmit: function() {
        event.preventDefault();
        if (this.state.name.length < 1) {
            this.setState({
                messages: [{Type: "danger", Content: "Name is required."}]
            });
            return ;
        }
        if (this.state.email.length < 1) {
            this.setState({
                messages: [{Type: "danger", Content: "Email is required."}]
            });
            return ;
        }
        if (this.validateNameState() !== "success" || this.validateEmailState() !== "success") {
            this.setState({
                messages: [{Type: "danger", Content: "Fix errors"}]
            });
            return ;
        }
        contacts.add(this.state.name, this.state.email, function(success, messages) {
            if (success == true) {
                this.setState({messages: messages, name: "", email: ""});
            } else {
                this.setState({
                    messages: messages,
                    name: this.state.name,
                    email: this.state.email
                });
            }
        }.bind(this));
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div className="col-md-4">
                <form onSubmit={this.handleSubmit} className="text-left">
                    <h2 className="page-header">Add Contact</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="Jane Smart" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Input label="Email" type="text" ref="email"
                           placeholder="test@example.com" value={this.state.email}
                           hasFeedback bsStyle={this.validateEmailState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        Add Contact
                    </Button>
                    <Link to="s_contacts_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});
