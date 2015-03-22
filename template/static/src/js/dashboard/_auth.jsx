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
            email: '',
            password: ''
        };
    },
    validationEmailState: function() {
        if (this.state.email.length == "") {
            return "error";
        } else if (validateEmail(this.state.email) !== true) {
            return "error";
        }
        return "success";
    },
    validationPasswordState: function() {
        if (this.state.password.length > 8) {
            return "success";
        } else if (this.state.password.length > 0) {
            return "error";
        }
    },
    handleChange: function() {
        this.setState({
            error: this.state.error,
            email: this.refs.email.getValue(),
            password: this.refs.password.getValue()
        });
    },
    handleSubmit: function(event) {
        event.preventDefault();
        var { router } = this.context;
        var nextPath = router.getCurrentQuery().nextPath;
        var email = this.refs.email.getDOMNode().value;
        var password = this.refs.password.getDOMNode().value;
        // Prevent form being submitted till elements are in valid state
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
        var Input = ReactBootstrap.Input,
            Button = ReactBootstrap.Button;
        return (
            <div className="col-md-3 col-md-offset-3">
                <h1>Sign In</h1>
                <form onSubmit={this.handleSubmit} className="text-left">
                    <Input label="Email Address" type="email" ref="email"
                           placeholder="Enter email" value={this.state.email}
                           autoFocus hasFeedback bsStyle={this.validationEmailState()}
                           onChange={this.handleChange} />
                    <Input label="Password" type="password" ref="password"
                           placeholder="Password" value={this.state.password}
                           hasFeedback bsStyle={this.validationPasswordState()}
                           onChange={this.handleChange} />
                   <Button type="submit">Sign In</Button>
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
        if (email === undefined || pass == undefined) {
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
