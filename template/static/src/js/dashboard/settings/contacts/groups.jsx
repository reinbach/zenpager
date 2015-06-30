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
        contactgroups.getAll(function(data, messages) {
            this.setState({
                groups: data,
                messages: messages
            });
        }.bind(this));
    },
    removeGroup: function(group) {
        contactgroups.remove(group.id, function(messages) {
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
                <SettingsContactsGroupLine key={group.id} group={group}
                                           removeGroup={this.removeGroup} />
            );
        }.bind(this));
        return (
            <div>
                <h2>Groups</h2>
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
                <td>
                    <Link to="s_contacts_group_contacts"
                          params={{"groupId": this.props.group.id}}
                          >{this.props.group.name}</Link>
                </td>
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
        settingsSideMenu.active("contacts");
        if (this.props.params.groupId != undefined) {
            var id = this.props.params.groupId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            contactgroups.get(id, this.handleGet);
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
            contactgroups.update(this.state.id, this.state.name,
                            this.handleFormResponse);
        } else {
            contactgroups.add(this.state.name,
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

var SettingsContactsGroupContacts = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    getInitialState: function() {
        return {
            group: {},
            messages: [],
            contacts: []
        };
    },
    componentDidMount: function() {
        settingsSideMenu.active("contacts");
        contactgroups.getContacts(
            this.props.params.groupId,
            function(data, messages) {
                if (this.isMounted()) {
                    this.setState({
                        group: data,
                        messages: messages
                    });
                }
            }.bind(this)
        );
        contacts.getAll(function(data, messages) {
            if (this.isMounted()) {
                this.setState({
                    contacts: data,
                    messages: messages
                });
            }
        }.bind(this));
    },
    removeContact: function(contact) {
        console.log("remove: " + contact);
    },
    addContact: function(contact) {
        contactgroups.addContact(
            this.props.params.groupId,
            contact.id,
            function(data, messages) {
                console.log("mmmm");
                this.setState({
                    contacts: removeFromList(this.state.contacts, contact),
                    messages: messages
                });
            }.bind(this)
        );
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        var current_contacts = [];
        var used_contact_ids = [];
        if (this.state.group.contacts !== undefined) {
            this.state.group.contacts.forEach(function(contact) {
                used_contact_ids.push(contact.id);
                current_contacts.push(
                    <SettingsContactsGroupContactLine key={contact.id}
                                                      contact={contact}
                                                      state="current"
                                                      removeContact={this.removeContact}
                                                      addContact={this.addContact} />
                );
            }.bind(this));
        }
        var available_contacts = [];
        this.state.contacts.forEach(function(contact) {
            if (used_contact_ids.indexOf(contact.id) === -1) {
                available_contacts.push(
                    <SettingsContactsGroupContactLine key={contact.id}
                                                      contact={contact}
                                                      state="available"
                                                      removeContact={this.removeContact}
                                                      addContact={this.addContact} />
                );
            }
        }.bind(this));
        return(
            <div>
                {msgs}
                <div className="row">
                    <div className="col-md-6">
                        <h2>Group: {this.state.group.name}</h2>
                        {current_contacts}
                    </div>
                    <div className="col-md-6">
                        <h3>Available Contacts</h3>
                        {available_contacts}
                    </div>
                </div>
                <Link to="s_contacts_list"
                      className="btn btn-default">Back to Contacts</Link>
            </div>
        )
    }
});

var SettingsContactsGroupContactLine = React.createClass({
    handleRemove: function() {
        this.props.removeContact(this.props.contact);
    },
    handleAdd: function() {
        this.props.addContact(this.props.contact);
    },
    render: function() {
        var button = [];
        if (this.props.state === "available") {
            button.push(
                <Button key={this.props.contact.id} bsSize="xsmall"
                        bsStyle="success"
                        onClick={this.handleAdd}>Add</Button>
            );
        } else {
            button.push(
                <Button key={this.props.contact.id} bsSize="xsmall"
                        bsStyle="danger"
                        onClick={this.handleRemove}>Remove</Button>
            );
        }
        return (
            <div>
                {this.props.contact.name}
                {button}
            </div>
        )
    }
})