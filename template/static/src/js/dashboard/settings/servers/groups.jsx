var SettingsServersGroups = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
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
        servergroups.getAll(function(data, messages) {
            this.setState({
                groups: data,
                messages: messages
            });
        }.bind(this));
    },
    removeGroup: function(group) {
        servergroups.remove(group.id, function(messages) {
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
                <SettingsServersGroupLine key={group.id} group={group}
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

                <Link to="s_servers_group_add"
                      className="btn btn-primary">Add Group</Link>
            </div>
        )
    }
});

var SettingsServersGroupLine = React.createClass({
    handleDelete: function() {
        this.props.removeGroup(this.props.group);
    },
    render: function() {
        return (
            <tr>
                <td>
                    <Link to="s_servers_group_servers"
                          params={{"groupId": this.props.group.id}}
                          >{this.props.group.name}</Link>
                </td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_servers_group_update"
                          params={{"groupId": this.props.group.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsServersGroupForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("servers");
        if (this.props.params.groupId != undefined) {
            var id = this.props.params.groupId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            servergroups.get(id, this.handleGet);
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
            servergroups.update(this.state.id, this.state.name,
                            this.handleFormResponse);
        } else {
            servergroups.add(this.state.name,
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
                    <h2 className="page-header">{this.state.action} Server Group</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="Database Servers" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        {this.state.action} Server Group
                    </Button>
                    <Link to="s_servers_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});

var SettingsServersGroupServers = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    getInitialState: function() {
        return {
            group: {},
            messages: [],
            all_servers: [],
            current_servers: []
        };
    },
    componentDidMount: function() {
        settingsSideMenu.active("servers");
        servergroups.getServers(
            this.props.params.groupId,
            function(data, messages) {
                if (this.isMounted()) {
                    this.setState({
                        group: data,
                        messages: messages
                    });
                    this.setServers(data.servers.map(function(obj) {
                        return this.renderServer(obj, "current");
                    }.bind(this)), undefined);
                }
            }.bind(this)
        );
        servers.getAll(function(data, messages) {
            if (this.isMounted()) {
                this.setState({messages: messages});
                this.setServers(undefined, data.map(function(obj) {
                    return this.renderServer(obj, "available");
                }.bind(this)));
            }
        }.bind(this));
    },
    removeServer: function(server) {
        servergroups.removeServer(
            this.props.params.groupId,
            server.id,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = removeFromListByKey(this.state.current_servers,
                                          server);
        var all = this.state.all_servers.concat(
            this.renderServer(server, "available")
        );
        this.setServers(current, all);
    },
    addServer: function(server) {
        servergroups.addServer(
            this.props.params.groupId,
            server.id,
            function(data, messages) {
                this.setState({
                    messages: messages
                });
            }.bind(this)
        );
        var current = this.state.current_servers.concat(
            this.renderServer(server, "current")
        );
        var all = removeFromListByKey(this.state.all_servers, server);
        this.setServers(current, all);
    },
    renderServer: function(server, state) {
        return (
            <SettingsServersGroupServerLine key={server.id}
                                              server={server}
                                              state={state}
                                              removeServer={this.removeServer}
                                              addServer={this.addServer} />
        )
    },
    setServers: function(current, all) {
        if (current === undefined) {
            current = this.state.current_servers;
        }
        if (all === undefined) {
            all = this.state.all_servers;
        }
        this.setState({
            all_servers: excludeByKey(all, current),
            current_servers: current
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
                            {this.state.current_servers}
                          </tbody>
                        </table>
                    </div>
                    <div className="col-md-6">
                        <h3 className="side-list-header">Available Servers</h3>
                        <table className="table table-striped table-condensed table-hover">
                          <tbody>
                            {this.state.all_servers}
                          </tbody>
                        </table>
                    </div>
                </div>
                <Link to="s_servers_list"
                      className="btn btn-default">Back to Servers</Link>
            </div>
        )
    }
});

var SettingsServersGroupServerLine = React.createClass({
    handleRemove: function() {
        this.props.removeServer(this.props.server);
    },
    handleAdd: function() {
        this.props.addServer(this.props.server);
    },
    render: function() {
        var button = [];
        if (this.props.state === "available") {
            button.push(
                <Button key={this.props.server.id} bsSize="xsmall"
                        bsStyle="success"
                        onClick={this.handleAdd}>Add</Button>
            );
        } else {
            button.push(
                <Button key={this.props.server.id} bsSize="xsmall"
                        bsStyle="danger"
                        onClick={this.handleRemove}>Remove</Button>
            );
        }
        return (
            <tr>
                <td>{this.props.server.name}</td>
                <td>{button}</td>
            </tr>
        )
    }
})
