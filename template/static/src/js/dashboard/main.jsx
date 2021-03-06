var Router = ReactRouter,
    Route = Router.Route,
    Link = Router.Link,
    RouteHandler = Router.RouteHandler,
    DefaultRoute = Router.DefaultRoute,
    Redirect = Router.Redirect,
    NotFoundRoute = Router.NotFoundRoute,

    // bootstrap
    Input = ReactBootstrap.Input,
    Button = ReactBootstrap.Button,
    Table = ReactBootstrap.Table;

var App = React.createClass({
    getInitialState: function() {
        return {
            loggedIn: auth.loggedIn(),
            messages: []
        };
    },
    setStateOnAuth: function(loggedIn, messages) {
        this.setState({loggedIn: loggedIn, messages: messages});
    },
    componentWillMount: function() {
        auth.onChange = this.setStateOnAuth;
        auth.login();
    },
    render: function() {
        var Navbar = ReactBootstrap.Navbar,
            Nav = ReactBootstrap.Nav,
            NavItem = ReactBootstrap.NavItem;
        if (this.state.loggedIn) {
            accessLink = <NavItem href="#/logout">Sign Out</NavItem>;
        } else {
            accessLink = <NavItem href="#/login">Sign In</NavItem>;
        }
        return (
            <div>
                <Navbar brand="ZenPager" fixedTop fluid inverse>
                    <Nav right>
                        <NavItem href="#/dashboard">Dashboard</NavItem>
                        <NavItem href="#/settings">Settings</NavItem>
                        <NavItem href="#/profile">Profile</NavItem>
                        {accessLink}
                    </Nav>
                </Navbar>
                <RouteHandler messages={this.state.messages} />
            </div>
        );
    }
});

var Messages = React.createClass({
    getInitialState: function() {
        return {
            visible: true
        };
    },
    handleDismiss: function() {
        this.setState({visible: false});
    },
    render: function() {
        if (this.state.visible) {
            var Alert = ReactBootstrap.Alert;
            return (
                <Alert bsStyle={this.props.type}
                       onDismiss={this.handleDismiss}>
                    {this.props.message}
                </Alert>
            );
        }
        return <div></div>;
    }
});

var NotFound = React.createClass({
    render: function() {
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-9 col-sm-offset-3 col-md-10
                                    col-md-offset-2 main">
                        <h1>404 Not Found</h1>
                    </div>
                </div>
            </div>
        );
    }
});

var routes = (
    <Route name="app" path="/" handler={App}>
        <NotFoundRoute handler={NotFound} />
        <Route name="login" handler={Login} />
        <Route name="logout" handler={Logout} />
        <Route name="dashboard" handler={DashboardHolder}>
            <Route name="d_overview" path="overview"
                   handler={DashboardOverview} />
            <Route name="d_servers" path="servers"
                   handler={DashboardServers} />
            <Route name="d_apps" path="apps" handler={DashboardApps} />
            <Redirect from="/dashboard" to="d_overview" />
        </Route>
        // Settings
        <Route name="settings" handler={SettingsHolder}>
            // Commands
            <Route name="commands" handler={SettingsCommandsHolder}>
                <Route name="s_commands_list" path="list"
                       handler={SettingsCommands} />
                <Route name="s_commands_add" path="add"
                       handler={SettingsCommandsForm} />
                <Route name="s_commands_update" path="update/:commandId"
                       handler={SettingsCommandsForm} />
                <Route name="s_commands_view" path=":commandId"
                       handler={SettingsCommandsView} />
                <Route name="s_commands_group_add" path="group/add"
                       handler={SettingsCommandsGroupForm} />
                <Route name="s_commands_group_update"
                       path="group/update/:groupId"
                       handler={SettingsCommandsGroupForm} />
                <Route name="s_commands_group_commands" path="group/:groupId"
                       handler={SettingsCommandsGroupCommands} />
                <Redirect from="/commands" to="s_commands_list" />
            </Route>
            // Contacts
            <Route name="contacts" handler={SettingsContactsHolder}>
                <Route name="s_contacts_list" path="list"
                       handler={SettingsContacts} />
                <Route name="s_contacts_add" path="add"
                       handler={SettingsContactsForm} />
                <Route name="s_contacts_update" path="update/:contactId"
                       handler={SettingsContactsForm} />
                <Route name="s_contacts_view" path=":contactId"
                       handler={SettingsContactsView} />
                <Route name="s_contacts_group_add" path="group/add"
                       handler={SettingsContactsGroupForm} />
                <Route name="s_contacts_group_update"
                       path="group/update/:groupId"
                       handler={SettingsContactsGroupForm} />
                <Route name="s_contacts_group_contacts" path="group/:groupId"
                       handler={SettingsContactsGroupContacts} />
                <Redirect from="/contacts" to="s_contacts_list" />
            </Route>
            // Servers
            <Route name="servers" handler={SettingsServersHolder}>
                <Route name="s_servers_list" path="list"
                       handler={SettingsServers} />
                <Route name="s_servers_add" path="add"
                       handler={SettingsServersForm} />
                <Route name="s_servers_update" path="update/:serverId"
                       handler={SettingsServersForm} />
                <Route name="s_servers_view" path=":serverId"
                       handler={SettingsServersView} />
                <Route name="s_servers_group_add" path="group/add"
                       handler={SettingsServersGroupForm} />
                <Route name="s_servers_group_update"
                       path="group/update/:groupId"
                       handler={SettingsServersGroupForm} />
                <Route name="s_servers_group_servers" path="group/:groupId"
                       handler={SettingsServersGroupServers} />
                <Redirect from="/servers" to="s_servers_list" />
            </Route>
            <Route name="s_timeperiods" path="timeperiods"
                   handler={SettingsTimePeriods} />
            <Redirect from="/settings" to="s_commands_list" />
        </Route>
        <Route name="profile" handler={ProfileHolder}>
            <Route name="p_password" path="password"
                   handler={ProfilePassword} />
            <Redirect from="/profile" to="p_password" />
        </Route>
        <Redirect from="/" to="d_overview" />
    </Route>
);

Router.run(routes, function(Handler) {
    React.render(<Handler />, document.body);
});
