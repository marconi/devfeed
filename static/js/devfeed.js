(function() {
  define(["backbone", "marionette", "underscore"], function(Backbone, Marionette, _) {
    var Devfeed;
    Devfeed = new Marionette.Application();
    Devfeed.addRegions({
      headerRegion: "#header-region",
      contentRegion: "#content-region"
    });
    Devfeed.navigate = function(route, options) {
      return Backbone.history.navigate(route, options || {});
    };
    Devfeed.redirect = function(route) {
      return window.location = route;
    };
    Devfeed.getCurrentRoute = function() {
      return Backbone.history.fragment;
    };
    Devfeed.messages = {
      500: "Something went wrong while trying to process your request and we are already looking at it"
    };
    Devfeed.isLoggedIn = function(authCallback, unauthCallback) {
      var loggingIn;
      if (authCallback == null) {
        authCallback = null;
      }
      if (unauthCallback == null) {
        unauthCallback = null;
      }
      loggingIn = Devfeed.request("session:isloggedin");
      return $.when(loggingIn).done(function(isLoggedIn) {
        if (isLoggedIn && authCallback) {
          return authCallback();
        } else if (!isLoggedIn && unauthCallback) {
          return unauthCallback();
        }
      });
    };
    Devfeed.on("home", function() {
      return Devfeed.isLoggedIn(function() {
        return Devfeed.trigger("projects:list");
      }, function() {
        return Devfeed.trigger("auth:login:show");
      });
    });
    Devfeed.on("initialize:after", function() {
      $(document).foundation();
      $('#notification').miniNotification({
        time: 5000,
        opacity: 1,
        closeButton: true,
        closeButtonText: "&times;"
      });
      if (Backbone.history) {
        Backbone.history.start({
          pushState: true
        });
      }
      if (Devfeed.getCurrentRoute() === "") {
        return Devfeed.trigger("home");
      }
    });
    return Devfeed;
  });

}).call(this);
