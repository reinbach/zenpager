var Router = ReactRouter,
    Route = Router.Route,
    Link = Router.Link,
    RouteHandler = Router.RouteHandler,
    DefaultRoute = Router.DefaultRoute,
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
        return (
            <div>
                <Navbar brand="ZenPager" fixedTop fluid inverse>
                    <Nav right>
                        <NavItem href="#">Dashboard</NavItem>
                        <NavItem href="#/settings">Settings</NavItem>
                        <NavItem href="#/profile">Profile</NavItem>
                        <NavItem href="#/logout">Sign Out</NavItem>
                    </Nav>
                </Navbar>
                <RouteHandler/>
            </div>
        );
    }
});

var Messages = React.createClass({
    render: function() {
        return (
            <div className="alert alert-info">{this.props.message}</div>
        );
    }
});

var DashboardHolder = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-3 col-md-2 sidebar">
                        <ul className="nav nav-sidebar">
                            <li><Link to="overview">Overview</Link></li>
                            <li><Link to="servers">Servers</Link></li>
                            <li><Link to="apps">Applications</Link></li>
                        </ul>
                    </div>
                    <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                        <Messages />
                        <RouteHandler />
                    </div>
                </div>
            </div>
        );
    }
});

var DashboardOverview = React.createClass({
    render: function() {
        return (
            <div>
                Dashboard Overview...
            </div>
        );
    }
});

var DashboardServers = React.createClass({
    render: function() {
        return (
            <div>
                Dashboard Servers...
            </div>
        );
    }
});

var DashboardApps = React.createClass({
    render: function() {
        return (
            <div>
                Dashboard Apps...
            </div>
        );
    }
});

var NotFound = React.createClass({
    render: function() {
        return (
            <div className="inner cover">
                <h1 className="cover-heading">404 Not Found</h1>
            </div>
        );
    }
});

var routes = (
    <Route name="app" path="/" handler={App}>
        <Route name="login" handler={Login} />
        <Route name="logout" handler={Logout} />
        <Route name="dashboard" handler={DashboardHolder}>
            <Route name="overview" handler={DashboardOverview} />
            <Route name="servers" handler={DashboardServers} />
            <Route name="apps" handler={DashboardApps} />
            <DefaultRoute handler={DashboardOverview} />
        </Route>
        <NotFoundRoute handler={NotFound} />
        <DefaultRoute handler={DashboardHolder} />
    </Route>
);

Router.run(routes, function(Handler) {
    React.render(<Handler />, document.body);
});
