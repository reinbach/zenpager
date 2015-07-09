var SettingsCommandsMixin = {
    componentDidMount: function() {
        settingsSideMenu.active("commands");
    }
};

var SettingsCommandsHolder = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    render: function() {
        return (
            <div>
                <h1 className="page-header">Commands</h1>
                <RouteHandler />
            </div>
        );
    }
});

var SettingsCommands = React.createClass({
    mixins: [AuthenticationMixin, SettingsCommandsMixin],
    render: function() {
        return (
            <div className="row">
                <div className="col-md-6">
                    <SettingsCommandsGroups />
                </div>
                <div className="col-md-6">
                    <SettingsCommandsList />
                </div>
            </div>
        )
    }
});
