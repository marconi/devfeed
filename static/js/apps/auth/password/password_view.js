(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_view", "tpl!apps/auth/password/templates/retrieve.tpl", "syphon", "jquery_spin"], function(Devfeed, CommonView, retrieveTpl) {
    Devfeed.module("AuthApp.Password.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref;
      return View.ForgotPasswordForm = (function(_super) {
        __extends(ForgotPasswordForm, _super);

        function ForgotPasswordForm() {
          _ref = ForgotPasswordForm.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        ForgotPasswordForm.prototype.id = "forgot-password";

        ForgotPasswordForm.prototype.className = "row collapse";

        ForgotPasswordForm.prototype.template = retrieveTpl;

        ForgotPasswordForm.prototype.events = {
          "click .retrieve": "retrieveClicked",
          "click .login": "loginClicked"
        };

        ForgotPasswordForm.prototype.ui = {
          emailInput: "input[name=email]",
          primaryBtn: ".retrieve"
        };

        ForgotPasswordForm.prototype.retrieveClicked = function(e) {
          var data;
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          data = Backbone.Syphon.serialize(this);
          return this.trigger("auth:forgotpass", data);
        };

        ForgotPasswordForm.prototype.loginClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:login:show");
        };

        ForgotPasswordForm.prototype.clearFields = function() {
          return this.ui.emailInput.val("").focus();
        };

        return ForgotPasswordForm;

      })(CommonView.FormViewMixin);
    });
    return Devfeed.AuthApp.Password.View;
  });

}).call(this);
