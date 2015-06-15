var SettingsHolder = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-3 col-md-2 sidebar">
                        <h1>Settings</h1>
                        <ul className="nav nav-sidebar">
                            <li><Link to="s_commands">Commands</Link></li>
                            <li><Link to="s_contacts" className="contacts-link">Contacts</Link></li>
                            <li><Link to="s_servers">Servers</Link></li>
                            <li>
                                <Link to="s_timeperiods">Time Periods</Link>
                            </li>
                        </ul>
                    </div>
                    <div className="col-sm-9 col-sm-offset-3 col-md-10
                                    col-md-offset-2 main">
                        {/*<Messages type="success" message="Hello World" />*/}
                        <RouteHandler />
                    </div>
                </div>
            </div>
        );
    }
});

var SettingsCommands = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div>
                Settings Commands...
            </div>
        );
    }
});

var SettingsServers = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div>
                Settings Servers...
            </div>
        );
    }
});

var SettingsTimePeriods = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div>
                Settings Time Periods...
            </div>
        );
    }
});
