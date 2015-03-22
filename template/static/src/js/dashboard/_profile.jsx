var ProfileHolder = React.createClass({
    mixins: [Authentication],
    render: function() {
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-sm-3 col-md-2 sidebar">
                        <h1>Profile</h1>
                        <ul className="nav nav-sidebar">
                            <li><Link to="p_password">Password</Link></li>
                        </ul>
                    </div>
                    <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                        {/*<Messages type="success" message="Hello World" />*/}
                        <RouteHandler />
                    </div>
                </div>
            </div>
        );
    }
});

var ProfilePassword = React.createClass({
    render: function() {
        return (
            <div>
                Reset Password...
            </div>
        );
    }
});
