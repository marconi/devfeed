(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_view", "tpl!apps/auth/login/templates/login.tpl", "syphon", "jquery_spin"], function(Devfeed, CommonView, loginTpl) {
    Devfeed.module("AuthApp.Login.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref;
      return View.LoginForm = (function(_super) {
        __extends(LoginForm, _super);

        function LoginForm() {
          _ref = LoginForm.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        LoginForm.prototype.id = "login";

        LoginForm.prototype.className = "row collapse";

        LoginForm.prototype.template = loginTpl;

        LoginForm.prototype.events = {
          "click .signin": "signinClicked",
          "click .register": "registerClicked",
          "click .forgot-pass": "forgotPassClicked",
          "click .resend-activation": "resendActivationClicked"
        };

        LoginForm.prototype.ui = {
          primaryBtn: ".signin"
        };

        LoginForm.prototype.signinClicked = function(e) {
          var data;
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          data = Backbone.Syphon.serialize(this);
          return this.trigger("auth:login", data);
        };

        LoginForm.prototype.registerClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:register:show");
        };

        LoginForm.prototype.forgotPassClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:forgotpass:show");
        };

        LoginForm.prototype.resendActivationClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:activation:show");
        };

        return LoginForm;

      })(CommonView.FormViewMixin);
    });
    return Devfeed.AuthApp.Login.View;
  });

}).call(this);
