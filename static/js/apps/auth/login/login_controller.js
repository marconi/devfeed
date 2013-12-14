(function() {
  define(["devfeed", "login_view"], function(Devfeed, LoginView) {
    Devfeed.module("AuthApp.Login", function(Login, Devfeed, Backbone, Marionette, $, _) {
      return Login.Controller = {
        showLogin: function() {
          var loginView;
          loginView = new LoginView.LoginForm();
          loginView.on("auth:login", function(data) {
            var loggingIn;
            loggingIn = Devfeed.request("session:login", data.email, data.password);
            return $.when(loggingIn).done(function(result) {
              if (result && (result.errors != null)) {
                return loginView.triggerMethod("form:data:invalid", result);
              } else {
                Devfeed.execute("show:userinfo");
                return Devfeed.trigger("projects:list");
              }
            });
          });
          return Devfeed.contentRegion.show(loginView);
        }
      };
    });
    return Devfeed.AuthApp.Login.Controller;
  });

}).call(this);
