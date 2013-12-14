(function() {
  define(["devfeed", "settings_config_view"], function(Devfeed, SettingsConfigView) {
    Devfeed.module("SettingsApp.Config", function(Config, Devfeed, Backbone, Marionette, $, _) {
      return Config.Controller = {
        showGeneralSettings: function() {
          var generalSettingsView, userSettings;
          userSettings = Devfeed.request("user:session");
          generalSettingsView = new SettingsConfigView.General({
            model: userSettings
          });
          generalSettingsView.on("settings:general:update", function(form) {
            var updating;
            updating = Devfeed.request("settings:update", form, "general");
            return $.when(updating).done(function(result) {
              if (result && (result.errors != null)) {
                return generalSettingsView.triggerMethod("form:data:invalid", result);
              } else {
                return generalSettingsView.triggerMethod("form:data:valid", result);
              }
            });
          });
          return Devfeed.contentRegion.show(generalSettingsView);
        },
        showPivotalSettings: function() {
          var pivotalSettingsView, userSettings;
          userSettings = Devfeed.request("user:session");
          pivotalSettingsView = new SettingsConfigView.Pivotal({
            model: userSettings
          });
          pivotalSettingsView.on("settings:pivotal:update", function(data) {
            var updating;
            updating = Devfeed.request("settings:update:pivotal", data);
            return $.when(updating).done(function(result) {
              if (result && (result.errors != null)) {
                return pivotalSettingsView.triggerMethod("form:data:invalid", result);
              } else {
                return pivotalSettingsView.triggerMethod("form:data:valid", result);
              }
            });
          });
          return Devfeed.contentRegion.show(pivotalSettingsView);
        }
      };
    });
    return Devfeed.SettingsApp.Config.Controller;
  });

}).call(this);
