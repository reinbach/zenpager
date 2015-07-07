var SettingsServersMixin = {
    componentDidMount: function() {
        settingsSideMenu.active("servers");
    }
};

var SettingsServersHolder = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    render: function() {
        return (
            <div>
                <h1 className="page-header">Servers</h1>
                <RouteHandler />
            </div>
        );
    }
});

var SettingsServers = React.createClass({
    mixins: [AuthenticationMixin, SettingsServersMixin],
    render: function() {
        return (
            <div className="row">
                <div className="col-md-6">
                    <SettingsServersGroups />
                </div>
                <div className="col-md-6">
                    <SettingsServersList />
                </div>
            </div>
        )
    }
});
