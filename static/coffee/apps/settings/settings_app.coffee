define ["devfeed", "settings_config_controller"], (Devfeed, SettingsConfigController) ->

  Devfeed.module "SettingsApp", (SettingsApp, Devfeed, Backbone, Marionette, $, _) ->

    class SettingsApp.Router extends Marionette.AppRouter
      appRoutes:
        "settings": "showGeneralSettings"
        "settings/pivotal": "showPivotalSettings"

    API =
      showGeneralSettings: ->
        Devfeed.isLoggedIn -> SettingsConfigController.showGeneralSettings()
      showPivotalSettings: ->
        Devfeed.isLoggedIn -> SettingsConfigController.showPivotalSettings()

    Devfeed.on "settings:general", ->
      Devfeed.navigate("settings")
      API.showGeneralSettings()

    Devfeed.on "settings:pivotal", ->
      Devfeed.navigate("settings/pivotal")
      API.showPivotalSettings()

    Devfeed.addInitializer ->
      new SettingsApp.Router
        controller: API

  return Devfeed.SettingsApp
