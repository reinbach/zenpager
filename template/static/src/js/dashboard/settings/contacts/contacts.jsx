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
            contacts.push(
                <SettingsContactsLine key={contact.id} contact={contact}
                                      removeContact={this.removeContact} />
            );
        }.bind(this));
        return (
            <div>
                <h2>Individuals</h2>
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
                <td>
                    <Link to="s_contacts_view"
                          params={{contactId: this.props.contact.id}}>
                    {this.props.contact.name}</Link>
                </td>
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
        settingsSideMenu.active("contacts");
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

var SettingsContactsView = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    getInitialState: function() {
        return {
            contact: {},
            groups: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("contacts");
        contacts.get(this.props.params.contactId, function(data, messages) {
            this.setState({
                contact: data,
                messages: messages
            })
        }.bind(this));
    },
    render: function() {
        var msg = [];
        var current_groups = [];
        var available_groups = [];
        return (
            <div>
                {msg}
                <div className="row">
                    <div className="col-md-6">
                        <h2>{this.state.contact.name}</h2>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {current_groups}
                          </tbody>
                        </table>
                    </div>
                    <div className="col-md-6">
                        <h2>Available Groups</h2>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {available_groups}
                          </tbody>
                        </table>
                    </div>
                </div>
                <Link to="s_contacts_list"
                      className="btn btn-default">Back to List of Contacts</Link>
            </div>
        )
    }
});
