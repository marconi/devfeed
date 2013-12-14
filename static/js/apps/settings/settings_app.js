(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "jquery_form", "settings_config_controller"], function(Devfeed, jQueryForm, SettingsConfigController) {
    Devfeed.module("SettingsApp", function(SettingsApp, Devfeed, Backbone, Marionette, $, _) {
      var API, _ref;
      SettingsApp.Router = (function(_super) {
        __extends(Router, _super);

        function Router() {
          _ref = Router.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Router.prototype.appRoutes = {
          "settings": "showGeneralSettings",
          "settings/pivotal": "showPivotalSettings"
        };

        return Router;

      })(Marionette.AppRouter);
      API = {
        showGeneralSettings: function() {
          return Devfeed.isLoggedIn(function() {
            return SettingsConfigController.showGeneralSettings();
          });
        },
        showPivotalSettings: function() {
          return Devfeed.isLoggedIn(function() {
            return SettingsConfigController.showPivotalSettings();
          });
        }
      };
      Devfeed.on("settings:general", function() {
        Devfeed.navigate("settings");
        return API.showGeneralSettings();
      });
      Devfeed.on("settings:pivotal", function() {
        Devfeed.navigate("settings/pivotal");
        return API.showPivotalSettings();
      });
      return Devfeed.addInitializer(function() {
        return new SettingsApp.Router({
          controller: API
        });
      });
    });
    return Devfeed.SettingsApp;
  });

}).call(this);
