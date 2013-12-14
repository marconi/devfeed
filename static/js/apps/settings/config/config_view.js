(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "alert", "common_view", "tpl!apps/settings/config/templates/general.tpl", "tpl!apps/settings/config/templates/pivotal.tpl"], function(Devfeed, Alert, CommonView, generalTpl, pivotalTpl) {
    Devfeed.module("SettingsApp.Config.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1;
      View.General = (function(_super) {
        __extends(General, _super);

        function General() {
          _ref = General.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        General.prototype.id = "general-settings";

        General.prototype.className = "row collapse";

        General.prototype.template = generalTpl;

        General.prototype.events = {
          "click .update": "updateClicked",
          "click .general .title a": "generalClicked",
          "click .pivotal .title a": "pivotalClicked"
        };

        General.prototype.ui = {
          form: "form",
          passwordInput: "input[name=password]",
          primaryBtn: ".update"
        };

        General.prototype.updateClicked = function(e) {
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          return this.trigger("settings:general:update", this.ui.form);
        };

        General.prototype.generalClicked = function(e) {
          return e.preventDefault();
        };

        General.prototype.pivotalClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("settings:pivotal");
        };

        General.prototype.clearFields = function() {
          return this.ui.passwordInput.val("");
        };

        return General;

      })(CommonView.FormViewMixin);
      return View.Pivotal = (function(_super) {
        __extends(Pivotal, _super);

        function Pivotal() {
          _ref1 = Pivotal.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        Pivotal.prototype.id = "pivotal-settings";

        Pivotal.prototype.className = "row collapse";

        Pivotal.prototype.template = pivotalTpl;

        Pivotal.prototype.events = {
          "click .update": "updateClicked",
          "click .general .title a": "generalClicked",
          "click .pivotal .title a": "pivotalClicked"
        };

        Pivotal.prototype.ui = {
          primaryBtn: ".update"
        };

        Pivotal.prototype.updateClicked = function(e) {
          var data;
          e.preventDefault();
          if (this.ui.primaryBtn.hasClass("disabled")) {
            return;
          }
          this.clearErrors();
          this.showPreloader();
          data = Backbone.Syphon.serialize(this);
          return this.trigger("settings:pivotal:update", data);
        };

        Pivotal.prototype.generalClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("settings:general");
        };

        Pivotal.prototype.pivotalClicked = function(e) {
          return e.preventDefault();
        };

        Pivotal.prototype.clearFields = function() {};

        return Pivotal;

      })(CommonView.FormViewMixin);
    });
    return Devfeed.SettingsApp.Config.View;
  });

}).call(this);
