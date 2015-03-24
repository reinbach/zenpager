var ProfileHolder = React.createClass({
    mixins: [Authentication],
    getDefaultProps: function() {
        return {
            messages: []
        };
    },
    render: function() {
        var msgs = [];
        this.props.messages.forEach(function(msg) {
            msgs.push(<Messages type="danger" message={msg} />);
        });
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
                        {msgs}
                        <RouteHandler />
                    </div>
                </div>
            </div>
        );
    }
});

var ProfilePassword = React.createClass({
    getInitialState: function() {
        return {
            password: '',
            password_confirm: ''
        };
    },
    validationPasswordState: function() {
        if (this.state.password.length > 0) {
            if (passwordValid(this.state.password)) {
                return "success";
            }
            return "error";
        }
    },
    validationPasswordConfirmState: function() {
        if (this.state.password_confirm.length > 0) {
            if (this.state.password_confirm === this.state.password) {
                return "success";
            }
            return "error";
        }
    },
    handleChange: function() {
        this.setState({
            password: this.refs.password.getValue(),
            password_confirm: this.refs.password_confirm.getValue()
        });
    },
    handleSubmit: function(event) {
        event.preventDefault();
        // make call to server to update password
        console.log("call server")
    },
    render: function() {
        return (
            <div className="col-md-3">
                <form onSubmit={this.handleSubmit} className="text-left">
                    <h1 className="page-header">Reset Password</h1>
                    <Input label="New Password" type="password" ref="password"
                           placeholder="New password" value={this.state.password}
                           autoFocus hasFeedback bsStyle={this.validationPasswordState()}
                           onChange={this.handleChange} />
                    <Input label="Confirm Password" type="password" ref="password_confirm"
                           placeholder="Confirm Password" value={this.state.password_confirm}
                           hasFeedback bsStyle={this.validationPasswordConfirmState()}
                           onChange={this.handleChange} />
                    <Button type="submit" bsStyle="success">
                        Update Password
                    </Button>
                </form>
            </div>
        );
    }
});
