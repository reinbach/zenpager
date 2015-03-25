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
            password_confirm: '',
            messages: []
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
        if (this.state.password_confirm.length > 0 || this.state.password.length > 0) {
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
        if (this.state.password.length < 1) {
            this.setState({"messages": [{Type: "danger", Content: "Password is required."}]});
            return ;
        }
        if (this.validationPasswordState() !== "success" || this.validationPasswordConfirmState() !== "success") {
            this.setState({"messages": [{Type: "danger", Content: "Fix errors"}]});
            return ;
        }
        user.updatePassword(this.state.password, function(success, messages) {
            if (success === true) {
                this.setState({messages: messages, password: "", password_confirm: ""});
            } else {
                this.setState({
                    messages: messages,
                    password: this.state.password,
                    password_confirm: this.state.password_confirm
                });
            }
        }.bind(this));
    },
    render: function() {
        var msgs = [];
        this.state.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div className="col-md-3">
                <form onSubmit={this.handleSubmit} className="text-left">
                    <h1 className="page-header">Reset Password</h1>
                    {msgs}
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

var user = {
    updatePassword: function(password, cb) {
        if (!localStorage.id) {
            if (cb) cb(false);
            return ;
        }
        if (password.length < 1) {
            if (cb) cb(false, ["Password is required."]);
            return ;
        }
        updatePassword(localStorage.id, password, function(res) {
            if (res.update) {
                if (cb) cb(true, [{Type: "success", Content: "Successfully updated password."}]);
            } else {
                if (cb) cb(false, res.errors);
            }
        });
    }
};

function updatePassword(userId, password, cb) {
    var r = new XMLHttpRequest();
    r.open("PATCH", "/api/v1/auth/user/" + userId, true);
    r.setRequestHeader("Content-Type", "application/json");
    r.onreadystatechange = function() {
        if (r.readyState === 4) {
            data = JSON.parse(r.responseText);
            if (data.Result === "success") {
                cb({update: true});
            } else {
                cb({update: false, errors: data.Messages});
            }
        }
    };
    r.send(JSON.stringify({password: password}));
}
