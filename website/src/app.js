document.addEventListener('polymer-ready', function() {
    var DEFAULT_ROUTE = "dashboard";
    var template = document.querySelector('template[is="auto-binding"]');
    template.pages = [
        {name: "Dashboard", hash: "dashboard", icon: "dashboard"},
        {name: "Settings", hash: "settings", icon: "settings"},
        {name: "Profile", hash: "profile", icon: "profile"}
    ];
    template.addEventListener('template-bound', function(e) {
        this.route = this.route || DEFAULT_ROUTE;
    });

    template.menuItemSelected = function(e, detail, sender) {
        var title = document.getElementById("title");
        title.innerText = detail.item.title;
        if (detail.isSelected) {
            document.querySelector('#scaffold').closeDrawer();
        }
    };
});
