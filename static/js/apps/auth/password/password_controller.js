(function() {
  define(["devfeed", "password_view"], function(Devfeed, PasswordView) {
    Devfeed.module("AuthApp.Password", function(Password, Devfeed, Backbone, Marionette, $, _) {
      return Password.Controller = {
        showForgotPassword: function() {
          var forgotPassView;
          forgotPassView = new PasswordView.ForgotPasswordForm();
          forgotPassView.on("auth:forgotpass", function(data) {
            var retrieving;
            retrieving = Devfeed.request("session:password:retrieve", data.email);
            return $.when(retrieving).done(function(result) {
              if (result && (result.errors != null)) {
                return forgotPassView.triggerMethod("form:data:invalid", result);
              } else {
                return forgotPassView.triggerMethod("form:data:valid", result);
              }
            });
          });
          return Devfeed.contentRegion.show(forgotPassView);
        }
      };
    });
    return Devfeed.AuthApp.Password.Controller;
  });

}).call(this);
