var Authentication = {
    statics: {
        willTransitionTo: function(transition) {
            var nextPath = transition.path;
            if (!auth.loggedIn()) {
                transition.redirect("/login", {}, {"nextPath": nextPath})
            }
        }
    }
}

var Login = React.createClass({
    contextTypes: {
        router: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            error: false
        };
    },
    handleSubmit: function(event) {
        event.preventDefault();
        var { router } = this.context;
        var nextPath = router.getCurrentQuery().nextPath;
        var email = this.refs.email.getDOMNode().value;
        var password = this.refs.password.getDOMNode().value;
        auth.login(email, password, function(loggedIn) {
            if (!loggedIn) {
                return this.setState({error: true});
            }
            if (nextPath) {
                router.replaceWith(nextPath);
            } else {
                router.replaceWith("/");
            }
        }.bind(this));
    },
    render: function() {
        var errors = this.state.error ? <p>Bad Login Information</p> : '';
        return (
            <div className="col-md-3 col-md-offset-3">
                <h1>Sign In</h1>
                <div className="alert alert-danger">{errors}</div>
                <form onSubmit={this.handleSubmit} className="text-left">
                    <div className="form-group">
                        <label>Email Address</label>
                        <input type="email" className="form-control"
                               ref="email" placeholder="Enter email"
                               value={this.props.email} autofocus />
                        <p className="help-block">{this.props.email_errors}</p>
                    </div>
                    <div className="form-group">
                        <label>Password</label>
                        <input type="password" className="form-control"
                               ref="password" placeholder="Password" />
                        <p className="help-block">{this.props.password_errors}</p>
                    </div>
                    <button type="submit" className="btn btn-default">Sign In</button>
                </form>
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
    login: function(email, pass, cb) {
        cb = arguments[arguments.length - 1];
        if (localStorage.token) {
            if (cb) cb(true);
            this.onChange(true);
            return ;
        }
        authenticate(email, pass, function(res) {
            if (res.authenticated) {
                localStorage.token = res.token;
                if (cb) cb(true);
                this.onChange(true);
            } else {
                if (cb) cb(false);
                this.onChange(false);
            }
        }.bind(this));
    },
    getToken: function() {
        return localStorage.token;
    },
    logout: function(cb) {
        delete localStorage.token;
        if (cb) cb();
        this.onChange(false);
    },
    loggedIn: function() {
        return !!localStorage.token;
    },
    onChange: function() {}
};

function authenticate(email, pass, cb) {
    setTimeout(function() {
        // make call to server and attempt to log user in
        // result from server at this point
        if (true) {
            cb({
                authenticated: true,
                token: Math.random().toString(36).substring(7)
            });
        } else {
            cb({authenticated: false});
        }
    }, 0);
}
