var contacts = {
    init: function() {
        $(".contacts-link").addClass("active");
    },
    close: function() {
        $(".contacts-link").removeClass("active");
    },
    get: function(id, cb) {
        callback = cb;
        request.get("/api/v1/contacts/" + id, this.processGet);
    },
    processGet: function(cb) {
        if (data.Result === "success") {
            callback(data.Data, []);
        } else {
            callback([], data.Messages);
        }
    },
    getAll: function(cb) {
        callback = cb;
        request.get("/api/v1/contacts/", this.processGetAll);
    },
    processGetAll: function(data) {
        if (data.Result === "success") {
            callback(data.Data, []);
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
    },
    remove: function(id, cb) {
        callback = cb
        request.remove(
            "/api/v1/contacts/" + id,
            this.processRemove
        );
    },
    processRemove: function(res) {
        if (res.Result == "success") {
            if (callback) callback([{
                Type: "success",
                Content: "Successfully removed contact."
            }]);
        } else {
            if (callback) callback(res.Messages);
        }
    },
    update: function(id, name, email, cb) {
        callback = cb;
        request.put(
            "/api/v1/contacts/" + id,
            {name: name, email: email},
            this.processUpdate
        );
    },
    processUpdate: function(res) {
        if (res.Result == "success") {
            if (callback) callback(true, [{
                Type: "success",
                Content: "Successfully updated contact."
            }]);
        } else {
            if (callback) callback(false, res.Messages);
        }
    }
}

var groups = {
    get: function(id, cb) {
        callback = cb;
        request.get("/api/v1/contacts/groups/" + id, this.processGet);
    },
    processGet: function(cb) {
        if (data.Result === "success") {
            callback(data.Data, []);
        } else {
            console.log(data);
            callback([], data.Messages);
        }
    },
    getAll: function(cb) {
        callback = cb;
        request.get("/api/v1/contacts/groups/", this.processGetAll);
    },
    processGetAll: function(data) {
        if (data.Result === "success") {
            callback(data.Data, []);
        } else {
            callback({data: [], errors: data.Messages});
        }
    },
    add: function(name, email, cb) {
        callback = cb;
        request.post(
            "/api/v1/contacts/groups/",
            {name: name},
            this.processAdd
        );
    },
    processAdd: function(res) {
        if (res.Result == "success") {
            if (callback) callback(true, [{
                Type: "success",
                Content: "Successfully added contact group."
            }]);
        } else {
            if (callback) callback(false, res.Messages);
        }
    },
    remove: function(id, cb) {
        callback = cb
        request.remove(
            "/api/v1/contacts/groups/" + id,
            this.processRemove
        );
    },
    processRemove: function(res) {
        if (res.Result == "success") {
            if (callback) callback([{
                Type: "success",
                Content: "Successfully removed contact group."
            }]);
        } else {
            if (callback) callback(res.Messages);
        }
    },
    update: function(id, name, cb) {
        callback = cb;
        request.put(
            "/api/v1/contacts/groups/" + id,
            {name: name},
            this.processUpdate
        );
    },
    processUpdate: function(res) {
        if (res.Result == "success") {
            if (callback) callback(true, [{
                Type: "success",
                Content: "Successfully updated contact group."
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
    render: function() {
        return (
            <div>
                <SettingsContactsGroups />
                <SettingsContactsList />
            </div>
        )
    }
});

var SettingsContactsGroups = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    propTypes: {
        groups: React.PropTypes.array,
        messages: React.PropTypes.array
    },
    getInitialState: function() {
        return {
            groups: [],
            messages: []
        };
    },
    componentWillMount: function() {
        groups.getAll(function(data, messages) {
            this.setState({
                groups: data,
                messages: messages
            });
        }.bind(this));
    },
    removeGroup: function(contact) {
        groups.remove(group.id, function(messages) {
            this.setState({messages: messages})
        }.bind(this));
        this.setState({
            groups: removeFromList(this.state.groups, group)
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        var groups = [];
        this.state.groups.forEach(function(group) {
            groups.push(
                <SettingsContactsGroupLine group={group}
                                           removeGroup={this.removeGroup} />
            );
        }.bind(this));
        return (
            <div>
                {msgs}
                <Table striped hover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {groups}
                    </tbody>
                </Table>

                <Link to="s_contacts_group_add"
                      className="btn btn-primary">Add Group</Link>
            </div>
        )
    }
});

var SettingsContactsGroupLine = React.createClass({
    handleDelete: function() {
        this.props.removeGroup(this.props.group);
    },
    render: function() {
        return (
            <tr>
                <td>{this.props.group.name}</td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_contacts_group_update"
                          params={{"groupId": this.props.group.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsContactsList = React.createClass({
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
        contacts.getAll(function(data, messages) {
            this.setState({
                contacts: data,
                messages: messages
            });
        }.bind(this));
    },
    removeContact: function(contact) {
        contacts.remove(contact.id, function(messages) {
            this.setState({messages: messages})
        }.bind(this));
        this.setState({
            contacts: removeFromList(this.state.contacts, contact)
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        var contacts = [];
        this.state.contacts.forEach(function(contact) {
            contacts.push(<SettingsContactsLine contact={contact} removeContact={this.removeContact} />);
        }.bind(this));
        return (
            <div>
                {msgs}
                <Table striped hover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {contacts}
                    </tbody>
                </Table>

                <Link to="s_contacts_add" className="btn btn-primary">Add Contact</Link>
            </div>
        );
    }
});

var SettingsContactsLine = React.createClass({
    handleDelete: function() {
        this.props.removeContact(this.props.contact);
    },
    render: function() {
        return (
            <tr>
                <td>{this.props.contact.name}</td>
                <td>{this.props.contact.user.email}</td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_contacts_update"
                          params={{"contactId": this.props.contact.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsContactsForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            email: "",
            messages: []
        }
    },
    componentDidMount: function() {
        if (this.props.params.contactId != undefined) {
            var id = this.props.params.contactId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            contacts.get(id, this.handleGet);
        }
    },
    handleGet: function(data, messages) {
        if (messages.length > 0) {
            this.setState({messages: messages, name: "", email: ""});
        } else {
            this.setState({
                messages: messages,
                name: data.name,
                email: data.user.email
            });
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

        if (this.state.id != "") {
            contacts.update(this.state.id, this.state.name, this.state.email,
                            this.handleFormResponse);
        } else {
            contacts.add(this.state.name, this.state.email,
                         this.handleFormResponse);
        }
    },
    handleFormResponse: function(success, messages) {
        if (success == true) {
            if (this.state.id) {
                this.setState({messages: messages});
            } else {
                this.setState({messages: messages, name: "", email: ""});
            }
        } else {
            this.setState({
                messages: messages,
                name: this.state.name,
                email: this.state.email
            });
        }
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div className="col-md-4">
                <form onSubmit={this.handleSubmit} className="text-left">
                    <h2 className="page-header">{this.state.action} Contact</h2>
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
                        {this.state.action} Contact
                    </Button>
                    <Link to="s_contacts_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});

var SettingsContactsGroupForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            messages: []
        }
    },
    componentDidMount: function() {
        if (this.props.params.groupId != undefined) {
            var id = this.props.params.groupId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            groups.get(id, this.handleGet);
        }
    },
    handleGet: function(data, messages) {
        if (messages.length > 0) {
            this.setState({messages: messages, name: ""});
        } else {
            this.setState({
                messages: messages,
                name: data.name
            });
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
    handleChange: function() {
        this.setState({
            name: this.refs.name.getValue(),
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
        if (this.validateNameState() !== "success") {
            this.setState({
                messages: [{Type: "danger", Content: "Fix errors"}]
            });
            return ;
        }

        if (this.state.id != "") {
            groups.update(this.state.id, this.state.name,
                            this.handleFormResponse);
        } else {
            groups.add(this.state.name,
                         this.handleFormResponse);
        }
    },
    handleFormResponse: function(success, messages) {
        if (success == true) {
            if (this.state.id) {
                this.setState({messages: messages});
            } else {
                this.setState({messages: messages, name: ""});
            }
        } else {
            this.setState({
                messages: messages,
                name: this.state.name
            });
        }
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div className="col-md-4">
                <form onSubmit={this.handleSubmit} className="text-left">
                    <h2 className="page-header">{this.state.action} Contact Group</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="Admins" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        {this.state.action} Contact Group
                    </Button>
                    <Link to="s_contacts_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});
