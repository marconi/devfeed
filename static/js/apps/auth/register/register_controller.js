(function() {
  define(["devfeed", "register_view"], function(Devfeed, RegisterView) {
    Devfeed.module("AuthApp.Register", function(Register, Devfeed, Backbone, Marionette, $, _) {
      return Register.Controller = {
        showRegister: function() {
          var registerView;
          registerView = new RegisterView.RegisterForm();
          registerView.on("auth:register", function(data) {
            var registering;
            registering = Devfeed.request("session:register", data.name, data.email, data.password);
            return $.when(registering).done(function(result) {
              if (result && (result.errors != null)) {
                return registerView.triggerMethod("form:data:invalid", result);
              } else {
                return registerView.triggerMethod("form:data:valid", result);
              }
            });
          });
          return Devfeed.contentRegion.show(registerView);
        }
      };
    });
    return Devfeed.AuthApp.Register.Controller;
  });

}).call(this);
