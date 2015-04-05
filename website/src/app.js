document.addEventListener('polymer-ready', function() {
    var DEFAULT_ROUTE = "dashboard/overview";
    var template = document.querySelector('template[is="auto-binding"]');
    template.sections = [
        {name: "Dashboard", hash: "dashboard", icon: "dashboard", pages: [
            {name: "Overview", hash: "overview"},
            {name: "Servers", hash: "servers"},
            {name: "Applications", hash: "applications"}
        ]},
        {name: "Settings", hash: "settings", icon: "settings", pages: [
            {name: "Contacts", hash: "contacts"},
            {name: "Servers", hash: "servers"},
            {name: "Time Periods", hash: "timeperiods"}
        ]},
        {name: "Profile", hash: "profile", icon: "account-box", pages: [
            {name: "Reset Password", hash: "password"}
        ]}
    ];
    template.addEventListener('template-bound', function(e) {
        //test
        console.log(auth.loggedIn());
        if (auth.loggedIn()) {
            //test
            console.log("heading off the reservation...");
            this.route = this.route || DEFAULT_ROUTE;
        } else {
            //test
            console.log("yesssss...");
            this.route = "signin";
        }
    });
    template.menuItemSelected = function(e, detail, sender) {
        if (auth.loggedIn()) {
            var sectionTitle = document.getElementById("sectionTitle");
            var titleDivider = document.getElementById("titleDivider");
            var pageTitle = document.getElementById("pageTitle");
            sectionTitle.innerText = detail.item.parentNode.label;
            pageTitle.innerText = detail.item.title;
            titleDivider.style.display = "";
            if (detail.isSelected) {
                document.querySelector('#scaffold').closeDrawer();
            }
        } else {
            this.route = "signin";
        }
    };
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
                    token: Math.random().toString(36).substring(7),
                    id: data.ID
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
