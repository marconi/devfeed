(function() {
  define(["devfeed", "header_controller"], function(Devfeed, HeaderController) {
    Devfeed.module("HeaderApp", function(HeaderApp, Devfeed, Backbone, Marionette, $, _) {
      var API;
      API = {
        showHeader: function() {
          return HeaderController.showHeader();
        },
        removeUserinfo: function() {
          return HeaderController.removeUserinfo();
        },
        showUserinfo: function() {
          return HeaderController.showUserinfo();
        }
      };
      Devfeed.on("auth:logout", function() {
        return API.removeUserinfo();
      });
      Devfeed.commands.setHandler("show:userinfo", function() {
        return API.showUserinfo();
      });
      return HeaderApp.on("start", function() {
        var loggingIn;
        API.showHeader();
        loggingIn = Devfeed.request("session:isloggedin");
        return $.when(loggingIn).done(function(isLoggedIn) {
          if (isLoggedIn) {
            return API.showUserinfo();
          }
        });
      });
    });
    return Devfeed.HeaderApp;
  });

}).call(this);
