var DashboardHolder = React.createClass({
    mixins: [Authentication],
    getInitialState: function() {
        return {
            messages: []
        };
    },
    render: function() {
        msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type="error" message={msg} />);
        });
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-3 col-md-2 sidebar">
                        <ul className="nav nav-sidebar">
                            <li><Link to="d_overview">Overview</Link></li>
                            <li><Link to="d_servers">Servers</Link></li>
                            <li><Link to="d_apps">Applications</Link></li>
                        </ul>
                    </div>
                    <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                        {msgs}
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
