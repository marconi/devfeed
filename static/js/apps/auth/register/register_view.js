(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_view", "tpl!apps/auth/register/templates/register.tpl", "syphon", "jquery_spin"], function(Devfeed, CommonView, registerTpl) {
    Devfeed.module("AuthApp.Register.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref;
      return View.RegisterForm = (function(_super) {
        __extends(RegisterForm, _super);

        function RegisterForm() {
          _ref = RegisterForm.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        RegisterForm.prototype.id = "register";

        RegisterForm.prototype.className = "row collapse";

        RegisterForm.prototype.template = registerTpl;

        RegisterForm.prototype.events = {
          "click .login": "loginClicked",
          "click .signup": "signupClicked"
        };

        RegisterForm.prototype.ui = {
          nameInput: "input[name=name]",
          emailInput: "input[name=email]",
          passwordInput: "input[name=password]",
          primaryBtn: ".signup"
        };

        RegisterForm.prototype.loginClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:login:show");
        };

        RegisterForm.prototype.signupClicked = function(e) {
          var data;
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          data = Backbone.Syphon.serialize(this);
          return this.trigger("auth:register", data);
        };

        RegisterForm.prototype.clearFields = function() {
          this.ui.nameInput.val("");
          this.ui.emailInput.val("");
          this.ui.passwordInput.val("");
          return this.ui.nameInput.focus();
        };

        return RegisterForm;

      })(CommonView.FormViewMixin);
    });
    return Devfeed.AuthApp.Register.View;
  });

}).call(this);
