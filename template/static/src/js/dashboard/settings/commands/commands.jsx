var SettingsCommandsList = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    propTypes: {
        commands: React.PropTypes.array,
        messages: React.PropTypes.array
    },
    getInitialState: function() {
        return {
            commands: [],
            messages: []
        };
    },
    componentWillMount: function() {
        commands.getAll(function(data, messages) {
            this.setState({
                commands: data,
                messages: messages
            });
        }.bind(this));
    },
    removeCommand: function(command) {
        commands.remove(command.id, function(messages) {
            this.setState({messages: messages})
        }.bind(this));
        this.setState({
            commands: removeFromList(this.state.commands, command)
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        var commands = [];
        this.state.commands.forEach(function(command) {
            commands.push(
                <SettingsCommandsLine key={command.id} command={command}
                                      removeCommand={this.removeCommand} />
            );
        }.bind(this));
        return (
            <div>
                <h2>Commands</h2>
                {msgs}
                <Table striped hover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Command</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {commands}
                    </tbody>
                </Table>

                <Link to="s_commands_add" className="btn btn-primary">Add Command</Link>
            </div>
        );
    }
});

var SettingsCommandsLine = React.createClass({
    handleDelete: function() {
        this.props.removeCommand(this.props.command);
    },
    render: function() {
        return (
            <tr>
                <td>
                    <Link to="s_commands_view"
                          params={{commandId: this.props.command.id}}>
                    {this.props.command.name}</Link>
                </td>
                <td>{this.props.command.command}</td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_commands_update"
                          params={{"commandId": this.props.command.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsCommandsForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            command: "",
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("commands");
        if (this.props.params.commandId != undefined) {
            var id = this.props.params.commandId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            commands.get(id, this.handleGet);
        }
    },
    handleGet: function(data, messages) {
        if (messages.length > 0) {
            this.setState({messages: messages, name: "", command: ""});
        } else {
            console.log(data);
            this.setState({
                messages: messages,
                name: data.name,
                command: data.command
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
    validateCommandState: function() {
        if (this.state.command.length > 0) {
            if (validateCommand(this.state.command) === true) {
                return "success";
            }
            return "error"
        }
    },
    handleChange: function() {
        this.setState({
            name: this.refs.name.getValue(),
            command: this.refs.command.getValue()
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
        if (this.state.command.length < 1) {
            this.setState({
                messages: [{Type: "danger", Content: "Command is required."}]
            });
            return ;
        }
        if (this.validateNameState() !== "success" || this.validateCommandState() !== "success") {
            this.setState({
                messages: [{Type: "danger", Content: "Fix errors"}]
            });
            return ;
        }

        if (this.state.id != "") {
            commands.update(this.state.id, this.state.name, this.state.command,
                            this.handleFormResponse);
        } else {
            commands.add(this.state.name, this.state.command,
                         this.handleFormResponse);
        }
    },
    handleFormResponse: function(success, messages) {
        if (success == true) {
            if (this.state.id) {
                this.setState({messages: messages});
            } else {
                this.setState({messages: messages, name: "", command: ""});
            }
        } else {
            this.setState({
                messages: messages,
                name: this.state.name,
                command: this.state.command
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
                    <h2 className="page-header">{this.state.action} Command</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="Ping" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Input label="Command" type="text" ref="command"
                           placeholder="/path/command $HOSTADDRESS$ $ARG1$"
                           value={this.state.command} hasFeedback
                           bsStyle={this.validateCommandState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        {this.state.action} Command
                    </Button>
                    <Link to="s_commands_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});

var SettingsCommandsView = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    getInitialState: function() {
        return {
            command: {},
            all_groups: [],
            current_groups: [],
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("commands");
        commands.getGroups(
            this.props.params.commandId,
            function(data, messages) {
                if (this.isMounted()) {
                    this.setState({
                        command: data,
                        messages: messages
                    });
                    this.setGroups(data.groups.map(function(obj) {
                        return this.renderGroup(obj, "current");
                    }.bind(this)), undefined);
                }
            }.bind(this)
        );
        commandgroups.getAll(function(data, messages) {
            if (this.isMounted()) {
                this.setState({messages: messages});
                this.setGroups(undefined, data.map(function(obj) {
                    return this.renderGroup(obj, "available");
                }.bind(this)));
            }
        }.bind(this));
    },
    removeGroup: function(group) {
        commandgroups.removeCommand(
            group.id,
            this.props.params.commandId,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = removeFromListByKey(this.state.current_groups,
                                          group);
        var all = this.state.all_groups.concat(
            this.renderGroup(group, "available")
        );
        this.setGroups(current, all);
    },
    addGroup: function(group) {
        commandgroups.addCommand(
            group.id,
            this.props.params.commandId,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = this.state.current_groups.concat(
            this.renderGroup(group, "current")
        );
        var all = removeFromListByKey(this.state.all_groups, group);
        this.setGroups(current, all);
    },
    renderGroup: function(group, state) {
        return (
            <SettingsCommandsGroupElement key={group.id}
                                          group={group}
                                          state={state}
                                          removeGroup={this.removeGroup}
                                          addGroup={this.addGroup} />
        )
    },
    setGroups: function(current, all) {
        if (current === undefined) {
            current = this.state.current_groups;
        }
        if (all === undefined) {
            all = this.state.all_groups;
        }
        this.setState({
            all_groups: excludeByKey(all, current),
            current_groups: current
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div>
                {msgs}
                <div className="row">
                    <div className="col-md-6">
                        <h2>{this.state.command.name} <small>({this.state.command.command})</small></h2>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {this.state.current_groups}
                          </tbody>
                        </table>
                    </div>
                    <div className="col-md-6">
                        <h3 className="side-list-header">Available Groups</h3>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {this.state.all_groups}
                          </tbody>
                        </table>
                    </div>
                </div>
                <Link to="s_commands_list"
                      className="btn btn-default">Back to List of Commands</Link>
            </div>
        )
    }
});

var SettingsCommandsGroupElement = React.createClass({
    handleRemove: function() {
        this.props.removeGroup(this.props.group);
    },
    handleAdd: function() {
        this.props.addGroup(this.props.group);
    },
    render: function() {
        var button = [];
        if (this.props.state === "available") {
            button.push(
                <Button key={this.props.group.id} bsSize="xsmall"
                        bsStyle="success"
                        onClick={this.handleAdd}>Add</Button>
            );
        } else {
            button.push(
                <Button key={this.props.group.id} bsSize="xsmall"
                        bsStyle="danger"
                        onClick={this.handleRemove}>Remove</Button>
            );
        }
        return (
            <tr>
                <td>{this.props.group.name}</td>
                <td>{button}</td>
            </tr>
        )
    }
})
