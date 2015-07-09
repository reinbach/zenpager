var SettingsCommandsGroups = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
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
        commandgroups.getAll(function(data, messages) {
            this.setState({
                groups: data,
                messages: messages
            });
        }.bind(this));
    },
    removeGroup: function(group) {
        commandgroups.remove(group.id, function(messages) {
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
                <SettingsCommandsGroupLine key={group.id} group={group}
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

                <Link to="s_commands_group_add"
                      className="btn btn-primary">Add Group</Link>
            </div>
        )
    }
});

var SettingsCommandsGroupLine = React.createClass({
    handleDelete: function() {
        this.props.removeGroup(this.props.group);
    },
    render: function() {
        return (
            <tr>
                <td>
                    <Link to="s_commands_group_commands"
                          params={{"groupId": this.props.group.id}}
                          >{this.props.group.name}</Link>
                </td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_commands_group_update"
                          params={{"groupId": this.props.group.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsCommandsGroupForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("commands");
        if (this.props.params.groupId != undefined) {
            var id = this.props.params.groupId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            commandgroups.get(id, this.handleGet);
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
            commandgroups.update(this.state.id, this.state.name,
                            this.handleFormResponse);
        } else {
            commandgroups.add(this.state.name,
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
                    <h2 className="page-header">{this.state.action} Command Group</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="Database Commands" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        {this.state.action} Command Group
                    </Button>
                    <Link to="s_commands_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});

var SettingsCommandsGroupCommands = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    getInitialState: function() {
        return {
            group: {},
            messages: [],
            all_commands: [],
            current_commands: []
        };
    },
    componentDidMount: function() {
        settingsSideMenu.active("commands");
        commandgroups.getCommands(
            this.props.params.groupId,
            function(data, messages) {
                if (this.isMounted()) {
                    this.setState({
                        group: data,
                        messages: messages
                    });
                    this.setCommands(data.commands.map(function(obj) {
                        return this.renderCommand(obj, "current");
                    }.bind(this)), undefined);
                }
            }.bind(this)
        );
        commands.getAll(function(data, messages) {
            if (this.isMounted()) {
                this.setState({messages: messages});
                this.setCommands(undefined, data.map(function(obj) {
                    return this.renderCommand(obj, "available");
                }.bind(this)));
            }
        }.bind(this));
    },
    removeCommand: function(command) {
        commandgroups.removeCommand(
            this.props.params.groupId,
            command.id,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = removeFromListByKey(this.state.current_commands,
                                          command);
        var all = this.state.all_commands.concat(
            this.renderCommand(command, "available")
        );
        this.setCommands(current, all);
    },
    addCommand: function(command) {
        commandgroups.addCommand(
            this.props.params.groupId,
            command.id,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = this.state.current_commands.concat(
            this.renderCommand(command, "current")
        );
        var all = removeFromListByKey(this.state.all_commands, command);
        this.setCommands(current, all);
    },
    renderCommand: function(command, state) {
        return (
            <SettingsCommandsGroupCommandLine key={command.id}
                                              command={command}
                                              state={state}
                                              removeCommand={this.removeCommand}
                                              addCommand={this.addCommand} />
        )
    },
    setCommands: function(current, all) {
        if (current === undefined) {
            current = this.state.current_commands;
        }
        if (all === undefined) {
            all = this.state.all_commands;
        }
        this.setState({
            all_commands: excludeByKey(all, current),
            current_commands: current
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return(
            <div>
                {msgs}
                <div className="row">
                    <div className="col-md-6">
                        <h2>{this.state.group.name}</h2>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {this.state.current_commands}
                          </tbody>
                        </table>
                    </div>
                    <div className="col-md-6">
                        <h3 className="side-list-header">Available Commands</h3>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {this.state.all_commands}
                          </tbody>
                        </table>
                    </div>
                </div>
                <Link to="s_commands_list"
                      className="btn btn-default">Back to Commands</Link>
            </div>
        )
    }
});

var SettingsCommandsGroupCommandLine = React.createClass({
    handleRemove: function() {
        this.props.removeCommand(this.props.command);
    },
    handleAdd: function() {
        this.props.addCommand(this.props.command);
    },
    render: function() {
        var button = [];
        if (this.props.state === "available") {
            button.push(
                <Button key={this.props.command.id} bsSize="xsmall"
                        bsStyle="success"
                        onClick={this.handleAdd}>Add</Button>
            );
        } else {
            button.push(
                <Button key={this.props.command.id} bsSize="xsmall"
                        bsStyle="danger"
                        onClick={this.handleRemove}>Remove</Button>
            );
        }
        return (
            <tr>
                <td>{this.props.command.name}</td>
                <td>{button}</td>
            </tr>
        )
    }
})
