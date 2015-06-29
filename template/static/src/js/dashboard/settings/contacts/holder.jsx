var SettingsContactsMixin = {
    componentDidMount: function() {
        settingsSideMenu.active("contacts");
    }
};

var SettingsContactsHolder = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    render: function() {
        return (
            <div>
                <h1 className="page-header">Contacts</h1>
                <RouteHandler />
            </div>
        );
    }
});

var SettingsContacts = React.createClass({
    mixins: [AuthenticationMixin, SettingsContactsMixin],
    render: function() {
        return (
            <div className="row">
                <div className="col-md-6">
                    <SettingsContactsGroups />
                </div>
                <div className="col-md-6">
                    <SettingsContactsList />
                </div>
            </div>
        )
    }
});
