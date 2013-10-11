define ["devfeed", "settings_config_view"], (Devfeed, SettingsConfigView) ->

  Devfeed.module "SettingsApp.Config", (Config, Devfeed, Backbone, Marionette, $, _) ->

    Config.Controller =
      showGeneralSettings: ->
        userSettings = Devfeed.request("user:session")
        generalSettingsView = new SettingsConfigView.General
          model: userSettings
        generalSettingsView.on "settings:general:update", (data) ->
          updating = Devfeed.request("settings:update", data, "general")
          $.when(updating).done (result) ->
            if result and result.errors?
              generalSettingsView.triggerMethod("form:data:invalid", result)
            else
              generalSettingsView.triggerMethod("form:data:valid", result)
        Devfeed.contentRegion.show(generalSettingsView)

      showPivotalSettings: ->
        userSettings = Devfeed.request("user:session")
        pivotalSettingsView = new SettingsConfigView.Pivotal
          model: userSettings
        pivotalSettingsView.on "settings:pivotal:update", (data) ->
          updating = Devfeed.request("settings:update", data, "pivotal")
          $.when(updating).done (result) ->
            if result and result.errors?
              pivotalSettingsView.triggerMethod("form:data:invalid", result)
            else
              pivotalSettingsView.triggerMethod("form:data:valid", result)
        Devfeed.contentRegion.show(pivotalSettingsView)

  return Devfeed.SettingsApp.Config.Controller
