var SettingsServersList = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    propTypes: {
        servers: React.PropTypes.array,
        messages: React.PropTypes.array
    },
    getInitialState: function() {
        return {
            servers: [],
            messages: []
        };
    },
    componentWillMount: function() {
        servers.getAll(function(data, messages) {
            this.setState({
                servers: data,
                messages: messages
            });
        }.bind(this));
    },
    removeServer: function(server) {
        servers.remove(server.id, function(messages) {
            this.setState({messages: messages})
        }.bind(this));
        this.setState({
            servers: removeFromList(this.state.servers, server)
        });
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        var servers = [];
        this.state.servers.forEach(function(server) {
            servers.push(
                <SettingsServersLine key={server.id} server={server}
                                      removeServer={this.removeServer} />
            );
        }.bind(this));
        return (
            <div>
                <h2>Servers</h2>
                {msgs}
                <Table striped hover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Url</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {servers}
                    </tbody>
                </Table>

                <Link to="s_servers_add" className="btn btn-primary">Add Server</Link>
            </div>
        );
    }
});

var SettingsServersLine = React.createClass({
    handleDelete: function() {
        this.props.removeServer(this.props.server);
    },
    render: function() {
        return (
            <tr>
                <td>
                    <Link to="s_servers_view"
                          params={{serverId: this.props.server.id}}>
                    {this.props.server.name}</Link>
                </td>
                <td>{this.props.server.url}</td>
                <td className="action-cell">
                    <Button bsSize="xsmall" bsStyle="danger"
                            onClick={this.handleDelete}>
                        Delete
                    </Button>
                    <Link to="s_servers_update"
                          params={{"serverId": this.props.server.id}}
                          className="btn btn-xs btn-default">Edit</Link>
                </td>
            </tr>
        );
    }
});

var SettingsServersForm = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    getInitialState: function() {
        return {
            id: "",
            action: "Add",
            name: "",
            url: "",
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("servers");
        if (this.props.params.serverId != undefined) {
            var id = this.props.params.serverId;
            this.setState({
                "id": id,
                "action": "Update"
            });
            servers.get(id, this.handleGet);
        }
    },
    handleGet: function(data, messages) {
        if (messages.length > 0) {
            this.setState({messages: messages, name: "", url: ""});
        } else {
            console.log(data);
            this.setState({
                messages: messages,
                name: data.name,
                url: data.url.Host
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
    validateUrlState: function() {
        if (this.state.url.length > 0) {
            if (validateUrl(this.state.url) === true) {
                return "success";
            }
            return "error"
        }
    },
    handleChange: function() {
        this.setState({
            name: this.refs.name.getValue(),
            url: this.refs.url.getValue()
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
        if (this.state.url.length < 1) {
            this.setState({
                messages: [{Type: "danger", Content: "Url is required."}]
            });
            return ;
        }
        if (this.validateNameState() !== "success" || this.validateUrlState() !== "success") {
            this.setState({
                messages: [{Type: "danger", Content: "Fix errors"}]
            });
            return ;
        }

        if (this.state.id != "") {
            servers.update(this.state.id, this.state.name, this.state.url,
                            this.handleFormResponse);
        } else {
            servers.add(this.state.name, this.state.url,
                         this.handleFormResponse);
        }
    },
    handleFormResponse: function(success, messages) {
        if (success == true) {
            if (this.state.id) {
                this.setState({messages: messages});
            } else {
                this.setState({messages: messages, name: "", url: ""});
            }
        } else {
            this.setState({
                messages: messages,
                name: this.state.name,
                url: this.state.url
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
                    <h2 className="page-header">{this.state.action} Server</h2>
                    {msgs}
                    <Input label="Name" type="text" ref="name"
                           placeholder="LAMP Server" value={this.state.name}
                           autoFocus hasFeedback bsStyle={this.validateNameState()}
                           onChange={this.handleChange} />
                    <Input label="Url" type="text" ref="url"
                           placeholder="127.0.0.1" value={this.state.url}
                           hasFeedback bsStyle={this.validateUrlState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        {this.state.action} Server
                    </Button>
                    <Link to="s_servers_list" className="btn btn-default">Cancel</Link>
                </form>
            </div>
        );
    }
});

var SettingsServersView = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    getInitialState: function() {
        return {
            server: {},
            all_groups: [],
            current_groups: [],
            messages: []
        }
    },
    componentDidMount: function() {
        settingsSideMenu.active("servers");
        servers.getGroups(
            this.props.params.serverId,
            function(data, messages) {
                if (this.isMounted()) {
                    this.setState({
                        server: data,
                        messages: messages
                    });
                    this.setGroups(data.groups.map(function(obj) {
                        return this.renderGroup(obj, "current");
                    }.bind(this)), undefined);
                }
            }.bind(this)
        );
        servergroups.getAll(function(data, messages) {
            if (this.isMounted()) {
                this.setState({messages: messages});
                this.setGroups(undefined, data.map(function(obj) {
                    return this.renderGroup(obj, "available");
                }.bind(this)));
            }
        }.bind(this));
    },
    removeGroup: function(group) {
        servergroups.removeServer(
            group.id,
            this.props.params.serverId,
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
        servergroups.addServer(
            group.id,
            this.props.params.serverId,
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
            <SettingsServersGroupElement key={group.id}
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
                        <h2>{this.state.server.name}</h2>
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
                <Link to="s_servers_list"
                      className="btn btn-default">Back to List of Servers</Link>
            </div>
        )
    }
});

var SettingsServersGroupElement = React.createClass({
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
