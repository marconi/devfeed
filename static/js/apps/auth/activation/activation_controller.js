(function() {
  define(["devfeed", "activation_view"], function(Devfeed, ActivationView) {
    Devfeed.module("AuthApp.Activation", function(Activation, Devfeed, Backbone, Marionette, $, _) {
      return Activation.Controller = {
        showActivation: function() {
          var resendView;
          resendView = new ActivationView.ResendForm();
          resendView.on("auth:activation", function(data) {
            var resending;
            resending = Devfeed.request("session:activation:resend", data.email);
            return $.when(resending).done(function(result) {
              if (result && (result.errors != null)) {
                return resendView.triggerMethod("form:data:invalid", result);
              } else {
                return resendView.triggerMethod("form:data:valid", result);
              }
            });
          });
          return Devfeed.contentRegion.show(resendView);
        }
      };
    });
    return Devfeed.AuthApp.Activation.Controller;
  });

}).call(this);
