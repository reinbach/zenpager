var AuthenticationMixin = {
    statics: {
        willTransitionTo: function(transition) {
            var nextPath = transition.path;
            if (!auth.loggedIn()) {
                transition.redirect("/login", {}, {"nextPath": nextPath});
            }
        }
    }
};

var Login = React.createClass({
    contextTypes: {
        router: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            email: '',
            password: ''
        };
    },
    getDefaultProps: function() {
        return {
            messages: []
        }
    },
    validationEmailState: function() {
        if (this.state.email.length > 0) {
            if (validateEmail(this.state.email) === true) {
                return "success";
            }
            return "error"
        }
    },
    validationPasswordState: function() {
        if (this.state.password.length > 0) {
            if (passwordValid(this.state.password)) {
                return "success";
            }
            return "error";
        }
    },
    handleChange: function() {
        this.setState({
            email: this.refs.email.getValue(),
            password: this.refs.password.getValue()
        });
    },
    handleSubmit: function(event) {
        event.preventDefault();
        var { router } = this.context;
        var nextPath = router.getCurrentQuery().nextPath;
        // Prevent form being submitted till elements are in valid state
        auth.login(this.state.email, this.state.password, function(loggedIn) {
            if (nextPath) {
                router.replaceWith(nextPath);
            } else {
                router.replaceWith("/");
            }
        }.bind(this));
    },
    render: function() {
        var msgs = [];
        this.props.messages.forEach(function(msg) {
            msgs.push(<Messages type={msg.Type} message={msg.Content} />);
        });
        return (
            <div className="container-fluid">
                <div className="row">
                    <div className="col-md-3 col-md-offset-5 main">
                        {msgs}
                        <h1 className="page-header">Sign In</h1>
                        <form onSubmit={this.handleSubmit} className="text-left">
                            <Input label="Email Address" type="email" ref="email"
                                   placeholder="Enter email" value={this.state.email}
                                   autoFocus hasFeedback bsStyle={this.validationEmailState()}
                                   onChange={this.handleChange} />
                            <Input label="Password" type="password" ref="password"
                                   placeholder="Password" value={this.state.password}
                                   hasFeedback bsStyle={this.validationPasswordState()}
                                   onChange={this.handleChange} />
                            <Button type="submit" bsStyle="success">
                                Sign In
                            </Button>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
});

var Logout = React.createClass({
    componentDidMount: function() {
        auth.logout();
    },
    render: function() {
        return (
            <div className="col-md-3 col-md-offset-3">
                <h1>Signed Out</h1>
                <p>You are now signed out!</p>
            </div>
        );
    }
});

var auth = {
    login: function(email, password, cb) {
        cb = arguments[arguments.length - 1];
        if (localStorage.token) {
            if (cb) cb(true);
            this.onChange(true);
            return ;
        }
        if (email === undefined || password == undefined) {
            if (cb) cb(false);
            return ;
        }
        authenticate(email, password, function(res) {
            if (res.authenticated) {
                localStorage.token = res.token;
                localStorage.id = res.id;
                if (cb) cb(true);
                this.onChange(true, res.errors);
            } else {
                if (cb) cb(false);
                this.onChange(false, res.errors);
            }
        }.bind(this));
    },
    getToken: function() {
        return localStorage.token;
    },
    logout: function(cb) {
        logout(localStorage.token);
        delete localStorage.token;
        if (cb) cb(false);
        this.onChange(false);
    },
    loggedIn: function() {
        return !!localStorage.token;
    },
    onChange: function() {}
};

function authenticate(email, password, cb) {
    var r = new XMLHttpRequest();
    r.open("POST", "/api/v1/auth/login", true);
    r.setRequestHeader("Content-Type", "application/json");
    r.onreadystatechange = function() {
        if (r.readyState === 4) {
            data = JSON.parse(r.responseText);
            if (data.Result === "success") {
                cb({
                    authenticated: true,
                    token: r.getResponseHeader("X-Access-Token"),
                    id: data.data
                });
            } else {
                cb({
                    authenticated: false,
                    errors: data.Messages
                });
            }
        }
    };
    r.send(JSON.stringify({email: email, password: password}));
}

function logout(token) {
    var r = new XMLHttpRequest();
    r.open("GET", "/api/v1/auth/logout", true);
    r.setRequestHeader("Content-Type", "application/json");
    r.setRequestHeader("X-Access-Token", token);
    r.onreadystatechange = function() {
        if (r.readyState === 4) {
            data = JSON.parse(r.responseText);
            if (data.Result === "success") {
                console.log("successfully logged out");
            } else {
                console.log("failed to logout!");
            }
        }
    };
    r.send();
}
