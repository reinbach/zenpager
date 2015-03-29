document.addEventListener('polymer-ready', function() {
    var DEFAULT_ROUTE = "home";
    var template = document.querySelector('template[is="auto-binding"]');
    template.pages = [
        {name: "Home", hash: "home"},
        {name: "Contact", hash: "contact"},
        {name: "Dashboard", hash: "dashboard"}
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
