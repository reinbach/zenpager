var Router = ReactRouter;
var Route = Router.Route;
var Link = Router.Link;
var RouteHandler = Router.RouteHandler;
var DefaultRoute = Router.DefaultRoute;
var NotFoundRoute = Router.NotFoundRoute;

var App = React.createClass({
    render: function() {
        return (
            <div className="site-wrapper">
                <div className="site-wrapper-inner">
                    <div className="cover-container">
                        <div className="masthead clearfix">
                            <div className="inner">
                                <h3 className="masthead-brand"><a href="/">ZenPager</a></h3>
                                <nav>
                                    <ul className="nav masthead-nav">
                                        <li>
                                            <Link to="home">Home</Link>
                                        </li>
                                        <li>
                                            <Link to="contact">Contact</Link>
                                        </li>
                                        <li>
                                            <a href="/dashboard/">Dashboard</a>
                                        </li>
                                    </ul>
                                </nav>
                            </div>
                        </div>

                        <RouteHandler />

                        <div className="mastfoot">
                            <div className="inner">
                                <p>Created by <a href="http://www.ironlabs.com">IRON Labs, Inc.</a></p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
});

var Home = React.createClass({
    render: function() {
        return (
            <div className="inner cover">
                <h1 className="cover-heading">Monitor What You Will</h1>
                <p className="lead">A single place to monitor your systems, track the performance of your applications, and be alerted when it matters.</p>
                <p className="lead">
                    <a href="/dashboard/" className="btn btn-lg btn-success">Dashboard</a>
                </p>
            </div>
        );
    }
});

var Contact = React.createClass({
    render: function() {
        return (
            <div className="inner cover">
                <h1 className="cover-heading">Contact</h1>
            </div>
        );
    }
});

var NotFound = React.createClass({
    render: function() {
        return (
            <div className="inner cover">
                <h1 className="cover-heading">404 Not Found</h1>
            </div>
        );
    }
});

var routes = (
    <Route name="app" path="/" handler={App}>
        <Route name="home" handler={Home} />
        <Route name="contact" handler={Contact} />
        <NotFoundRoute handler={NotFound} />
        <DefaultRoute handler={Home} />
    </Route>
);

Router.run(routes, function(Handler) {
    React.render(<Handler />, document.body);
});
