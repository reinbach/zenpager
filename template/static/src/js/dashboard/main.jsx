var Router = ReactRouter,
    Route = Router.Route,
    Link = Router.Link,
    RouteHandler = Router.RouteHandler,
    DefaultRoute = Router.DefaultRoute,
    Redirect = Router.Redirect,
    NotFoundRoute = Router.NotFoundRoute;

var App = React.createClass({
    getInitialState: function() {
        return {
            loggedIn: auth.loggedIn()
        };
    },
    setStateOnAuth: function(loggedIn) {
        this.setState({loggedIn: loggedIn});
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
                <RouteHandler />
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
                <Alert bsStyle={this.props.type} onDismiss={this.handleDismiss}
                       dismissAfter={2000}>{this.props.message}</Alert>
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
                    <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
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
        <Route name="settings" handler={SettingsHolder}>
            <Route name="s_commands" path="commands"
                   handler={SettingsCommands} />
            <Route name="s_contacts" path="contacts"
                   handler={SettingsContacts} />
            <Route name="s_servers" path="servers" handler={SettingsServers} />
            <Route name="s_timeperiods" path="timeperiods"
                   handler={SettingsTimePeriods} />
            <Redirect from="/settings" to="s_commands" />
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
