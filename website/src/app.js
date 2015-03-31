document.addEventListener('polymer-ready', function() {
    var DEFAULT_ROUTE = "dashboard/overview";
    var template = document.querySelector('template[is="auto-binding"]');
    template.sections = [
        {name: "Dashboard", hash: "dashboard", icon: "dashboard", pages: [
            {name: "Overview", hash: "overview", icon: "overview"},
            {name: "Servers", hash: "servers", icon: "servers"},
            {name: "Applications", hash: "applications", icon: "applications"}
        ]},
        {name: "Settings", hash: "settings", icon: "settings", pages: [
            {name: "Contacts", hash: "contacts", icon: "contacts"},
            {name: "Servers", hash: "servers", icon: "servers"},
            {name: "Time Periods", hash: "timeperiods", icon: "timeperiods"}
        ]},
        {name: "Profile", hash: "profile", icon: "account-box", pages: [
            {name: "Reset Password", hash: "password", icon: "password"}
        ]}
    ];
    template.addEventListener('template-bound', function(e) {
        this.route = this.route || DEFAULT_ROUTE;
    });

    template.menuItemSelected = function(e, detail, sender) {
        var sectionTitle = document.getElementById("sectionTitle");
        var pageTitle = document.getElementById("pageTitle");
        sectionTitle.innerText = detail.item.parentNode.label;
        pageTitle.innerText = detail.item.title;
        if (detail.isSelected) {
            document.querySelector('#scaffold').closeDrawer();
        }
    };
});
