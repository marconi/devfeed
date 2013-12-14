(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_view", "tpl!apps/auth/activation/templates/resend.tpl", "syphon", "jquery_spin"], function(Devfeed, CommonView, resendTpl) {
    Devfeed.module("AuthApp.Activation.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref;
      return View.ResendForm = (function(_super) {
        __extends(ResendForm, _super);

        function ResendForm() {
          _ref = ResendForm.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        ResendForm.prototype.id = "activation";

        ResendForm.prototype.className = "row collapse";

        ResendForm.prototype.template = resendTpl;

        ResendForm.prototype.events = {
          "click .send": "sendClicked",
          "click .login": "loginClicked"
        };

        ResendForm.prototype.ui = {
          emailInput: "input[name=email]",
          primaryBtn: ".send"
        };

        ResendForm.prototype.sendClicked = function(e) {
          var data;
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          data = Backbone.Syphon.serialize(this);
          return this.trigger("auth:activation", data);
        };

        ResendForm.prototype.loginClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("auth:login:show");
        };

        ResendForm.prototype.clearFields = function() {
          return this.ui.emailInput.val("").focus();
        };

        return ResendForm;

      })(CommonView.FormViewMixin);
    });
    return Devfeed.AuthApp.Activation.View;
  });

}).call(this);
