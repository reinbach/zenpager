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
            error: false,
            is_valid: false,
            errors: {
                email: [],
                password: []
            }
        };
    },
    handleValidate: function(email, password) {
        this.state.errors.email = [];
        this.state.errors.password = [];
        this.state.is_valid = true;
        if (email.length == "") {
            this.state.is_valid = false;
            this.state.errors.email.push(
                <FormElementError error="Email Address is required" />
            );
        } else if (validateEmail(email) !== true) {
            this.state.is_valid = false;
            this.state.errors.email.push(
                <FormElementError error="Invalid Email Address" />
            );
        }
        if (password.length == "") {
            this.state.is_valid = false;
            this.state.errors.password.push(
                <FormElementError error="Password is required" />
            );
        }
        this.setState({is_valid: this.state.is_valid,
                       error: this.state.error, errors: this.state.errors});
    },
    handleSubmit: function(event) {
        event.preventDefault();
        var { router } = this.context;
        var nextPath = router.getCurrentQuery().nextPath;
        var email = this.refs.email.getDOMNode().value;
        var password = this.refs.password.getDOMNode().value;
        this.handleValidate(email, password);
        if (!this.state.is_valid) {
            return;
        }
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
                    <div className={this.state.errors.email.length ? 'form-group has-error' : 'form-group '}>
                        <label>Email Address</label>
                        <input type="email" className="form-control"
                               ref="email" placeholder="Enter email"
                               value={this.props.email} autofocus />
                        {this.state.errors.email}
                    </div>
                    <div className={this.state.errors.email.length ? 'form-group has-error' : 'form-group '}>
                        <label>Password</label>
                        <input type="password" className="form-control"
                               ref="password" placeholder="Password" />
                        {this.state.errors.password}
                    </div>
                    <button type="submit" className="btn btn-default">Sign In</button>
                </form>
            </div>
        );
    }
});

var FormElementError = React.createClass({
    render: function() {
        return (
            <p className="help-block">{this.props.error}</p>
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

function validateEmail(email) {
    var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
}
