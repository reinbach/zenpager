document.addEventListener('polymer-ready', function() {
    var DEFAULT_ROUTE = "home";
    var template = document.querySelector('template[is="auto-binding"]');
    template.pages = [
        {name: "Home", hash: "home"},
        {name: "Contact", hash: "contact"}
    ];
    template.addEventListener('template-bound', function(e) {
        this.route = this.route || DEFAULT_ROUTE;
    });
});
