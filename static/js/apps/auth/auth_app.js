(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "login_controller", "register_controller", "activation_controller", "password_controller"], function(Devfeed, LoginController, RegisterController, ActivationController, PasswordController) {
    Devfeed.module("AuthApp", function(AuthApp, Devfeed, Backbone, Marionette, $, _) {
      var API, _ref;
      AuthApp.Router = (function(_super) {
        __extends(Router, _super);

        function Router() {
          _ref = Router.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Router.prototype.appRoutes = {
          "login": "showLogin",
          "register": "showRegister",
          "activation/resend": "showActivation",
          "password/retrieve": "showForgotPassword"
        };

        return Router;

      })(Marionette.AppRouter);
      API = {
        showLogin: function() {
          return LoginController.showLogin();
        },
        showRegister: function() {
          return RegisterController.showRegister();
        },
        showActivation: function() {
          return ActivationController.showActivation();
        },
        showForgotPassword: function() {
          return PasswordController.showForgotPassword();
        }
      };
      Devfeed.on("auth:login:show", function() {
        Devfeed.navigate("login");
        return API.showLogin();
      });
      Devfeed.on("auth:register:show", function() {
        Devfeed.navigate("register");
        return API.showRegister();
      });
      Devfeed.on("auth:activation:show", function() {
        Devfeed.navigate("activation/resend");
        return API.showActivation();
      });
      Devfeed.on("auth:forgotpass:show", function() {
        Devfeed.navigate("password/retrieve");
        return API.showForgotPassword();
      });
      return Devfeed.addInitializer(function() {
        return new AuthApp.Router({
          controller: API
        });
      });
    });
    return Devfeed.AuthApp;
  });

}).call(this);
