(function() {
    Polymer("login-button", {
        ready: function() {
            if (localStorage.token !== undefined) {
                this.hash = "signout";
                this.title = "Sign Out";
            } else {
                this.hash = "signin";
                this.title = "Sign In";
            }
        }
    });
}());
