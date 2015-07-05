var settingsSideMenu = {
    active: function(link) {
        $(".nav-sidebar li a").removeClass("active");
        var elem = $("." + link + "-link");
        elem.addClass("active");
    }
}

var SettingsHolder = React.createClass({
    mixins: [AuthenticationMixin],
    render: function() {
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-3 col-md-2 sidebar">
                        <h1>Settings</h1>
                        <ul className="nav nav-sidebar">
                            <li>
                                <Link to="s_commands"
                                      className="commands-link">Commands</Link>
                            </li>
                            <li>
                                <Link to="s_contacts_list"
                                      className="contacts-link">Contacts</Link>
                            </li>
                            <li>
                                <Link to="s_servers"
                                      className="servers-link">Servers</Link>
                            </li>
                            <li>
                                <Link to="s_timeperiods"
                                      className="timeperiods-link">Time Periods</Link>
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
    mixins: [AuthenticationMixin],
    componentDidMount: function() {
        settingsSideMenu.active("commands");
    },
    render: function() {
        return (
            <div>
                Settings Commands...
            </div>
        );
    }
});

var SettingsServers = React.createClass({
    mixins: [AuthenticationMixin],
    componentDidMount: function() {
        settingsSideMenu.active("servers");
    },
    render: function() {
        return (
            <div>
                Settings Servers...
            </div>
        );
    }
});

var SettingsTimePeriods = React.createClass({
    mixins: [AuthenticationMixin],
    componentDidMount: function() {
        settingsSideMenu.active("timeperiods");
    },
    render: function() {
        return (
            <div>
                Settings Time Periods...
            </div>
        );
    }
});
